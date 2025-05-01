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

type CustomerHandler struct {
	customer *grpc.CustomerClient
}

func NewCustomerHandler(customer *grpc.CustomerClient) *CustomerHandler {
	return &CustomerHandler{customer: customer}
}

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "customer_handler").
		Str("method", "CreateCustomer").
		Logger()

	log.Info().Msg("request_received")

	req := &customerv1.CreateCustomerRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Debug().Interface("request", req).Msg("calling_customer_service")

	resp, err := h.customer.Api.CreateCustomer(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Int64("customer_id", resp.Customer.GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusCreated, resp)
}

func (h *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "customer_handler").
		Str("method", "GetCustomer").
		Logger()

	log.Info().Msg("request_received")

	customerIDStr := chi.URLParam(r, "customer_id")
	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		log.Error().Str("customer_id", customerIDStr).Err(err).Msg("invalid_customer_id")
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid customer_id"))
		return
	}

	req := &customerv1.GetCustomerRequest{
		Lookup: &customerv1.GetCustomerRequest_Id{Id: customerID},
	}

	log.Debug().Int64("customer_id", customerID).Msg("calling_customer_service")

	resp, err := h.customer.Api.GetCustomer(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusNotFound, err)
		return
	}

	log.Info().Int64("customer_id", resp.Customer.GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "customer_handler").
		Str("method", "DeleteCustomer").
		Logger()

	log.Info().Msg("request_received")

	customerIDStr := chi.URLParam(r, "customer_id")
	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		log.Error().Str("customer_id", customerIDStr).Err(err).Msg("invalid_customer_id")
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid customer_id"))
		return
	}

	req := &customerv1.DeleteCustomerRequest{Id: customerID}
	log.Debug().Int64("customer_id", customerID).Msg("calling_customer_service")

	resp, err := h.customer.Api.DeleteCustomer(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Int64("customer_id", customerID).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *CustomerHandler) PartialUpdateCustomer(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "customer_handler").
		Str("method", "PartialUpdateCustomer").
		Logger()

	log.Info().Msg("request_received")

	customerIDStr := chi.URLParam(r, "customer_id")
	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		log.Error().Str("customer_id", customerIDStr).Err(err).Msg("invalid_customer_id")
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid customer_id"))
		return
	}

	req := &customerv1.PartialUpdateCustomerRequest{
		Customer: &customerv1.Customer{Id: &customerID},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Debug().Int64("customer_id", customerID).Msg("calling_customer_service")

	resp, err := h.customer.Api.PartialUpdateCustomer(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Int64("customer_id", resp.Customer.GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "customer_handler").
		Str("method", "UpdateCustomer").
		Logger()

	log.Info().Msg("request_received")

	customerIDStr := chi.URLParam(r, "customer_id")
	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		log.Error().Str("customer_id", customerIDStr).Err(err).Msg("invalid_customer_id")
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid customer_id"))
		return
	}

	req := &customerv1.UpdateCustomerRequest{
		Customer: &customerv1.Customer{Id: &customerID},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Debug().Int64("customer_id", customerID).Msg("calling_customer_service")

	resp, err := h.customer.Api.UpdateCustomer(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Int64("customer_id", resp.Customer.GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}
