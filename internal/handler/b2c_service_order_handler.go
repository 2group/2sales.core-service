package handler

import (
	"log/slog"
	"net/http"

	"github.com/2group/2sales.core-service/internal/grpc"
	b2c_service_orderv1 "github.com/2group/2sales.core-service/pkg/gen/go/b2c_service_order"
	"github.com/2group/2sales.core-service/pkg/json"
)

type B2CServiceOrderHandler struct {
	log               *slog.Logger
	b2c_service_order *grpc.B2CServiceOrderClient
}

func NewB2CServiceOrderHandler(log *slog.Logger, b2c_service_order *grpc.B2CServiceOrderClient) *B2CServiceOrderHandler {
	return &B2CServiceOrderHandler{
		log:               log,
		b2c_service_order: b2c_service_order,
	}
}

func (h *B2CServiceOrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	req := &b2c_service_orderv1.CreateOrderRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.b2c_service_order.Api.CreateOrder(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *B2CServiceOrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	req := &b2c_service_orderv1.UpdateOrderRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.b2c_service_order.Api.UpdateOrder(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *B2CServiceOrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	req := &b2c_service_orderv1.GetOrderRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.b2c_service_order.Api.GetOrder(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}
