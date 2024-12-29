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
	authId string,
	duration time.Duration,
) *AuthClaims {
	sessionId := cuid2.Generate()

	return &AuthClaims{
		Email: email,
		Phone: phone,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        sessionId,
			Subject:   authId,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
}
