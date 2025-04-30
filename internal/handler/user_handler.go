package handler

//
//import (
//	"github.com/2group/2sales.core-service/pkg/middleware"
//	"log/slog"
//	"net/http"
//	"strconv"
//
//	"github.com/2group/2sales.core-service/internal/grpc"
//	userv1 "github.com/2group/2sales.core-service/pkg/gen/go/user"
//	"github.com/2group/2sales.core-service/pkg/json"
//	"github.com/go-chi/chi/v5"
//)
//
//type UserHandler struct {
//	log  *slog.Logger
//	user *grpc.UserClient
//}
//
//func NewUserHandler(user *grpc.UserClient) *UserHandler {
//	return &UserHandler{
//		user: user,
//	}
//}
//
//func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
//	log := middleware.LoggerFromContext(r.Context()).With(
//		"component", "user_handler",
//		"method", "Login",
//	)
//
//	log.Info("request_received")
//
//	req := &userv1.LoginRequest{}
//	if err := json.ParseJSON(r, req); err != nil {
//		log.Error("invalid_payload", "error", err)
//		json.WriteError(w, http.StatusBadRequest, err)
//		return
//	}
//
//	log.Debug("calling_user_service", "request", req)
//	resp, err := h.user.Api.Login(r.Context(), req)
//	if err != nil {
//		log.Error("gRPC_call_failed", "error", err)
//		json.WriteError(w, http.StatusBadRequest, err)
//		return
//	}
//
//	log.Info("succeeded")
//	json.WriteJSON(w, http.StatusOK, resp)
//}
//
//func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
//	log := middleware.LoggerFromContext(r.Context()).With(
//		"component", "user_handler",
//		"method", "UpdateUser",
//	)
//
//	log.Info("request_received")
//
//	req := &userv1.UpdateUserRequest{}
//	if err := json.ParseJSON(r, req); err != nil {
//		log.Error("invalid_payload", "error", err)
//		json.WriteError(w, http.StatusBadRequest, err)
//		return
//	}
//
//	log.Debug("calling_user_service", "request", req)
//	resp, err := h.user.Api.UpdateUser(r.Context(), req)
//	if err != nil {
//		log.Error("gRPC_call_failed", "error", err)
//		json.WriteError(w, http.StatusInternalServerError, err)
//		return
//	}
//
//	log.Info("succeeded", "user_id", resp.GetUser().GetId())
//	json.WriteJSON(w, http.StatusOK, resp)
//}
//
//func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
//	log := middleware.LoggerFromContext(r.Context()).With(
//		"component", "user_handler",
//		"method", "CreateUser",
//	)
//
//	log.Info("request_received")
//
//	req := &userv1.CreateUserRequest{}
//	if err := json.ParseJSON(r, req); err != nil {
//		log.Error("invalid_payload", "error", err)
//		json.WriteError(w, http.StatusBadRequest, err)
//		return
//	}
//
//	log.Debug("calling_user_service", "request", req)
//	resp, err := h.user.Api.CreateUser(r.Context(), req)
//	if err != nil {
//		log.Error("gRPC_call_failed", "error", err)
//		json.WriteError(w, http.StatusInternalServerError, err)
//		return
//	}
//
//	log.Info("succeeded", "user_id", resp.GetUser().GetId())
//	json.WriteJSON(w, http.StatusOK, resp)
//}
//
//// func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
//// 	req := &userv1.GetUserRequest{}
//// 	err := json.ParseJSON(r, &req)
//// 	if err != nil {
//// 		json.WriteError(w, http.StatusBadRequest, err)
//// 		return
//// 	}
//
//// 	response, err := h.user.Api.GetUser(r.Context(), req)
//// 	if err != nil {
//// 		json.WriteError(w, http.StatusInternalServerError, err)
//// 		return
//// 	}
//
//// 	json.WriteJSON(w, http.StatusOK, response)
//// 	return
//// }
//
//func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
//	log := middleware.LoggerFromContext(r.Context()).With(
//		"component", "user_handler",
//		"method", "ListUsers",
//	)
//
//	log.Info("request_received")
//
//	limitStr := chi.URLParam(r, "limit")
//	offsetStr := chi.URLParam(r, "offset")
//
//	limit, err := strconv.ParseInt(limitStr, 10, 32)
//	if err != nil {
//		log.Warn("invalid_limit", "limit", limitStr, "error", err)
//		limit = 20
//	}
//	offset, err := strconv.ParseInt(offsetStr, 10, 64)
//	if err != nil {
//		log.Warn("invalid_offset", "offset", offsetStr, "error", err)
//		offset = 0
//	}
//
//	req := &userv1.ListUsersRequest{
//		Limit:  int32(limit),
//		Offset: int32(offset),
//	}
//
//	log.Debug("calling_user_service", "request", req)
//	resp, err := h.user.Api.ListUsers(r.Context(), req)
//	if err != nil {
//		log.Error("gRPC_call_failed", "error", err)
//		json.WriteError(w, http.StatusInternalServerError, err)
//		return
//	}
//
//	log.Info("succeeded", "users_count", len(resp.Users))
//	json.WriteJSON(w, http.StatusOK, resp)
//}
