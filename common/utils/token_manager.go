package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yash91989201/superfast-delivery-api/common/types"
)

type TokenManager struct {
	secretKey string
}

type AuthClaims struct {
	Email         *string        `json:"email"`
	EmailVerified bool           `json:"email_verified"`
	Phone         *string        `json:"phone"`
	Role          types.AuthRole `json:"role"`
	jwt.RegisteredClaims
}

func NewTokenManager(secretKey string) *TokenManager {
	return &TokenManager{secretKey}
}

func (manager *TokenManager) GenerateAccessToken(auth *types.Auth, sessionID string) (string, error) {
	authClaims := &AuthClaims{
		Email: auth.Email,
		Phone: auth.Phone,
		Role:  auth.AuthRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        sessionID,
			Subject:   auth.ID,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)
	accessTokenStr, err := accessToken.SignedString([]byte(manager.secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to create access token")
	}

	return accessTokenStr, nil
}

func (manager *TokenManager) GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	rawToken := base64.URLEncoding.EncodeToString(b)

	return rawToken, nil
}

func (manager *TokenManager) VerifyAccessToken(accessToken string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid token signing method")
		}

		return []byte(manager.secretKey), nil
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
