package handler

// import (
// 	"fmt"
// 	"log/slog"
// 	"net/http"
// 	"strconv"

// 	"github.com/2group/2sales.core-service/internal/grpc"
// 	"github.com/2group/2sales.core-service/internal/lib"
// 	organizationv1 "github.com/2group/2sales.core-service/pkg/gen/go/organization"
// 	userv1 "github.com/2group/2sales.core-service/pkg/gen/go/user"
// 	"github.com/2group/2sales.core-service/pkg/json"
// 	"github.com/2group/2sales.core-service/templates/page/brand"
// 	"github.com/2group/2sales.core-service/templates/page/dashboard"
// 	"github.com/2group/2sales.core-service/templates/page/login"
// 	"github.com/go-chi/chi/v5"
// )

// type AdminHandler struct {
// 	log          *slog.Logger
// 	user         *grpc.UserClient
// 	organization *grpc.OrganizationClient
// }

// func NewAdminHandler(user *grpc.UserClient, organization *grpc.OrganizationClient) *AdminHandler {
// 	return &AdminHandler{
// 		user:         user,
// 		organization: organization,
// 	}
// }

// func (h *AdminHandler) DashboardPage(w http.ResponseWriter, r *http.Request) {
// 	lib.Render(w, r, dashboard.Index())
// }

// func (h *AdminHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
// 	lib.Render(w, r, login.Login(nil))
// }

// func (h *AdminHandler) Login(w http.ResponseWriter, r *http.Request) {
// 	req := &userv1.LoginRequest{}

// 	req.Login = r.FormValue("email")
// 	req.Password = r.FormValue("password")

// 	_, err := h.user.Api.Login(r.Context(), req)
// 	if err != nil {
// 		errString := err.Error()
// 		lib.Render(w, r, login.Login(&errString))
// 		return
// 	}

// 	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
// 	return
// }

// func (h *AdminHandler) BrandPage(w http.ResponseWriter, r *http.Request) {
// 	// Parse the brand ID from the URL
// 	brandIDStr := chi.URLParam(r, "brand_id")
// 	brandID, err := strconv.ParseInt(brandIDStr, 10, 64)
// 	if err != nil {
// 		h.log.Error("Invalid brand_id format", "brand_id", brandIDStr, "error", err)
// 		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid brand_id: %v", err))
// 		return
// 	}

// 	// Create the gRPC request
// 	req := &organizationv1.GetOrganizationRequest{
// 		Id: brandID,
// 	}

// 	// Fetch the organization details
// 	response, err := h.organization.Api.GetOrganization(r.Context(), req)
// 	if err != nil {
// 		h.log.Error("Failed to fetch organization", "brand_id", brandID, "error", err)
// 		http.Error(w, "Failed to fetch organization details", http.StatusInternalServerError)
// 		return
// 	}

// 	// Map Images
// 	var images []brand.ImageModel
// 	if response.Organization.Images != nil {
// 		for _, img := range response.Organization.Images {
// 			images = append(images, brand.ImageModel{
// 				ImageUrl:   img.ImageUrl,
// 				ImageIndex: *img.ImageIndex,
// 			})
// 		}
// 	}

// 	// Map Contacts
// 	var contacts []brand.Contact
// 	if response.Organization.Contacts != nil {
// 		for _, contact := range response.Organization.Contacts {
// 			contacts = append(contacts, brand.Contact{
// 				ContactID:     contact.ContactId,
// 				ContactType:   contact.ContactType,
// 				ContactPerson: contact.ContactPerson,
// 				PhoneNumber:   contact.PhoneNumber,
// 				Email:         contact.Email,
// 			})
// 		}
// 	}

// 	// Map response to the Organization struct
// 	organization := brand.Organization{
// 		ID:          *response.Organization.Id,
// 		Name:        *response.Organization.Name,
// 		LegalName:   response.Organization.LegalName,
// 		ImageUrl:    response.Organization.ImageUrl,
// 		AddressLine: response.Organization.Address.AddressLine,
// 		Images:      &images,
// 		Contacts:    &contacts,
// 	}

// 	// Render the Brand template with organization data
// 	lib.Render(w, r, brand.Brand(organization))
// }

// func (h *AdminHandler) BrandsPage(w http.ResponseWriter, r *http.Request) {
// 	query := r.URL.Query()

// 	limit := 10
// 	offset := 0
// 	currentPage := 1

// 	if query.Get("page") != "" {
// 		if page, err := strconv.Atoi(query.Get("page")); err == nil && page > 0 {
// 			currentPage = page
// 			offset = (currentPage - 1) * limit
// 		}
// 	}

// 	req := &organizationv1.ListOrganizationsRequest{
// 		Limit:  int32(limit),
// 		Offset: int32(offset),
// 		Type:   "manufacturer",
// 	}

// 	response, err := h.organization.Api.ListOrganizations(r.Context(), req)
// 	if err != nil {
// 		http.Error(w, "Failed to fetch organizations", http.StatusInternalServerError)
// 		return
// 	}

