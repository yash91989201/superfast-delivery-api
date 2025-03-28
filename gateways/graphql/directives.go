package graphql

import (
	"context"
	"fmt"
	"slices"

	"github.com/99designs/gqlgen/graphql"
	"github.com/casbin/casbin/v2"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/yash91989201/superfast-delivery-api/gateways/graphql/middleware"
)

type GQLDirective struct {
	casbinEnforcer *casbin.Enforcer
}

func NewGQLDirective(casbinEnforcer *casbin.Enforcer) *GQLDirective {
	return &GQLDirective{casbinEnforcer: casbinEnforcer}
}

func (d *GQLDirective) RequireAuthRole(ctx context.Context, obj any, next graphql.Resolver, roles []AuthRole) (any, error) {
	auth, err := middleware.GetCtxAuth(ctx)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: "Authentication required",
			Extensions: map[string]any{
				"code":    "UNAUTHENTICATED",
				"details": "No valid authentication token provided",
			},
		}
	}

	gqlCtx := graphql.GetOperationContext(ctx)
	if gqlCtx == nil {
		return nil, &gqlerror.Error{
			Message: "Internal server error",
			Extensions: map[string]any{
				"code": "INTERNAL_ERROR",
			},
		}
	}

	operationType := gqlCtx.Operation.Operation
	operationName := gqlCtx.Operation.Name
	casbinObj := fmt.Sprintf("%s/%s", operationType, operationName)
	authRole := TypesToGQAuthRole(auth.AuthRole)
	stringAuthRole := string(authRole)

	if !slices.Contains(roles, authRole) {
		return nil, &gqlerror.Error{
			Message: "Insufficient permissions",
			Extensions: map[string]any{
				"code":          "FORBIDDEN",
				"requiredRoles": roles,
				"userRole":      stringAuthRole,
				"operation":     casbinObj,
			},
		}
	}

	allowed, err := d.casbinEnforcer.Enforce(stringAuthRole, casbinObj)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: "Authorization check failed",
			Extensions: map[string]any{
				"code":  "INTERNAL_ERROR",
				"error": err.Error(),
			},
		}
	}

	if !allowed {
		return nil, &gqlerror.Error{
			Message: "Operation not permitted",
			Extensions: map[string]any{
				"code":      "FORBIDDEN",
				"userRole":  stringAuthRole,
				"operation": casbinObj,
			},
		}
	}

	return next(ctx)
}
