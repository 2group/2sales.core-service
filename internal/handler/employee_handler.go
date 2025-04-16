package handler

import (
	"log/slog"

	"github.com/2group/2sales.core-service/internal/grpc"
)

type EmployeeHandler struct {
	log      *slog.Logger
	employee *grpc.EmployeeClient
}

func NewEmployeeHandler(log *slog.Logger, customer *grpc.CustomerClient) *CustomerHandler {
	return &CustomerHandler{
		log:      log,
		customer: customer,
	}
}
