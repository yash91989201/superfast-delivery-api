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

func (r *queryResolver) Profile(ctx context.Context, in GetProfileInput) (*Profile, error) {
	return nil, nil
}

func (r *queryResolver) GetShop(ctx context.Context, id string) (*Shop, error) {
	shop, err := r.server.shopClient.GetShop(ctx, &pb.GetShopReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQShop(shop), nil
}

func (r *queryResolver) ListShops(ctx context.Context, in *ListShopsInput) (*ListShopsOutput, error) {
	res, err := r.server.shopClient.ListShops(ctx, ToPbListShopReq(in))
	if err != nil {
		return nil, err
	}

	return &ListShopsOutput{
		Shops: ToGQShops(res.Shops),
		Total: int32(len(res.Shops)),
	}, nil
}
