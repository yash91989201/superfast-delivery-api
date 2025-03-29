package graphql

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"github.com/yash91989201/superfast-delivery-api/common/utils"
	customMiddleware "github.com/yash91989201/superfast-delivery-api/gateways/graphql/middleware"
)

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) SignInWithEmail(ctx context.Context, in SignInWithEmailInput) (*SignInOutput, error) {
	signInRes, err := r.server.AuthenticationClient.SignInWithEmail(
		ctx,
		&pb.SignInWithEmailReq{
			Email:    in.Email,
			AuthRole: ToPbAuthRole(in.AuthRole),
			Otp:      in.Otp,
		},
	)
	if err != nil {
		return nil, err
	}

	if signInRes.Auth == nil {
		return &SignInOutput{VerifyOtp: true}, nil
	}

	profile, err := r.server.UserClient.GetProfile(ctx, &pb.GetProfileReq{AuthId: signInRes.Auth.Id})
	res := &SignInOutput{
		Auth:      ToGQAuth(signInRes.Auth),
		Session:   ToGQSession(signInRes.Session),
		VerifyOtp: false,
	}

	if err != nil {
		res.CreateProfile = true
		return res, nil
	}

	res.Profile = ToGQProfile(profile)

	if err = setSessionCookies(ctx, signInRes.Session.AccessToken, signInRes.Session.RefreshToken); err != nil {
		return &SignInOutput{}, err
	}

	return res, nil
}

func (r *mutationResolver) SignInWithPhone(ctx context.Context, in SignInWithPhoneInput) (*SignInOutput, error) {
	return nil, nil
}

func (r *mutationResolver) SignInWithGoogle(ctx context.Context, in SignInWithGoogleInput) (*SignInOutput, error) {
	return nil, nil
}

func (r *mutationResolver) RefreshAccessToken(ctx context.Context, refreshToken string) (*SignInOutput, error) {
	signInRes, err := r.server.AuthenticationClient.RefreshAccessToken(ctx, &pb.RefreshAccessTokenReq{RefreshToken: refreshToken})
	if err != nil {
		return nil, err
	}

	profile, _ := r.server.UserClient.GetProfile(ctx, &pb.GetProfileReq{AuthId: signInRes.Auth.Id})

	if err = setSessionCookies(ctx, signInRes.Session.AccessToken, signInRes.Session.RefreshToken); err != nil {
		return &SignInOutput{}, err
	}

	return &SignInOutput{
		Auth:          ToGQAuth(signInRes.Auth),
		Profile:       ToGQProfile(profile),
		Session:       ToGQSession(signInRes.Session),
		CreateProfile: profile == nil,
		VerifyOtp:     false,
	}, nil
}

func (r *mutationResolver) LogOut(ctx context.Context) (*LogOutOutput, error) {
	sessionID, err := customMiddleware.GetCtxSessionId(ctx)
	if err != nil {
		return nil, err
	}

	_, err = r.server.AuthenticationClient.LogOut(ctx, &pb.LogOutReq{SessionId: sessionID})
	if err != nil {
		return nil, err
	}

	deleteSessionCookies(ctx)

	return &LogOutOutput{
		Success: true,
		Message: "Successfully logged out",
	}, nil
}

