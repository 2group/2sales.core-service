package middleware

import (
	"net/http"

	jwtv1 "github.com/2group/2sales.core-service/pkg/jwt"
)

const (
	RoleOrgAdmin    int64 = 34
	RoleBranchAdmin int64 = 35
)

// ScopeChecker holds the role→scope data for the current request.
type ScopeChecker struct {
	Scopes []jwtv1.RoleScope
}

// NewScopeChecker extracts RoleScope slice from request context.
func NewScopeChecker(r *http.Request) *ScopeChecker {
	scopes, _ := GetScopes(r)
	return &ScopeChecker{Scopes: scopes}
}

// HasOrgAdmin returns true if user has ORG_ADMIN on given organization.
func (c *ScopeChecker) HasOrgAdmin(orgID int64) bool {
	for _, s := range c.Scopes {
		if s.RoleID == RoleOrgAdmin && s.OrganizationID != nil && *s.OrganizationID == orgID {
			return true
		}
	}
	return false
}

// HasBranchAdmin returns true if user has BRANCH_ADMIN on given branch.
func (c *ScopeChecker) HasBranchAdmin(branchID int64) bool {
	for _, s := range c.Scopes {
		if s.RoleID == RoleBranchAdmin && s.BranchID != nil && *s.BranchID == branchID {
			return true
		}
	}
	return false
}
