package grpc

import (
	"context"
	"time"

	orderv1 "github.com/2group/2sales.core-service/pkg/gen/go/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderClient struct {
	Api orderv1.OrderServiceClient
}

func NewOrderClient(ctx context.Context, addr string, timeout time.Duration, retriesCount int) (*OrderClient, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &OrderClient{
		Api: orderv1.NewOrderServiceClient(cc),
	}, nil
}
