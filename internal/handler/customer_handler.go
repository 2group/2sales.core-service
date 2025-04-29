package handler

import (
	"errors"
	"github.com/2group/2sales.core-service/pkg/middleware"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	customerv1 "github.com/2group/2sales.core-service/pkg/gen/go/customer"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
)

type CustomerHandler struct {
	customer *grpc.CustomerClient
}

func NewCustomerHandler(customer *grpc.CustomerClient) *CustomerHandler {
	return &CustomerHandler{customer: customer}
}

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "customer_handler",
		"method", "CreateCustomer",
	)
	log.Info("request_received")

	req := &customerv1.CreateCustomerRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Debug("calling_customer_service", "request", req)
	resp, err := h.customer.Api.CreateCustomer(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("succeeded", "customer_id", resp.Customer.GetId())
	json.WriteJSON(w, http.StatusCreated, resp)
}

func (h *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "customer_handler",
		"method", "GetCustomer",
	)
	log.Info("request_received")

	customerIDStr := chi.URLParam(r, "customer_id")
	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		log.Error("invalid_customer_id", "customer_id", customerIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid customer_id"))
		return
	}

	req := &customerv1.GetCustomerRequest{
		Lookup: &customerv1.GetCustomerRequest_Id{Id: customerID},
	}
	log.Debug("calling_customer_service", "customer_id", customerID)

	resp, err := h.customer.Api.GetCustomer(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusNotFound, err)
		return
	}

	log.Info("succeeded", "customer_id", resp.Customer.GetId())
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "customer_handler",
		"method", "DeleteCustomer",
	)
	log.Info("request_received")

	customerIDStr := chi.URLParam(r, "customer_id")
	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		log.Error("invalid_customer_id", "customer_id", customerIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid customer_id"))
		return
	}

	req := &customerv1.DeleteCustomerRequest{Id: customerID}
	log.Debug("calling_customer_service", "customer_id", customerID)

	resp, err := h.customer.Api.DeleteCustomer(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("succeeded", "customer_id", customerID)
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *CustomerHandler) PartialUpdateCustomer(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "customer_handler",
		"method", "PartialUpdateCustomer",
	)
	log.Info("request_received")

	customerIDStr := chi.URLParam(r, "customer_id")
	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		log.Error("invalid_customer_id", "customer_id", customerIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid customer_id"))
		return
	}

	req := &customerv1.PartialUpdateCustomerRequest{
		Customer: &customerv1.Customer{Id: &customerID},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Debug("calling_customer_service", "customer_id", customerID)
	resp, err := h.customer.Api.PartialUpdateCustomer(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("succeeded", "customer_id", resp.Customer.GetId())
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "customer_handler",
		"method", "UpdateCustomer",
	)
	log.Info("request_received")

	customerIDStr := chi.URLParam(r, "customer_id")
	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		log.Error("invalid_customer_id", "customer_id", customerIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid customer_id"))
		return
	}

	req := &customerv1.UpdateCustomerRequest{
		Customer: &customerv1.Customer{Id: &customerID},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Debug("calling_customer_service", "customer_id", customerID)
	resp, err := h.customer.Api.UpdateCustomer(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("succeeded", "customer_id", resp.Customer.GetId())
	json.WriteJSON(w, http.StatusOK, resp)
}
