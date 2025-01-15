package handler

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	organizationv1 "github.com/2group/2sales.core-service/pkg/gen/go/organization"
	"github.com/2group/2sales.core-service/pkg/json"
	middleware "github.com/2group/2sales.core-service/pkg/middeware"
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
	user_id, ok := middleware.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized"))
		return
	}

	req := &organizationv1.CreateOrganizationRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Println(req)

	req.UserId = &user_id

	response, err := h.organization.Api.CreateOrganization(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) GetOrganization(w http.ResponseWriter, r *http.Request) {
	organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &organizationv1.GetOrganizationRequest{
		Id: int64(organization_id),
	}

	response, err := h.organization.Api.GetOrganization(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) ListOrganizations(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	name := ""
	orgType := ""
	name = queryParams.Get("name")
	orgType = queryParams.Get("type")

	page, err := strconv.Atoi(queryParams.Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(queryParams.Get("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	req := &organizationv1.ListOrganizationsRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Type:     orgType,
		Name:     name,
	}

	response, err := h.organization.Api.ListOrganizations(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *OrganizationHandler) PatchOrganization(w http.ResponseWriter, r *http.Request) {
	organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &organizationv1.PatchOrganizationRequest{}

	json.ParseJSON(r, &req)

	req.Organization.Id = &organization_id

	response, err := h.organization.Api.PatchOrganization(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &organizationv1.UpdateOrganizationRequest{}

	json.ParseJSON(r, &req)

	req.Organization.Id = &organization_id

	response, err := h.organization.Api.UpdateOrganization(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) CreateRelationshipType(w http.ResponseWriter, r *http.Request) {
	organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &organizationv1.CreateRelationshipTypeRequest{}

	json.ParseJSON(r, &req)

	req.RelationshipType.OwningOrganizationId = &organization_id

	response, err := h.organization.Api.CreateRelationshipType(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) UpdateRelationshipType(w http.ResponseWriter, r *http.Request) {
	relationshipTypeIDStr := chi.URLParam(r, "relationship_type_id")

	relationshipTypeID, err := strconv.ParseInt(relationshipTypeIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid lead_id format", "lead_id", relationshipTypeIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.UpdateRelationshipTypeRequest{}

	json.ParseJSON(r, &req)

	req.RelationshipType.Id = &relationshipTypeID

	response, err := h.organization.Api.UpdateRelationshipType(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) GetRelationshipType(w http.ResponseWriter, r *http.Request) {
	relationshipTypeIDStr := chi.URLParam(r, "relationship_type_id")

	relationshipTypeID, err := strconv.ParseInt(relationshipTypeIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid lead_id format", "lead_id", relationshipTypeIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.GetRelationshipTypeRequest{Id: relationshipTypeID}

	json.ParseJSON(r, &req)

	response, err := h.organization.Api.GetRelationshipType(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) ListRelationshipTypes(w http.ResponseWriter, r *http.Request) {
	organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &organizationv1.ListRelationshipTypesRequest{OrganizationId: organization_id}

	json.ParseJSON(r, &req)

	response, err := h.organization.Api.ListRelationshipTypes(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}