// 	// Map organizations from the response
// 	organizations := make([]brand.Organization, len(response.Organizations))
// 	for i, org := range response.Organizations {
// 		organizations[i] = brand.Organization{
// 			Name:      *org.Name,
// 			ID:        *org.Id,
// 			LegalName: org.LegalName,
// 			ImageUrl:  org.ImageUrl,
// 		}
// 	}

// 	// Calculate total pages
// 	totalCount := int(response.TotalCount)
// 	totalPages := (totalCount + limit - 1) / limit // Round up for pages

// 	// Render the template with organizations and pagination info
// 	lib.Render(w, r, brand.Brands(organizations, currentPage, totalPages))
// }

// func (h *AdminHandler) CreateBrandPage(w http.ResponseWriter, r *http.Request) {
// 	// Render the CreateBrand template
// 	lib.Render(w, r, brand.CreateBrand())
// }

// func (h *AdminHandler) PatchBrandImagesPage(w http.ResponseWriter, r *http.Request) {
// 	// Parse the brand ID from the URL
// 	brandIDStr := chi.URLParam(r, "brand_id")
// 	brandID, err := strconv.ParseInt(brandIDStr, 10, 64)
// 	if err != nil {
// 		h.log.Error("Invalid brand_id format", slog.String("brand_id", brandIDStr), slog.Any("error", err))
// 		http.Error(w, "Invalid brand_id format", http.StatusBadRequest)
// 		return
// 	}

// 	// Create gRPC request to fetch organization details
// 	req := &organizationv1.GetOrganizationRequest{Id: brandID}

// 	// Fetch the organization details
// 	response, err := h.organization.Api.GetOrganization(r.Context(), req)
// 	if err != nil {
// 		h.log.Error("Failed to fetch organization details", slog.Int64("brand_id", brandID), slog.Any("error", err))
// 		http.Error(w, "Failed to fetch organization details", http.StatusInternalServerError)
// 		return
// 	}

// 	// Map images into ImageModel
// 	var existingImages []brand.ImageModel
// 	if response.Organization.Images != nil {
// 		for _, img := range response.Organization.Images {
// 			existingImages = append(existingImages, brand.ImageModel{
// 				ImageUrl:   img.ImageUrl,
// 				ImageIndex: *img.ImageIndex,
// 			})
// 		}
// 	}

// 	// Render the PatchImagesBrand template
// 	lib.Render(w, r, brand.PatchBrandImages(brandID, existingImages))
// }

// func (h *AdminHandler) CreateBrand(w http.ResponseWriter, r *http.Request) {
// 	req := &organizationv1.CreateOrganizationRequest{}
// 	if err := json.ParseJSON(r, req); err != nil {
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	organizationType := "office"

// 	req.Organization.Address.Type = &organizationType

// 	response, err := h.organization.Api.CreateOrganization(r.Context(), req)
// 	if err != nil {
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	json.WriteJSON(w, http.StatusCreated, response)
// 	return
// }

// func (h *AdminHandler) PatchBrand(w http.ResponseWriter, r *http.Request) {
// 	brandIDStr := chi.URLParam(r, "brand_id")

// 	brandID, err := strconv.ParseInt(brandIDStr, 10, 64)
// 	if err != nil {
// 		h.log.Error("invalid brand id format", "brand_id", brandIDStr, "error", err)
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	req := &organizationv1.PatchOrganizationRequest{}
// 	if err := json.ParseJSON(r, req); err != nil {
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	req.Organization.Id = &brandID

// 	response, err := h.organization.Api.PatchOrganization(r.Context(), req)
// 	if err != nil {
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	json.WriteJSON(w, http.StatusCreated, response)
// 	return
// }

// func (h *AdminHandler) ListBrands(w http.ResponseWriter, r *http.Request) {
// 	queryParams := r.URL.Query()
// 	name := ""

// 	name = queryParams.Get("name")

// 	limit, err := strconv.ParseInt(queryParams.Get("limit"), 10, 64)

// 	offset, err := strconv.ParseInt(queryParams.Get("offset"), 10, 64)

// 	req := &organizationv1.ListOrganizationsRequest{
// 		Limit:  int32(limit),
// 		Offset: int32(offset),
// 		Type:   "manufacturer",
// 		Name:   name,
// 	}

// 	response, err := h.organization.Api.ListOrganizations(r.Context(), req)
// 	if err != nil {
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	json.WriteJSON(w, http.StatusOK, response)
// 	return
// }

// func (h *AdminHandler) GeneratePresignedURLs(w http.ResponseWriter, r *http.Request) {
// 	req := &organizationv1.GeneratePresignedURLsRequest{}

// 	if err := json.ParseJSON(r, req); err != nil {
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	response, err := h.organization.Api.GeneratePresignedURLs(r.Context(), req)
// 	if err != nil {
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	json.WriteJSON(w, http.StatusCreated, response)
// 	return
// }
