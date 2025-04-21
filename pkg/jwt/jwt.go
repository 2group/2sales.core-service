package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID         *int64   `json:"user_id,omitempty"`
	EmployeeID     *int64   `json:"employee_id,omitempty"`
	CustomerID     *int64   `json:"customer_id,omitempty"`
	OrganizationID *int64   `json:"organization_id,omitempty"`
	RoleIDs        *[]int64 `json:"role_ids,omitempty"`
	BranchIDs      *[]int64 `json:"branch_ids,omitempty"`

	jwt.RegisteredClaims
}

var Secret = []byte("yourSigningKey")

func NewToken(
	userID, employeeID, customerID, organizationID *int64,
	roleIDs, branchIDs *[]int64,
) (string, error) {
	claims := Claims{
		UserID:         userID,
		EmployeeID:     employeeID,
		CustomerID:     customerID,
		OrganizationID: organizationID,
		RoleIDs:        roleIDs,
		BranchIDs:      branchIDs,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret)
}
