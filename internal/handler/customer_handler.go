package handler

import (
	"log/slog"
	"net/http"

	"github.com/2group/2sales.core-service/internal/grpc"
	customerv1 "github.com/2group/2sales.core-service/pkg/gen/go/customer"
	"github.com/2group/2sales.core-service/pkg/json"
)

type CustomerHandler struct {
	log      *slog.Logger
	customer *grpc.CustomerClient
}

func NewCustomerHandler(log *slog.Logger, customer *grpc.CustomerClient) *CustomerHandler {
	return &CustomerHandler{
		log:      log,
		customer: customer,
	}
}

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to create customer")

	req := &customerv1.CreateCustomerRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.customer.Api.CreateCustomer(r.Context(), req)
	if err != nil {
		h.log.Error("Error creating customer", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Customer created successfully", "response", response)

	json.WriteJSON(w, http.StatusCreated, response)
	h.log.Info("Response sent", "status", http.StatusCreated)
}

func (h *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to get customer")

	req := &customerv1.GetCustomerRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.customer.Api.GetCustomer(r.Context(), req)
	if err != nil {
		h.log.Error("Error getting customer", "error", err)
		json.WriteError(w, http.StatusNotFound, err)
		return
	}
	h.log.Info("Customer retrieved successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to delete customer")

	req := &customerv1.DeleteCustomerRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.customer.Api.DeleteCustomer(r.Context(), req)
	if err != nil {
		h.log.Error("Error deleting customer", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Customer deleted successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *CustomerHandler) PatchCustomer(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to patch customer")

	req := &customerv1.PartialUpdateCustomerRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.customer.Api.PartialUpdateCustomer(r.Context(), req)
	if err != nil {
		h.log.Error("Error patching customer", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Customer patched successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *CustomerHandler) PutCustomer(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to update customer")

	req := &customerv1.UpdateCustomerRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.customer.Api.UpdateCustomer(r.Context(), req)
	if err != nil {
		h.log.Error("Error updating customer", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Customer updated successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}
