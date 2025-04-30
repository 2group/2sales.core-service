package grpc

import (
	"context"

	"github.com/2group/2sales.core-service/pkg/logging"
	"github.com/2group/2sales.core-service/pkg/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func CorrelationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// 1) pull the correlation id from incoming metadata
		md, _ := metadata.FromIncomingContext(ctx)
		var cid string
		if vals := md.Get("X-Correlation-Id"); len(vals) > 0 {
			cid = vals[0]
		}

		// 2) stash the raw id in context (for FromContext)
		ctx = context.WithValue(ctx, middleware.CorrelationIDKey, cid)

		// 3) build a per-RPC zerolog.Logger and put it into context
		rpcLogger := logging.L().With().
			Str("correlation_id", cid).
			Logger()
		ctx = rpcLogger.WithContext(ctx)

		// 4) invoke handler with enriched context
		return handler(ctx, req)
	}
}
