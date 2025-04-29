package grpc

import (
	"context"
	"time"

	b2c_service_orderv1 "github.com/2group/2sales.core-service/pkg/gen/go/b2c_service_order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type B2CServiceOrderClient struct {
	Api b2c_service_orderv1.B2CServiceOrderServiceClient
}

func NewB2CServiceOrderClient(ctx context.Context, addr string, timeout time.Duration, retriesCount int) (*B2CServiceOrderClient, error) {
	cc, err := grpc.DialContext(ctx, addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.WithUnaryInterceptor(CorrelationUnaryInterceptor),
	)
	if err != nil {
		return nil, err
	}

	return &B2CServiceOrderClient{
		Api: b2c_service_orderv1.NewB2CServiceOrderServiceClient(cc),
	}, nil
}
