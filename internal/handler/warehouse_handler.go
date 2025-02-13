package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	warehousev1 "github.com/2group/2sales.core-service/pkg/gen/go/warehouse"
	"github.com/2group/2sales.core-service/pkg/json"
	middleware "github.com/2group/2sales.core-service/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

type WarehouseHandler struct {
	log       *slog.Logger
	warehouse *grpc.WarehouseClient
}

func NewWarehouseHandler(log *slog.Logger, warehouse *grpc.WarehouseClient) *WarehouseHandler {
	return &WarehouseHandler{
        log: log.With("component", "warehouse_handler"), warehouse: warehouse}
}

func (h *WarehouseHandler) UpdateWarehouse(w http.ResponseWriter, r *http.Request) {
    ci, _ := middleware.GetCorrelationID(r.Context())
    log := h.log.With("method", "UpdateWarehouse", "correlation_id", ci)


	warehouse_id_str := chi.URLParam(r, "warehouse_id")
	warehouse_id, err := strconv.Atoi(warehouse_id_str)
    log.Info("Getting Warehouse Id", "warehouse_id", warehouse_id)
	if err != nil {
        log.Error("Error getting warehouse id", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &warehousev1.UpdateWarehouseRequest{
		Id: int64(warehouse_id),
	}
	json.ParseJSON(r, &req)
    log.Info("Request", "request", req)

	response, err := h.warehouse.Api.UpdateWarehouse(r.Context(), req)
    log.Info("Response", "response", response)
	if err != nil {
        log.Error("Error updating warehouse", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

    log.Info("Warehouse updated successfully")
	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *WarehouseHandler) ListWarehouses(w http.ResponseWriter, r *http.Request) {
    ci, _ := middleware.GetCorrelationID(r.Context())
    log := h.log.With("method", "ListWarehouses", "correlation_id", ci)

	organization_id, ok := middleware.GetOrganizationID(r)
    log.Info("Getting Organization Id", "organization_id", organization_id)
	if !ok {
        log.Error("Unauthorized")
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

    log.Info("Getting query parameters")
	query := r.URL.Query()

	include_address := false
	include_address_str := query.Get("include_address")
	if include_address_str == "true" {
		include_address = true
	}

	req := &warehousev1.ListWarehousesRequest{
		OrganizationId: organization_id,
		IncludeAddress: include_address,
	}
    log.Info("Request", "request", req)

	response, err := h.warehouse.Api.ListWarehouses(r.Context(), req)
    log.Info("Response", "response", response)
	if err != nil {
        log.Error("Error listing warehouses", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

    log.Info("Warehouses listed successfully")
	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *WarehouseHandler) GetWarehouse(w http.ResponseWriter, r *http.Request) {
    ci, _ := middleware.GetCorrelationID(r.Context())
    log := h.log.With("method", "GetWarehouse", "correlation_id", ci)

	warehouse_id_str := chi.URLParam(r, "warehouse_id")
	is_all := false
	var warehouse_id int64
	if warehouse_id_str != "all" {
		new_warehouse_id, err := strconv.Atoi(warehouse_id_str)
		if err != nil {
			json.WriteError(w, http.StatusBadRequest, err)
			return
		}
		warehouse_id = int64(new_warehouse_id)
	} else {
		warehouse_id = 0
		is_all = true
	}
    log.Info("Getting Warehouse Id", "warehouse_id", warehouse_id)

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
    log.Info("Getting query parameters", "limit", limit, "offset", offset)

	organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
        log.Error("Unauthorized")
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}
    log.Info("Getting Organization Id", "organization_id", organization_id)

	req := &warehousev1.GetProductsInWarehouseRequest{
		WarehouseId:    int64(warehouse_id),
		Page:           int64(offset),
		PageSize:       int64(limit),
		IsAll:          is_all,
		OrganizationId: organization_id,
	}

    log.Info("Request", "request", req)
	response, err := h.warehouse.Api.GetProductsInWarehouse(r.Context(), req)
    log.Info("Response", "response", response)
	if err != nil {
        log.Error("Error getting warehouse", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

    log.Info("Warehouse retrieved successfully")
	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *WarehouseHandler) CreateWarehouse(w http.ResponseWriter, r *http.Request) {
    ci, _ := middleware.GetCorrelationID(r.Context())
    log := h.log.With("method", "CreateWarehouse", "correlation_id", ci)

	organization_id, ok := middleware.GetOrganizationID(r)
    log.Info("Getting Organization Id", "organization_id", organization_id)
	if !ok {
        log.Error("Unauthorized")
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &warehousev1.CreateWarehouseRequest{
		OrganizationId: organization_id,
	}

	json.ParseJSON(r, &req)
    log.Info("Request", "request", req)

	response, err := h.warehouse.Api.CreateWarehouse(r.Context(), req)
    log.Info("Response", "response", response)
	if err != nil {
        log.Error("Error creating warehouse", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

    log.Info("Warehouse created successfully")
	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *WarehouseHandler) GetCountProducts(w http.ResponseWriter, r *http.Request) {
	warehouse_id_str := chi.URLParam(r, "warehouse_id")
	warehouse_id, err := strconv.Atoi(warehouse_id_str)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	var product_ids []int64
	for key, values := range r.URL.Query() {
		if key == "product_id" {
			for _, value := range values {
				product_id, err := strconv.Atoi(value)
				if err != nil {
					json.WriteError(w, http.StatusBadRequest, err)
					return
				}
				product_ids = append(product_ids, int64(product_id))
			}
		}
	}

	req := &warehousev1.GetCountProductsRequest{
		WarehouseId: int64(warehouse_id),
		ProductIds:  product_ids,
	}

	response, err := h.warehouse.Api.GetCountProducts(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}
