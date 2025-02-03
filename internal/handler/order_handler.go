package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	orderv1 "github.com/2group/2sales.core-service/pkg/gen/go/order"
	organizationv1 "github.com/2group/2sales.core-service/pkg/gen/go/organization"
	"github.com/2group/2sales.core-service/pkg/json"
	middleware "github.com/2group/2sales.core-service/pkg/middeware"
	"github.com/go-chi/chi/v5"
)

type OrderHandler struct {
	log   *slog.Logger
	order *grpc.OrderClient
}

func NewOrderHandler(log *slog.Logger, order *grpc.OrderClient) *OrderHandler {
	return &OrderHandler{log: log, order: order}
}

func (h *OrderHandler) CreateSubOrder(w http.ResponseWriter, r *http.Request) {
	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	organizationType, ok := middleware.GetOrganizationType(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	req := &orderv1.CreateSubOrderRequest{
                SubOrder: &orderv1.SubOrder{
                        FromOrganization: &organizationv1.Organization{},
                        ToOrganization: &organizationv1.Organization{},
                },
        }

	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if organizationType == "retailer" {
		req.SubOrder.FromOrganization = &organizationv1.Organization{
                        Id: &organizationID,
                }
	} else {
                req.SubOrder.FromOrganization.Id = req.SubOrder.ToOrganization.Id
                req.SubOrder.ToOrganization = &organizationv1.Organization{
                        Id: &organizationID,
                }
        }

	response, err := h.order.Api.CreateSubOrder(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *OrderHandler) UpdateSubOrder(w http.ResponseWriter, r *http.Request) {
        suborder_id_str := chi.URLParam(r, "suborder_id")
	suborder_id, err := strconv.Atoi(suborder_id_str)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
        
        req := &orderv1.UpdateSubOrderRequest{
                Id: int64(suborder_id),
        }

        response, err := h.order.Api.UpdateSubOrder(r.Context(), req)
        if err != nil {
                json.WriteError(w, http.StatusInternalServerError, err)
                return
        }

        json.WriteJSON(w, http.StatusOK, response)
        return
}

func (h *OrderHandler) GetSubOrder(w http.ResponseWriter, r *http.Request) {
	suborder_id_str := chi.URLParam(r, "suborder_id")
	suborder_id, err := strconv.Atoi(suborder_id_str)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	req := &orderv1.GetSubOrderRequest{
		Id: int64(suborder_id),
	}

	response, err := h.order.Api.GetSubOrder(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *OrderHandler) ListSubOrder(w http.ResponseWriter, r *http.Request) {
	organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	query := r.URL.Query()

	limit := 10
	offset := 0

	if limitStr := query.Get("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if offsetStr := query.Get("offset"); offsetStr != "" {
		if parsedOffset, err := strconv.Atoi(offsetStr); err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	req := &orderv1.ListSubOrderRequest{
		OrganizationId: organization_id,
		Limit:          int64(limit),
		Offset:         int64(offset),
	}

	response, err := h.order.Api.ListSubOrder(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *OrderHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	req := &orderv1.ListCartRequest{
		OrganizationId: organizationID,
	}

	response, err := h.order.Api.ListCart(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *OrderHandler) AddProductToCart(w http.ResponseWriter, r *http.Request) {
	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	req := &orderv1.AddProductToCartRequest{
		Cart: &orderv1.Cart{
			Organization: &organizationv1.Organization{
				Id: &organizationID,
			},
		},
	}

	json.ParseJSON(r, &req)

	response, err := h.order.Api.AddProductToCart(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *OrderHandler) DeleteProductFromCart(w http.ResponseWriter, r *http.Request) {
	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	req := &orderv1.DeleteProductFromCartRequest{
		Cart: &orderv1.Cart{
			Organization: &organizationv1.Organization{
				Id: &organizationID,
			},
		},
	}

	json.ParseJSON(r, &req)

	response, err := h.order.Api.DeleteProductFromCart(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}
