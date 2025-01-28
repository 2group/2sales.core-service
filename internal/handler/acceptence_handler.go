package handler

import (
	"fmt"
	"net/http"
	"strconv"

	warehousev1 "github.com/2group/2sales.core-service/pkg/gen/go/warehouse"
	"github.com/2group/2sales.core-service/pkg/json"
	middleware "github.com/2group/2sales.core-service/pkg/middeware"
	"github.com/go-chi/chi/v5"
)

func (h *WarehouseHandler) CreateAcceptance(w http.ResponseWriter, r *http.Request) {
    organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}
    user_id, ok := middleware.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}
    req := &warehousev1.CreateAcceptenceRequest{}
    json.ParseJSON(r, &req)
    req.OrganizationId = organization_id
    req.UserId = user_id

    response, err := h.warehouse.Api.CreateAcceptence(r.Context(), req)
    if err != nil {
        json.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    json.WriteJSON(w, http.StatusCreated, response)
    return
}

func (h *WarehouseHandler) GetAcceptance(w http.ResponseWriter, r *http.Request) {
    acceptance_id_str := chi.URLParam(r, "acceptance_id")
	acceptance_id, err := strconv.Atoi(acceptance_id_str)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

    req := &warehousev1.GetAcceptenceRequest{
        Id: int64(acceptance_id),
    }

    response, err := h.warehouse.Api.GetAcceptence(r.Context(), req)
    if err != nil {
        json.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    json.WriteJSON(w, http.StatusOK, response)
    return
}

func (h *WarehouseHandler) ListAcceptances(w http.ResponseWriter, r *http.Request) {
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

    req := &warehousev1.ListAcceptenceRequest{
        OrganizationId: organization_id,
        Page: int64(offset),
        PageSize: int64(limit),
    }

    response, err := h.warehouse.Api.ListAcceptence(r.Context(), req)
    if err != nil {
        json.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    json.WriteJSON(w, http.StatusOK, response)
    return
}
