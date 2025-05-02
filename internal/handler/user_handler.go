package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	userv1 "github.com/2group/2sales.core-service/pkg/gen/go/user"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type UserHandler struct {
	log  *slog.Logger
	user *grpc.UserClient
}

func NewUserHandler(user *grpc.UserClient) *UserHandler {
	return &UserHandler{
		user: user,
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "user_handler").
		Str("method", "Login").
		Logger()

	log.Info().Msg("request_received")

	req := &userv1.LoginRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Debug().Any("request", req).Msg("calling_user_service")
	resp, err := h.user.Api.Login(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "organization_handler").
		Str("method", "GetOrganization").
		Logger()

	log.Info().Msg("request_received")

	req := &userv1.UpdateUserRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Debug().Any("request", req).Msg("calling_user_service")
	resp, err := h.user.Api.UpdateUser(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int64("user_id", *resp.User.Id).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "organization_handler").
		Str("method", "GetOrganization").
		Logger()

	log.Info().Msg("request_received")

	req := &userv1.CreateUserRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Debug().Any("request", req).Msg("calling_user_service")
	resp, err := h.user.Api.CreateUser(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int64("user_id", *resp.User.Id).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
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

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "organization_handler").
		Str("method", "GetOrganization").
		Logger()

	log.Info().Msg("request_received")

	limitStr := chi.URLParam(r, "limit")
	offsetStr := chi.URLParam(r, "offset")

	limit, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		log.Warn().Err(err).Int64("limit", limit).Msg("invalid_limit")
		limit = 20
	}
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		log.Warn().Err(err).Int64("offset", offset).Msg("invalid_offset")
		offset = 0
	}

	req := &userv1.ListUsersRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	log.Debug().Any("request", req).Msg("calling_user_service")
	resp, err := h.user.Api.ListUsers(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int("users_count", len(resp.Users)).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *UserHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "user_handler").
		Str("method", "RefreshToken").
		Logger()

	log.Info().Msg("request_received")

	req := &userv1.RefreshTokenRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.user.Api.RefreshToken(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("grpc_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "user_handler").
		Str("method", "Logout").
		Logger()

	log.Info().Msg("request_received")

	req := &userv1.LogoutRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err := h.user.Api.Logout(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("grpc_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Msg("succeeded")
	w.WriteHeader(http.StatusOK)
}
