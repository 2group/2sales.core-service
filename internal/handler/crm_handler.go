package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/2group/2sales.core-service/internal/grpc"
	crmv1 "github.com/2group/2sales.core-service/pkg/gen/go/crm"
	"github.com/2group/2sales.core-service/pkg/json"
	middleware "github.com/2group/2sales.core-service/pkg/middeware"
)

type CrmHandler struct {
	log *slog.Logger
	crm *grpc.CrmClient
}

func NewCrmHandler(log *slog.Logger, crm *grpc.CrmClient) *CrmHandler {
	return &CrmHandler{log: log, crm: crm}
}

func (h *CrmHandler) CreateLead(w http.ResponseWriter, r *http.Request) {
	// 1. Retrieve organization ID from middleware
	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

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
