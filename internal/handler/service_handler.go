package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	servicev1 "github.com/2group/2sales.core-service/pkg/gen/go/service"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
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

	serviceIDStr := chi.URLParam(r, "customer_id")

	serviceID, err := strconv.ParseInt(serviceIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid service_id format", "service_id", serviceIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &servicev1.GetServiceRequest{
		Id: serviceID,
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
	h.log.Info("Received request to delete customer")

	serviceIDStr := chi.URLParam(r, "service_id")

	serviceID, err := strconv.ParseInt(serviceIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid service_id format", "service_id", serviceIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &servicev1.DeleteServiceRequest{
		Id: serviceID,
	}

	response, err := h.service.Api.DeleteService(r.Context(), req)
	if err != nil {
		h.log.Error("Error deleting customer", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Service deleted successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *ServiceHandler) PartialUpdateService(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to partial update service")

	req := &servicev1.PartialUpdateServiceRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.service.Api.PartialUpdateService(r.Context(), req)
	if err != nil {
		h.log.Error("Error patching customer", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Customer patched successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *ServiceHandler) UpdateService(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to update customer")

	req := &servicev1.UpdateServiceRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.service.Api.UpdateService(r.Context(), req)
	if err != nil {
		h.log.Error("Error updating customer", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Customer updated successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}
