package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/2group/2sales.core-service/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type APIServer struct {
	cfg *config.Config
	log *slog.Logger
}

func NewAPIServer(cfg *config.Config, log *slog.Logger) *APIServer {
	return &APIServer{cfg: cfg, log: log}
}

func (s *APIServer) Run() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	context := context.Background()
    _ = context

    return http.ListenAndServe(fmt.Sprintf("localhost:%d", s.cfg.REST.Port), router)
}
