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
	auth "github.com/2group/2sales.core-service/pkg/middeware"
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

	organizationgrpc, err := grpc.NewOrganizationClient(context, s.cfg.GRPC.Organization, time.Hour, 2)
	if err != nil {
		panic(err)
	}

	productgrpc, err := grpc.NewProductClient(context, s.cfg.GRPC.Product, time.Hour, 2)
	if err != nil {
		panic(err)
	}

	crmgrpc, err := grpc.NewCrmClient(context, s.cfg.GRPC.CRM, time.Hour, 2)
	if err != nil {
		panic(err)
	}

	userHandler := handler.NewUserHandler(usergrpc)
	organizationHandler := handler.NewOrganizationHandler(organizationgrpc)
	productHandler := handler.NewProductHandler(s.log, productgrpc)
	crmHandler := handler.NewCrmHandler(s.log, crmgrpc)

	router.Route("/api/v1", func(apiRouter chi.Router) {
		apiRouter.Route("/user", func(userRouter chi.Router) {
			userRouter.Post("/login", userHandler.HandleLogin)
			userRouter.Post("/register", userHandler.HandleRegister)

			userRouter.Group(func(authRouter chi.Router) {
				authRouter.Use(auth.AuthMiddleware)
				authRouter.Get("/", userHandler.HandleGetUser)
				authRouter.Put("/", userHandler.HandleUpdateUser)
				authRouter.Patch("/", userHandler.HandlePatchUser)
			})
		})
		apiRouter.Route("/category", func(categoryRouter chi.Router) {
			categoryRouter.Post("/", productHandler.CreateCategory)
			categoryRouter.Get("/", productHandler.GetFirstLevelCategories)
			categoryRouter.Get("/{category_id}", productHandler.GetCategory)

			categoryRouter.Group(func(authRouter chi.Router) {
				authRouter.Use(auth.AuthMiddleware)
			})
		})
		apiRouter.Route("/organization", func(organizationRouter chi.Router) {
			organizationRouter.Group(func(authRouter chi.Router) {
				authRouter.Use(auth.AuthMiddleware)
				authRouter.Post("/", organizationHandler.CreateOrganization)
				authRouter.Get("/", organizationHandler.GetOrganization)
				authRouter.Patch("/", organizationHandler.PatchOrganization)
				authRouter.Get("/list", organizationHandler.ListOrganizations)
			})
		})
		apiRouter.Route("/crm", func(crmRouter chi.Router) {
			crmRouter.Group(func(authRouter chi.Router) {
				authRouter.Use(auth.AuthMiddleware)
				authRouter.Post("/leads", crmHandler.CreateLead)
				authRouter.Get("/leads", crmHandler.ListLeads)
				authRouter.Get("/leads/{lead_id}", crmHandler.GetLead)
				authRouter.Put("/leads/{lead_id}", crmHandler.UpdateLead)
				authRouter.Patch("/leads/{lead_id}", crmHandler.PatchLead)
				authRouter.Delete("/leads/{lead_id}", crmHandler.DeleteLead)
			})
		})
	})

	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", s.cfg.REST.Port), router)
}
