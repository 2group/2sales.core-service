package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	orderv1 "github.com/2group/2sales.core-service/pkg/gen/go/b2c_service_order"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	req := &orderv1.CreateOrderRequest{}

	if err := json.ParseProtoJSON(r.Body, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if req.Order == nil {
		h.log.Error("Parsed request has nil order field")
		json.WriteError(w, http.StatusBadRequest, errors.New("order data is missing in request"))
		return
	}

	if req.Order.CustomerId == nil || *req.Order.CustomerId == 0 {
		h.log.Error("Customer ID is nil or 0 after parsing")
		json.WriteError(w, http.StatusBadRequest, errors.New("customerId cannot be 0"))
		return
	}

	h.log.Info("Customer ID after protojson parsing in handler", "customer_id", *req.Order.CustomerId)

	response, err := h.b2c_service_order.Api.CreateOrder(r.Context(), req)
	if err != nil {
		h.log.Error("gRPC call to CreateOrder failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, errors.New("internal server error during order creation"))
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	h.log.Info("CreateOrder response sent", "status", http.StatusCreated)
}

func (h *B2CServiceOrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	orderIDStr := chi.URLParam(r, "order_id")
	if orderIDStr == "" {
		json.WriteError(w, http.StatusBadRequest, errors.New("missing order_id in URL path"))
		return
	}

	orderID, err := strconv.ParseInt(orderIDStr, 10, 64)
	if err != nil || orderID <= 0 {
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid order_id format in URL path"))
		return
	}

	req := &orderv1.GetOrderRequest{
		Id: orderID,
	}

	h.log.Info("Calling GetOrder gRPC method", "order_id", orderID)

	response, err := h.b2c_service_order.Api.GetOrder(r.Context(), req)
	if err != nil {
		h.log.Error("gRPC call to GetOrder failed", "error", err, "order_id", orderID)

		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.NotFound:
				json.WriteError(w, http.StatusNotFound, errors.New("order not found"))
			case codes.InvalidArgument:
				json.WriteError(w, http.StatusBadRequest, errors.New(st.Message()))
			default:
				json.WriteError(w, http.StatusInternalServerError, errors.New("internal server error retrieving order"))
			}
		} else {
			json.WriteError(w, http.StatusInternalServerError, errors.New("internal server error retrieving order"))
		}
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("GetOrder response sent", "status", http.StatusOK)
}
