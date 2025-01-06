package grpc

import (
	"context"
	"time"

	productv1 "github.com/2group/2sales.core-service/pkg/gen/go/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductClient struct {
    Api productv1.ProductServiceClient 
}

func NewProductClient(ctx context.Context, addr string, timeout time.Duration, retriesCount int) (*ProductClient, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &ProductClient{
		Api: productv1.NewProductServiceClient(cc),
	}, nil
}


