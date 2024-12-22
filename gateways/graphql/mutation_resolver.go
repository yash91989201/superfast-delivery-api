package graphql

import (
	"context"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"github.com/yash91989201/superfast-delivery-api/common/types"
)

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) SignInWithEmail(ctx context.Context, in SignInWithEmailInput) (*SignInOutput, error) {
	auth, err := r.server.authenticationClient.SignInWithEmail(ctx, &pb.SignInWithEmailReq{Email: in.Email, Otp: in.Otp})
	if err != nil {
		return nil, err
	}

	if auth.Id == "" {
		return &SignInOutput{
			Auth:      nil,
			Profile:   nil,
			VerifyOtp: types.ToBoolPtr(true),
		}, nil
	}

	profile, err := r.server.userClient.GetProfile(ctx, &pb.GetProfileReq{AuthId: auth.Id})
	if err != nil {
		return &SignInOutput{
			Auth:      ToAuth(auth),
			Profile:   nil,
			VerifyOtp: types.ToBoolPtr(false),
		}, nil
	}

	return &SignInOutput{
		Auth:      ToAuth(auth),
		Profile:   ToProfile(profile),
		VerifyOtp: types.ToBoolPtr(false),
	}, nil
}

func (r *mutationResolver) SignInWithPhone(ctx context.Context, in SignInWithPhoneInput) (*SignInOutput, error) {
	auth, err := r.server.authenticationClient.SignInWithPhone(ctx, &pb.SignInWithPhoneReq{Phone: in.Phone, Otp: in.Otp})
	if err != nil {
		return nil, err
	}

	if auth.Id == "" {
		return &SignInOutput{
			Auth:      nil,
			Profile:   nil,
			VerifyOtp: types.ToBoolPtr(true),
		}, nil
	}

	profile, err := r.server.userClient.GetProfile(ctx, &pb.GetProfileReq{AuthId: auth.Id})
	if err != nil {
		return nil, err
	}

	return &SignInOutput{
		Auth:      ToAuth(auth),
		Profile:   ToProfile(profile),
		VerifyOtp: types.ToBoolPtr(false),
	}, nil
}

func (r *mutationResolver) SignInWithGoogle(ctx context.Context, in SignInWithGoogleInput) (*SignInOutput, error) {
	auth, err := r.server.authenticationClient.SignInWithGoogle(ctx, &pb.SignInWithGoogleReq{IdToken: in.IDToken})
	if err != nil {
		return nil, err
	}

	// get profile
	profile, err := r.server.userClient.GetProfile(ctx, &pb.GetProfileReq{AuthId: auth.Id})
	if err != nil {
		return nil, err
	}

	return &SignInOutput{
		Auth:      ToAuth(auth),
		Profile:   ToProfile(profile),
		VerifyOtp: types.ToBoolPtr(false),
	}, nil
}
