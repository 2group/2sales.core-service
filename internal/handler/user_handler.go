package handler

import (
	"log/slog"
	"net/http"

	"github.com/2group/2sales.core-service/internal/grpc"
	userv1 "github.com/2group/2sales.core-service/pkg/gen/go/user"
	"github.com/2group/2sales.core-service/pkg/json"
)

type UserHandler struct {
	log  *slog.Logger
	user *grpc.UserClient
}

func NewUserHandler(user *grpc.UserClient) *UserHandler {
	return &UserHandler{user: user}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
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

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
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

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	req := &userv1.UpdateUserRequest{}
	err := json.ParseJSON(r, &req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.user.Api.UpdateUser(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	req := &userv1.CreateUserRequest{}
	err := json.ParseJSON(r, &req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.user.Api.CreateUser(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

// func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
// 	req := &userv1.GetUserRequest{}
// 	err := json.ParseJSON(r, &req)
// 	if err != nil {
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	response, err := h.user.Api.GetUser(r.Context(), req)
// 	if err != nil {
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	json.WriteJSON(w, http.StatusOK, response)
// 	return
// }
