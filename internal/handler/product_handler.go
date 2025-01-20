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

func (h *ProductHandler) GetFirstLevelCategories(w http.ResponseWriter, r *http.Request) {
	req := &productv1.GetZeroLevelCategoriesRequest{}

	response, err := h.product.Api.GetZeroLevelCategories(r.Context(), req)
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

	req := &productv1.ListProductRequest{
		PageSize:       int32(limit),
		Page:           int32(offset),
		PdfUrl:         url,
		PriceFrom:      float32(price_from),
		PriceTo:        float32(price_to),
		OrganizationId: organization_id,
		CategoryId:     category_id,
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
	product_id_str := chi.URLParam(r, "product_id")
	product_id, err := strconv.Atoi(product_id_str)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &productv1.GetProductRequest{
		Id: int64(product_id),
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
	req.CreatedBy = user_id
    req.OrganizationId = organization_id
    req.BrandId = organization_id

	response, err := h.product.Api.CreateProduct(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	product_id_str := chi.URLParam(r, "product_id")
	product_id, err := strconv.Atoi(product_id_str)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &productv1.UpdateProductRequest{
		Id: int64(product_id),
	}

	response, err := h.product.Api.UpdateProduct(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}