func (r *mutationResolver) CreateProfile(ctx context.Context, in CreateProfileInput) (*Profile, error) {
	res, err := r.server.UserClient.CreateProfile(ctx, &pb.CreateProfileReq{
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

	return ToGQProfile(res), nil
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, in UpdateProfileInput) (*Profile, error) {
	res, err := r.server.UserClient.UpdateProfile(ctx, &pb.UpdateProfileReq{
		Id:          in.ID,
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

	return ToGQProfile(res), nil
}

func (r *mutationResolver) CreateDeliveryAddress(ctx context.Context, in CreateDeliveryAddressInput) (*DeliveryAddress, error) {
	newDeliveryAddress, err := r.server.UserClient.CreateDeliveryAddress(ctx, ToPbCreateDeliveryAddress(in))
	if err != nil {
		return nil, err
	}

	_, err = r.server.GeolocationClient.ReverseGeocode(ctx, &pb.ReverseGeocodeReq{
		Latitude:  in.Latitude,
		Longitude: in.Longitude,
		AddressId: newDeliveryAddress.Id,
	})

	if err != nil {
		return nil, err
	}

	return ToGQDeliveryAddress(newDeliveryAddress), nil
}

func (r *mutationResolver) UpdateDeliveryAddress(ctx context.Context, in UpdateDeliveryAddressInput) (*DeliveryAddress, error) {
	return nil, nil
}

func (r *mutationResolver) UpdateDefaultDeliveryAddress(ctx context.Context, in UpdateDefaultDeliveryAddressInput) (*UpdateOutput, error) {
	_, err := r.server.UserClient.UpdateDefaultDeliveryAddress(ctx, &pb.UpdateDefaultDeliveryAddressReq{
		DeliveryAddressId: in.DeliveryAddressID,
		AuthId:            in.AuthID,
	})

	if err != nil {
		return nil, err
	}

	return &UpdateOutput{
		Message: "default delivery address updated",
	}, nil
}

func (r *mutationResolver) DeleteDeliveryAddress(ctx context.Context, addressId string) (*DeleteOutput, error) {
	_, err := r.server.UserClient.DeleteDeliveryAddress(ctx, &pb.DeleteDeliveryAddressReq{
		Id: addressId,
	})

	if err != nil {
		return nil, err
	}

	return &DeleteOutput{
		Message: "default delivery address updated",
	}, nil
}

func (r *mutationResolver) CreateShop(ctx context.Context, in CreateShopInput) (*Shop, error) {
	res, err := r.server.ShopClient.CreateShop(ctx, ToPbCreateShopReq(in))
	if err != nil {
		return nil, err
	}

	if _, err = r.server.GeolocationClient.ReverseGeocode(ctx, &pb.ReverseGeocodeReq{
		AddressId: res.Address.Id,
		Latitude:  in.Address.Latitude,
		Longitude: in.Address.Longitude,
	}); err != nil {
		fmt.Printf("failed to store address detail: %+v", err)
	}

	return ToGQShop(res), nil
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

func (r *mutationResolver) CreateRestaurantMenu(ctx context.Context, in CreateRestaurantMenuInput) (*RestaurantMenu, error) {
	res, err := r.server.ProductClient.CreateRestaurantMenu(ctx, ToPbCreateRestaurantMenuReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQRestaurantMenu(res), nil
}

func (r *mutationResolver) CreateMenuItem(ctx context.Context, in CreateMenuItemInput) (*MenuItem, error) {
	res, err := r.server.ProductClient.CreateMenuItem(ctx, ToPbCreateMenuItemReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQMenuItem(res), nil
}

func (r *mutationResolver) CreateMenuItemVariant(ctx context.Context, in CreateItemVariantInput) (*ItemVariant, error) {
	res, err := r.server.ProductClient.CreateMenuItemVariant(ctx, ToPbCreateItemVariantReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQItemVariant(res), nil
}

func (r *mutationResolver) CreateMenuItemAddon(ctx context.Context, in CreateItemAddonInput) (*ItemAddon, error) {
	res, err := r.server.ProductClient.CreateMenuItemAddon(ctx, ToPbCreateItemAddonReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQItemAddon(res), nil
}

func (r *mutationResolver) CreateRetailCategory(ctx context.Context, in CreateRetailCategoryInput) (*RetailCategory, error) {
	res, err := r.server.ProductClient.CreateRetailCategory(ctx, ToPbCreateRetailCategoryReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQRetailCategory(res), nil
}

func (r *mutationResolver) CreateRetailItem(ctx context.Context, in CreateRetailItemInput) (*RetailItem, error) {
	res, err := r.server.ProductClient.CreateRetailItem(ctx, ToPbCreateRetailItemReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQRetailItem(res), nil
}

func (r *mutationResolver) CreateRetailItemVariant(ctx context.Context, in CreateItemVariantInput) (*ItemVariant, error) {
	res, err := r.server.ProductClient.CreateRetailItemVariant(ctx, ToPbCreateItemVariantReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQItemVariant(res), nil
}

func (r *mutationResolver) CreateMedicineCategory(ctx context.Context, in CreateMedicineCategoryInput) (*MedicineCategory, error) {
	res, err := r.server.ProductClient.CreateMedicineCategory(ctx, ToPbCreateMedicineCategoryReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQMedicineCategory(res), nil
}

func (r *mutationResolver) CreateMedicineItem(ctx context.Context, in CreateMedicineItemInput) (*MedicineItem, error) {
	res, err := r.server.ProductClient.CreateMedicineItem(ctx, ToPbCreateMedicineItemReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQMedicineItem(res), nil
}

func (r *mutationResolver) UpdateRestaurantMenu(ctx context.Context, in UpdateRestaurantMenuInput) (*UpdateOutput, error) {
	req := &pb.UpdateRestaurantMenuReq{
		Id:       in.ID,
		MenuName: in.MenuName,
		ImageUrl: in.ImageURL,
	}

	_, err := r.server.ProductClient.UpdateRestaurantMenu(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Restaurant menu updated successfully"}, nil
}

func (r *mutationResolver) UpdateMenuItem(ctx context.Context, in UpdateMenuItemInput) (*UpdateOutput, error) {
	req := &pb.UpdateMenuItemReq{
		Id:          in.ID,
		Name:        in.Name,
		Price:       in.Price,
		ImageUrl:    in.ImageURL,
		Description: in.Description,
	}

	_, err := r.server.ProductClient.UpdateMenuItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Menu item updated successfully"}, nil
}

func (r *mutationResolver) UpdateMenuItemVariant(ctx context.Context, in UpdateItemVariantInput) (*UpdateOutput, error) {
	req := &pb.UpdateItemVariantReq{
		Id:              in.ID,
		VariantName:     in.VariantName,
		RelativePricing: in.RelativePricing,
		RelativePrice:   in.RelativePrice,
		Price:           in.Price,
		ImageUrl:        in.ImageURL,
		Description:     in.Description,
	}

	_, err := r.server.ProductClient.UpdateMenuItemVariant(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Item variant updated successfully"}, nil
}

func (r *mutationResolver) UpdateMenuItemAddon(ctx context.Context, in UpdateItemAddonInput) (*UpdateOutput, error) {
	req := &pb.UpdateItemAddonReq{
		Id:          in.ID,
		AddonName:   in.AddonName,
		AddonPrice:  in.AddonPrice,
		ImageUrl:    in.ImageURL,
		Description: in.Description,
	}

	_, err := r.server.ProductClient.UpdateMenuItemAddon(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Item addon updated successfully"}, nil
}

func (r *mutationResolver) UpdateRetailCategory(ctx context.Context, in UpdateRetailCategoryInput) (*UpdateOutput, error) {
	req := &pb.UpdateRetailCategoryReq{
		Id:           in.ID,
		CategoryName: in.CategoryName,
		ImageUrl:     in.ImageURL,
	}

	_, err := r.server.ProductClient.UpdateRetailCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Retail category updated successfully"}, nil
}

func (r *mutationResolver) UpdateRetailItem(ctx context.Context, in UpdateRetailItemInput) (*UpdateOutput, error) {
	req := &pb.UpdateRetailItemReq{
		Id:          in.ID,
		Name:        in.Name,
		Price:       in.Price,
		ImageUrl:    in.ImageURL,
		Description: in.Description,
	}

	_, err := r.server.ProductClient.UpdateRetailItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Retail item updated successfully"}, nil
}

func (r *mutationResolver) UpdateRetailItemVariant(ctx context.Context, in UpdateItemVariantInput) (*UpdateOutput, error) {
	req := &pb.UpdateItemVariantReq{
		Id:              in.ID,
		VariantName:     in.VariantName,
		RelativePricing: in.RelativePricing,
		RelativePrice:   in.RelativePrice,
		Price:           in.Price,
		ImageUrl:        in.ImageURL,
		Description:     in.Description,
		ItemId:          in.ItemID,
	}

	_, err := r.server.ProductClient.UpdateRetailItemVariant(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Item variant updated successfully"}, nil
}

func (r *mutationResolver) UpdateMedicineCategory(ctx context.Context, in UpdateMedicineCategoryInput) (*UpdateOutput, error) {
	req := &pb.UpdateMedicineCategoryReq{
		Id:           in.ID,
		CategoryName: in.CategoryName,
		ImageUrl:     in.ImageURL,
	}

	_, err := r.server.ProductClient.UpdateMedicineCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Medicine category updated successfully"}, nil
}

func (r *mutationResolver) UpdateMedicineItem(ctx context.Context, in UpdateMedicineItemInput) (*UpdateOutput, error) {
	req := &pb.UpdateMedicineItemReq{
		Id:          in.ID,
		Name:        in.Name,
		Price:       in.Price,
		ImageUrl:    in.ImageURL,
		Description: in.Description,
	}

	_, err := r.server.ProductClient.UpdateMedicineItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Medicine item updated successfully"}, nil
}

func (r *mutationResolver) DeleteRestaurantMenu(ctx context.Context, menuID string) (*DeleteOutput, error) {
	req := &pb.DeleteRestaurantMenuReq{Id: menuID}

	_, err := r.server.ProductClient.DeleteRestaurantMenu(ctx, req)
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Restaurant menu deleted successfully"}, nil
}

func (r *mutationResolver) DeleteMenuItem(ctx context.Context, itemID string) (*DeleteOutput, error) {
	req := &pb.DeleteMenuItemReq{Id: itemID}

	_, err := r.server.ProductClient.DeleteMenuItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Menu item deleted successfully"}, nil
}

func (r *mutationResolver) DeleteMenuItemVariant(ctx context.Context, in DeleteItemVariantInput) (*DeleteOutput, error) {
	_, err := r.server.ProductClient.DeleteMenuItemVariant(ctx, &pb.DeleteItemVariantReq{
		ItemId:    in.ItemID,
		VariantId: in.VariantID,
	})

	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Retail category deleted successfully"}, nil
}

func (r *mutationResolver) DeleteMenuItemAddon(ctx context.Context, in DeleteItemAddonInput) (*DeleteOutput, error) {
	_, err := r.server.ProductClient.DeleteMenuItemAddon(ctx, &pb.DeleteItemAddonReq{
		ItemId:  in.ItemID,
		AddonId: in.AddonID,
	})

	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Retail category deleted successfully"}, nil
}

func (r *mutationResolver) DeleteRetailCategory(ctx context.Context, categoryID string) (*DeleteOutput, error) {
	req := &pb.DeleteRetailCategoryReq{Id: categoryID}

	_, err := r.server.ProductClient.DeleteRetailCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Retail category deleted successfully"}, nil
}

func (r *mutationResolver) DeleteRetailItem(ctx context.Context, itemID string) (*DeleteOutput, error) {
	_, err := r.server.ProductClient.DeleteRetailItem(ctx, &pb.DeleteRetailItemReq{Id: itemID})
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Retail item deleted successfully"}, nil
}

func (r *mutationResolver) DeleteRetailItemVariant(ctx context.Context, in DeleteItemVariantInput) (*DeleteOutput, error) {
	_, err := r.server.ProductClient.DeleteRetailItemVariant(ctx, &pb.DeleteItemVariantReq{
		ItemId:    in.ItemID,
		VariantId: in.VariantID,
	})

	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Retail category deleted successfully"}, nil
}

func (r *mutationResolver) DeleteMedicineCategory(ctx context.Context, categoryID string) (*DeleteOutput, error) {
	req := &pb.DeleteMedicineCategoryReq{Id: categoryID}

	_, err := r.server.ProductClient.DeleteMedicineCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Medicine category deleted successfully"}, nil
}

func (r *mutationResolver) DeleteMedicineItem(ctx context.Context, itemID string) (*DeleteOutput, error) {
	req := &pb.DeleteMedicineItemReq{Id: itemID}

	_, err := r.server.ProductClient.DeleteMedicineItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Medicine item deleted successfully"}, nil
}

func (r *mutationResolver) CreateItemStock(ctx context.Context, in CreateItemStockInput) (*ItemStock, error) {
	res, err := r.server.InventoryClient.CreateItemStock(ctx, ToPbCreateItemStockReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQItemStock(res), nil
}

func (r *mutationResolver) CreateVariantStock(ctx context.Context, in CreateVariantStockInput) (*VariantStock, error) {
	res, err := r.server.InventoryClient.CreateVariantStock(ctx, ToPbCreateVariantStockReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQVariantStock(res), nil
}

func (r *mutationResolver) CreateAddonStock(ctx context.Context, in CreateAddonStockInput) (*AddonStock, error) {
	res, err := r.server.InventoryClient.CreateAddonStock(ctx, ToPbCreateAddonStockReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQAddonStock(res), nil
}

func (r *mutationResolver) UpdateItemStock(ctx context.Context, in UpdateItemStockInput) (*ItemStock, error) {
	res, err := r.server.InventoryClient.UpdateItemStock(ctx, ToPbUpdateItemStockReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQItemStock(res), nil
}

func (r *mutationResolver) UpdateVariantStock(ctx context.Context, in UpdateVariantStockInput) (*VariantStock, error) {
	res, err := r.server.InventoryClient.UpdateVariantStock(ctx, ToPbUpdateVariantStockReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQVariantStock(res), nil
}

func (r *mutationResolver) UpdateAddonStock(ctx context.Context, in UpdateAddonStockInput) (*AddonStock, error) {
	res, err := r.server.InventoryClient.UpdateAddonStock(ctx, ToPbUpdateAddonStockReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQAddonStock(res), nil
}

func (r *mutationResolver) DeleteItemStock(ctx context.Context, id string) (*DeleteOutput, error) {
	_, err := r.server.InventoryClient.DeleteItemStock(ctx, &pb.DeleteItemStockReq{Id: id})
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Item stock deleted successfully"}, nil
}

func (r *mutationResolver) DeleteVariantStock(ctx context.Context, id string) (*DeleteOutput, error) {
	_, err := r.server.InventoryClient.DeleteVariantStock(ctx, &pb.DeleteVariantStockReq{Id: id})
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Variant stock deleted successfully"}, nil
}

func (r *mutationResolver) DeleteAddonStock(ctx context.Context, id string) (*DeleteOutput, error) {
	_, err := r.server.InventoryClient.DeleteAddonStock(ctx, &pb.DeleteAddonStockReq{Id: id})
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Addon stock deleted successfully"}, nil
}

func setSessionCookies(ctx context.Context, accessToken, refreshToken string) error {
	cookieManager, err := customMiddleware.GetCookieManager(ctx)
	if err != nil {
		return err
	}

	if accessToken == "" || refreshToken == "" {
		return fmt.Errorf("access token/ refresh token empty")
	}

	cookieManager.SetCookie(
		"access_token",
		accessToken,
		utils.CookieOptions{
			Path:     "/",
			MaxAge:   int((15 * time.Minute).Seconds()),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		},
	)

	cookieManager.SetCookie(
		"refresh_token",
		refreshToken,
		utils.CookieOptions{
			Path:     "/",
			MaxAge:   int((30 * 24 * time.Hour).Seconds()),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		},
	)

	return nil
}

func deleteSessionCookies(ctx context.Context) {
	cookieManager, err := customMiddleware.GetCookieManager(ctx)
	if err != nil {
		return
	}

	cookieManager.DeleteCookie("access_token", utils.CookieOptions{
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	cookieManager.DeleteCookie("refresh_token", utils.CookieOptions{
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

}
