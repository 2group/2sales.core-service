package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	organizationv1 "github.com/2group/2sales.core-service/pkg/gen/go/organization"
	userv1 "github.com/2group/2sales.core-service/pkg/gen/go/user"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
)

type AdminHandler struct {
	log          *slog.Logger
	user         *grpc.UserClient
	organization *grpc.OrganizationClient
}

func NewAdminHandler(user *grpc.UserClient, organization *grpc.OrganizationClient) *AdminHandler {
	return &AdminHandler{
		user:         user,
		organization: organization,
	}
}

func (h *AdminHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := &userv1.LoginRequest{}

	err := json.ParseJSON(r, &req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.user.Api.Login(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *AdminHandler) CreateBrand(w http.ResponseWriter, r *http.Request) {
	req := &organizationv1.CreateOrganizationRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	organizationType := "office"

	req.Organization.Address.Type = &organizationType

	response, err := h.organization.Api.CreateOrganization(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *AdminHandler) PatchBrand(w http.ResponseWriter, r *http.Request) {
	brandIDStr := chi.URLParam(r, "brand_id")

	brandID, err := strconv.ParseInt(brandIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid brand id format", "brand_id", brandIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &organizationv1.PatchOrganizationRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.Organization.Id = &brandID

	response, err := h.organization.Api.PatchOrganization(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *AdminHandler) ListBrands(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	name := ""

	name = queryParams.Get("name")

	limit, err := strconv.ParseInt(queryParams.Get("limit"), 10, 64)

	offset, err := strconv.ParseInt(queryParams.Get("offset"), 10, 64)

	req := &organizationv1.ListOrganizationsRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
		Type:   "manufacturer",
		Name:   name,
	}

	response, err := h.organization.Api.ListOrganizations(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *AdminHandler) GeneratePresignedURLs(w http.ResponseWriter, r *http.Request) {
	req := &organizationv1.GeneratePresignedURLsRequest{}

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
