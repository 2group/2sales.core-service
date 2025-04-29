package grpc

import (
	"context"
	"time"

	servicev1 "github.com/2group/2sales.core-service/pkg/gen/go/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Api servicev1.ServiceServiceClient
}

func NewServiceClient(ctx context.Context, addr string, timeout time.Duration, retriesCount int) (*ServiceClient, error) {
	cc, err := grpc.DialContext(ctx, addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(CorrelationUnaryInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return &ServiceClient{
		Api: servicev1.NewServiceServiceClient(cc),
	}, nil
}
