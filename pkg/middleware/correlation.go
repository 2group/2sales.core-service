// pkg/middleware/correlation.go
package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	// CorrelationIDKey is the key for the correlation ID stored in context.
	CorrelationIDKey contextKey = "correlationID"
)

// CorrelationMiddleware extracts the correlation ID from the incoming request header,
// or generates a new one if none exists, and then attaches it to the request context.
func CorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Try to get an incoming correlation ID from the request header.
		cid := r.Header.Get("X-Correlation-ID")
		if cid == "" {
			// If not present, generate a new one.
			cid = uuid.NewString()
		}

		// Optionally, add the correlation ID to the response header so clients can see it.
		w.Header().Set("X-Correlation-ID", cid)

		// Store the correlation ID in the request's context.
		ctx := context.WithValue(r.Context(), CorrelationIDKey, cid)

		// Pass the request with the new context to the next handler.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func WithCorrelationID(ctx context.Context, cid string) context.Context {
	return context.WithValue(ctx, CorrelationIDKey, cid)
}

// GetCorrelationID retrieves the correlation ID from the context.
func GetCorrelationID(ctx context.Context) (string, bool) {
	cid, ok := ctx.Value(CorrelationIDKey).(string)
	return cid, ok
}
