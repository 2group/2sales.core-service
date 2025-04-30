package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	jwtv1 "github.com/2group/2sales.core-service/pkg/jwt"
	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("2group.kz")

type contextKey string

const (
	ContextUserIDKey     contextKey = "user_id"
	ContextEmployeeIDKey contextKey = "employee_id"
	ContextCustomerIDKey contextKey = "customer_id"
	ContextScopesKey     contextKey = "scopes"
)

// extractToken pulls the Bearer token out of the Authorization header.
func extractToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid authorization header format")
	}
	return parts[1], nil
}

// validateToken parses and validates the JWT, returning its claims.
func validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}
	return claims, nil
}

// AuthMiddleware validates the JWT and injects user/context fields.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1) extract the raw token
		tokenString, err := extractToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// 2) validate & get claims
		claims, err := validateToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// 3) user_id
		rawUser, ok := claims["user_id"].(float64)
		if !ok {
			http.Error(w, "invalid or missing user_id", http.StatusUnauthorized)
			return
		}
		userID := int64(rawUser)
		//ctx := context.WithValue(r.Context(), LoggerKey,
		//	LoggerFromContext(r.Context()).With("user_id", userID),
		//)

		// 4) employee_id or customer_id
		var employeeID *int64
		if raw, exists := claims["employee_id"]; exists && raw != nil {
			if f, ok := raw.(float64); ok {
				v := int64(f)
				employeeID = &v
			} else {
				http.Error(w, "invalid employee_id", http.StatusUnauthorized)
				return
			}
		}
		var customerID *int64
		if raw, exists := claims["customer_id"]; exists && raw != nil {
			if f, ok := raw.(float64); ok {
				v := int64(f)
				customerID = &v
			} else {
				http.Error(w, "invalid customer_id", http.StatusUnauthorized)
				return
			}
		}
		if employeeID == nil && customerID == nil {
			http.Error(w, "neither employee_id nor customer_id present", http.StatusUnauthorized)
			return
		}

		// 5) scopes
		var scopes []jwtv1.RoleScope
		if rawScopes, exists := claims["scopes"]; exists && rawScopes != nil {
			items, ok := rawScopes.([]interface{})
			if !ok {
				http.Error(w, "invalid scopes format", http.StatusUnauthorized)
				return
			}
			for _, item := range items {
				m, ok := item.(map[string]interface{})
				if !ok {
					http.Error(w, "invalid scope entry", http.StatusUnauthorized)
					return
				}
				f, ok := m["role_id"].(float64)
				if !ok {
					http.Error(w, "invalid role_id in scope", http.StatusUnauthorized)
					return
				}
				rs := jwtv1.RoleScope{RoleID: int64(f)}
				if rawOrg, ok := m["organization_id"]; ok && rawOrg != nil {
					if of, ok := rawOrg.(float64); ok {
						v := int64(of)
						rs.OrganizationID = &v
					} else {
						http.Error(w, "invalid organization_id in scope", http.StatusUnauthorized)
						return
					}
				}
				if rawBr, ok := m["branch_id"]; ok && rawBr != nil {
					if bf, ok := rawBr.(float64); ok {
						v := int64(bf)
						rs.BranchID = &v
					} else {
						http.Error(w, "invalid branch_id in scope", http.StatusUnauthorized)
						return
					}
				}
				scopes = append(scopes, rs)
			}
		}

		ctx := context.WithValue(r.Context(), ContextUserIDKey, userID)
		if employeeID != nil {
			ctx = context.WithValue(ctx, ContextEmployeeIDKey, *employeeID)
		}
		if customerID != nil {
			ctx = context.WithValue(ctx, ContextCustomerIDKey, *customerID)
		}
		ctx = context.WithValue(ctx, ContextScopesKey, scopes)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserID(r *http.Request) (int64, bool) {
	uid, ok := r.Context().Value(ContextUserIDKey).(int64)
	return uid, ok
}

// GetEmployeeID returns the employee_id from context, if present.
func GetEmployeeID(r *http.Request) (int64, bool) {
	eid, ok := r.Context().Value(ContextEmployeeIDKey).(int64)
	return eid, ok
}

// GetCustomerID returns the customer_id from context, if present.
func GetCustomerID(r *http.Request) (int64, bool) {
	cid, ok := r.Context().Value(ContextCustomerIDKey).(int64)
	return cid, ok
}

// GetScopes returns the parsed RoleScope slice from context, if present.
func GetScopes(r *http.Request) ([]jwtv1.RoleScope, bool) {
	scopes, ok := r.Context().Value(ContextScopesKey).([]jwtv1.RoleScope)
	return scopes, ok
}
