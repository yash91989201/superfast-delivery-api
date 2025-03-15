package graphql

import (
	"context"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/yash91989201/superfast-delivery-api/common/pb"
)

type contextKey string

const (
	authContextKey = contextKey("auth")
)

var publicOperations = map[string]bool{
	"SignInWithEmail":  true,
	"SignInWithPhone":  true,
	"SignInWithGoogle": true,
	"RefreshToken":     true,

	"__schema":           true,
	"__type":             true,
	"IntrospectionQuery": true,
}

func (s *Server) AuthenticationMiddleware(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	oc := graphql.GetOperationContext(ctx)
	if oc == nil {
		return func(ctx context.Context) *graphql.Response {
			return graphql.ErrorResponse(ctx, "invalid request context")
		}
	}

	operationName := oc.Operation.Name
	if _, isPublic := publicOperations[operationName]; isPublic {
		return next(ctx)
	}

	authHeader := oc.Headers.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return func(ctx context.Context) *graphql.Response {
			return graphql.ErrorResponse(ctx, "unauthenticated: missing or invalid token")
		}
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return func(ctx context.Context) *graphql.Response {
			return graphql.ErrorResponse(ctx, "unauthenticated: missing token")
		}
	}

	validateSessionRes, err := s.authenticationClient.ValidateSession(ctx, &pb.ValidateSessionReq{AuthToken: token})
	if err != nil || !validateSessionRes.Valid {
		return func(ctx context.Context) *graphql.Response {
			return graphql.ErrorResponse(ctx, "unauthorized: Invalid session, please login again")
		}
	}

	ctx = context.WithValue(ctx, authContextKey, validateSessionRes.Auth)

	return next(ctx)
}
