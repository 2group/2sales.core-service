package grpc

import (
	"context"
	"time"

	employeev1 "github.com/2group/2sales.core-service/pkg/gen/go/employee"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type EmployeeClient struct {
	Api employeev1.EmployeeServiceClient
}

func NewEmployeeClient(ctx context.Context, addr string, timeout time.Duration, retriesCount int) (*EmployeeClient, error) {
	cc, err := grpc.DialContext(ctx, addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.WithUnaryInterceptor(CorrelationUnaryInterceptor),
	)
	if err != nil {
		return nil, err
	}

	return &EmployeeClient{
		Api: employeev1.NewEmployeeServiceClient(cc),
	}, nil
}
