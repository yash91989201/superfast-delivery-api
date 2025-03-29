package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/yash91989201/superfast-delivery-api/common/clients"
	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"github.com/yash91989201/superfast-delivery-api/common/types"
	"github.com/yash91989201/superfast-delivery-api/common/utils"
)

type authCtxKey string

const AuthCtxKey authCtxKey = "auth"
const SessionIdCtxKey authCtxKey = "session_id"

type gqlRequest struct {
	Query         string         `json:"query"`
	OperationName string         `json:"operationName"`
	Variables     map[string]any `json:"variables"`
}

var publicOperations = map[string]bool{
	"SignInWithEmail":    true,
	"SignInWithPhone":    true,
	"SignInWithGoogle":   true,
	"RefreshAccessToken": true,
	"IntrospectionQuery": true,
}

func Authentication(authClient *clients.AuthenticationClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.URL.Path == "/playground" || (r.URL.Path == "/graphql" && r.Method == http.MethodGet) {
				next.ServeHTTP(w, r)
				return
			}

			operationName, err := extractOperationName(r)
			if err != nil {
				http.Error(w, "invalid request body", http.StatusBadRequest)
				return
			}

			if isPublicOperation(operationName) {
				next.ServeHTTP(w, r)
				return
			}

			accessToken, err := extractAccessToken(r)
			if err != nil {
				log.Print(err)
				writeGraphQLAuthError(w, "UNAUTHENTICATED: Missing or invalid access token", operationName)
				return
			}

			res, err := authClient.ValidateSession(r.Context(), &pb.ValidateSessionReq{
				AccessToken: accessToken,
			})

			if err != nil {
				auth, err := tryRefreshingAccessToken(r.Context(), authClient)
				if err != nil {
					writeGraphQLAuthError(w, "UNAUTHENTICATED: Failed to refresh access token", operationName)
					return
				}

				ctx := context.WithValue(r.Context(), AuthCtxKey, auth)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			ctx := context.WithValue(r.Context(), AuthCtxKey, types.ToAuth(res.Auth))
			ctx = context.WithValue(ctx, SessionIdCtxKey, res.SessionId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func extractOperationName(r *http.Request) (string, error) {
	if r.Method == http.MethodGet {
		return r.URL.Query().Get("operationName"), nil
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	r.Body.Close()

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var req gqlRequest
	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		return "", err
	}

	return req.OperationName, nil
}

func extractAccessToken(r *http.Request) (string, error) {
	if authHeader := r.Header.Get("Authorization"); authHeader != "" {
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			return "", fmt.Errorf("invalid authorization header format")
		}

		return tokenParts[1], nil
	}

	cookieManager, err := GetCookieManager(r.Context())
	if err != nil {
		return "", err
	}

	accessTokenCookie, err := cookieManager.GetCookie("access_token")
	if err != nil {
		return "", err
	}

	if accessTokenCookie.Value == "" {
		return "", fmt.Errorf("access_token cookie is empty")
	}

	return accessTokenCookie.Value, nil
}

func tryRefreshingAccessToken(ctx context.Context, authClient *clients.AuthenticationClient) (*types.Auth, error) {
	cookieManager, err := GetCookieManager(ctx)
	if err != nil {
		return nil, err
	}

	refreshTokenCookie, err := cookieManager.GetCookie("refresh_token")
	if err != nil {
		return nil, err
	}

	refreshAccessTokenRes, err := authClient.RefreshAccessToken(ctx, &pb.RefreshAccessTokenReq{
		RefreshToken: refreshTokenCookie.Value,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to refresh access token")
	}

	newAccessToken := refreshAccessTokenRes.AccessToken
	newRefreshToken := refreshAccessTokenRes.RefreshToken

	cookieManager.SetCookie(
		"access_token",
		newAccessToken,
		utils.CookieOptions{
			Path:     "/",
			MaxAge:   int((15 * time.Minute).Seconds()),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		},
	)

	cookieManager.SetCookie(
		"refresh_token",
		newRefreshToken,
		utils.CookieOptions{
			Path:     "/",
			MaxAge:   int((30 * 24 * time.Hour).Seconds()),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		},
	)

	return types.ToAuth(refreshAccessTokenRes.Auth), nil
}

func isPublicOperation(operationName string) bool {
	return publicOperations[operationName]
}

func writeGraphQLAuthError(w http.ResponseWriter, message, operationName string) {
	errResponse := &gqlerror.Error{
		Message: message,
		Extensions: map[string]any{
			"operationName": operationName,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	if err := json.NewEncoder(w).Encode(map[string]any{
		"errors": []*gqlerror.Error{errResponse},
	}); err != nil {
		http.Error(w, "failed to encode error response", http.StatusInternalServerError)
	}
}

func GetCtxAuth(ctx context.Context) (*types.Auth, error) {
	auth, ok := ctx.Value(AuthCtxKey).(*types.Auth)
	if !ok {
		return nil, errors.New("unauthenticated request")
	}

	return auth, nil
}

func GetCtxSessionId(ctx context.Context) (string, error) {
	sessionID, ok := ctx.Value(SessionIdCtxKey).(string)
	if !ok {
		return "", errors.New("unauthenticated request")
	}

	return sessionID, nil
}
