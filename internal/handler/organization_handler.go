package handler

import (
	"log/slog"

	"github.com/2group/2sales.core-service/internal/grpc"
)

type OrganizationHandler struct {
    log *slog.Logger
    organization *grpc.OrganizationClient
}

func NewOrganizationHandler() *OrganizationHandler{
    return &OrganizationHandler{}
}
