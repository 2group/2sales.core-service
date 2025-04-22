package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RoleScope struct {
	RoleID         int64
	OrganizationID *int64
	BranchID       *int64
}

var Secret = "2group.kz"

func NewToken(
	userID int64,
	employeeID, customerID *int64,
	roleScopes []RoleScope,
) (string, error) {
	rawScopes := make([]map[string]interface{}, len(roleScopes))
	for i, rs := range roleScopes {
		entry := map[string]interface{}{
			"role_id": rs.RoleID,
		}
		if rs.BranchID != nil {
			entry["branch_id"] = *rs.BranchID
		} else if rs.OrganizationID != nil {
			entry["organization_id"] = *rs.OrganizationID
		}
		rawScopes[i] = entry
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"scopes":  rawScopes,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 24 * 365).Unix(),
	}

	if employeeID != nil {
		claims["employee_id"] = *employeeID
	} else if customerID != nil {
		claims["customer_id"] = *customerID
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret)
}
