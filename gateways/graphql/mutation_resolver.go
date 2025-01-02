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
	signInRes, err := r.server.authenticationClient.SignInWithEmail(ctx, &pb.SignInWithEmailReq{Email: in.Email, Otp: in.Otp})
	if err != nil {
		return nil, err
	}

	if signInRes.Auth.Id == "" {
		return &SignInOutput{
			VerifyOtp: true,
		}, nil
	}

	profile, _ := r.server.userClient.GetProfile(ctx, &pb.GetProfileReq{AuthId: signInRes.Auth.Id})

	return &SignInOutput{
		Auth:                 ToAuth(signInRes.Auth),
		SessionID:            &signInRes.SessionId,
		AccessToken:          &signInRes.AccessToken,
		AccessTokenExpiresAt: types.PbTimeStampToStrPtr(signInRes.AccessTokenExpiresAt),
		Profile:              ToProfile(profile),
		CreateProfile:        profile == nil,
		VerifyOtp:            false,
	}, nil
}

func (r *mutationResolver) SignInWithPhone(ctx context.Context, in SignInWithPhoneInput) (*SignInOutput, error) {
	return nil, nil
}

func (r *mutationResolver) SignInWithGoogle(ctx context.Context, in SignInWithGoogleInput) (*SignInOutput, error) {
	return nil, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, session_id string) (*SignInOutput, error) {
	signInRes, err := r.server.authenticationClient.RefreshToken(ctx, &pb.RefreshTokenReq{SessionId: session_id})
	if err != nil {
		return nil, err
	}

	profile, _ := r.server.userClient.GetProfile(ctx, &pb.GetProfileReq{AuthId: signInRes.Auth.Id})

	return &SignInOutput{
		Auth:                 ToAuth(signInRes.Auth),
		SessionID:            &signInRes.SessionId,
		AccessToken:          &signInRes.AccessToken,
		AccessTokenExpiresAt: types.PbTimeStampToStrPtr(signInRes.AccessTokenExpiresAt),
		Profile:              ToProfile(profile),
		CreateProfile:        profile == nil,
		VerifyOtp:            false,
	}, nil
}

func (r *mutationResolver) LogOut(ctx context.Context, session_id string) (*SignInOutput, error) {
	_, err := r.server.authenticationClient.LogOut(ctx, &pb.LogOutReq{SessionId: session_id})
	if err != nil {
		return nil, err
	}

	return &SignInOutput{}, nil
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

func (r *mutationResolver) CreateShop(ctx context.Context, in CreateShopInput) (*CreateShopOutput, error) {
	res, err := r.server.shopClient.CreateShop(ctx, ToPbCreateShopReq(in))
	if err != nil {
		return nil, err
	}

	return ToGQCreateShopOutput(res), nil
}

func (r *mutationResolver) UpdateShop(ctx context.Context, in UpdateShopInput) (*UpdateShopOutput, error) {
	return nil, nil
}

func (r *mutationResolver) UpdateShopAddress(ctx context.Context, in UpdateShopAddressInput) (*UpdateShopOutput, error) {
	return nil, nil
}

func (r *mutationResolver) UpdateShopContact(ctx context.Context, in UpdateShopContactInput) (*UpdateShopOutput, error) {
	return nil, nil
}

func (r *mutationResolver) UpdateShopImages(ctx context.Context, in []*UpdateShopImageInput) (*UpdateShopOutput, error) {
	return nil, nil
}

func (r *mutationResolver) UpdateShopTimings(ctx context.Context, in []*UpdateShopTimingInput) (*UpdateShopOutput, error) {
	return nil, nil
}

func (r *mutationResolver) DeleteShop(ctx context.Context, in string) (*UpdateShopOutput, error) {
	return nil, nil
}

func (r *mutationResolver) CreateItemVariant(ctx context.Context, in CreateItemVariantReq) (*ItemVariant, error) {
	res, err := r.server.productClient.CreateItemVariant(ctx, ToPbCreateItemVariantReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQItemVariant(res), nil
}

func (r *mutationResolver) CreateItemAddon(ctx context.Context, in CreateItemAddonReq) (*ItemAddon, error) {
	res, err := r.server.productClient.CreateItemAddon(ctx, ToPbCreateItemAddonReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQItemAddon(res), nil
}

func (r *mutationResolver) CreateRestaurantMenu(ctx context.Context, in CreateRestaurantMenuReq) (*RestaurantMenu, error) {
	res, err := r.server.productClient.CreateRestaurantMenu(ctx, ToPbCreateRestaurantMenuReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQRestaurantMenu(res), nil
}

func (r *mutationResolver) CreateMenuItem(ctx context.Context, in CreateMenuItemReq) (*MenuItem, error) {
	res, err := r.server.productClient.CreateMenuItem(ctx, ToPbCreateMenuItemReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQMenuItem(res), nil
}

func (r *mutationResolver) CreateRetailCategory(ctx context.Context, in CreateRetailCategoryReq) (*RetailCategory, error) {
	return nil, nil
}

func (r *mutationResolver) CreateRetailItem(ctx context.Context, in CreateRetailItemReq) (*RetailItem, error) {
	return nil, nil
}

func (r *mutationResolver) CreateMedicineCategory(ctx context.Context, in CreateMedicineCategoryReq) (*MedicineCategory, error) {
	return nil, nil
}

func (r *mutationResolver) CreateMedicineItem(ctx context.Context, in CreateMedicineItemReq) (*MedicineItem, error) {
	return nil, nil
}
