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

// func (h *WarehouseHandler) CreateAcceptance(w http.ResponseWriter, r *http.Request) {
// 	ci, _ := middleware.GetCorrelationID(r.Context())
// 	log := h.log.With("method", "CreateAcceptance", "correlation_id", ci)
// 	organization_id, ok := middleware.GetOrganizationID(r)
// 	if !ok {
// 		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
// 		return
// 	}
// 	log.Info("Getting Organization Id", "organization_id", organization_id)
// 	user_id, ok := middleware.GetUserID(r)
// 	if !ok {
// 		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
// 		return
// 	}

// 	log.Info("Getting User Id", "user_id", user_id)
// 	req := &warehousev1.CreateAcceptanceRequest{}
// 	json.ParseJSON(r, &req)

// 	req.OrganizationId = organization_id
// 	req.UserId = user_id

// 	log.Info("Creating Acceptance", "request", req)

// 	response, err := h.warehouse.Api.CreateAcceptance(r.Context(), req)
// 	if err != nil {
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	log.Info("Acceptance Created", "response", response)

// 	json.WriteJSON(w, http.StatusCreated, response)
// 	return
// }

// func (h *WarehouseHandler) GetAcceptance(w http.ResponseWriter, r *http.Request) {
// 	ci, _ := middleware.GetCorrelationID(r.Context())
// 	log := h.log.With("method", "UpdateAcceptance", "correlation_id", ci)
// 	acceptance_id_str := chi.URLParam(r, "acceptance_id")
// 	acceptance_id, err := strconv.Atoi(acceptance_id_str)
// 	if err != nil {
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	log.Info("Getting Acceptance Id", "acceptance_id", acceptance_id)

// 	req := &warehousev1.GetAcceptanceRequest{
// 		Id: int64(acceptance_id),
// 	}
// 	log.Info("Getting Acceptance", "request", req)

// 	response, err := h.warehouse.Api.GetAcceptance(r.Context(), req)
// 	if err != nil {
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	log.Info("Acceptance Found", "response", response)

// 	json.WriteJSON(w, http.StatusOK, response)
// 	return
// }

// func (h *WarehouseHandler) ListAcceptances(w http.ResponseWriter, r *http.Request) {
// 	ci, _ := middleware.GetCorrelationID(r.Context())
// 	log := h.log.With("method", "ListAcceptances", "correlation_id", ci)
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

// 	req := &warehousev1.ListAcceptancesRequest{
// 		OrganizationId: organization_id,
// 		Page:           int64(offset),
// 		PageSize:       int64(limit),
// 	}
// 	log.Info("Listing Acceptances", "request", req)

// 	response, err := h.warehouse.Api.ListAcceptances(r.Context(), req)
// 	if err != nil {
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	log.Info("Acceptances Listed", "response", response)

// 	json.WriteJSON(w, http.StatusOK, response)
// 	return
// }
