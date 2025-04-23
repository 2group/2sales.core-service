package handler

import (
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
	h.log.Info("Received request to create service")

	req := &servicev1.CreateServiceRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.service.Api.CreateService(r.Context(), req)
	if err != nil {
		h.log.Error("Error creating service", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Service created successfully", "response", response)

	json.WriteJSON(w, http.StatusCreated, response)
	h.log.Info("Response sent", "status", http.StatusCreated)
}

func (h *ServiceHandler) GetService(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to get service")

	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.log.Error("invalid id format", "id", idStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &servicev1.GetServiceRequest{
		Id: id,
	}

	response, err := h.service.Api.GetService(r.Context(), req)
	if err != nil {
		h.log.Error("Error getting service", "error", err)
		json.WriteError(w, http.StatusNotFound, err)
		return
	}
	h.log.Info("Service retrieved successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *ServiceHandler) DeleteService(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to delete service")

	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.log.Error("invalid id format", "id", idStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &servicev1.DeleteServiceRequest{
		Id: id,
	}

	response, err := h.service.Api.DeleteService(r.Context(), req)
	if err != nil {
		h.log.Error("Error deleting service", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
	}
	h.log.Info("Service deleted successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *ServiceHandler) PartialUpdateService(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to partial update service")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.log.Error("invalid id format", "id", idStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &servicev1.PartialUpdateServiceRequest{
		Service: &servicev1.Service{
			Id: proto.Int64(id),
		},
	}

	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.service.Api.PartialUpdateService(r.Context(), req)
	if err != nil {
		h.log.Error("Error patching service", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Service patched successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *ServiceHandler) UpdateService(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to update service")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.log.Error("invalid id format", "id", idStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &servicev1.UpdateServiceRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if req.Service == nil {
		req.Service = &servicev1.Service{}
	}
	req.Service.Id = proto.Int64(id)

	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.service.Api.UpdateService(r.Context(), req)
	if err != nil {
		h.log.Error("Error updating service", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	h.log.Info("Service updated successfully", "response", response)
	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *ServiceHandler) GeneratePresignedURLs(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to generate presigned URLs")

	req := &servicev1.GeneratePresignedURLsRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.service.Api.GeneratePresignedURLs(r.Context(), req)
	if err != nil {
		h.log.Error("Error generating presigned URLs", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Presigned URLs generated successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *ServiceHandler) ListServices(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to list services")

	limitStr := chi.URLParam(r, "limit")
	offsetStr := chi.URLParam(r, "offset")

	limit, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		h.log.Warn("invalid limit format", "limit", limitStr, "error", err)
		limit = 20
	}

	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		h.log.Warn("invalid offset format", "offset", offsetStr, "error", err)
		offset = 0
	}

	req := &servicev1.ListServicesRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	response, err := h.serviceApi.ListServices(r.Context(), req)
	if err != nil {
		h.log.Error("Error listing services", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	h.log.Info("Services listed successfully", "response", response)
	json.WriteJSON(w, http.StatusOK, response)
}
