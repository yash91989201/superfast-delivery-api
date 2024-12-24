package graphql

import (
	"context"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
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
			Auth:          nil,
			Profile:       nil,
			CreateProfile: false,
			VerifyOtp:     true,
		}, nil
	}

	profile, err := r.server.userClient.GetProfile(ctx, &pb.GetProfileReq{AuthId: auth.Id})
	if err != nil {
		return &SignInOutput{
			Auth:          ToAuth(auth),
			Profile:       nil,
			CreateProfile: true,
			VerifyOtp:     false,
		}, nil
	}

	return &SignInOutput{
		Auth:          ToAuth(auth),
		Profile:       ToProfile(profile),
		CreateProfile: false,
		VerifyOtp:     false,
	}, nil
}

func (r *mutationResolver) SignInWithPhone(ctx context.Context, in SignInWithPhoneInput) (*SignInOutput, error) {
	auth, err := r.server.authenticationClient.SignInWithPhone(ctx, &pb.SignInWithPhoneReq{Phone: in.Phone, Otp: in.Otp})
	if err != nil {
		return nil, err
	}

	if auth.Id == "" {
		return &SignInOutput{
			Auth:          nil,
			Profile:       nil,
			CreateProfile: false,
			VerifyOtp:     true,
		}, nil
	}

	profile, err := r.server.userClient.GetProfile(ctx, &pb.GetProfileReq{AuthId: auth.Id})
	if err != nil {
		return &SignInOutput{
			Auth:          ToAuth(auth),
			Profile:       nil,
			CreateProfile: true,
			VerifyOtp:     false,
		}, nil
	}

	return &SignInOutput{
		Auth:          ToAuth(auth),
		Profile:       ToProfile(profile),
		CreateProfile: false,
		VerifyOtp:     false,
	}, nil
}

func (r *mutationResolver) SignInWithGoogle(ctx context.Context, in SignInWithGoogleInput) (*SignInOutput, error) {
	// TODO: different flow required for google because profile will be created in the user service by fetching from google
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
		VerifyOtp: false,
	}, nil
}

func (r *mutationResolver) CreateProfile(ctx context.Context, in CreateProfileInput) (*Profile, error) {
	res, err := r.server.userClient.CreateProfile(ctx, &pb.CreateProfileReq{
		Name:        in.Name,
		ImageUrl:    in.ImageURL,
		Dob:         ToPbDate(in.Dob),
		Anniversary: ToPbDate(in.Anniversary),
		Gender:      ToPbGenderPtr(in.Gender),
		AuthId:      in.AuthID,
	})

	if err != nil {
		return nil, err
	}

	return ToProfile(res), nil
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, in UpdateProfileInput) (*Profile, error) {
	res, err := r.server.userClient.UpdateProfile(ctx, &pb.UpdateProfileReq{
		Name:        in.Name,
		ImageUrl:    in.ImageURL,
		Dob:         ToPbDate(in.Dob),
		Anniversary: ToPbDate(in.Anniversary),
		Gender:      ToPbGenderPtr(in.Gender),
		AuthId:      in.AuthID,
	})

	if err != nil {
		return nil, err
	}

	return ToProfile(res), nil
}
