package handler

import (
	"net/http"

	"github.com/2group/2sales.core-service/internal/lib"
	"github.com/2group/2sales.core-service/templates/page/home"
)

type HomeHandler struct {
}

func NewHomeHandler() *HomeHandler {
    return &HomeHandler{}
}

func (h *HomeHandler) HandleHome(w http.ResponseWriter, r *http.Request) {
	lib.Render(w, r, home.Index())
}
