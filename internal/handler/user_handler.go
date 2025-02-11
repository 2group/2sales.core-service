package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	userv1 "github.com/2group/2sales.core-service/pkg/gen/go/user"
	"github.com/2group/2sales.core-service/pkg/json"
	middleware "github.com/2group/2sales.core-service/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	log  *slog.Logger
	user *grpc.UserClient
}

func NewUserHandler(user *grpc.UserClient) *UserHandler {
	return &UserHandler{user: user}
}

func (h *UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
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

func (h *UserHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	req := &userv1.RegisterRequest{}
	err := json.ParseJSON(r, &req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.user.Api.Register(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *UserHandler) HandleGetMyProfile(w http.ResponseWriter, r *http.Request) {
	user_id, ok := middleware.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}
	req := &userv1.GetUserRequest{
		UserId: user_id,
	}

	response, err := h.user.Api.GetUser(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *UserHandler) HandleUpdateMyProfile(w http.ResponseWriter, r *http.Request) {
	req := &userv1.UpdateUserRequest{}
	err := json.ParseJSON(r, req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user_id, ok := middleware.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized"))
		return
	}

	userIDPtr := int64(user_id)
	req.User.Id = &userIDPtr

	response, err := h.user.Api.UpdateUser(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *UserHandler) HandlePatchMyProfile(w http.ResponseWriter, r *http.Request) {
	user_id, ok := middleware.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &userv1.PatchUserRequest{}

	json.ParseJSON(r, req)

	req.User.Id = &user_id
	req.AssignedBy = &user_id

	response, err := h.user.Api.PatchUser(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *UserHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	userIDStr := chi.URLParam(r, "user_id")
	patchedUserID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &userv1.PatchUserRequest{}

	json.ParseJSON(r, req)

	req.User.Id = &patchedUserID
	req.AssignedBy = &userID
	req.User.OrganizationId = &organizationID

	response, err := h.user.Api.PatchUser(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	organizationID, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &userv1.CreateUserRequest{}

	json.ParseJSON(r, req)

	req.AssignedBy = &userID
	req.User.OrganizationId = &organizationID

	response, err := h.user.Api.CreateUser(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *UserHandler) ListMyOrganizationUsers(w http.ResponseWriter, r *http.Request) {
	organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &userv1.ListUsersRequest{
		OrganizationId: organization_id,
	}

	response, err := h.user.Api.ListUsers(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *UserHandler) ListMyOrganizationRoles(w http.ResponseWriter, r *http.Request) {
	organization_id, ok := middleware.GetOrganizationID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &userv1.ListRolesRequest{
		OrganizationId: organization_id,
	}

	response, err := h.user.Api.ListRoles(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}
