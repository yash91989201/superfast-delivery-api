package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yash91989201/superfast-delivery-api/common/types"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) *JWTMaker {
	return &JWTMaker{secretKey}
}

func (maker *JWTMaker) CreateToken(email *string, email_verified bool, phone *string, role types.AuthRole, auth_id string, duration time.Duration) (string, *AuthClaims, error) {
	claims := NewAuthClaims(email, email_verified, phone, role, auth_id, duration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(maker.secretKey))

	if err != nil {
		return "", nil, fmt.Errorf("Error creating token :%w", err)
	}

	return tokenStr, claims, nil
}

func (maker *JWTMaker) VerifyToken(tokenStr string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid token signing method")
		}

		return []byte(maker.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("Error parsing token :%w", err)
	}

	claims, ok := token.Claims.(*AuthClaims)
	if !ok {
		return nil, fmt.Errorf("Invalid token claims")
	}

	return claims, nil
}
