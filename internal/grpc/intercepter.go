// internal/grpc/interceptors.go
package grpc

import (
	"context"
	"log/slog"

	"github.com/2group/2sales.core-service/pkg/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ctxKeyLogger struct{}

var loggerKey = &ctxKeyLogger{}

func CorrelationUnaryServerInterceptor(baseLogger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// 1) pull metadata
		md, _ := metadata.FromIncomingContext(ctx)
		var cid string
		if vals := md.Get("x-correlation-id"); len(vals) > 0 {
			cid = vals[0]
		}

		// 2) build a per-RPC logger
		rpcLogger := baseLogger.With("correlation_id", cid)

		// 3) stash it in the context
		ctx = context.WithValue(ctx, loggerKey, rpcLogger)

		// 4) call the handler
		return handler(ctx, req)
	}
}

func CorrelationUnaryInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// pull the cid
		if cid, ok := ctx.Value(middleware.CorrelationIDKey).(string); ok {
			// append to outgoing metadata
			ctx = metadata.AppendToOutgoingContext(ctx, "X-Correlation-Id", cid)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
