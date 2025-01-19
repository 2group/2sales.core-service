package grpc

import (
	"context"
	"time"

	userv1 "github.com/2group/2sales.core-service/pkg/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	Api userv1.UserServiceClient
}

func NewUserClient(ctx context.Context, addr string, timeout time.Duration, retriesCount int) (*UserClient, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &UserClient{
		Api: userv1.NewUserServiceClient(cc),
	}, nil
}
