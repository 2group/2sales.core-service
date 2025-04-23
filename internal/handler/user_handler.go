package handler

import (
	"github.com/2group/2sales.core-service/internal/grpc"
	userv1 "github.com/2group/2sales.core-service/pkg/gen/go/user"
	"github.com/2group/2sales.core-service/pkg/json"
	"log/slog"
	"net/http"
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

// func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
// 	h.log.Info("Received request to list users")

// 	limitStr := chi.URLParam(r, "limit")
// 	offsetStr := chi.URLParam(r, "offset")

// 	limit, err := strconv.ParseInt(limitStr, 10, 32)
// 	if err != nil {
// 		h.log.Warn("invalid limit format", "limit", limitStr, "error", err)
// 		limit = 20
// 	}
// 	offset, err := strconv.ParseInt(offsetStr, 10, 32)
// 	if err != nil {
// 		h.log.Warn("invalid offset format", "offset", offsetStr, "error", err)
// 		offset = 0
// 	}

// 	req := &userv1.ListUsersRequest{
// 		Limit:  int32(limit),
// 		Offset: int32(offset),
// 	}

// 	resp, err := h.user.ListUsers(r.Context(), req)
// 	if err != nil {
// 		h.log.Error("Error listing users", "error", err)
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	h.log.Info("Users listed successfully", "response", resp)
// 	json.WriteJSON(w, http.StatusOK, resp)
// 	h.log.Info("Response sent", "status", http.StatusOK)
// }
