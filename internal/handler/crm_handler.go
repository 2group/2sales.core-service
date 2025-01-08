package handler

import (
	"log/slog"

	"github.com/2group/2sales.core-service/internal/grpc"
)

type CrmHandler struct {
	log *slog.Logger
	crm *grpc.CrmClient
}

func NewCrmHandler(log *slog.Logger, crm *grpc.CrmClient) *CrmHandler {
	return &CrmHandler{log: log, crm: crm}
}
