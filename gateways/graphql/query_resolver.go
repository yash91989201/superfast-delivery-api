package graphql

import (
	"context"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
)

type queryResolver struct {
	server *Server
}

func (r *queryResolver) AuthByID(ctx context.Context, in GetAuthByIDInput) (*Auth, error) {
	auth, err := r.server.authenticationClient.GetAuthById(ctx, &pb.GetAuthByIdReq{Id: in.ID})
	if err != nil {
		return nil, err
	}

	return ToAuth(auth), nil
}

func (r *queryResolver) Auth(ctx context.Context, in GetAuthInput) (*Auth, error) {
	auth, err := r.server.authenticationClient.GetAuth(ctx, &pb.GetAuthReq{Email: in.Email, Phone: in.Phone})
	if err != nil {
		return nil, err
	}

	return ToAuth(auth), nil
}
