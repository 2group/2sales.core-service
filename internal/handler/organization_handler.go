package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	organizationv1 "github.com/2group/2sales.core-service/pkg/gen/go/organization"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
)

type OrganizationHandler struct {
	log          *slog.Logger
	organization *grpc.OrganizationClient
}

func NewOrganizationHandler(log *slog.Logger, organization *grpc.OrganizationClient) *OrganizationHandler {
	return &OrganizationHandler{log: log, organization: organization}
}

func (h *OrganizationHandler) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	req := &organizationv1.CreateOrganizationRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.organization.Api.CreateOrganization(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) GetOrganization(w http.ResponseWriter, r *http.Request) {
	organizationIDStr := chi.URLParam(r, "organization_id")

	organizationID, err := strconv.ParseInt(organizationIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid organization_id format", "organization_id", organizationIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	req := &organizationv1.GetOrganizationRequest{
		Id: organizationID,
	}

	response, err := h.organization.Api.GetOrganization(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	json.WriteJSON(w, http.StatusFound, response)
	return
}
