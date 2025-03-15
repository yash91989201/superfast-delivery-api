package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yash91989201/superfast-delivery-api/common/types"
)

type AuthClaims struct {
	Email         *string        `json:"email"`
	EmailVerified bool           `json:"email_verified"`
	Phone         *string        `json:"phone"`
	Role          types.AuthRole `json:"role"`
	jwt.RegisteredClaims
}

type NewAuthClaim struct {
	Email         *string
	EmailVerified bool
	Phone         *string
	Role          types.AuthRole
	AuthId        string
	SessionId     string
	Duration      time.Duration
}

func NewAuthClaims(newAuthClaim NewAuthClaim) *AuthClaims {

	return &AuthClaims{
		Email: newAuthClaim.Email,
		Phone: newAuthClaim.Phone,
		Role:  newAuthClaim.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        newAuthClaim.SessionId,
			Subject:   newAuthClaim.AuthId,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(newAuthClaim.Duration)),
		},
	}
}
