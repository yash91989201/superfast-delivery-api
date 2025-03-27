package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/yash91989201/superfast-delivery-api/common/clients"
	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"github.com/yash91989201/superfast-delivery-api/common/types"
)

type contextKey string

const AuthCtxKey contextKey = "auth"
const SessionIdCtxKey contextKey = "session_id"

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

func AuthenticationMiddleware(authClient *clients.AuthenticationClient) func(http.Handler) http.Handler {
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

			accessToken, err := extractToken(r)
			if err != nil {
				writeGraphQLAuthError(w, "UNAUTHORIZED", operationName)
				return
			}

			res, err := authClient.ValidateSession(r.Context(), &pb.ValidateSessionReq{
				AccessToken: accessToken,
			})

			if err != nil {
				writeGraphQLAuthError(w, "UNAUTHORIZED", operationName)
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

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var req gqlRequest
	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		return "", err
	}

	return req.OperationName, nil
}

func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header required")
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return "", errors.New("invalid authorization header")
	}

	return tokenParts[1], nil
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
