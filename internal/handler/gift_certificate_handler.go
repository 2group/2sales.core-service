package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	customerv1 "github.com/2group/2sales.core-service/pkg/gen/go/customer"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type GiftCertificateHandler struct {
	client *grpc.CustomerClient
}

func NewGiftCertificateHandler(client *grpc.CustomerClient) *GiftCertificateHandler {
	return &GiftCertificateHandler{client: client}
}

func (h *GiftCertificateHandler) CreateGiftCertificate(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "gift_certificate_handler").
		Str("method", "CreateGiftCertificate").
		Logger()

	log.Info().Msg("request_received")

	req := &customerv1.CreateGiftCertificateRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Debug().Interface("request", req).Msg("calling_customer_service")

	resp, err := h.client.Api.CreateGiftCertificate(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Int64("certificate_id", resp.Certificate.GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusCreated, resp)
}

func (h *GiftCertificateHandler) GetGiftCertificate(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "gift_certificate_handler").
		Str("method", "GetGiftCertificate").
		Logger()

	log.Info().Msg("request_received")

	certificateIDStr := chi.URLParam(r, "certificate_id")
	certificateID, err := strconv.ParseInt(certificateIDStr, 10, 64)
	if err != nil {
		log.Error().Str("certificate_id", certificateIDStr).Err(err).Msg("invalid_certificate_id")
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid certificate_id"))
		return
	}

	req := &customerv1.GetGiftCertificateRequest{Id: certificateID}
	log.Debug().Int64("certificate_id", certificateID).Msg("calling_customer_service")

	resp, err := h.client.Api.GetGiftCertificate(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusNotFound, err)
		return
	}

	log.Info().Int64("certificate_id", resp.Certificate.GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *GiftCertificateHandler) ListGiftCertificates(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "gift_certificate_handler").
		Str("method", "ListGiftCertificates").
		Logger()

	log.Info().Msg("request_received")

	query := r.URL.Query()

	var (
		customerID     *int64
		organizationID *int64
		limit          int32 = 20
		offset         int32 = 0
	)

	if customerIDStr := query.Get("customer_id"); customerIDStr != "" {
		if id, err := strconv.ParseInt(customerIDStr, 10, 64); err == nil {
			customerID = &id
		} else {
			log.Warn().Str("customer_id", customerIDStr).Err(err).Msg("invalid_customer_id")
		}
	}

	if organizationIDStr := query.Get("organization_id"); organizationIDStr != "" {
		if id, err := strconv.ParseInt(organizationIDStr, 10, 64); err == nil {
			organizationID = &id
		} else {
			log.Warn().Str("organization_id", organizationIDStr).Err(err).Msg("invalid_organization_id")
		}
	}

	if limitStr := query.Get("limit"); limitStr != "" {
		if l, err := strconv.ParseInt(limitStr, 10, 32); err == nil {
			limit = int32(l)
		} else {
			log.Warn().Str("limit", limitStr).Err(err).Msg("invalid_limit")
		}
	}

	if offsetStr := query.Get("offset"); offsetStr != "" {
		if o, err := strconv.ParseInt(offsetStr, 10, 32); err == nil {
			offset = int32(o)
		} else {
			log.Warn().Str("offset", offsetStr).Err(err).Msg("invalid_offset")
		}
	}

	req := &customerv1.ListGiftCertificatesRequest{
		CustomerId:     customerID,
		OrganizationId: organizationID,
		Limit:          limit,
		Offset:         offset,
	}

	log.Debug().Interface("request", req).Msg("calling_customer_service")

	resp, err := h.client.Api.ListGiftCertificates(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int("certificates_count", len(resp.Certificates)).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *GiftCertificateHandler) ListGiftCertificateLabels(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "gift_certificate_handler").
		Str("method", "ListGiftCertificateLabels").
		Logger()

	log.Info().Msg("request_received")

	resp, err := h.client.Api.ListGiftCertificateLabels(r.Context(), &customerv1.ListGiftCertificateLabelsRequest{})
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int("labels_count", len(resp.Labels)).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *GiftCertificateHandler) ListGiftCertificateIcons(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "gift_certificate_handler").
		Str("method", "ListGiftCertificateIcons").
		Logger()

	log.Info().Msg("request_received")

	resp, err := h.client.Api.ListGiftCertificateIcons(r.Context(), &customerv1.ListGiftCertificateIconsRequest{})
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int("icons_count", len(resp.Icons)).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *GiftCertificateHandler) ListGiftCertificateBackgrounds(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "gift_certificate_handler").
		Str("method", "ListGiftCertificateBackgrounds").
		Logger()

	log.Info().Msg("request_received")

	resp, err := h.client.Api.ListGiftCertificateBackgrounds(r.Context(), &customerv1.ListGiftCertificateBackgroundsRequest{})
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int("backgrounds_count", len(resp.Backgrounds)).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *GiftCertificateHandler) ListGiftCertificateDesigns(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "gift_certificate_handler").
		Str("method", "ListGiftCertificateDesigns").
		Logger()

	log.Info().Msg("request_received")

	resp, err := h.client.Api.ListGiftCertificateDesigns(r.Context(), &customerv1.ListGiftCertificateDesignsRequest{})
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().
		Int("labels", len(resp.Labels)).
		Int("icons", len(resp.Icons)).
		Int("backgrounds", len(resp.Backgrounds)).
		Msg("succeeded")

	json.WriteJSON(w, http.StatusOK, resp)
}
