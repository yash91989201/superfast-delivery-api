package rest

import (
	"context"
)

type handler struct {
	ctx context.Context
}

func NewHandler(ctx context.Context) *handler {
	return &handler{
		ctx: ctx,
	}
}
