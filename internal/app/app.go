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
	auth "github.com/2group/2sales.core-service/pkg/middleware"
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
	router.Use(auth.CorrelationMiddleware)
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

	warehousegrpc, err := grpc.NewWarehouseClient(context, s.cfg.GRPC.Warehouse, time.Hour, 2)
	if err != nil {
		panic(err)
	}

	ordergrpc, err := grpc.NewOrderClient(context, s.cfg.GRPC.Order, time.Hour, 2)
	if err != nil {
		panic(err)
	}

	advertisementgrpc, err := grpc.NewAdvertisementClient(context, s.cfg.GRPC.Advertisement, time.Hour, 2)
	if err != nil {
		panic(err)
	}

	userHandler := handler.NewUserHandler(usergrpc)
	organizationHandler := handler.NewOrganizationHandler(s.log, organizationgrpc)
	productHandler := handler.NewProductHandler(s.log, productgrpc)
	crmHandler := handler.NewCrmHandler(s.log, crmgrpc)
	warehouseHandler := handler.NewWarehouseHandler(s.log, warehousegrpc)
	orderHandler := handler.NewOrderHandler(s.log, ordergrpc)
	advertisementHandler := handler.NewAdvertisementHandler(s.log, advertisementgrpc)
	adminHandler := handler.NewAdminHandler(usergrpc, organizationgrpc)

	router.Route("/admin/api", func(adminRouter chi.Router) {
		adminRouter.Post("/login", adminHandler.Login)
		adminRouter.Get("/brands", adminHandler.ListBrands)
		adminRouter.Post("/brands", adminHandler.CreateBrand)
		adminRouter.Patch("/brands/{brand_id}/images", adminHandler.PatchBrand)
		adminRouter.Post("/brands/presigned-urls", adminHandler.GeneratePresignedURLs)
	})

	router.Route("/api/v1", func(apiRouter chi.Router) {
		apiRouter.Route("/user", func(userRouter chi.Router) {
			userRouter.Post("/login", userHandler.HandleLogin)
			userRouter.Post("/register", userHandler.HandleRegister)
			userRouter.Group(func(authRouter chi.Router) {
				authRouter.Use(auth.AuthMiddleware)
				authRouter.Get("/me", userHandler.HandleGetMyProfile)
				authRouter.Put("/me", userHandler.HandleUpdateMyProfile)
				authRouter.Patch("/me", userHandler.HandlePatchMyProfile)
				authRouter.Route("/roles", func(rolesRouter chi.Router) {
					rolesRouter.Get("/my", userHandler.ListMyOrganizationRoles)
				})
				authRouter.Route("/users", func(usersRouter chi.Router) {
					usersRouter.Post("/", userHandler.CreateUser)
					usersRouter.Get("/my", userHandler.ListMyOrganizationUsers)
					usersRouter.Patch("/{user_id}", userHandler.PatchUser)
				})
			})
		})
		apiRouter.Route("/product", func(productRouter chi.Router) {
			productRouter.Group(func(authRouter chi.Router) {
				authRouter.Use(auth.AuthMiddleware)
				authRouter.Get("/", productHandler.ListProducts)
				authRouter.Post("/", productHandler.CreateProduct)
				authRouter.Get("/{product_id}", productHandler.GetProduct)
				authRouter.Patch("/{product_id}", productHandler.PatchProduct)
				authRouter.Route("/categories", func(categoryRouter chi.Router) {
					categoryRouter.Get("/", productHandler.ListCategories)
					categoryRouter.Post("/", productHandler.CreateCategory)
					categoryRouter.Get("/{category_id}", productHandler.GetCategory)
					categoryRouter.Group(func(authRouter chi.Router) {
						authRouter.Use(auth.AuthMiddleware)
					})
				})
				authRouter.Route("/group", func(productGroupRouter chi.Router) {
					productGroupRouter.Get("/{product_group_id}", productHandler.GetProductGroup)
					productGroupRouter.Get("/", productHandler.ListProductGroup)
					productGroupRouter.Post("/", productHandler.CreateProductGroup)
					productGroupRouter.Put("/{product_group_id}", productHandler.UpdateProductGroup)
					productGroupRouter.Delete("/{product_group_id}", productHandler.DeleteProductGroup)
				})
			})
		})
		apiRouter.Route("/organizations", func(orgRouter chi.Router) {
			orgRouter.Group(func(authRouter chi.Router) {
				authRouter.Use(auth.AuthMiddleware)
				authRouter.Get("/", organizationHandler.ListOrganizations)
				authRouter.Post("/my", organizationHandler.CreateOrganization)
				authRouter.Get("/my", organizationHandler.GetMyOrganization)
				authRouter.Put("/my", organizationHandler.UpdateMyOrganization)
				authRouter.Patch("/my", organizationHandler.PatchMyOrganization)
				authRouter.Post("/presigned-urls", organizationHandler.GeneratePresignedURLs)
				authRouter.Get("/{organization_id}", organizationHandler.GetOrganization)
				authRouter.Route("/addresses", func(aRouter chi.Router) {
					aRouter.Post("/", organizationHandler.CreateAddress)
					aRouter.Put("/{address_id}", organizationHandler.UpdateAddress)
					aRouter.Get("/{address_id}", organizationHandler.GetAddress)
					aRouter.Patch("/{address_id}", organizationHandler.PatchAddress)
					aRouter.Delete("/{address_id}", organizationHandler.DeleteAddress)
				})
				authRouter.Route("/relationships", func(rRouter chi.Router) {
					rRouter.Post("/", organizationHandler.CreateRelationship)
					rRouter.Get("/my", organizationHandler.ListRelationships)
					rRouter.Put("/{relationship_id}", organizationHandler.UpdateRelationship)
				})
				authRouter.Route("/relationship-types", func(rtRouter chi.Router) {
					rtRouter.Post("/", organizationHandler.CreateRelationshipType)
					rtRouter.Get("/my", organizationHandler.ListRelationshipTypes)
					rtRouter.Get("/{relationship_type_id}", organizationHandler.GetRelationshipType)
					rtRouter.Put("/{relationship_type_id}", organizationHandler.UpdateRelationshipType)
				})
				authRouter.Route("/contacts", func(cRouter chi.Router) {
					cRouter.Post("/", organizationHandler.CreateContact)
					cRouter.Get("/{contact_id}", organizationHandler.GetContact)
					cRouter.Put("/{contact_id}", organizationHandler.UpdateContact)
					cRouter.Delete("/{contact_id}", organizationHandler.DeleteContact)
				})
				authRouter.Route("/counterparties", func(coRouter chi.Router) {
					coRouter.Post("/", organizationHandler.CreateCounterparty)
					coRouter.Get("/my", organizationHandler.ListCounterparties)
					coRouter.Get("/{counterparty_id}", organizationHandler.GetCounterparty)
					coRouter.Patch("/{counterparty_id}", organizationHandler.PatchMyCounterparty)
				})
				authRouter.Route("/bank_accounts", func(baRouter chi.Router) {
					baRouter.Get("/my", organizationHandler.ListMyBankAccounts)
					baRouter.Put("/my", organizationHandler.UpdateMyBankAccounts)
					baRouter.Post("/", organizationHandler.CreateBankAccount)
					baRouter.Put("/{bank_account_id}", organizationHandler.UpdateBankAccount)
					baRouter.Delete("/{bank_account_id}", organizationHandler.DeleteBankAccount)
				})
				authRouter.Route("/settings", func(baRouter chi.Router) {
					baRouter.Get("/", organizationHandler.ListSaleSettings)
					baRouter.Put("/", organizationHandler.UpdateSaleSettings)
					baRouter.Post("/", organizationHandler.CreateSaleSettings)
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
		apiRouter.Route("/warehouse", func(warehouseRouter chi.Router) {
			warehouseRouter.Group(func(authRouter chi.Router) {
				authRouter.Use(auth.AuthMiddleware)
				authRouter.Get("/", warehouseHandler.ListWarehouses)
				authRouter.Post("/", warehouseHandler.CreateWarehouse)
				authRouter.Get("/{warehouse_id}", warehouseHandler.GetWarehouse)
				authRouter.Get("/product/{warehouse_id}", warehouseHandler.GetCountProducts)
				authRouter.Put("/{warehouse_id}", warehouseHandler.UpdateWarehouse)
				authRouter.Route("/acceptance", func(acpRouter chi.Router) {
					acpRouter.Use(auth.AuthMiddleware)
					acpRouter.Get("/", warehouseHandler.ListAcceptances)
					acpRouter.Get("/{acceptance_id}", warehouseHandler.GetAcceptance)
					acpRouter.Post("/", warehouseHandler.CreateAcceptance)
				})
			})
			warehouseRouter.Route("/write_off", func(writeOffRouter chi.Router) {
				writeOffRouter.Group(func(authRouter chi.Router) {
					authRouter.Use(auth.AuthMiddleware)
					authRouter.Get("/", warehouseHandler.ListWriteOff)
					authRouter.Get("/{write_off_id}", warehouseHandler.GetWriteOff)
					authRouter.Post("/", warehouseHandler.CreateWriteOff)
				})
			})
			warehouseRouter.Route("/moving", func(movingRouter chi.Router) {
				movingRouter.Group(func(authRouter chi.Router) {
					authRouter.Use(auth.AuthMiddleware)
					authRouter.Get("/", warehouseHandler.ListMoving)
					authRouter.Get("/{moving_id}", warehouseHandler.GetMoving)
					authRouter.Post("/", warehouseHandler.CreateMoving)
				})
			})
			warehouseRouter.Route("/inventory", func(movingRouter chi.Router) {
				movingRouter.Group(func(authRouter chi.Router) {
					authRouter.Use(auth.AuthMiddleware)
					authRouter.Get("/", warehouseHandler.ListInventory)
					authRouter.Get("/{inventory_id}", warehouseHandler.GetInventory)
					authRouter.Post("/", warehouseHandler.CreateInventory)
				})
			})
		})
		apiRouter.Route("/orders", func(orderRouter chi.Router) {
			orderRouter.Route("/sub-orders", func(suborderRouter chi.Router) {
				suborderRouter.Group(func(authRouter chi.Router) {
					authRouter.Use(auth.AuthMiddleware)
					authRouter.Post("/", orderHandler.CreateSubOrder)
					authRouter.Post("/order", orderHandler.CreateOrder)
					authRouter.Get("/", orderHandler.ListSubOrder)
					authRouter.Get("/{suborder_id}", orderHandler.GetSubOrder)
					authRouter.Put("/{suborder_id}", orderHandler.UpdateSubOrder)
				})
			})
			orderRouter.Route("/carts", func(cartRouter chi.Router) {
				cartRouter.Group(func(authRouter chi.Router) {
					authRouter.Use(auth.AuthMiddleware)
					authRouter.Get("/", orderHandler.GetCart)
					authRouter.Post("/add", orderHandler.AddProductToCart)
					authRouter.Post("/delete", orderHandler.DeleteProductFromCart)
				})
			})
		})
		apiRouter.Route("/advertisements", func(advertisementRouter chi.Router) {
			advertisementRouter.Group(func(authRouter chi.Router) {
				authRouter.Use(auth.AuthMiddleware)
				authRouter.Post("/banners", advertisementHandler.CreateBanner)
				authRouter.Get("/banners", advertisementHandler.ListBanners)
			})
		})
	})

	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", s.cfg.REST.Port), router)
}
