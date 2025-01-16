package grpc

import (
	"context"
	"time"

	warehousev1 "github.com/2group/2sales.core-service/pkg/gen/go/warehouse"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type WarehouseClient struct {
	Api warehousev1.WarehouseServiceClient
}

func NewWarehouseClient(ctx context.Context, addr string, timeout time.Duration, retriesCount int) (*WarehouseClient, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &WarehouseClient{
		Api: warehousev1.NewWarehouseServiceClient(cc),
	}, nil
}
