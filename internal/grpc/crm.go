package grpc

import (
	"context"
	"time"

	crmv1 "github.com/2group/2sales.core-service/pkg/gen/go/crm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CrmClient struct {
	Api crmv1.CRMServiceClient
}

func NewCrmClient(ctx context.Context, addr string, timeout time.Duration, retriesCount int) (*CrmClient, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &CrmClient{
		Api: crmv1.NewCRMServiceClient(cc),
	}, nil
}
