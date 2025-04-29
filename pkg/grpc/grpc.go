package grpc

import (
	"context"

	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ctxKeyLogger struct{}

var loggerKey = &ctxKeyLogger{}

// CorrelationUnaryServerInterceptor extracts the correlation ID from metadata,
// attaches it to the slog logger, and stores that in the RPC context.
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
