package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	warehousev1 "github.com/2group/2sales.core-service/pkg/gen/go/warehouse"
	"github.com/2group/2sales.core-service/pkg/json"
	middleware "github.com/2group/2sales.core-service/pkg/middeware"
	"github.com/go-chi/chi/v5"
)

type WarehouseHandler struct {
	log       *slog.Logger
	warehouse *grpc.WarehouseClient
}

func NewWarehouseHandler(log *slog.Logger, warehouse *grpc.WarehouseClient) *WarehouseHandler {
	return &WarehouseHandler{log: log, warehouse: warehouse}
}

func (h *WarehouseHandler) ListWarehouses(w http.ResponseWriter, r *http.Request) {
	organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &warehousev1.ListWarehousesRequest{
		OrganizationId: organization_id,
	}

	response, err := h.warehouse.Api.ListWarehouses(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *WarehouseHandler) GetWarehouse(w http.ResponseWriter, r *http.Request) {
	warehouse_id_str := chi.URLParam(r, "warehouse_id")
	warehouse_id, err := strconv.Atoi(warehouse_id_str)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &warehousev1.GetProductsInWarehouseRequest{
		WarehouseId: int64(warehouse_id),
	}

	response, err := h.warehouse.Api.GetProductsInWarehouse(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}
