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
	h.log.Info("", limit, offset)
	organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}
	
	req := &warehousev1.GetProductsInWarehouseRequest{
		WarehouseId: int64(warehouse_id),
		Page:        int64(offset),
		PageSize:    int64(limit),
		IsAll: is_all,
		OrganizationId: organization_id,
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
                ProductIds: product_ids,
        }

        response, err := h.warehouse.Api.GetCountProducts(r.Context(), req)
        if err != nil {
                json.WriteError(w, http.StatusInternalServerError, err)
                return
        }

        json.WriteJSON(w, http.StatusOK, response)
        return
}
