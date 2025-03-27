package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func HasAuthRoleDirective(ctx context.Context, obj any, next graphql.Resolver, authRole AuthRole) (any, error) {
	return nil, nil
}
