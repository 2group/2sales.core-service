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

// func (h *WarehouseHandler) CreateWriteOff(w http.ResponseWriter, r *http.Request) {
// 	ci, _ := middleware.GetCorrelationID(r.Context())
// 	log := h.log.With("method", "CreateWriteOff", "correlation_id", ci)
// 	organization_id, ok := middleware.GetOrganizationID(r)
// 	if !ok {
// 		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized"))
// 		return
// 	}
// 	user_id, ok := middleware.GetUserID(r)
// 	if !ok {
// 		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized"))
// 		return
// 	}

// 	req := &warehousev1.CreateWriteOffRequest{
// 		WriteOff: &warehousev1.WriteOffModel{
// 			OrganizationId: organization_id,
// 			UserId:         user_id,
// 		},
// 	}
// 	log.Info("Creating Write Off", "request", req)

// 	json.ParseJSON(r, &req)

// 	response, err := h.warehouse.Api.CreateWriteOff(r.Context(), req)
// 	if err != nil {
// 		log.Error("Failed to create write off", "error", err)
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	log.Info("Write Off created succesfully", "response", response)

// 	json.WriteJSON(w, http.StatusCreated, response)
// 	return
// }

// func (h *WarehouseHandler) ListWriteOff(w http.ResponseWriter, r *http.Request) {
// 	ci, _ := middleware.GetCorrelationID(r.Context())
// 	log := h.log.With("method", "ListWriteOff", "correlation_id", ci)
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
// 		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
// 		return
// 	}

// 	req := &warehousev1.ListWriteOffRequest{
// 		OrganizationId: organization_id,
// 		Page:           int64(offset),
// 		PageSize:       int64(limit),
// 	}
// 	log.Info("Listing write offs", "request", req)

// 	response, err := h.warehouse.Api.ListWriteOff(r.Context(), req)
// 	if err != nil {
// 		log.Error("Failed to list write offs", "error", err)
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	log.Info("Succefully listed write offs", "response", response)

// 	json.WriteJSON(w, http.StatusOK, response)
// 	return
// }

// func (h *WarehouseHandler) GetWriteOff(w http.ResponseWriter, r *http.Request) {
// 	ci, _ := middleware.GetCorrelationID(r.Context())
// 	log := h.log.With("method", "GetWriteOff", "correlation_id", ci)

// 	acceptance_id_str := chi.URLParam(r, "write_off_id")
// 	acceptance_id, err := strconv.Atoi(acceptance_id_str)
// 	if err != nil {
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	req := &warehousev1.GetWriteOffRequest{
// 		WriteOffId: int64(acceptance_id),
// 	}
// 	log.Info("Getting Write Off", "request", req)

// 	response, err := h.warehouse.Api.GetWriteOff(r.Context(), req)
// 	if err != nil {
// 		log.Error("Error Getting write off", "error", err)
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	log.Info("Succesfully get the write off", "response", response)

// 	json.WriteJSON(w, http.StatusOK, response)
// 	return
// }
