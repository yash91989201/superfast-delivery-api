package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
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

const (
	AuthCtxKey      authCtxKey = "auth"
	SessionIdCtxKey authCtxKey = "session_id"
)

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
			operationName, err := getOperationName(r)
			if err != nil {
				http.Error(w, "invalid request body", http.StatusBadRequest)
				return
			}

			if isPublicOperation(operationName) {
				next.ServeHTTP(w, r)
				return
			}

			ctx := r.Context()
			auth, sessionID, err := authenticate(ctx, r, authClient)
			if err != nil {
				writeGraphQLAuthError(w, err.Error(), operationName)
				return
			}

			ctx = context.WithValue(ctx, AuthCtxKey, auth)
			if sessionID != "" {
				ctx = context.WithValue(ctx, SessionIdCtxKey, sessionID)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getOperationName(r *http.Request) (string, error) {
	if r.Method == http.MethodGet {
		return r.URL.Query().Get("operationName"), nil
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	var req gqlRequest
	if err := json.Unmarshal(body, &req); err != nil {
		return "", err
	}
	return req.OperationName, nil
}

func authenticate(ctx context.Context, r *http.Request, authClient *clients.AuthenticationClient) (*types.Auth, string, error) {
	auth, sessionID, err := validateAccessToken(ctx, r, authClient)
	if err == nil {
		return auth, sessionID, nil
	}

	auth, err = refreshAccessToken(ctx, authClient)
	if err != nil {
		return nil, "", err
	}

	return auth, "", nil
}

func extractAccessToken(r *http.Request) (string, error) {
	if authHeader := r.Header.Get("Authorization"); authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && strings.EqualFold(parts[0], "bearer") {
			return parts[1], nil
		}
		return "", errors.New("invalid authorization header")
	}

	cookieManager, err := GetCookieManager(r.Context())
	if err != nil {
		return "", errors.New("missing cookie manager")
	}

	cookie, err := cookieManager.GetCookie("access_token")
	if err != nil || cookie.Value == "" {
		return "", errors.New("missing access token")
	}
	return cookie.Value, nil
}

func validateAccessToken(ctx context.Context, r *http.Request, authClient *clients.AuthenticationClient) (*types.Auth, string, error) {
	accessToken, err := extractAccessToken(r)
	if err != nil {
		return nil, "", err
	}

	res, err := authClient.ValidateSession(ctx, &pb.ValidateSessionReq{AccessToken: accessToken})
	if err != nil {
		return nil, "", errors.New("session validation failed")
	}
	return types.ToAuth(res.Auth), res.SessionId, nil
}

func refreshAccessToken(ctx context.Context, authClient *clients.AuthenticationClient) (*types.Auth, error) {
	cookieManager, err := GetCookieManager(ctx)
	if err != nil {
		return nil, errors.New("missing cookie manager")
	}

	refreshTokenCookie, err := cookieManager.GetCookie("refresh_token")
	if err != nil || refreshTokenCookie.Value == "" {
		return nil, errors.New("missing refresh token")
	}

	res, err := authClient.RefreshAccessToken(ctx, &pb.RefreshAccessTokenReq{RefreshToken: refreshTokenCookie.Value})
	if err != nil {
		return nil, errors.New("failed to refresh access token")
	}

	setAuthCookies(*cookieManager, res.AccessToken, res.RefreshToken)
	return types.ToAuth(res.Auth), nil
}

func setAuthCookies(cookieManager utils.CookieManager, accessToken, refreshToken string) {
	cookieManager.SetCookie("access_token", accessToken, utils.CookieOptions{
		Path:     "/",
		MaxAge:   int((15 * time.Minute).Seconds()),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	cookieManager.SetCookie("refresh_token", refreshToken, utils.CookieOptions{
		Path:     "/",
		MaxAge:   int((30 * 24 * time.Hour).Seconds()),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
}

func isPublicOperation(operationName string) bool {
	return publicOperations[operationName]
}

func writeGraphQLAuthError(w http.ResponseWriter, message, operationName string) {
	errResponse := &gqlerror.Error{
		Message:    message,
		Extensions: map[string]any{"operationName": operationName},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	if err := json.NewEncoder(w).Encode(map[string]any{"errors": []*gqlerror.Error{errResponse}}); err != nil {
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
