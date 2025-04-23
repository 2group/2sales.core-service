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
	middleware "github.com/2group/2sales.core-service/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

type CrmHandler struct {
	log *slog.Logger
	crm *grpc.CrmClient
}

func NewCrmHandler(log *slog.Logger, crm *grpc.CrmClient) *CrmHandler {
	return &CrmHandler{
		log: log,
		crm: crm,
	}
}

func (h *CrmHandler) CreateLead(w http.ResponseWriter, r *http.Request) {
	ci, _ := middleware.GetCorrelationID(r.Context())
	log := h.log.With("method", "CreateLead", "correlation_id", ci)

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

	req := &crmv1.CreateLeadRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error("Failed to parse json", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	req.Lead.CreatedByUser = &userv1.User{}

	req.Lead.CreatedByOrganizationId = &organizationID
	req.Lead.CreatedByUser.Id = &userID

	log.Info("Creating lead", "request", req)
	response, err := h.crm.Api.CreateLead(r.Context(), req)
	if err != nil {
		log.Error("Failed to create lead", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("Successfully created lead", "response", response)
	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *CrmHandler) GetLead(w http.ResponseWriter, r *http.Request) {
	ci, _ := middleware.GetCorrelationID(r.Context())
	log := h.log.With("method", "GetLead", "correlation_id", ci)
	leadIDStr := chi.URLParam(r, "lead_id")

	leadID, err := strconv.ParseInt(leadIDStr, 10, 64)
	if err != nil {
		log.Error("Invalid lead_id format", "lead_id", leadIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &crmv1.GetLeadRequest{LeadId: leadID}
	log.Info("Getting lead", "request", req)

	response, err := h.crm.Api.GetLead(r.Context(), req)
	if err != nil {
		log.Error("Failed to get lead", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("Successfuly created lead", "response", response)
	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *CrmHandler) UpdateLead(w http.ResponseWriter, r *http.Request) {
	ci, _ := middleware.GetCorrelationID(r.Context())
	log := h.log.With("method", "UpdateLead", "correlation_id", ci)
	leadIDStr := chi.URLParam(r, "lead_id")

	log.Info("Update lead handler", "lead_id", leadIDStr)

	leadID, err := strconv.ParseInt(leadIDStr, 10, 64)
	if err != nil {
		log.Error("Invalid lead_id format", "lead_id", leadIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &crmv1.UpdateLeadRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.Lead.LeadId = &leadID
	log.Info("Updating lead", "request", req)

	response, err := h.crm.Api.UpdateLead(r.Context(), req)
	if err != nil {
		log.Error("Failed to update lead", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("Successfully updated lead", "response", response)
	json.WriteJSON(w, http.StatusCreated, response)
}

func (h *CrmHandler) PatchLead(w http.ResponseWriter, r *http.Request) {
	ci, _ := middleware.GetCorrelationID(r.Context())
	log := h.log.With("method", "PatchLead", "correlation_id", ci)
	leadIDStr := chi.URLParam(r, "lead_id")

	log.Info("Get lead handler", "lead_id", leadIDStr)

	leadID, err := strconv.ParseInt(leadIDStr, 10, 64)
	if err != nil {
		log.Error("Invalid lead_id format", "lead_id", leadIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &crmv1.PatchLeadRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error("Failed to parse json", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.Lead.LeadId = &leadID
	log.Info("Patching lead", "request", req)

	response, err := h.crm.Api.PatchLead(r.Context(), req)
	if err != nil {
		log.Error("Failed to patch lead", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("Successfully patched lead", "response", response)
	json.WriteJSON(w, http.StatusCreated, response)
}

func (h *CrmHandler) DeleteLead(w http.ResponseWriter, r *http.Request) {
	ci, _ := middleware.GetCorrelationID(r.Context())
	log := h.log.With("method", "DeleteLead", "correlation_id", ci)
	leadIDStr := chi.URLParam(r, "lead_id")

	log.Info("delete lead handler", "lead_id", leadIDStr)

	leadID, err := strconv.ParseInt(leadIDStr, 10, 64)
	if err != nil {
		log.Error("invalid lead_id format", "lead_id", leadIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &crmv1.DeleteLeadRequest{LeadId: leadID}
	log.Info("Deleting lead", "request", req)

	response, err := h.crm.Api.DeleteLead(r.Context(), req)
	if err != nil {
		log.Error("Failed to delete the lead", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("Successfully deleted the lead", "response", response)
	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *CrmHandler) ListLeads(w http.ResponseWriter, r *http.Request) {
	ci, _ := middleware.GetCorrelationID(r.Context())
	log := h.log.With("method", "ListLeads", "correlation_id", ci)

	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}
	log.Info("list leads handler", "organization_id", organizationID)

	query := r.URL.Query()
	limitParam := query.Get("limit")
	offsetParam := query.Get("offset")

	var (
		limit  int32
		offset int32
	)

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

	req := &crmv1.ListLeadsRequest{
		OrganizationId: organizationID,
		Limit:          limit,
		Offset:         offset,
	}
	log.Info("Listing leads", "request", req)

	response, err := h.crm.Api.ListLeads(r.Context(), req)
	if err != nil {
		log.Error("Failed to list leads", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("Successfully listed leads", "response", response)
	json.WriteJSON(w, http.StatusOK, response)
}
