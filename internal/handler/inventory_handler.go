package handler

import (
	"net/http"
	"fmt"
	"strconv"

	warehousev1 "github.com/2group/2sales.core-service/pkg/gen/go/warehouse"
	"github.com/2group/2sales.core-service/pkg/json"
	middleware "github.com/2group/2sales.core-service/pkg/middeware"
	"github.com/go-chi/chi/v5"
)

func(h *WarehouseHandler) CreateInventory(w http.ResponseWriter, r *http.Request) {
	organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized"))
		return
	}
	req := &warehousev1.CreateInventoryRequest{
		Inventory: &warehousev1.InventoryModel{},
	}
	json.ParseJSON(r, &req)
	req.Inventory.OrganizationId = organization_id

	response, err := h.warehouse.Api.CreateInventory(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func(h *WarehouseHandler) ListInventory(w http.ResponseWriter, r *http.Request) {
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

	organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &warehousev1.ListInventoryRequest{
		OrganizationId: organization_id,
		Page: int64(offset),
		PageSize: int64(limit),
	}

	response, err := h.warehouse.Api.ListInventory(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func(h *WarehouseHandler) GetInventory(w http.ResponseWriter, r *http.Request) {
	inventory_id_str := chi.URLParam(r, "inventory_id")
	inventory_id, err := strconv.Atoi(inventory_id_str)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &warehousev1.GetInventoryRequest{
		Id: int64(inventory_id),
	}

	response, err := h.warehouse.Api.GetInventory(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}
