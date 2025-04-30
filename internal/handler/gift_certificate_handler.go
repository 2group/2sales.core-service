package handler

//
//import (
//	"errors"
//	"github.com/2group/2sales.core-service/pkg/middleware"
//	"net/http"
//	"strconv"
//
//	"github.com/2group/2sales.core-service/internal/grpc"
//	customerv1 "github.com/2group/2sales.core-service/pkg/gen/go/customer"
//	"github.com/2group/2sales.core-service/pkg/json"
//	"github.com/go-chi/chi/v5"
//)
//
//type GiftCertificateHandler struct {
//	client *grpc.CustomerClient
//}
//
//func NewGiftCertificateHandler(client *grpc.CustomerClient) *GiftCertificateHandler {
//	return &GiftCertificateHandler{client: client}
//}
//
//func (h *GiftCertificateHandler) CreateGiftCertificate(w http.ResponseWriter, r *http.Request) {
//	log := middleware.LoggerFromContext(r.Context()).With(
//		"component", "gift_certificate_handler",
//		"method", "CreateGiftCertificate",
//	)
//
//	log.Info("request_received")
//
//	req := &customerv1.CreateGiftCertificateRequest{}
//	if err := json.ParseJSON(r, req); err != nil {
//		log.Error("invalid_payload", "error", err)
//		json.WriteError(w, http.StatusBadRequest, err)
//		return
//	}
//
//	log.Debug("calling_customer_service", "request", req)
//	resp, err := h.client.Api.CreateGiftCertificate(r.Context(), req)
//	if err != nil {
//		log.Error("gRPC_call_failed", "error", err)
//		json.WriteError(w, http.StatusBadRequest, err)
//		return
//	}
//
//	log.Info("succeeded", "certificate_id", resp.Certificate.GetId())
//	json.WriteJSON(w, http.StatusCreated, resp)
//}
//
//func (h *GiftCertificateHandler) GetGiftCertificate(w http.ResponseWriter, r *http.Request) {
//	log := middleware.LoggerFromContext(r.Context()).With(
//		"component", "gift_certificate_handler",
//		"method", "GetGiftCertificate",
//	)
//
//	log.Info("request_received")
//
//	certificateIDStr := chi.URLParam(r, "certificate_id")
//	certificateID, err := strconv.ParseInt(certificateIDStr, 10, 64)
//	if err != nil {
//		log.Error("invalid_certificate_id", "certificate_id", certificateIDStr, "error", err)
//		json.WriteError(w, http.StatusBadRequest, errors.New("invalid certificate_id"))
//		return
//	}
//
//	req := &customerv1.GetGiftCertificateRequest{Id: certificateID}
//	log.Debug("calling_customer_service", "certificate_id", certificateID)
//
//	resp, err := h.client.Api.GetGiftCertificate(r.Context(), req)
//	if err != nil {
//		log.Error("gRPC_call_failed", "error", err)
//		json.WriteError(w, http.StatusNotFound, err)
//		return
//	}
//
//	log.Info("succeeded", "certificate_id", resp.Certificate.GetId())
//	json.WriteJSON(w, http.StatusOK, resp)
//}
//
//func (h *GiftCertificateHandler) ListGiftCertificates(w http.ResponseWriter, r *http.Request) {
//	log := middleware.LoggerFromContext(r.Context()).With(
//		"component", "gift_certificate_handler",
//		"method", "ListGiftCertificates",
//	)
//
//	log.Info("request_received")
//
//	query := r.URL.Query()
//
//	customerIDStr := query.Get("customer_id")
//	organizationIDStr := query.Get("organization_id")
//	limitStr := query.Get("limit")
//	offsetStr := query.Get("offset")
//
//	var (
//		customerID     *int64
//		organizationID *int64
//		limit          int32 = 20
//		offset         int32 = 0
//		err            error
//	)
//
//	if customerIDStr != "" {
//		id, err := strconv.ParseInt(customerIDStr, 10, 64)
//		if err == nil {
//			customerID = &id
//		}
//	}
//	if organizationIDStr != "" {
//		id, err := strconv.ParseInt(organizationIDStr, 10, 64)
//		if err == nil {
//			organizationID = &id
//		}
//	}
//	if limitStr != "" {
//		l, err := strconv.ParseInt(limitStr, 10, 32)
//		if err == nil {
//			limit = int32(l)
//		}
//	}
//	if offsetStr != "" {
//		o, err := strconv.ParseInt(offsetStr, 10, 32)
//		if err == nil {
//			offset = int32(o)
//		}
//	}
//
//	req := &customerv1.ListGiftCertificatesRequest{
//		CustomerId:     customerID,
//		OrganizationId: organizationID,
//		Limit:          limit,
//		Offset:         offset,
//	}
//
//	log.Debug("calling_customer_service", "request", req)
//
//	resp, err := h.client.Api.ListGiftCertificates(r.Context(), req)
//	if err != nil {
//		log.Error("gRPC_call_failed", "error", err)
//		json.WriteError(w, http.StatusInternalServerError, err)
//		return
//	}
//
//	log.Info("succeeded", "certificates_count", len(resp.Certificates))
//	json.WriteJSON(w, http.StatusOK, resp)
//}
