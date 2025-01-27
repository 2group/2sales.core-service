package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	orderv1 "github.com/2group/2sales.core-service/pkg/gen/go/order"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
)

type OrderHandler struct {
	log   *slog.Logger
	order *grpc.OrderClient
}

func NewOrderHandler(log *slog.Logger, order *grpc.OrderClient) *OrderHandler {
	return &OrderHandler{log: log, order: order}
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	orderIDStr := chi.URLParam(r, "order_id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &orderv1.GetOrderRequest{
		Id: int64(orderID),
	}

	response, err := h.order.Api.GetOrder(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}
