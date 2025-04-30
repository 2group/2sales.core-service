package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type CtxKeyCorrelationID struct{}

var CorrelationIDKey = &CtxKeyCorrelationID{}

// CorrelationMiddleware adds a correlation ID and request-scoped logger to the context.
func CorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cid := r.Header.Get("X-Correlation-ID")
		if cid == "" {
			cid = uuid.NewString()
		}
		w.Header().Set("X-Correlation-ID", cid)

		logger := log.Logger.With().Str("correlation_id", cid).Logger()

		ctx := logger.WithContext(r.Context())
		ctx = context.WithValue(ctx, CorrelationIDKey, cid)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LoggerFromContext(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}

func CorrelationIDFromContext(ctx context.Context) string {
	if val, ok := ctx.Value(CorrelationIDKey).(string); ok {
		return val
	}
	return ""
}
