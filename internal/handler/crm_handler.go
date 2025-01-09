package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	crmv1 "github.com/2group/2sales.core-service/pkg/gen/go/crm"
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

	h.log.Info("create lead handler", "organization_id", organizationID)

	req := &crmv1.CreateLeadRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.Lead.CreatedByOrganizationId = &organizationID

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
	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	h.log.Info("list leads handler", "organization_id", organizationID)

	req := &crmv1.ListLeadsRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.OrganizationId = organizationID

	response, err := h.crm.Api.ListLeads(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// 5. Respond with the newly created lead
	json.WriteJSON(w, http.StatusCreated, response)
}
