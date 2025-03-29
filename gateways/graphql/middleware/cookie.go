package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/yash91989201/superfast-delivery-api/common/utils"
)

type cookieManagerCtx struct{}

var cookieManagerCtxKey = cookieManagerCtx{}

func WithCookieManager(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieManager := utils.NewCookieManager(w, r)

		ctx := context.WithValue(r.Context(), cookieManagerCtxKey, cookieManager)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetCookieManager(ctx context.Context) (*utils.CookieManager, error) {
	cookieManager, ok := ctx.Value(cookieManagerCtxKey).(*utils.CookieManager)
	if !ok {
		return nil, fmt.Errorf("cookie manager not available in context")
	}

	return cookieManager, nil
}
