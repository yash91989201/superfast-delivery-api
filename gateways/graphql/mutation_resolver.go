package graphql

import (
	"context"
	"fmt"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
)

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) SignInWithEmail(ctx context.Context, in SignInWithEmailInput) (*SignInOutput, error) {
	signInRes, err := r.server.authenticationClient.SignInWithEmail(
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

	profile, err := r.server.userClient.GetProfile(ctx, &pb.GetProfileReq{AuthId: signInRes.Auth.Id})
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
	return res, nil
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
		Auth:          ToGQAuth(signInRes.Auth),
		Session:       ToGQSession(signInRes.Session),
		Profile:       ToGQProfile(profile),
		CreateProfile: profile == nil,
		VerifyOtp:     false,
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

	return ToGQProfile(res), nil
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, in UpdateProfileInput) (*Profile, error) {
	res, err := r.server.userClient.UpdateProfile(ctx, &pb.UpdateProfileReq{
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
	newDeliveryAddress, err := r.server.userClient.CreateDeliveryAddress(ctx, ToPbCreateDeliveryAddress(in))
	if err != nil {
		return nil, err
	}

	_, err = r.server.geolocationClient.ReverseGeocode(ctx, &pb.ReverseGeocodeReq{
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
	_, err := r.server.userClient.UpdateDefaultDeliveryAddress(ctx, &pb.UpdateDefaultDeliveryAddressReq{
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
	_, err := r.server.userClient.DeleteDeliveryAddress(ctx, &pb.DeleteDeliveryAddressReq{
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
	res, err := r.server.shopClient.CreateShop(ctx, ToPbCreateShopReq(in))
	if err != nil {
		return nil, err
	}

	if _, err = r.server.geolocationClient.ReverseGeocode(ctx, &pb.ReverseGeocodeReq{
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
	res, err := r.server.productClient.CreateRestaurantMenu(ctx, ToPbCreateRestaurantMenuReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQRestaurantMenu(res), nil
}

func (r *mutationResolver) CreateMenuItem(ctx context.Context, in CreateMenuItemInput) (*MenuItem, error) {
	res, err := r.server.productClient.CreateMenuItem(ctx, ToPbCreateMenuItemReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQMenuItem(res), nil
}

func (r *mutationResolver) CreateMenuItemVariant(ctx context.Context, in CreateItemVariantInput) (*ItemVariant, error) {
	res, err := r.server.productClient.CreateMenuItemVariant(ctx, ToPbCreateItemVariantReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQItemVariant(res), nil
}

func (r *mutationResolver) CreateMenuItemAddon(ctx context.Context, in CreateItemAddonInput) (*ItemAddon, error) {
	res, err := r.server.productClient.CreateMenuItemAddon(ctx, ToPbCreateItemAddonReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQItemAddon(res), nil
}

func (r *mutationResolver) CreateRetailCategory(ctx context.Context, in CreateRetailCategoryInput) (*RetailCategory, error) {
	res, err := r.server.productClient.CreateRetailCategory(ctx, ToPbCreateRetailCategoryReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQRetailCategory(res), nil
}

func (r *mutationResolver) CreateRetailItem(ctx context.Context, in CreateRetailItemInput) (*RetailItem, error) {
	res, err := r.server.productClient.CreateRetailItem(ctx, ToPbCreateRetailItemReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQRetailItem(res), nil
}

func (r *mutationResolver) CreateRetailItemVariant(ctx context.Context, in CreateItemVariantInput) (*ItemVariant, error) {
	res, err := r.server.productClient.CreateRetailItemVariant(ctx, ToPbCreateItemVariantReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQItemVariant(res), nil
}

func (r *mutationResolver) CreateMedicineCategory(ctx context.Context, in CreateMedicineCategoryInput) (*MedicineCategory, error) {
	res, err := r.server.productClient.CreateMedicineCategory(ctx, ToPbCreateMedicineCategoryReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQMedicineCategory(res), nil
}

func (r *mutationResolver) CreateMedicineItem(ctx context.Context, in CreateMedicineItemInput) (*MedicineItem, error) {
	res, err := r.server.productClient.CreateMedicineItem(ctx, ToPbCreateMedicineItemReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQMedicineItem(res), nil
}

// UpdateItemVariant resolver

// UpdateRestaurantMenu resolver
func (r *mutationResolver) UpdateRestaurantMenu(ctx context.Context, input UpdateRestaurantMenuInput) (*UpdateOutput, error) {
	req := &pb.UpdateRestaurantMenuReq{
		Id:       input.ID,
		MenuName: input.MenuName,
		ImageUrl: input.ImageURL,
	}

	_, err := r.server.productClient.UpdateRestaurantMenu(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Restaurant menu updated successfully"}, nil
}

// UpdateMenuItem resolver
func (r *mutationResolver) UpdateMenuItem(ctx context.Context, input UpdateMenuItemInput) (*UpdateOutput, error) {
	req := &pb.UpdateMenuItemReq{
		Id:          input.ID,
		Name:        input.Name,
		Price:       input.Price,
		ImageUrl:    input.ImageURL,
		Description: input.Description,
	}

	_, err := r.server.productClient.UpdateMenuItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Menu item updated successfully"}, nil
}

func (r *mutationResolver) UpdateMenuItemVariant(ctx context.Context, input UpdateItemVariantInput) (*UpdateOutput, error) {
	req := &pb.UpdateItemVariantReq{
		Id:              input.ID,
		VariantName:     input.VariantName,
		RelativePricing: input.RelativePricing,
		RelativePrice:   input.RelativePrice,
		Price:           input.Price,
		ImageUrl:        input.ImageURL,
		Description:     input.Description,
	}

	_, err := r.server.productClient.UpdateMenuItemVariant(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Item variant updated successfully"}, nil
}

// UpdateItemAddon resolver
func (r *mutationResolver) UpdateMenuItemAddon(ctx context.Context, input UpdateItemAddonInput) (*UpdateOutput, error) {
	req := &pb.UpdateItemAddonReq{
		Id:          input.ID,
		AddonName:   input.AddonName,
		AddonPrice:  input.AddonPrice,
		ImageUrl:    input.ImageURL,
		Description: input.Description,
	}

	_, err := r.server.productClient.UpdateMenuItemAddon(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Item addon updated successfully"}, nil
}

// UpdateRetailCategory resolver
func (r *mutationResolver) UpdateRetailCategory(ctx context.Context, input UpdateRetailCategoryInput) (*UpdateOutput, error) {
	req := &pb.UpdateRetailCategoryReq{
		Id:           input.ID,
		CategoryName: input.CategoryName,
		ImageUrl:     input.ImageURL,
	}

	_, err := r.server.productClient.UpdateRetailCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Retail category updated successfully"}, nil
}

// UpdateRetailItem resolver
func (r *mutationResolver) UpdateRetailItem(ctx context.Context, input UpdateRetailItemInput) (*UpdateOutput, error) {
	req := &pb.UpdateRetailItemReq{
		Id:          input.ID,
		Name:        input.Name,
		Price:       input.Price,
		ImageUrl:    input.ImageURL,
		Description: input.Description,
	}

	_, err := r.server.productClient.UpdateRetailItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Retail item updated successfully"}, nil
}

func (r *mutationResolver) UpdateRetailItemVariant(ctx context.Context, input UpdateItemVariantInput) (*UpdateOutput, error) {
	req := &pb.UpdateItemVariantReq{
		Id:              input.ID,
		VariantName:     input.VariantName,
		RelativePricing: input.RelativePricing,
		RelativePrice:   input.RelativePrice,
		Price:           input.Price,
		ImageUrl:        input.ImageURL,
		Description:     input.Description,
	}

	_, err := r.server.productClient.UpdateRetailItemVariant(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Item variant updated successfully"}, nil
}

// UpdateMedicineCategory resolver
func (r *mutationResolver) UpdateMedicineCategory(ctx context.Context, input UpdateMedicineCategoryInput) (*UpdateOutput, error) {
	req := &pb.UpdateMedicineCategoryReq{
		Id:           input.ID,
		CategoryName: input.CategoryName,
		ImageUrl:     input.ImageURL,
	}

	_, err := r.server.productClient.UpdateMedicineCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Medicine category updated successfully"}, nil
}

// UpdateMedicineItem resolver
func (r *mutationResolver) UpdateMedicineItem(ctx context.Context, input UpdateMedicineItemInput) (*UpdateOutput, error) {
	req := &pb.UpdateMedicineItemReq{
		Id:          input.ID,
		Name:        input.Name,
		Price:       input.Price,
		ImageUrl:    input.ImageURL,
		Description: input.Description,
	}

	_, err := r.server.productClient.UpdateMedicineItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Message: "Medicine item updated successfully"}, nil
}

func (r *mutationResolver) DeleteRestaurantMenu(ctx context.Context, menuID string) (*DeleteOutput, error) {
	req := &pb.DeleteRestaurantMenuReq{Id: menuID}

	_, err := r.server.productClient.DeleteRestaurantMenu(ctx, req)
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Restaurant menu deleted successfully"}, nil
}

func (r *mutationResolver) DeleteMenuItem(ctx context.Context, itemID string) (*DeleteOutput, error) {
	req := &pb.DeleteMenuItemReq{Id: itemID}

	_, err := r.server.productClient.DeleteMenuItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Menu item deleted successfully"}, nil
}

func (r *mutationResolver) DeleteMenuItemVariant(ctx context.Context, in DeleteItemVariantInput) (*DeleteOutput, error) {
	_, err := r.server.productClient.DeleteMenuItemVariant(ctx, &pb.DeleteItemVariantReq{
		ItemId:    in.ItemID,
		VariantId: in.VariantID,
	})

	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Retail category deleted successfully"}, nil
}

func (r *mutationResolver) DeleteMenuItemAddon(ctx context.Context, in DeleteItemAddonInput) (*DeleteOutput, error) {
	_, err := r.server.productClient.DeleteMenuItemAddon(ctx, &pb.DeleteItemAddonReq{
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

	_, err := r.server.productClient.DeleteRetailCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Retail category deleted successfully"}, nil
}

func (r *mutationResolver) DeleteRetailItem(ctx context.Context, itemID string) (*DeleteOutput, error) {
	_, err := r.server.productClient.DeleteRetailItem(ctx, &pb.DeleteRetailItemReq{Id: itemID})
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Retail item deleted successfully"}, nil
}

func (r *mutationResolver) DeleteRetailItemVariant(ctx context.Context, in DeleteItemVariantInput) (*DeleteOutput, error) {
	_, err := r.server.productClient.DeleteRetailItemVariant(ctx, &pb.DeleteItemVariantReq{
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

	_, err := r.server.productClient.DeleteMedicineCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Medicine category deleted successfully"}, nil
}

// DeleteMedicineItem resolver
func (r *mutationResolver) DeleteMedicineItem(ctx context.Context, itemID string) (*DeleteOutput, error) {
	req := &pb.DeleteMedicineItemReq{Id: itemID}

	_, err := r.server.productClient.DeleteMedicineItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Medicine item deleted successfully"}, nil
}

func (r *mutationResolver) CreateItemStock(ctx context.Context, in CreateItemStockInput) (*ItemStock, error) {
	res, err := r.server.inventoryClient.CreateItemStock(ctx, ToPbCreateItemStockReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQItemStock(res), nil
}

func (r *mutationResolver) CreateVariantStock(ctx context.Context, in CreateVariantStockInput) (*VariantStock, error) {
	res, err := r.server.inventoryClient.CreateVariantStock(ctx, ToPbCreateVariantStockReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQVariantStock(res), nil
}

func (r *mutationResolver) CreateAddonStock(ctx context.Context, in CreateAddonStockInput) (*AddonStock, error) {
	res, err := r.server.inventoryClient.CreateAddonStock(ctx, ToPbCreateAddonStockReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQAddonStock(res), nil
}

func (r *mutationResolver) UpdateItemStock(ctx context.Context, in UpdateItemStockInput) (*ItemStock, error) {
	res, err := r.server.inventoryClient.UpdateItemStock(ctx, ToPbUpdateItemStockReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQItemStock(res), nil
}

func (r *mutationResolver) UpdateVariantStock(ctx context.Context, in UpdateVariantStockInput) (*VariantStock, error) {
	res, err := r.server.inventoryClient.UpdateVariantStock(ctx, ToPbUpdateVariantStockReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQVariantStock(res), nil
}

func (r *mutationResolver) UpdateAddonStock(ctx context.Context, in UpdateAddonStockInput) (*AddonStock, error) {
	res, err := r.server.inventoryClient.UpdateAddonStock(ctx, ToPbUpdateAddonStockReq(&in))
	if err != nil {
		return nil, err
	}

	return ToGQAddonStock(res), nil
}

func (r *mutationResolver) DeleteItemStock(ctx context.Context, id string) (*DeleteOutput, error) {
	_, err := r.server.inventoryClient.DeleteItemStock(ctx, &pb.DeleteItemStockReq{Id: id})
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Item stock deleted successfully"}, nil
}

func (r *mutationResolver) DeleteVariantStock(ctx context.Context, id string) (*DeleteOutput, error) {
	_, err := r.server.inventoryClient.DeleteVariantStock(ctx, &pb.DeleteVariantStockReq{Id: id})
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Variant stock deleted successfully"}, nil
}

func (r *mutationResolver) DeleteAddonStock(ctx context.Context, id string) (*DeleteOutput, error) {
	_, err := r.server.inventoryClient.DeleteAddonStock(ctx, &pb.DeleteAddonStockReq{Id: id})
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{Message: "Addon stock deleted successfully"}, nil
}
