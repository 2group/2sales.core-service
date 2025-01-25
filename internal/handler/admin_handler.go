package handler

import (
	"log/slog"
	"net/http"

	"github.com/2group/2sales.core-service/internal/grpc"
	"github.com/2group/2sales.core-service/internal/lib"
	userv1 "github.com/2group/2sales.core-service/pkg/gen/go/user"
	"github.com/2group/2sales.core-service/templates/page/brands"
	"github.com/2group/2sales.core-service/templates/page/dashboard"
	"github.com/2group/2sales.core-service/templates/page/login"
)

type AdminHandler struct {
	log  *slog.Logger
	user *grpc.UserClient
}

func NewAdminHandler(user *grpc.UserClient) *AdminHandler {
	return &AdminHandler{user: user}
}

func (h *AdminHandler) DashboardPage(w http.ResponseWriter, r *http.Request) {
	lib.Render(w, r, dashboard.Index())
}

func (h *AdminHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	lib.Render(w, r, login.Index(nil))
}

func (h *AdminHandler) BrandsPage(w http.ResponseWriter, r *http.Request) {
	lib.Render(w, r, brands.Index())
}

func (h *AdminHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := &userv1.LoginRequest{}

	req.Login = r.FormValue("email")
	req.Password = r.FormValue("password")

	_, err := h.user.Api.Login(r.Context(), req)
	if err != nil {
		errString := err.Error()
		lib.Render(w, r, login.Index(&errString))
		return
	}

	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
	return
}
