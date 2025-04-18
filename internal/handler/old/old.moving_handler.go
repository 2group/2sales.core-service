package handler

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	warehousev1 "github.com/2group/2sales.core-service/pkg/gen/go/warehouse"
// 	"github.com/2group/2sales.core-service/pkg/json"
// 	middleware "github.com/2group/2sales.core-service/pkg/middleware"
// 	"github.com/go-chi/chi/v5"
// )

// func (h *WarehouseHandler) CreateMoving(w http.ResponseWriter, r *http.Request) {
// 	ci, _ := middleware.GetCorrelationID(r.Context())
// 	log := h.log.With("method", "CreateMoving", "correlation_id", ci)
// 	organization_id, ok := middleware.GetOrganizationID(r)
// 	if !ok {
// 		json.WriteJSON(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized"))
// 		return
// 	}

// 	req := &warehousev1.CreateMovingRequest{
// 		Moving: &warehousev1.MovingModel{},
// 	}

// 	json.ParseJSON(r, &req)
// 	req.Moving.OrganizationId = organization_id

// 	log.Info("Creating moving", "request", req)

// 	response, err := h.warehouse.Api.CreateMoving(r.Context(), req)
// 	if err != nil {
// 		log.Error("Failed to create moving", "error", err)
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	log.Info("Successfully created the moving", "response", response)

// 	json.WriteJSON(w, http.StatusCreated, response)
// 	return
// }

// func (h *WarehouseHandler) ListMoving(w http.ResponseWriter, r *http.Request) {
// 	ci, _ := middleware.GetCorrelationID(r.Context())
// 	log := h.log.With("method", "ListMoving", "correlation_id", ci)
// 	query := r.URL.Query()

// 	limit := 10
// 	offset := 0

// 	if limitStr := query.Get("limit"); limitStr != "" {
// 		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
// 			limit = parsedLimit
// 		}
// 	}

// 	if offsetStr := query.Get("offset"); offsetStr != "" {
// 		if parsedOffset, err := strconv.Atoi(offsetStr); err == nil && parsedOffset >= 0 {
// 			offset = parsedOffset
// 		}
// 	}

// 	organization_id, ok := middleware.GetOrganizationID(r)
// 	if !ok {
// 		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized"))
// 		return
// 	}
// 	req := &warehousev1.ListMovingRequest{
// 		OrganizationId: organization_id,
// 		Page:           int64(offset),
// 		PageSize:       int64(limit),
// 	}
// 	log.Info("Listing movings", "request", req)

// 	response, err := h.warehouse.Api.ListMoving(r.Context(), req)
// 	if err != nil {
// 		log.Error("Failed to list movings", "error", err)
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	log.Info("Successfully listed movings", "response", response)

// 	json.WriteJSON(w, http.StatusOK, response)
// 	return
// }

// func (h *WarehouseHandler) GetMoving(w http.ResponseWriter, r *http.Request) {
// 	ci, _ := middleware.GetCorrelationID(r.Context())
// 	log := h.log.With("method", "GetMoving", "correlation_id", ci)

// 	moving_id_str := chi.URLParam(r, "moving_id")
// 	moving_id, err := strconv.Atoi(moving_id_str)
// 	if err != nil {
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	req := &warehousev1.GetMovingRequest{
// 		Id: int64(moving_id),
// 	}
// 	log.Info("Getting moving", "request", req)

// 	response, err := h.warehouse.Api.GetMoving(r.Context(), req)
// 	if err != nil {
// 		log.Error("Failed to get moving", "error", err)
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	log.Info("Successfully get moving", "response", response)
// 	json.WriteJSON(w, http.StatusOK, response)
// 	return
// }
