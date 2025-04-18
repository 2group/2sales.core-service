package handler

import (
	"fmt"
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
	fmt.Println(response)
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

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *OrganizationHandler) DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to delete organization")

	organizationIDStr := chi.URLParam(r, "organization_id")

	organizationID, err := strconv.ParseInt(organizationIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid organization_id format", "organization_id", organizationIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.DeleteOrganizationRequest{
		Id: organizationID,
	}

	response, err := h.organization.Api.DeleteOrganization(r.Context(), req)
	if err != nil {
		h.log.Error("Error deleting organization", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Organization deleted successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *OrganizationHandler) PartialUpdateOrganization(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to partial update organization")
	organizationIDStr := chi.URLParam(r, "organization_id")

	organizationID, err := strconv.ParseInt(organizationIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid organization_id format", "organization_id", organizationIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.PartialUpdateOrganizationRequest{
		Organization: &organizationv1.Organization{
			Id: &organizationID,
		},
	}

	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.organization.Api.PartialUpdateOrganization(r.Context(), req)
	if err != nil {
		h.log.Error("Error patching customer", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Customer patched successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *OrganizationHandler) UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to update organization")

	organizationIDStr := chi.URLParam(r, "organization_id")

	organizationID, err := strconv.ParseInt(organizationIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid organization_id format", "organization_id", organizationIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.UpdateOrganizationRequest{
		Organization: &organizationv1.Organization{
			Id: &organizationID,
		},
	}

	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.organization.Api.UpdateOrganization(r.Context(), req)
	if err != nil {
		h.log.Error("Error updating customer", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Organization updated successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
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

func (h *OrganizationHandler) CreateBonusLevel(w http.ResponseWriter, r *http.Request) {
	req := &organizationv1.CreateBonusLevelRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.organization.Api.CreateBonusLevel(r.Context(), req)
	if err != nil {
		h.log.Error("Failed to create bonus level", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, resp)
}

func (h *OrganizationHandler) GetBonusLevel(w http.ResponseWriter, r *http.Request) {
	bonusIDStr := chi.URLParam(r, "bonus_level_id")
	bonusID, err := strconv.ParseInt(bonusIDStr, 10, 64)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid bonus_level_id: %w", err))
		return
	}

	req := &organizationv1.GetBonusLevelRequest{Id: bonusID}
	resp, err := h.organization.Api.GetBonusLevel(r.Context(), req)
	if err != nil {
		h.log.Error("Failed to get bonus level", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *OrganizationHandler) UpdateBonusLevel(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to update bonus level")

	// Получаем ID из URL
	bonusLevelIDStr := chi.URLParam(r, "bonus_level_id")
	bonusLevelID, err := strconv.ParseInt(bonusLevelIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid bonus_level_id format", "bonus_level_id", bonusLevelIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Парсим JSON
	req := &organizationv1.UpdateBonusLevelRequest{
		BonusLevel: &organizationv1.BonusLevel{
			Id: &bonusLevelID,
		},
	}
	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	h.log.Info("Parsed request JSON successfully", "request", req)

	// Вызов gRPC
	resp, err := h.organization.Api.UpdateBonusLevel(r.Context(), req)
	if err != nil {
		h.log.Error("Failed to update bonus level", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	h.log.Info("Bonus level updated successfully", "response", resp)
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *OrganizationHandler) ListBonusLevelsByOrganization(w http.ResponseWriter, r *http.Request) {
	orgIDStr := chi.URLParam(r, "organization_id")
	orgID, err := strconv.ParseInt(orgIDStr, 10, 64)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid organization_id: %w", err))
		return
	}

	req := &organizationv1.ListBonusLevelsByOrganizationRequest{OrganizationId: orgID}
	resp, err := h.organization.Api.ListBonusLevelsByOrganization(r.Context(), req)
	if err != nil {
		h.log.Error("Failed to list bonus levels", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Make sure we are passing the correct response type
	if resp.BonusLevels == nil {
		resp.BonusLevels = []*organizationv1.BonusLevel{}
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

	response, err := h.organization.Api.UpdateStory(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

// func (h *OrganizationHandler) ListStories(w http.ResponseWriter, r *http.Request) {
// 	req := &organizationv1.ListStoriesRequest{}
// 	if err := json.ParseJSON(r, req); err != nil {
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	response, err := h.organization.Api.ListStories(r.Context(), req)
// 	if err != nil {
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	json.WriteJSON(w, http.StatusCreated, response)
// 	return
// }

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

	response, err := h.organization.Api.UpdateBanner(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

// func (h *OrganizationHandler) ListBanners(w http.ResponseWriter, r *http.Request) {
// 	req := &organizationv1.ListBannersRequest{}
// 	if err := json.ParseJSON(r, req); err != nil {
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	response, err := h.organization.Api.ListBanners(r.Context(), req)
// 	if err != nil {
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	json.WriteJSON(w, http.StatusCreated, response)
// 	return
// }
