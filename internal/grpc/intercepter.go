// internal/grpc/interceptor.go
package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/2group/2sales.core-service/pkg/middleware"
)

// CorrelationUnaryInterceptor is a client interceptor that attaches the correlation ID
// from the context (if any) to the outgoing gRPC metadata.
func CorrelationUnaryInterceptor(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	if cid, ok := middleware.GetCorrelationID(ctx); ok {
		// Extract or create outgoing metadata.
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}
		md.Set("x-correlation-id", cid)
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return invoker(ctx, method, req, reply, cc, opts...)
}

// CorrelationUnaryServerInterceptor is a server interceptor that extracts the correlation ID
// from incoming metadata and attaches it to the context.
func CorrelationUnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if ids := md.Get("x-correlation-id"); len(ids) > 0 {
			// Attach the correlation ID into the context.
			ctx = middleware.WithCorrelationID(ctx, ids[0])
		}
	}
	return handler(ctx, req)
}
