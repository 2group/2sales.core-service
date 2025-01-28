package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/2group/2sales.core-service/internal/grpc"
	orderv1 "github.com/2group/2sales.core-service/pkg/gen/go/order"
	organizationv1 "github.com/2group/2sales.core-service/pkg/gen/go/organization"
	"github.com/2group/2sales.core-service/pkg/json"
	middleware "github.com/2group/2sales.core-service/pkg/middeware"
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

	req := &orderv1.CreateSubOrderRequest{}

	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.SubOrder.FromOrganization = &organizationv1.Organization{Id: &organizationID}

	response, err := h.order.Api.CreateSubOrder(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}
