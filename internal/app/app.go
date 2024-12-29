package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/2group/2sales.core-service/internal/config"
	"github.com/2group/2sales.core-service/internal/grpc"
	"github.com/2group/2sales.core-service/internal/handler"
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

    usergrpc, err := grpc.NewUserClient(context, s.cfg.GRPC.User, time.Hour, 2)
    if err != nil {
        panic(err)
    }

    userHandler := handler.NewUserHandler(usergrpc)

    router.Route("/api/v1", func(apiRouter chi.Router) {
        apiRouter.Route("/user", func(userRouter chi.Router) {
			userRouter.Post("/login", userHandler.HandleLogin)
        })
    })

    return http.ListenAndServe(fmt.Sprintf("localhost:%d", s.cfg.REST.Port), router)
}
