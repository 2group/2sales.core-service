package grpc

import (
	"context"
	"time"

	userv1 "github.com/2group/2sales.core-service/pkg/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OtpClient struct {
	Api userv1.OtpServiceClient
}

func NewOtpClient(ctx context.Context, addr string, timeout time.Duration, retriesCount int) (*OtpClient, error) {
	cc, err := grpc.DialContext(ctx, addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.WithUnaryInterceptor(CorrelationUnaryInterceptor),
	)
	if err != nil {
		return nil, err
	}

	return &OtpClient{
		Api: userv1.NewOtpServiceClient(cc),
	}, nil
}
