package middleware

import (
	"context"
	"net/http"

	"log/slog"

	"github.com/google/uuid"
)

// context key for request-scoped logger
type CtxKeyLogger struct{}

var LoggerKey = &CtxKeyLogger{}

type CtxKeyCorrelationID struct{}

var CorrelationIDKey = &CtxKeyCorrelationID{}

func CorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid := r.Header.Get("X-Correlation-ID")
		if cid == "" {
			cid = uuid.NewString()
		}
		w.Header().Set("X-Correlation-Id", cid)

		// 1) stash in logger
		reqLogger := slog.Default().With("correlation_id", cid)
		ctx := context.WithValue(r.Context(), LoggerKey, reqLogger)

		// 2) also stash raw cid for gRPC propagation
		ctx = context.WithValue(ctx, CorrelationIDKey, cid)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// LoggerFromContext returns the request-scoped logger or falls back to default.
func LoggerFromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(LoggerKey).(*slog.Logger); ok {
		return l
	}
	return slog.Default()
}
