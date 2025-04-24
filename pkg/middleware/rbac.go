package middleware

import (
	"net/http"

	jwtv1 "github.com/2group/2sales.core-service/pkg/jwt"
)

const (
	RoleOrgAdmin    int64 = 34
	RoleBranchAdmin int64 = 35
	RoleSuperAdmin  int64 = 36
)

// ScopeChecker holds the role→scope data for the current request.
type RoleChecker struct {
	Roles []jwtv1.RoleScope
}

// NewScopeChecker extracts RoleScope slice from request context.
func NewRoleChecker(r *http.Request) *RoleChecker {
	scopes, _ := GetScopes(r)
	return &RoleChecker{Roles: scopes}
}

// HasOrgAdmin returns true if user has ORG_ADMIN on given organization.
func (c *RoleChecker) HasOrgAdmin(orgID int64) bool {
	for _, s := range c.Roles {
		if s.RoleID == RoleOrgAdmin && s.OrganizationID != nil && *s.OrganizationID == orgID {
			return true
		}
	}
	return false
}

func (c *RoleChecker) HasBranchAdmin(branchID int64) bool {
	for _, s := range c.Roles {
		if s.RoleID == RoleBranchAdmin && s.BranchID != nil && *s.BranchID == branchID {
			return true
		}
	}
	return false
}

func (c *RoleChecker) HasSuperAdmin() bool {
	for _, s := range c.Roles {
		if s.RoleID == RoleSuperAdmin {
			return true
		}
	}
	return false
}
