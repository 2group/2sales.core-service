package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/2group/2sales.core-service/internal/grpc"
	organizationv1 "github.com/2group/2sales.core-service/pkg/gen/go/organization"
	"github.com/2group/2sales.core-service/pkg/json"
	middleware "github.com/2group/2sales.core-service/pkg/middeware"
)

type OrganizationHandler struct {
	log          *slog.Logger
	organization *grpc.OrganizationClient
}

func NewOrganizationHandler(organization *grpc.OrganizationClient) *OrganizationHandler {
	return &OrganizationHandler{organization: organization}
}

func (h *OrganizationHandler) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	user_id, ok := middleware.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &organizationv1.CreateOrganizationRequest{}
	json.ParseJSON(r, &req)

    req.UserId = int64(user_id)

	response, err := h.organization.Api.CreateOrganization(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, response)
	return
}
