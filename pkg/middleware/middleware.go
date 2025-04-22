package middleware

import (
	"context"
	"net/http"
	"strings"

	jwtv1 "github.com/2group/2sales.core-service/pkg/jwt"
	"github.com/golang-jwt/jwt/v5"
)

var secret = "2group.kz"

type contextKey string

const (
	ContextUserIDKey     contextKey = "user_id"
	ContextEmployeeIDKey contextKey = "employee_id"
	ContextCustomerIDKey contextKey = "customer_id"
	ContextScopesKey     contextKey = "scopes"
)

// AuthMiddleware validates the JWT and injects user/context fields.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]

		// Parse and validate
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// 1) user_id
		rawUser, ok := claims["user_id"].(float64)
		if !ok {
			http.Error(w, "Invalid or missing user_id", http.StatusUnauthorized)
			return
		}
		userID := int64(rawUser)

		// 2) employee_id or customer_id
		var employeeID *int64
		if raw, exists := claims["employee_id"]; exists && raw != nil {
			if f, ok := raw.(float64); ok {
				id := int64(f)
				employeeID = &id
			} else {
				http.Error(w, "Invalid employee_id", http.StatusUnauthorized)
				return
			}
		}
		var customerID *int64
		if raw, exists := claims["customer_id"]; exists && raw != nil {
			if f, ok := raw.(float64); ok {
				id := int64(f)
				customerID = &id
			} else {
				http.Error(w, "Invalid customer_id", http.StatusUnauthorized)
				return
			}
		}
		if employeeID == nil && customerID == nil {
			http.Error(w, "Neither employee_id nor customer_id present", http.StatusUnauthorized)
			return
		}

		// 3) scopes array
		var scopes []jwtv1.RoleScope
		if rawScopes, exists := claims["scopes"]; exists && rawScopes != nil {
			// rawScopes is []interface{}
			items, ok := rawScopes.([]interface{})
			if !ok {
				http.Error(w, "Invalid scopes format", http.StatusUnauthorized)
				return
			}
			for _, item := range items {
				m, ok := item.(map[string]interface{})
				if !ok {
					http.Error(w, "Invalid scope entry", http.StatusUnauthorized)
					return
				}
				// parse role_id
				f, ok := m["role_id"].(float64)
				if !ok {
					http.Error(w, "Invalid role_id in scope", http.StatusUnauthorized)
					return
				}
				rs := jwtv1.RoleScope{RoleID: int64(f)}
				// optional organization_id
				if rawOrg, exists := m["organization_id"]; exists && rawOrg != nil {
					if of, ok := rawOrg.(float64); ok {
						o := int64(of)
						rs.OrganizationID = &o
					} else {
						http.Error(w, "Invalid organization_id in scope", http.StatusUnauthorized)
						return
					}
				}
				// optional branch_id
				if rawBr, exists := m["branch_id"]; exists && rawBr != nil {
					if bf, ok := rawBr.(float64); ok {
						b := int64(bf)
						rs.BranchID = &b
					} else {
						http.Error(w, "Invalid branch_id in scope", http.StatusUnauthorized)
						return
					}
				}
				scopes = append(scopes, rs)
			}
		}

		// 4) Inject into context
		ctx := r.Context()
		ctx = context.WithValue(ctx, ContextUserIDKey, userID)
		if employeeID != nil {
			ctx = context.WithValue(ctx, ContextEmployeeIDKey, *employeeID)
		}
		if customerID != nil {
			ctx = context.WithValue(ctx, ContextCustomerIDKey, *customerID)
		}
		ctx = context.WithValue(ctx, ContextScopesKey, scopes)

		// Call next with enriched context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
