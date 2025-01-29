package grpc

import (
	"context"
	"time"

	advertisementv1 "github.com/2group/2sales.core-service/pkg/gen/go/advertisement"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AdvertisementClient struct {
	Api advertisementv1.AdvertisementServiceClient
}

func NewAdvertisementClient(ctx context.Context, addr string, timeout time.Duration, retriesCount int) (*AdvertisementClient, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &AdvertisementClient{
		Api: advertisementv1.NewAdvertisementServiceClient(cc),
	}, nil
}
