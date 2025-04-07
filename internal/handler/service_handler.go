package handler

import (
	"log/slog"
	"net/http"

	"github.com/2group/2sales.core-service/internal/grpc"
	servicev1 "github.com/2group/2sales.core-service/pkg/gen/go/service"
	"github.com/2group/2sales.core-service/pkg/json"
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
