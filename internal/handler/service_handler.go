package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	servicev1 "github.com/2group/2sales.core-service/pkg/gen/go/service"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/proto"
)

type ServiceHandler struct {
	log     *slog.Logger
	service *grpc.ServiceClient
}

func NewServiceHandler(log *slog.Logger, service *grpc.ServiceClient) *ServiceHandler {
	return &ServiceHandler{
		log:     log,
		service: service,
	}
}

func (h *ServiceHandler) CreateService(w http.ResponseWriter, r *http.Request) {
	log := h.log.With("method", "CreateService")
	log.Info("request_received")

	req := &servicev1.CreateServiceRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid payload"))
		return
	}

	log.Info("calling_service_microservice", "request", req)
	resp, err := h.service.Api.CreateService(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info("succeeded", "service_id", resp.GetService().GetId())
	json.WriteJSON(w, http.StatusCreated, resp)
}

func (h *ServiceHandler) GetService(w http.ResponseWriter, r *http.Request) {
	log := h.log.With("method", "GetService")
	log.Info("request_received")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error("invalid_id", "id", idStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("calling_service_microservice", "id", id)
	req := &servicev1.GetServiceRequest{Id: id}
	resp, err := h.service.Api.GetService(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusNotFound, err)
		return
	}

	log.Info("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *ServiceHandler) DeleteService(w http.ResponseWriter, r *http.Request) {
	log := h.log.With("method", "DeleteService")
	log.Info("request_received")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error("invalid_id", "id", idStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("calling_service_microservice", "id", id)
	req := &servicev1.DeleteServiceRequest{Id: id}
	resp, err := h.service.Api.DeleteService(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *ServiceHandler) PartialUpdateService(w http.ResponseWriter, r *http.Request) {
	log := h.log.With("method", "PartialUpdateService")
	log.Info("request_received")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error("invalid_id", "id", idStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &servicev1.PartialUpdateServiceRequest{
		Service: &servicev1.Service{Id: proto.Int64(id)},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("calling_service_microservice", "request", req)
	resp, err := h.service.Api.PartialUpdateService(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *ServiceHandler) UpdateService(w http.ResponseWriter, r *http.Request) {
	log := h.log.With("method", "UpdateService")
	log.Info("request_received")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error("invalid_id", "id", idStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &servicev1.UpdateServiceRequest{
		Service: &servicev1.Service{Id: proto.Int64(id)},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("calling_service_microservice", "request", req)
	resp, err := h.service.Api.UpdateService(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *ServiceHandler) GeneratePresignedURLs(w http.ResponseWriter, r *http.Request) {
	log := h.log.With("method", "GeneratePresignedURLs")
	log.Info("request_received")

	req := &servicev1.GeneratePresignedURLsRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("calling_service_microservice", "request", req)
	resp, err := h.service.Api.GeneratePresignedURLs(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *ServiceHandler) ListServices(w http.ResponseWriter, r *http.Request) {
	log := h.log.With("method", "ListServices")
	log.Info("request_received")

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	orgIDStr := r.URL.Query().Get("organization_id")
	branchIDStr := r.URL.Query().Get("branch_id")

	limit, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		log.Warn("invalid_limit", "limit", limitStr, "error", err)
		limit = 20
	}
	offset, err := strconv.ParseInt(offsetStr, 10, 32)
	if err != nil {
		log.Warn("invalid_offset", "offset", offsetStr, "error", err)
		offset = 0
	}

	var organizationID *int64
	if orgIDStr != "" {
		orgID, err := strconv.ParseInt(orgIDStr, 10, 64)
		if err == nil {
			organizationID = proto.Int64(orgID)
		}
	}

	var branchID *int64
	if branchIDStr != "" {
		brID, err := strconv.ParseInt(branchIDStr, 10, 64)
		if err == nil {
			branchID = proto.Int64(brID)
		}
	}

	req := &servicev1.ListServicesRequest{
		OrganizationId: organizationID,
		BranchId:       branchID,
		Limit:          int32(limit),
		Offset:         int32(offset),
	}

	log.Info("calling_service_microservice", "organization_id", organizationID, "branch_id", branchID, "limit", limit, "offset", offset)
	resp, err := h.service.Api.ListServices(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info("succeeded", "services_count", len(resp.Services))
	json.WriteJSON(w, http.StatusOK, resp)
}
