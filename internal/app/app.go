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
	router.Use(auth.CorsMiddleware)
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
	organizationHandler := handler.NewOrganizationHandler(s.log, organizationgrpc)
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
		apiRouter.Route("/product", func(productRouter chi.Router) {
			productRouter.Group(func(authRouter chi.Router) {
				authRouter.Use(auth.AuthMiddleware)
				authRouter.Get("/{product_id}", productHandler.GetProduct)
				authRouter.Get("/", productHandler.ListProducts)
				authRouter.Post("/", productHandler.CreateProduct)
				authRouter.Put("/{product_id}", productHandler.UpdateProduct)
			})
		})
		apiRouter.Route("/organizations", func(orgRouter chi.Router) {
			orgRouter.Group(func(authRouter chi.Router) {
				authRouter.Use(auth.AuthMiddleware)
				authRouter.Post("/", organizationHandler.CreateOrganization)
				authRouter.Get("/", organizationHandler.ListOrganizations)
				authRouter.Get("/my", organizationHandler.GetOrganization)
				authRouter.Put("/my", organizationHandler.UpdateOrganization)
				authRouter.Patch("/my", organizationHandler.PatchOrganization)
				authRouter.Route("/relationship-types", func(rtRouter chi.Router) {
					rtRouter.Post("/", organizationHandler.CreateRelationshipType)
					rtRouter.Get("/{relationship_type_id}", organizationHandler.GetRelationshipType)
					rtRouter.Put("/{relationship_type_id}", organizationHandler.UpdateRelationshipType)
				})
			})
		})
		apiRouter.Route("/crm", func(crmRouter chi.Router) {
			crmRouter.Group(func(authRouter chi.Router) {
				authRouter.Use(auth.AuthMiddleware)
				authRouter.Route("/leads", func(lRouter chi.Router) {
					lRouter.Post("/", crmHandler.CreateLead)
					lRouter.Get("/my", crmHandler.ListLeads)
					lRouter.Get("/{lead_id}", crmHandler.GetLead)
					lRouter.Put("/{lead_id}", crmHandler.UpdateLead)
					lRouter.Patch("/{lead_id}", crmHandler.PatchLead)
					lRouter.Delete("/{lead_id}", crmHandler.DeleteLead)
				})
			})
		})
	})

	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", s.cfg.REST.Port), router)
}
