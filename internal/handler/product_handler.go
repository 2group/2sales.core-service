package handler

import (
	"log/slog"
	"net/http"

	"github.com/2group/2sales.core-service/internal/grpc"
	productv1 "github.com/2group/2sales.core-service/pkg/gen/go/product"
	"github.com/2group/2sales.core-service/pkg/json"
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
