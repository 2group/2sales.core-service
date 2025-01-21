package handler

import (
	"fmt"
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
	userID, ok := middleware.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized"))
		return
	}

	req := &organizationv1.CreateOrganizationRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.UserId = &userID

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
		h.log.Error("invalid organization_id format", "relationship_type_id", organizationIDStr, "error", err)
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

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) GetMyOrganization(w http.ResponseWriter, r *http.Request) {
	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
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

func (h *OrganizationHandler) PatchMyOrganization(w http.ResponseWriter, r *http.Request) {
	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &organizationv1.PatchOrganizationRequest{}

	json.ParseJSON(r, &req)

	req.Organization.Id = &organizationID

	response, err := h.organization.Api.PatchOrganization(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) UpdateMyOrganization(w http.ResponseWriter, r *http.Request) {
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

func (h *OrganizationHandler) CreateRelationship(w http.ResponseWriter, r *http.Request) {
	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &organizationv1.CreateRelationshipRequest{}

	json.ParseJSON(r, &req)

	req.Relationship.SourceOrganizationId = &organizationID

	response, err := h.organization.Api.CreateRelationship(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) ListRelationships(w http.ResponseWriter, r *http.Request) {
	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &organizationv1.ListRelationshipsRequest{}

	req.OrganizationId = organizationID

	response, err := h.organization.Api.ListRelationships(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) UpdateRelationship(w http.ResponseWriter, r *http.Request) {
	relationshipIDStr := chi.URLParam(r, "relationship_id")

	relationshipID, err := strconv.ParseInt(relationshipIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid relationship id format", "relationship_id", relationshipIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.UpdateRelationshipRequest{}

	json.ParseJSON(r, &req)

	req.Relationship.Id = &relationshipID

	response, err := h.organization.Api.UpdateRelationship(r.Context(), req)
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
		h.log.Error("invalid relationship_type_id format", "relationship_type_id", relationshipTypeIDStr, "error", err)
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
		h.log.Error("invalid relationship_type_id format", "relationship_type_id", relationshipTypeIDStr, "error", err)
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

func (h *OrganizationHandler) CreateAddress(w http.ResponseWriter, r *http.Request) {
	req := &organizationv1.CreateAddressRequest{}

	json.ParseJSON(r, &req)

	response, err := h.organization.Api.CreateAddress(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	addressIDStr := chi.URLParam(r, "address_id")

	addressID, err := strconv.ParseInt(addressIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid address_id format", "address_id", addressIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.UpdateAddressRequest{}

	json.ParseJSON(r, &req)

	req.Address.Id = &addressID

	response, err := h.organization.Api.UpdateAddress(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) PatchAddress(w http.ResponseWriter, r *http.Request) {
	addressIDStr := chi.URLParam(r, "address_id")

	addressID, err := strconv.ParseInt(addressIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid address_id format", "address_id", addressIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.PatchAddressRequest{}

	json.ParseJSON(r, &req)

	req.Address.Id = &addressID

	response, err := h.organization.Api.PatchAddress(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	addressIDStr := chi.URLParam(r, "address_id")

	addressID, err := strconv.ParseInt(addressIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid address_id format", "address_id", addressIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.DeleteAddressRequest{}

	req.Id = addressID

	response, err := h.organization.Api.DeleteAddress(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) GetAddress(w http.ResponseWriter, r *http.Request) {
	addressIDStr := chi.URLParam(r, "address_id")

	addressID, err := strconv.ParseInt(addressIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid address_id format", "address_id", addressIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.GetAddressRequest{}

	req.Id = addressID

	response, err := h.organization.Api.GetAddress(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	req := &organizationv1.CreateContactRequest{}

	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.organization.Api.CreateContact(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	contactIDStr := chi.URLParam(r, "contact_id")

	contactID, err := strconv.ParseInt(contactIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid contact_id format", "contact_id", contactIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.GetContactRequest{}

	req.Id = contactID

	response, err := h.organization.Api.GetContact(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	contactIDStr := chi.URLParam(r, "contact_id")

	contactID, err := strconv.ParseInt(contactIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid contact_id format", "contact_id", contactIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.UpdateContactRequest{}

	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.Contact.ContactId = &contactID

	response, err := h.organization.Api.UpdateContact(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	contactIDStr := chi.URLParam(r, "contact_id")

	contactID, err := strconv.ParseInt(contactIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid contact_id format", "contact_id", contactIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.DeleteContactRequest{}

	req.Id = contactID

	response, err := h.organization.Api.DeleteContact(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) GeneratePresignedURLs(w http.ResponseWriter, r *http.Request) {
	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &organizationv1.GeneratePresignedURLsRequest{}

	req.OrganizationId = organizationID

	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.organization.Api.GeneratePresignedURLs(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) CreateCounterparty(w http.ResponseWriter, r *http.Request) {
	req := &organizationv1.CreateCounterpartyRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req.SourceOrganizationId = organizationID

	response, err := h.organization.Api.CreateCounterparty(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) GetCounterparty(w http.ResponseWriter, r *http.Request) {
	counterpartyIDStr := chi.URLParam(r, "counterparty_id")

	counterpartyID, err := strconv.ParseInt(counterpartyIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid counterparty_id format", "counterparty_id", counterpartyIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.GetCounterpartyRequest{
		Id: counterpartyID,
	}

	response, err := h.organization.Api.GetCounterparty(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) ListCounterparties(w http.ResponseWriter, r *http.Request) {
	sourceOrganizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

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

	req := &organizationv1.ListCounterpartiesRequest{
		SourceOrganizationId: sourceOrganizationID,
		Limit:                limit,
		Offset:               offset,
	}

	response, err := h.organization.Api.ListCounterparties(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}
