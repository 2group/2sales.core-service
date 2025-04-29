package handler

import (
	"errors"
	"github.com/2group/2sales.core-service/pkg/middleware"
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
	b2c_service_order *grpc.B2CServiceOrderClient
}

func NewB2CServiceOrderHandler(b2c_service_order *grpc.B2CServiceOrderClient) *B2CServiceOrderHandler {
	return &B2CServiceOrderHandler{
		b2c_service_order: b2c_service_order,
	}
}

func (h *B2CServiceOrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "b2c_service_order_handler",
		"method", "CreateOrder",
	)
	log.Info("request_received")

	req := &orderv1.CreateOrderRequest{}
	if err := json.ParseProtoJSON(r.Body, req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if req.Order == nil {
		log.Error("missing_order_field")
		json.WriteError(w, http.StatusBadRequest, errors.New("order data is missing in request"))
		return
	}

	if req.Order.CustomerId == nil || *req.Order.CustomerId == 0 {
		log.Error("invalid_customer_id", "customer_id", req.Order.CustomerId)
		json.WriteError(w, http.StatusBadRequest, errors.New("customerId cannot be 0"))
		return
	}

	log.Info("calling_b2c_service_order_microservice", "customer_id", *req.Order.CustomerId)

	resp, err := h.b2c_service_order.Api.CreateOrder(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, errors.New("internal server error during order creation"))
		return
	}

	log.Info("succeeded", "order_id", resp.GetOrderDetail().GetId())
	json.WriteJSON(w, http.StatusCreated, resp)
}

func (h *B2CServiceOrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "b2c_service_order_handler",
		"method", "GetOrder",
	)
	log.Info("request_received")

	orderIDStr := chi.URLParam(r, "order_id")
	orderID, err := strconv.ParseInt(orderIDStr, 10, 64)
	if err != nil || orderID <= 0 {
		log.Error("invalid_order_id", "order_id", orderIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid order_id"))
		return
	}

	log.Info("calling_b2c_service_order_microservice", "order_id", orderID)

	req := &orderv1.GetOrderRequest{Id: orderID}
	resp, err := h.b2c_service_order.Api.GetOrder(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err, "order_id", orderID)

		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.NotFound:
				json.WriteError(w, http.StatusNotFound, errors.New("order not found"))
			case codes.InvalidArgument:
				json.WriteError(w, http.StatusBadRequest, errors.New(st.Message()))
			default:
				json.WriteError(w, http.StatusInternalServerError, errors.New("internal error"))
			}
		} else {
			json.WriteError(w, http.StatusInternalServerError, errors.New("internal error"))
		}
		return
	}

	log.Info("succeeded", "order_id", resp.GetOrderDetail().GetId())
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *B2CServiceOrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "b2c_service_order_handler",
		"method", "ListOrders",
	)
	log.Info("request_received")

	query := r.URL.Query()
	orgIDStr := query.Get("organization_id")
	branchIDStr := query.Get("branch_id")
	limitStr := query.Get("limit")
	offsetStr := query.Get("offset")

	var (
		orgID    int64 = 0
		branchID int64 = 0
		limit    int64 = 50
		offset   int64 = 0
		err      error
	)

	if orgIDStr != "" {
		orgID, err = strconv.ParseInt(orgIDStr, 10, 64)
		if err != nil || orgID < 0 {
			log.Error("invalid_organization_id", "organization_id", orgIDStr, "error", err)
			json.WriteError(w, http.StatusBadRequest, errors.New("invalid organization_id"))
			return
		}
	}

	if branchIDStr != "" {
		branchID, err = strconv.ParseInt(branchIDStr, 10, 64)
		if err != nil || branchID < 0 {
			log.Error("invalid_branch_id", "branch_id", branchIDStr, "error", err)
			json.WriteError(w, http.StatusBadRequest, errors.New("invalid branch_id"))
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.ParseInt(limitStr, 10, 64)
		if err != nil || limit < 0 {
			log.Error("invalid_limit", "limit", limitStr, "error", err)
			json.WriteError(w, http.StatusBadRequest, errors.New("invalid limit"))
			return
		}
	}

	if offsetStr != "" {
		offset, err = strconv.ParseInt(offsetStr, 10, 64)
		if err != nil || offset < 0 {
			log.Error("invalid_offset", "offset", offsetStr, "error", err)
			json.WriteError(w, http.StatusBadRequest, errors.New("invalid offset"))
			return
		}
	}

	log.Info("calling_b2c_service_order_microservice",
		"organization_id", orgID,
		"branch_id", branchID,
		"limit", limit,
		"offset", offset,
	)

	req := &orderv1.ListB2CServiceOrdersRequest{
		OrganizationId: orgID,
		BranchId:       branchID,
		Limit:          limit,
		Offset:         offset,
	}

	resp, err := h.b2c_service_order.Api.ListOrders(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, errors.New("failed to fetch orders"))
		return
	}

	log.Info("succeeded", "orders_count", len(resp.Orders))
	json.WriteJSON(w, http.StatusOK, resp)
}
