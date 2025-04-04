package grpc

import (
	"context"
	"time"

	customerv1 "github.com/2group/2sales.core-service/pkg/gen/go/customer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CustomerClient struct {
	Api customerv1.CustomerServiceClient
}

func NewCustomerClient(ctx context.Context, addr string, timeout time.Duration, retriesCount int) (*CustomerClient, error) {
	cc, err := grpc.DialContext(ctx, addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.WithUnaryInterceptor(CorrelationUnaryInterceptor),
	)
	if err != nil {
		return nil, err
	}

	return &CustomerClient{
		Api: customerv1.NewCustomerServiceClient(cc),
	}, nil
}
