package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	crmv1 "github.com/2group/2sales.core-service/pkg/gen/go/crm"
	userv1 "github.com/2group/2sales.core-service/pkg/gen/go/user"
	"github.com/2group/2sales.core-service/pkg/json"
	middleware "github.com/2group/2sales.core-service/pkg/middeware"
	"github.com/go-chi/chi/v5"
)

type CrmHandler struct {
	log *slog.Logger
	crm *grpc.CrmClient
}

func NewCrmHandler(log *slog.Logger, crm *grpc.CrmClient) *CrmHandler {
	return &CrmHandler{log: log, crm: crm}
}

func (h *CrmHandler) CreateLead(w http.ResponseWriter, r *http.Request) {
	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	userID, ok := middleware.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	h.log.Info("create lead handler", "organization_id", organizationID, "user_id", userID)

	req := &crmv1.CreateLeadRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.Lead.CreatedByUser = &userv1.User{}

	req.Lead.CreatedByOrganizationId = &organizationID
	req.Lead.CreatedByUser.Id = &userID

	response, err := h.crm.Api.CreateLead(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// 5. Respond with the newly created lead
	json.WriteJSON(w, http.StatusCreated, response)
}

func (h *CrmHandler) GetLead(w http.ResponseWriter, r *http.Request) {
	leadIDStr := chi.URLParam(r, "lead_id")

	h.log.Info("get lead handler", "lead_id", leadIDStr)

	leadID, err := strconv.ParseInt(leadIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid lead_id format", "lead_id", leadIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Create the GetLeadRequest
	req := &crmv1.GetLeadRequest{LeadId: leadID}

	response, err := h.crm.Api.GetLead(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// 5. Respond with the newly created lead
	json.WriteJSON(w, http.StatusCreated, response)
}

func (h *CrmHandler) UpdateLead(w http.ResponseWriter, r *http.Request) {
	leadIDStr := chi.URLParam(r, "lead_id")

	h.log.Info("update lead handler", "lead_id", leadIDStr)

	leadID, err := strconv.ParseInt(leadIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid lead_id format", "lead_id", leadIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &crmv1.UpdateLeadRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.Lead.LeadId = &leadID

	response, err := h.crm.Api.UpdateLead(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
}

func (h *CrmHandler) PatchLead(w http.ResponseWriter, r *http.Request) {
	leadIDStr := chi.URLParam(r, "lead_id")

	h.log.Info("get lead handler", "lead_id", leadIDStr)

	leadID, err := strconv.ParseInt(leadIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid lead_id format", "lead_id", leadIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &crmv1.PatchLeadRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.Lead.LeadId = &leadID

	response, err := h.crm.Api.PatchLead(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
}

func (h *CrmHandler) DeleteLead(w http.ResponseWriter, r *http.Request) {
	leadIDStr := chi.URLParam(r, "lead_id")

	h.log.Info("delete lead handler", "lead_id", leadIDStr)

	leadID, err := strconv.ParseInt(leadIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid lead_id format", "lead_id", leadIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Create the GetLeadRequest
	req := &crmv1.DeleteLeadRequest{LeadId: leadID}

	response, err := h.crm.Api.DeleteLead(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// 5. Respond with the newly created lead
	json.WriteJSON(w, http.StatusCreated, response)
}

func (h *CrmHandler) ListLeads(w http.ResponseWriter, r *http.Request) {
	// 1. Get organization ID from your middleware
	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}
	h.log.Info("list leads handler", "organization_id", organizationID)

	// 2. Parse query parameters for limit and offset
	query := r.URL.Query()
	limitParam := query.Get("limit")
	offsetParam := query.Get("offset")

	var (
		limit  int32
		offset int32
	)

	// Default limit if none provided (e.g., 10)
	if limitParam == "" {
		limit = 10
	} else {
		parsedLimit, err := strconv.ParseInt(limitParam, 10, 32)
		if err != nil {
			json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid limit param"))
			return
		}
		limit = int32(parsedLimit)
	}

	if offsetParam != "" {
		parsedOffset, err := strconv.ParseInt(offsetParam, 10, 32)
		if err != nil {
			json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid offset param"))
			return
		}
		offset = int32(parsedOffset)
	}

	// 3. Build the gRPC request
	req := &crmv1.ListLeadsRequest{
		OrganizationId: organizationID,
		Limit:          limit,
		Offset:         offset,
	}

	// 4. Call the CRM service
	response, err := h.crm.Api.ListLeads(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// 5. Return the leads JSON response
	json.WriteJSON(w, http.StatusOK, response)
}
