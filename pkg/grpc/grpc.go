package grpc

import (
	"context"

	"github.com/2group/2sales.core-service/pkg/logging"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type CtxKeyGrpcLogger struct{}

var GrpcLoggerKey = &CtxKeyGrpcLogger{}

// CorrelationUnaryServerInterceptor extracts the correlation ID from metadata,
// attaches it to the slog logger, and stores that in the RPC context.
func CorrelationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// 1) Extract correlation ID from metadata
		md, _ := metadata.FromIncomingContext(ctx)
		var cid string
		if vals := md.Get("X-Correlation-Id"); len(vals) > 0 {
			cid = vals[0]
		}

		// 2) Build a per-RPC logger with the correlation_id field
		rpcLogger := logging.Slog().With().Str("correlation_id", cid)

		// 3) Store it in context so handlers can pull it out
		ctx = context.WithValue(ctx, GrpcLoggerKey, rpcLogger)

		// 4) Continue handling the RPC
		return handler(ctx, req)
	}
}
