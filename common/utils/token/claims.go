package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nrednav/cuid2"
	"github.com/yash91989201/superfast-delivery-api/common/types"
)

type AuthClaims struct {
	Email         *string        `json:"email"`
	EmailVerified bool           `json:"email_verified"`
	Phone         *string        `json:"phone"`
	Role          types.AuthRole `json:"role"`
	jwt.RegisteredClaims
}

func NewAuthClaims(
	email *string,
	emailVerfied bool,
	phone *string,
	role types.AuthRole,
	auth_id string,
	duration time.Duration,
) *AuthClaims {
	tokenId := cuid2.Generate()

	return &AuthClaims{
		Email: email,
		Phone: phone,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenId,
			Subject:   auth_id,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
}
