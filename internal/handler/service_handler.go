package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/2group/2sales.core-service/internal/grpc"
	servicev1 "github.com/2group/2sales.core-service/pkg/gen/go/service"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/proto"
)

type ServiceHandler struct {
	service *grpc.ServiceClient
}

func NewServiceHandler(service *grpc.ServiceClient) *ServiceHandler {
	return &ServiceHandler{service: service}
}

func (h *ServiceHandler) CreateService(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "service_handler").
		Str("method", "CreateService").
		Logger()

	log.Info().Msg("request_received")

	req := &servicev1.CreateServiceRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid payload"))
		return
	}

	log.Info().Interface("request", req).Msg("calling_service_microservice")
	resp, err := h.service.Api.CreateService(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int64("service_id", resp.GetService().GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusCreated, resp)
}

func (h *ServiceHandler) GetService(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "service_handler").
		Str("method", "GetService").
		Logger()

	log.Info().Msg("request_received")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error().Str("id", idStr).Err(err).Msg("invalid_id")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Int64("id", id).Msg("calling_service_microservice")
	req := &servicev1.GetServiceRequest{Id: id}
	resp, err := h.service.Api.GetService(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusNotFound, err)
		return
	}

	log.Info().Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *ServiceHandler) DeleteService(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "service_handler").
		Str("method", "DeleteService").
		Logger()

	log.Info().Msg("request_received")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error().Str("id", idStr).Err(err).Msg("invalid_id")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Int64("id", id).Msg("calling_service_microservice")
	req := &servicev1.DeleteServiceRequest{Id: id}
	resp, err := h.service.Api.DeleteService(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *ServiceHandler) PartialUpdateService(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "service_handler").
		Str("method", "PartialUpdateService").
		Logger()

	log.Info().Msg("request_received")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error().Str("id", idStr).Err(err).Msg("invalid_id")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &servicev1.PartialUpdateServiceRequest{
		Service: &servicev1.Service{Id: proto.Int64(id)},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Interface("request", req).Msg("calling_service_microservice")
	resp, err := h.service.Api.PartialUpdateService(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *ServiceHandler) UpdateService(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "service_handler").
		Str("method", "UpdateService").
		Logger()

	log.Info().Msg("request_received")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error().Str("id", idStr).Err(err).Msg("invalid_id")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &servicev1.UpdateServiceRequest{
		Service: &servicev1.Service{Id: proto.Int64(id)},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Interface("request", req).Msg("calling_service_microservice")
	resp, err := h.service.Api.UpdateService(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *ServiceHandler) ListServices(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "service_handler").
		Str("method", "ListServices").
		Logger()

	log.Info().Msg("request_received")

	query := r.URL.Query()
	limitStr := query.Get("limit")
	offsetStr := query.Get("offset")
	orgIDStr := query.Get("organization_id")
	branchIDStr := query.Get("branch_id")
	searchTextStr := query.Get("search_text")
	createdAtFromStr := query.Get("created_at_from")
	createdAtToStr := query.Get("created_at_to")
	priceFromStr := query.Get("price_from")
	priceToStr := query.Get("price_to")

	limit, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		log.Warn().Str("limit", limitStr).Err(err).Msg("invalid_limit")
		limit = 20
	}
	offset, err := strconv.ParseInt(offsetStr, 10, 32)
	if err != nil {
		log.Warn().Str("offset", offsetStr).Err(err).Msg("invalid_offset")
		offset = 0
	}

	var organizationID *int64
	if orgIDStr != "" {
		if id, err := strconv.ParseInt(orgIDStr, 10, 64); err == nil {
			organizationID = proto.Int64(id)
		}
	}

	var branchID *int64
	if branchIDStr != "" {
		if id, err := strconv.ParseInt(branchIDStr, 10, 64); err == nil {
			branchID = proto.Int64(id)
		}
	}

	var searchText *string
	if searchTextStr != "" {
		searchText = proto.String(searchTextStr)
	}

	var createdAtFrom *string
	if createdAtFromStr != "" {
		if _, err := time.Parse(time.RFC3339, createdAtFromStr); err == nil {
			createdAtFrom = proto.String(createdAtFromStr)
		} else if t, err := time.Parse("2006-01-02", createdAtFromStr); err == nil {
			s := t.Format(time.RFC3339)
			createdAtFrom = proto.String(s)
		} else {
			log.Warn().Str("created_at_from", createdAtFromStr).Msg("invalid_created_at_from")
		}
	}

	var createdAtTo *string
	if createdAtToStr != "" {
		if _, err := time.Parse(time.RFC3339, createdAtToStr); err == nil {
			createdAtTo = proto.String(createdAtToStr)
		} else if t, err := time.Parse("2006-01-02", createdAtToStr); err == nil {
			t = t.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
			s := t.Format(time.RFC3339)
			createdAtTo = proto.String(s)
		} else {
			log.Warn().Str("created_at_to", createdAtToStr).Msg("invalid_created_at_to")
		}
	}

	var priceFrom, priceTo *float64
	if priceFromStr != "" {
		if val, err := strconv.ParseFloat(priceFromStr, 64); err == nil {
			priceFrom = proto.Float64(val)
		} else {
			log.Warn().Str("price_from", priceFromStr).Err(err).Msg("invalid_price_from")
		}
	}
	if priceToStr != "" {
		if val, err := strconv.ParseFloat(priceToStr, 64); err == nil {
			priceTo = proto.Float64(val)
		} else {
			log.Warn().Str("price_to", priceToStr).Err(err).Msg("invalid_price_to")
		}
	}

	req := &servicev1.ListServicesRequest{
		OrganizationId: organizationID,
		BranchId:       branchID,
		SearchText:     searchText,
		CreatedAtFrom:  createdAtFrom,
		CreatedAtTo:    createdAtTo,
		PriceFrom:      priceFrom,
		PriceTo:        priceTo,
		Limit:          int32(limit),
		Offset:         int32(offset),
	}

	log.Info().
		Int64("limit", limit).
		Int64("offset", offset).
		Str("search_text", searchTextStr).
		Str("created_at_from", createdAtFromStr).
		Str("created_at_to", createdAtToStr).
		Str("price_from", priceFromStr).
		Str("price_to", priceToStr).
		Msg("calling_service_microservice")

	if organizationID != nil {
		log = log.With().Int64("organization_id", *organizationID).Logger()
	}
	if branchID != nil {
		log = log.With().Int64("branch_id", *branchID).Logger()
	}

	resp, err := h.service.Api.ListServices(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int("services_count", len(resp.Services)).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}
