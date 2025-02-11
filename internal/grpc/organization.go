package grpc

import (
	"context"
	"time"

	organizationv1 "github.com/2group/2sales.core-service/pkg/gen/go/organization"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrganizationClient struct {
	Api organizationv1.OrganizationServiceClient
}

func NewOrganizationClient(ctx context.Context, addr string, timeout time.Duration, retriesCount int) (*OrganizationClient, error) {
	cc, err := grpc.DialContext(ctx, addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(CorrelationUnaryInterceptor),
	)
	if err != nil {
		return nil, err
	}

	return &OrganizationClient{
		Api: organizationv1.NewOrganizationServiceClient(cc),
	}, nil
}
