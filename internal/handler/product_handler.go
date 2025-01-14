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
    log *slog.Logger
    product *grpc.ProductClient
}

func NewProductHandler(log *slog.Logger, product *grpc.ProductClient) *ProductHandler {
    return &ProductHandler{log: log, product: product}
}

func (h *ProductHandler) CreateCategory (w http.ResponseWriter, r *http.Request) {
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

    req := &productv1.ListProductRequest{
        PageSize: int32(limit),
        Page:     int32(offset),
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

    req := &productv1.CreateProductRequest{}
    json.ParseJSON(r, &req)
    req.CreatedBy = user_id

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
        Id:  int64(product_id),
    }

    response, err := h.product.Api.UpdateProduct(r.Context(), req)
    if err != nil {
        json.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    json.WriteJSON(w, http.StatusOK, response)
    return
}
