package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	advertisementv1 "github.com/2group/2sales.core-service/pkg/gen/go/advertisement"
	"github.com/2group/2sales.core-service/pkg/json"
)

type AdvertisementHandler struct {
	log           *slog.Logger
	advertisement *grpc.AdvertisementClient
}

func NewAdvertisementHandler(log *slog.Logger, advertisement *grpc.AdvertisementClient) *AdvertisementHandler {
	return &AdvertisementHandler{log: log, advertisement: advertisement}
}

func (h *AdvertisementHandler) CreateBanner(w http.ResponseWriter, r *http.Request) {
	req := &advertisementv1.CreateBannerRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.advertisement.Api.CreateBanner(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
}

func (h *AdvertisementHandler) ListBanners(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	isActiveStr := query.Get("is_active")
	isActive, err := strconv.ParseBool(isActiveStr)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &advertisementv1.ListBannersRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.IsActive = isActive

	response, err := h.advertisement.Api.ListBanners(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
}
