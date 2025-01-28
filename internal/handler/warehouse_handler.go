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

func (h *WarehouseHandler) UpdateWarehouse(w http.ResponseWriter, r *http.Request) {
        warehouse_id_str := chi.URLParam(r, "warehouse_id")
	warehouse_id, err := strconv.Atoi(warehouse_id_str)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

        req := &warehousev1.UpdateWarehouseRequest{
                Id: int64(warehouse_id),
        }
        json.ParseJSON(r, &req)

        response, err := h.warehouse.Api.UpdateWarehouse(r.Context(), req)
        if err != nil {
                json.WriteError(w, http.StatusInternalServerError, err)
                return
        }

        json.WriteJSON(w, http.StatusOK, response)
        return
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

	req := &warehousev1.GetProductsInWarehouseRequest{
		WarehouseId: int64(warehouse_id),
		Page:        int64(offset),
		PageSize:    int64(limit),
	}

	response, err := h.warehouse.Api.GetProductsInWarehouse(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *WarehouseHandler) CreateWarehouse(w http.ResponseWriter, r *http.Request) {
	organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &warehousev1.CreateWarehouseRequest{
		OrganizationId: organization_id,
	}

	json.ParseJSON(r, &req)

	response, err := h.warehouse.Api.CreateWarehouse(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}
