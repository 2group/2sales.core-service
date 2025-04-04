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
	req := &customerv1.CreateCustomerRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	response, err := h.customer.Api.CreateCustomer(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}
