package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	productv1 "github.com/2group/2sales.core-service/pkg/gen/go/product"
	"github.com/2group/2sales.core-service/pkg/json"
	middleware "github.com/2group/2sales.core-service/pkg/middeware"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	log     *slog.Logger
	product *grpc.ProductClient
}

func NewProductHandler(log *slog.Logger, product *grpc.ProductClient) *ProductHandler {
	return &ProductHandler{log: log, product: product}
}

func (h *ProductHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	req := &productv1.CreateProductCategoryRequest{}
	json.ParseJSON(r, &req)

	response, err := h.product.Api.CreateProductCategory(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *ProductHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	popularStr := queryParams.Get("popular")
	var level32 *int32

	if popularStr == "true" {
		level32 = nil // Set level32 to nil if conversion fails
	} else {
		level32Value := int32(1)
		level32 = &level32Value
	}

	req := &productv1.ListCategoriesRequest{
		Level: level32,
	}

	response, err := h.product.Api.ListCategories(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *ProductHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	category_id := chi.URLParam(r, "category_id")
	req := &productv1.GetProductCategoryWithChildrenRequest{
		Id: category_id,
	}

	response, err := h.product.Api.GetProductCategoryWithChildren(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	limit := 10
	offset := 0

	popular, err := strconv.ParseBool(query.Get("popular"))
	popularPtr := &popular

	if limitStr := query.Get("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if offsetStr := query.Get("offset"); offsetStr != "" {
		if parsedOffset, err := strconv.Atoi(offsetStr); err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	url := query.Get("url")
	category_id := query.Get("category_id")
	name := query.Get("name")

	price_from := 0
	if price_from_str := query.Get("price_form"); price_from_str != "" {
		parsed_price_from, err := strconv.Atoi(price_from_str)
		if err == nil {
			price_from = parsed_price_from
		}
	}

	price_to := 0
	if price_to_str := query.Get("price_to"); price_to_str != "" {
		parsed_price_to, err := strconv.Atoi(price_to_str)
		if err == nil {
			price_to = parsed_price_to
		}
	}

	organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	if organization_id_str := query.Get("organization_id"); organization_id_str != "" {
		parsed_ogranization_id, err := strconv.Atoi(organization_id_str)
		if err == nil {
			organization_id = int64(parsed_ogranization_id)
		}
	}

	organization_type, ok := middleware.GetOrganizationType(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	var exclude_product_ids []int64
	var exclude_product_ids_str []string

	filters := make(map[string]*productv1.Values)
	for key, value := range r.URL.Query() {
		switch key {
		case "name", "sort_by", "sort_order", "price_from", "price_to", "organization_id", "category_id", "url", "offset", "limit", "brand_id":
			continue
		case "exclude_product_id":
			exclude_product_ids_str = value
			continue
		}
		if len(value) > 0 {
			filters[key] = &productv1.Values{
				Values: value,
			}
		}
	}

	for _, value := range exclude_product_ids_str {
		if value != "" {
			product_id, _ := strconv.Atoi(value)
			exclude_product_ids = append(exclude_product_ids, int64(product_id))
		}
	}

	req := &productv1.ListProductsRequest{
		PageSize:          int32(limit),
		Page:              int32(offset) / int32(limit),
		PdfUrl:            url,
		PriceFrom:         float32(price_from),
		PriceTo:           float32(price_to),
		OrganizationId:    organization_id,
		CategoryId:        category_id,
		SearchQuery:       name,
		OrganizationType:  organization_type,
		Filter:            filters,
		ExcludeProductIds: exclude_product_ids,
		Popular:           popularPtr,
	}

	brand_id_str := query.Get("brand_id")
	if brand_id_str != "" {
		brand_id, err := strconv.Atoi(brand_id_str)
		if err == nil {
			req.BrandId = int64(brand_id)
			req.OrganizationId = 0
		}
	}

	response, err := h.product.Api.ListProducts(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productIDStr := chi.URLParam(r, "product_id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	queryParams := r.URL.Query()

	includeCharacteristicsStr := queryParams.Get("include_characteristics")
	includeCharacteristics, err := strconv.ParseBool(includeCharacteristicsStr)
	if err != nil {
		includeCharacteristics = false
	}

	includeImagesStr := queryParams.Get("include_images")
	includeImages, err := strconv.ParseBool(includeImagesStr)
	if err != nil {
		includeImages = false
	}

	includeProductGroupsStr := queryParams.Get("include_product_groups")
	includeProductGroups, err := strconv.ParseBool(includeProductGroupsStr)
	if err != nil {
		includeProductGroups = false
	}

	userID, ok := middleware.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	organizationType, ok := middleware.GetOrganizationType(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &productv1.GetProductRequest{
		Id:                     int64(productID),
		IncludeCharacteristics: includeCharacteristics,
		IncludeImages:          includeImages,
		IncludeProductGroups:   includeProductGroups,
		UserId:                 userID,
		OrganizationType:       organizationType,
	}

	response, err := h.product.Api.GetProduct(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	user_id, ok := middleware.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &productv1.CreateProductRequest{}
	json.ParseJSON(r, &req)
	req.Product.CreatedBy = &user_id
	req.Product.OrganizationId = &organization_id

	response, err := h.product.Api.CreateProduct(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *ProductHandler) PatchProduct(w http.ResponseWriter, r *http.Request) {
	productIDStr := chi.URLParam(r, "product_id")
	productID, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &productv1.PatchProductRequest{}
	json.ParseJSON(r, &req)

	req.Product.Id = &productID

	response, err := h.product.Api.PatchProduct(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *ProductHandler) CreateProductGroup(w http.ResponseWriter, r *http.Request) {
	req := &productv1.CreateProductGroupRequest{}
	json.ParseJSON(r, &req)

	response, err := h.product.Api.CreateProductGroup(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *ProductHandler) UpdateProductGroup(w http.ResponseWriter, r *http.Request) {
	req := &productv1.UpdateProductGroupRequest{}
	json.ParseJSON(r, &req)

	productGroupIDStr := chi.URLParam(r, "product_group_id")
	productGroupID, err := strconv.ParseInt(productGroupIDStr, 10, 64)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	req.ProductGroup.Id = productGroupID

	response, err := h.product.Api.UpdateProductGroup(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *ProductHandler) ListProductGroup(w http.ResponseWriter, r *http.Request) {
	req := &productv1.ListProductGroupsRequest{}

	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}
	req.OrganizationId = organizationID

	query := r.URL.Query()

	if name := query.Get("name"); name != "" {
		req.Name = name
	}

	if productID := query.Get("product_id"); productID != "" {
		id, err := strconv.ParseInt(productID, 10, 64)
		if err != nil {
			json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product_id: %w", err))
			return
		}
		req.ProductId = id
	}

	page := 1
	pageSize := 10

	if pageStr := query.Get("page"); pageStr != "" {
		p, err := strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid page number: %w", err))
			return
		}
		if p < 1 {
			json.WriteError(w, http.StatusBadRequest, fmt.Errorf("page number must be greater than 0"))
			return
		}
		page = int(p)
	}

	if pageSizeStr := query.Get("page_size"); pageSizeStr != "" {
		ps, err := strconv.ParseInt(pageSizeStr, 10, 64)
		if err != nil {
			json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid page size: %w", err))
			return
		}
		if ps < 1 {
			json.WriteError(w, http.StatusBadRequest, fmt.Errorf("page size must be greater than 0"))
			return
		}
		pageSize = int(ps)
	}

	req.Page = int64(page)
	req.PageSize = int64(pageSize)

	response, err := h.product.Api.ListProductGroups(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if response != nil && response.ProductGroups != nil {
		json.WriteJSON(w, http.StatusOK, response)
	} else {
		empty_response := map[string]interface{}{
			"product_groups": []interface{}{},
			"total":          0,
		}
		json.WriteJSON(w, http.StatusOK, empty_response)
	}
	return
}

func (h *ProductHandler) GetProductGroup(w http.ResponseWriter, r *http.Request) {
	req := &productv1.GetProductGroupRequest{}

	product_group_id_str := chi.URLParam(r, "product_group_id")
	product_group_id, err := strconv.Atoi(product_group_id_str)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	req.Id = int64(product_group_id)

	response, err := h.product.Api.GetProductGroup(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *ProductHandler) DeleteProductGroup(w http.ResponseWriter, r *http.Request) {
	req := &productv1.DeleteProductGroupRequest{}

	product_group_id_str := chi.URLParam(r, "product_group_id")
	product_group_id, err := strconv.Atoi(product_group_id_str)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	req.Id = int64(product_group_id)

	response, err := h.product.Api.DeleteProductGroup(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}
