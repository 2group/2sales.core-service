package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/2group/2sales.core-service/internal/grpc"
	userv1 "github.com/2group/2sales.core-service/pkg/gen/go/user"
	"github.com/2group/2sales.core-service/pkg/json"
	middleware "github.com/2group/2sales.core-service/pkg/middeware"
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

func (h *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
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

func (h *UserHandler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
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

func (h *UserHandler) HandlePatchUser(w http.ResponseWriter, r *http.Request) {
	req := &userv1.PatchUserRequest{}
	err := json.ParseJSON(r, req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user_id, ok := middleware.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	userIDPtr := int64(user_id)
	req.User.Id = &userIDPtr

	response, err := h.user.Api.PatchUser(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}
