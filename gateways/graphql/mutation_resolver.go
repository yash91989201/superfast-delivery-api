package graphql

import (
	"context"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
)

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) SignInWithEmail(ctx context.Context, in SignInWithEmailInput) (*User, error) {
	auth, err := r.server.authenticationClient.SignInWithEmail(ctx, &pb.SignInWithEmailReq{Email: in.Email, Otp: in.Otp})
	if err != nil {
		return nil, err
	}

	profile, err := r.server.userClient.GetProfile(ctx, &pb.GetProfileReq{AuthId: auth.Id})
	if err != nil {
		return nil, err
	}

	return &User{
		Auth:    ToAuth(auth),
		Profile: ToProfile(profile),
	}, nil

}

func (r *mutationResolver) SignInWithPhone(ctx context.Context, in SignInWithPhoneInput) (*User, error) {
	return nil, nil
}

func (r *mutationResolver) SignInWithGoogle(ctx context.Context, in SignInWithGoogleInput) (*User, error) {
	auth, err := r.server.authenticationClient.SignInWithGoogle(ctx, &pb.SignInWithGoogleReq{IdToken: in.IDToken})
	if err != nil {
		return nil, err
	}

	// get profile
	profile, err := r.server.userClient.GetProfile(ctx, &pb.GetProfileReq{AuthId: auth.Id})
	if err != nil {
		return nil, err
	}

	return &User{
		Auth:    ToAuth(auth),
		Profile: ToProfile(profile),
	}, nil
}
