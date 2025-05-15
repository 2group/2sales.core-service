package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/2group/2sales.core-service/internal/grpc"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	orderv1 "github.com/2group/2sales.core-service/pkg/gen/go/b2c_service_order"
)

type B2CServiceOrderHandler struct {
	b2c_service_order *grpc.B2CServiceOrderClient
}

func NewB2CServiceOrderHandler(b2c_service_order *grpc.B2CServiceOrderClient) *B2CServiceOrderHandler {
	return &B2CServiceOrderHandler{b2c_service_order: b2c_service_order}
}

func (h *B2CServiceOrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "b2c_service_order_handler").
		Str("method", "CreateOrder").
		Logger()

	log.Info().Msg("request_received")

	req := &orderv1.CreateOrderRequest{}
	if err := json.ParseProtoJSON(r.Body, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if req.Order == nil {
		log.Error().Msg("missing_order_field")
		json.WriteError(w, http.StatusBadRequest, errors.New("order data is missing in request"))
		return
	}

	if req.Order.CustomerId == nil || *req.Order.CustomerId == 0 {
		log.Error().Int64("customer_id", 0).Msg("invalid_customer_id")
		json.WriteError(w, http.StatusBadRequest, errors.New("customerId cannot be 0"))
		return
	}

	log.Info().Int64("customer_id", *req.Order.CustomerId).Msg("calling_microservice")

	resp, err := h.b2c_service_order.Api.CreateOrder(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("grpc_call_failed")

		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.InvalidArgument {
				json.WriteError(w, http.StatusBadRequest, errors.New(st.Message()))
			} else {
				json.WriteError(w, http.StatusInternalServerError, errors.New(st.Message()))
			}
		} else {
			json.WriteError(w, http.StatusInternalServerError, errors.New("internal server error during order creation"))
		}
		return
	}

	log.Info().Int64("order_id", resp.GetOrderDetail().GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusCreated, resp)
}

func (h *B2CServiceOrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "b2c_service_order_handler").
		Str("method", "GetOrder").
		Logger()

	log.Info().Msg("request_received")

	orderIDStr := chi.URLParam(r, "order_id")
	orderID, err := strconv.ParseInt(orderIDStr, 10, 64)
	if err != nil || orderID <= 0 {
		log.Error().Str("order_id", orderIDStr).Err(err).Msg("invalid_order_id")
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid order_id"))
		return
	}

	log.Info().Int64("order_id", orderID).Msg("calling_microservice")

	req := &orderv1.GetOrderRequest{Id: orderID}
	resp, err := h.b2c_service_order.Api.GetOrder(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Int64("order_id", orderID).Msg("grpc_call_failed")

		if st, ok := status.FromError(err); ok {
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

	log.Info().Int64("order_id", resp.GetOrderDetail().GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *B2CServiceOrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "b2c_service_order_handler").
		Str("method", "ListOrders").
		Logger()

	log.Info().Msg("request_received")

	query := r.URL.Query()
	orgIDStr := query.Get("organization_id")
	branchIDStr := query.Get("branch_id")
	customerIDStr := query.Get("customer_id")
	limitStr := query.Get("limit")
	offsetStr := query.Get("offset")
	searchTextStr := query.Get("search_text")
	createdAtFromStr := query.Get("created_at_from")
	createdAtToStr := query.Get("created_at_to")
	priceFromStr := query.Get("price_from")
	priceToStr := query.Get("price_to")

	var (
		orgID, branchID, customerID *int64
		priceFrom, priceTo          *float64
		createdAtFrom, createdAtTo  *string
	)

	// parse required limit and offset
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil || limit < 0 {
		log.Error().Str("limit", limitStr).Err(err).Msg("invalid_limit")
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid limit"))
		return
	}

	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil || offset < 0 {
		log.Error().Str("offset", offsetStr).Err(err).Msg("invalid_offset")
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid offset"))
		return
	}

	// optional parameters
	if orgIDStr != "" {
		if val, err := strconv.ParseInt(orgIDStr, 10, 64); err == nil {
			orgID = &val
		}
	}
	if branchIDStr != "" {
		if val, err := strconv.ParseInt(branchIDStr, 10, 64); err == nil {
			branchID = &val
		}
	}
	if customerIDStr != "" {
		if val, err := strconv.ParseInt(customerIDStr, 10, 64); err == nil {
			customerID = &val
		}
	}
	if priceFromStr != "" {
		if val, err := strconv.ParseFloat(priceFromStr, 64); err == nil {
			priceFrom = &val
		}
	}
	if priceToStr != "" {
		if val, err := strconv.ParseFloat(priceToStr, 64); err == nil {
			priceTo = &val
		}
	}

	// Handle both RFC3339 and YYYY-MM-DD for createdAtFrom
	if createdAtFromStr != "" {
		if _, err := time.Parse(time.RFC3339, createdAtFromStr); err == nil {
			createdAtFrom = &createdAtFromStr
		} else if t, err := time.Parse("2006-01-02", createdAtFromStr); err == nil {
			s := t.Format(time.RFC3339)
			createdAtFrom = &s
		} else {
			log.Warn().Str("created_at_from", createdAtFromStr).Msg("invalid_created_at_from_format")
		}
	}

	// Handle both RFC3339 and YYYY-MM-DD for createdAtTo
	if createdAtToStr != "" {
		if _, err := time.Parse(time.RFC3339, createdAtToStr); err == nil {
			createdAtTo = &createdAtToStr
		} else if t, err := time.Parse("2006-01-02", createdAtToStr); err == nil {
			t = t.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
			s := t.Format(time.RFC3339)
			createdAtTo = &s
		} else {
			log.Warn().Str("created_at_to", createdAtToStr).Msg("invalid_created_at_to_format")
		}
	}

	req := &orderv1.ListB2CServiceOrdersRequest{
		OrganizationId: orgID,
		BranchId:       branchID,
		CustomerId:     customerID,
		Limit:          limit,
		Offset:         offset,
		SearchText:     &searchTextStr,
		CreatedAtFrom:  createdAtFrom,
		CreatedAtTo:    createdAtTo,
		PriceFrom:      priceFrom,
		PriceTo:        priceTo,
	}

	log.Info().Interface("request", req).Msg("calling_microservice")

	resp, err := h.b2c_service_order.Api.ListOrders(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("grpc_call_failed")
		json.WriteError(w, http.StatusInternalServerError, errors.New("failed to fetch orders"))
		return
	}

	log.Info().Int("orders_count", len(resp.Orders)).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}
