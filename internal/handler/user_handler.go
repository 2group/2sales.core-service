package handler

import (
	"log/slog"
	"net/http"

	"github.com/2group/2sales.core-service/internal/grpc"
	userv1 "github.com/2group/2sales.core-service/pkg/gen/go/user"
	"github.com/2group/2sales.core-service/pkg/json"
)

type UserHandler struct {
    log *slog.Logger
    user *grpc.UserClient
}

func NewUserHandler(user *grpc.UserClient) *UserHandler {
    return &UserHandler{user: user}
}

func (h *UserHandler) HandleLogin (w http.ResponseWriter, r *http.Request) {
    req :=&userv1.LoginRequest{}
    err := json.ParseJSON(r, &req)
    if err != nil {
        h.log.Info("%s", err.Error()) 
        json.WriteError(w, http.StatusBadRequest, err)
        return
    }

    response, err := h.user.Api.Login(r.Context(), req)
    if err != nil {
        h.log.Info("%s", err.Error()) 
        json.WriteError(w, http.StatusBadRequest, err)
        return
    }

    json.WriteJSON(w, http.StatusOK, response)
    return
}
