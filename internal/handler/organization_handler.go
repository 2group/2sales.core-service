package handler

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	organizationv1 "github.com/2group/2sales.core-service/pkg/gen/go/organization"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/2group/2sales.core-service/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

type OrganizationHandler struct {
	log          *slog.Logger
	organization *grpc.OrganizationClient
}

func NewOrganizationHandler(organization *grpc.OrganizationClient) *OrganizationHandler {
	return &OrganizationHandler{organization: organization}
}

func (h *OrganizationHandler) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "organization_handler",
		"method", "CreateOrganization",
	)

	log.Info("request_received")

	rc := middleware.NewRoleChecker(r)
	if !rc.HasSuperAdmin() {
		log.Error("forbidden")
		json.WriteError(w, http.StatusForbidden, errors.New("forbidden"))
		return
	}

	req := &organizationv1.CreateOrganizationRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid payload"))
		return
	}

	log.Info("calling_organization_microservice", "request", req)
	resp, err := h.organization.Api.CreateOrganization(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, errors.New("unable to create organization"))
		return
	}

	log.Info("succeeded", "organization_id", resp.GetOrganization().GetId())
	json.WriteJSON(w, http.StatusCreated, resp)
}

func (h *OrganizationHandler) GetOrganization(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "organization_handler",
		"method", "GetOrganization",
	)
	log.Info("request_recieved")
	orgIDStr := chi.URLParam(r, "organization_id")
	orgID, err := strconv.ParseInt(orgIDStr, 10, 64)
	if err != nil {
		log.Error("invalid_id", "organization_id", orgIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	rc := middleware.NewRoleChecker(r)
	if !(rc.HasSuperAdmin() || rc.HasOrgAdmin(orgID)) {
		log.Error("forbidden")
		json.WriteError(w, http.StatusForbidden, errors.New("No permission"))
		return
	}

	log.Info("calling_organization_microservice", "organization_id", orgID)

	req := &organizationv1.GetOrganizationRequest{Id: orgID}
	resp, err := h.organization.Api.GetOrganization(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	log.Info("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)

}

func (h *OrganizationHandler) DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "organization_handler",
		"method", "DeleteOrganization",
	)

	log.Info("request_received")

	rc := middleware.NewRoleChecker(r)
	if !rc.HasSuperAdmin() {
		log.Error("forbidden")
		json.WriteError(w, http.StatusForbidden, errors.New("forbidden"))
		return
	}

	orgIDStr := chi.URLParam(r, "organization_id")
	orgID, err := strconv.ParseInt(orgIDStr, 10, 64)
	if err != nil {
		log.Error("invalid_id", "organization_id", orgIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid organization_id"))
		return
	}

	log.Info("calling_organization_microservice", "organization_id", orgID)
	req := &organizationv1.DeleteOrganizationRequest{Id: orgID}
	resp, err := h.organization.Api.DeleteOrganization(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, errors.New("unable to delete organization"))
		return
	}

	log.Info("succeeded", "organization_id", orgID)
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *OrganizationHandler) PartialUpdateOrganization(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "organization_handler",
		"method", "PartialUpdateOrganization",
	)

	log.Info("request_received")

	orgIDStr := chi.URLParam(r, "organization_id")
	orgID, err := strconv.ParseInt(orgIDStr, 10, 64)
	if err != nil {
		log.Error("invalid_id", "organization_id", orgIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid organization_id"))
		return
	}

	rc := middleware.NewRoleChecker(r)
	if !rc.HasSuperAdmin() && !rc.HasOrgAdmin(orgID) {
		log.Error("forbidden")
		json.WriteError(w, http.StatusForbidden, errors.New("forbidden"))
		return
	}

	req := &organizationv1.PartialUpdateOrganizationRequest{
		Organization: &organizationv1.Organization{Id: &orgID},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid payload"))
		return
	}
	log.Info("calling_organization_microservice", "request", req)

	resp, err := h.organization.Api.PartialUpdateOrganization(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, errors.New("unable to update organization"))
		return
	}

	log.Info("succeeded", "organization_id", orgID)
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *OrganizationHandler) UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "organization_handler",
		"method", "UpdateOrganization",
	)
	log.Info("request_received")

	orgIDStr := chi.URLParam(r, "organization_id")
	orgID, err := strconv.ParseInt(orgIDStr, 10, 64)
	if err != nil {
		log.Error("invalid_id", "organization_id", orgIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid organization_id"))
		return
	}

	rc := middleware.NewRoleChecker(r)
	if !rc.HasSuperAdmin() && !rc.HasOrgAdmin(orgID) {
		log.Error("forbidden")
		json.WriteError(w, http.StatusForbidden, errors.New("forbidden"))
		return
	}

	req := &organizationv1.UpdateOrganizationRequest{
		Organization: &organizationv1.Organization{Id: &orgID},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid payload"))
		return
	}
	log.Info("calling_organization_microservice", "request", req)

	resp, err := h.organization.Api.UpdateOrganization(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, errors.New("unable to update organization"))
		return
	}

	log.Info("succeeded", "organization_id", orgID)
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *OrganizationHandler) ListOrganizations(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "organization_handler",
		"method", "ListOrganizations",
	)
	log.Info("request_received")

	rc := middleware.NewRoleChecker(r)
	if !rc.HasSuperAdmin() {
		log.Error("forbidden")
		json.WriteError(w, http.StatusForbidden, errors.New("forbidden"))
		return
	}

	limitStr := chi.URLParam(r, "limit")
	offsetStr := chi.URLParam(r, "offset")

	limit, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		log.Warn("invalid_limit", "limit", limitStr, "error", err)
		limit = 20
	}

	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		log.Warn("invalid_offset", "offset", offsetStr, "error", err)
		offset = 0
	}

	log.Info("calling_organization_microservice", "limit", limit, "offset", offset)
	req := &organizationv1.ListOrganizationsRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	resp, err := h.organization.Api.ListOrganizations(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, errors.New("unable to list organizations"))
		return
	}

	log.Info("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *OrganizationHandler) CreateBranch(w http.ResponseWriter, r *http.Request) {
	req := &organizationv1.CreateBranchRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.organization.Api.CreateBranch(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *OrganizationHandler) GetBranch(w http.ResponseWriter, r *http.Request) {
	branchIDStr := chi.URLParam(r, "branch_id")

	branchID, err := strconv.ParseInt(branchIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid branch_id format", "branch_id", branchIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	req := &organizationv1.GetBranchRequest{
		Id: branchID,
	}

	response, err := h.organization.Api.GetBranch(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	json.WriteJSON(w, http.StatusAccepted, response)
	return
}

func (h *OrganizationHandler) DeleteBranch(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to delete branch")

	branchIDStr := chi.URLParam(r, "branch_id")

	branchID, err := strconv.ParseInt(branchIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid branch_id format", "branch_id", branchIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.DeleteBranchRequest{
		Id: branchID,
	}

	response, err := h.organization.Api.DeleteBranch(r.Context(), req)
	if err != nil {
		h.log.Error("Error deleting branch", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Branch deleted successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *OrganizationHandler) PartialUpdateBranch(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to partial update branch")
	branchIDStr := chi.URLParam(r, "branch_id")

	branchID, err := strconv.ParseInt(branchIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid branch_id format", "branch_id", branchIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.PartialUpdateBranchRequest{
		Branch: &organizationv1.Branch{
			Id: &branchID,
		},
	}

	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.organization.Api.PartialUpdateBranch(r.Context(), req)
	if err != nil {
		h.log.Error("Error patching customer", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Branch patched successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *OrganizationHandler) UpdateBranch(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to update branch")

	branchIDStr := chi.URLParam(r, "branch_id")

	branchID, err := strconv.ParseInt(branchIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid branch_id format", "branch_id", branchIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.UpdateBranchRequest{
		Branch: &organizationv1.Branch{
			Id: &branchID,
		},
	}

	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.organization.Api.UpdateBranch(r.Context(), req)
	if err != nil {
		h.log.Error("Error updating customer", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Organization updated successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *OrganizationHandler) CreateAddress(w http.ResponseWriter, r *http.Request) {
	req := &organizationv1.CreateAddressRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.organization.Api.CreateAddress(r.Context(), req)
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
		h.log.Error("invalid adress_id format", "branch_id", addressIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	req := &organizationv1.GetAddressRequest{
		Id: addressID,
	}

	response, err := h.organization.Api.GetAddress(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	json.WriteJSON(w, http.StatusAccepted, response)
	return
}

func (h *OrganizationHandler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to delete branch")

	addressIDStr := chi.URLParam(r, "address_id")

	addressID, err := strconv.ParseInt(addressIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid address_id format", "address_id", addressIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.DeleteAddressRequest{
		Id: addressID,
	}

	response, err := h.organization.Api.DeleteAddress(r.Context(), req)
	if err != nil {
		h.log.Error("Error deleting address", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Address deleted successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *OrganizationHandler) PartialUpdateAddress(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to partial update address")
	addressIDStr := chi.URLParam(r, "address_id")

	addressID, err := strconv.ParseInt(addressIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid address_id format", "address_id", addressIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.PartialUpdateAddressRequest{
		Address: &organizationv1.Address{
			Id: &addressID,
		},
	}

	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.organization.Api.PartialUpdateAddress(r.Context(), req)
	if err != nil {
		h.log.Error("Error patching address", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Address patched successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *OrganizationHandler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to update address")

	addressIDStr := chi.URLParam(r, "address_id")

	addressID, err := strconv.ParseInt(addressIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid address_id format", "address_id", addressIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.UpdateAddressRequest{
		Address: &organizationv1.Address{
			Id: &addressID,
		},
	}

	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.organization.Api.UpdateAddress(r.Context(), req)
	if err != nil {
		h.log.Error("Error updating customer", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Address updated successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *OrganizationHandler) CreateLoyaltyLevel(w http.ResponseWriter, r *http.Request) {
	req := &organizationv1.CreateLoyaltyLevelRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.organization.Api.CreateLoyaltyLevel(r.Context(), req)
	if err != nil {
		h.log.Error("Failed to create loyalty level", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, resp)
}

func (h *OrganizationHandler) GetLoyaltyLevel(w http.ResponseWriter, r *http.Request) {
	loyaltyIDStr := chi.URLParam(r, "loyalty_level_id")
	loyaltyID, err := strconv.ParseInt(loyaltyIDStr, 10, 64)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid loyalty_level_id: %w", err))
		return
	}

	req := &organizationv1.GetLoyaltyLevelRequest{Id: loyaltyID}
	resp, err := h.organization.Api.GetLoyaltyLevel(r.Context(), req)
	if err != nil {
		h.log.Error("Failed to get loyalty level", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *OrganizationHandler) UpdateLoyaltyLevel(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to update loyalty level")

	// Получаем ID из URL
	loyaltyLevelIDStr := chi.URLParam(r, "loyalty_level_id")
	loyaltyLevelID, err := strconv.ParseInt(loyaltyLevelIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid loyalty_level_id format", "loyalty_level_id", loyaltyLevelIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Парсим JSON
	req := &organizationv1.UpdateLoyaltyLevelRequest{
		LoyaltyLevel: &organizationv1.LoyaltyLevel{
			Id: &loyaltyLevelID,
		},
	}
	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	h.log.Info("Parsed request JSON successfully", "request", req)

	// Вызов gRPC
	resp, err := h.organization.Api.UpdateLoyaltyLevel(r.Context(), req)
	if err != nil {
		h.log.Error("Failed to update loyalty level", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	h.log.Info("Loyalty level updated successfully", "response", resp)
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *OrganizationHandler) ListLoyaltyLevelsByOrganization(w http.ResponseWriter, r *http.Request) {
	orgIDStr := chi.URLParam(r, "organization_id")
	orgID, err := strconv.ParseInt(orgIDStr, 10, 64)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid organization_id: %w", err))
		return
	}

	req := &organizationv1.ListLoyaltyLevelsByOrganizationRequest{OrganizationId: orgID}
	resp, err := h.organization.Api.ListLoyaltyLevelsByOrganization(r.Context(), req)
	if err != nil {
		h.log.Error("Failed to list loyalty levels", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Make sure we are passing the correct response type
	if resp.LoyaltyLevels == nil {
		resp.LoyaltyLevels = []*organizationv1.LoyaltyLevel{}
	}

	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *OrganizationHandler) CreateStory(w http.ResponseWriter, r *http.Request) {
	req := &organizationv1.CreateStoryRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.organization.Api.CreateStory(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}
func (h *OrganizationHandler) UpdateStory(w http.ResponseWriter, r *http.Request) {
	req := &organizationv1.UpdateStoryRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	idStr := chi.URLParam(r, "story_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	req.Story.Id = &id
	response, err := h.organization.Api.UpdateStory(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
}

func (h *OrganizationHandler) ListStories(w http.ResponseWriter, r *http.Request) {
	req := &organizationv1.ListStoryRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.organization.Api.ListStory(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
}

func (h *OrganizationHandler) CreateBanner(w http.ResponseWriter, r *http.Request) {
	req := &organizationv1.CreateBannerRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.organization.Api.CreateBanner(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}
func (h *OrganizationHandler) UpdateBanner(w http.ResponseWriter, r *http.Request) {
	req := &organizationv1.UpdateBannerRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	idStr := chi.URLParam(r, "banner_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	req.Banner.Id = &id
	response, err := h.organization.Api.UpdateBanner(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
}

func (h *OrganizationHandler) ListBanners(w http.ResponseWriter, r *http.Request) {
	req := &organizationv1.ListBannerRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.organization.Api.ListBanner(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
}
