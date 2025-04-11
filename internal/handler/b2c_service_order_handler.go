package handler

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	orderv1 "github.com/2group/2sales.core-service/pkg/gen/go/b2c_service_order"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
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

func parseProtoJSON(r *http.Request, m proto.Message) error {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.New("cannot read request body")
	}
	defer r.Body.Close()

	unmarshalOpts := protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}

	if err := unmarshalOpts.Unmarshal(bodyBytes, m); err != nil {
		return errors.New("invalid JSON format or data")
	}
	return nil
}

func (h *B2CServiceOrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	req := &orderv1.CreateOrderRequest{}

	if err := parseProtoJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if req.GetOrder() == nil {
		h.log.Error("Parsed request has nil order field")
		json.WriteError(w, http.StatusBadRequest, errors.New("order data is missing in request"))
		return
	}

	receivedCustomerID := req.GetOrder().GetCustomerId()
	h.log.Info("Customer ID after protojson parsing in handler", "customer_id", receivedCustomerID)

	if receivedCustomerID == 0 {
		h.log.Error("Customer ID is 0 after parsing")
		json.WriteError(w, http.StatusBadRequest, errors.New("customerId cannot be 0"))
		return
	}

	response, err := h.b2c_service_order.Api.CreateOrder(r.Context(), req)
	if err != nil {
		h.log.Error("gRPC call to CreateOrder failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, errors.New("internal server error during order creation"))
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
}

func (h *B2CServiceOrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	req := &orderv1.UpdateOrderRequest{}
	if err := parseProtoJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.b2c_service_order.Api.UpdateOrder(r.Context(), req)
	if err != nil {
		h.log.Error("gRPC call to UpdateOrder failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, errors.New("internal server error during order update"))
		return
	}
	json.WriteJSON(w, http.StatusOK, response)
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
}
