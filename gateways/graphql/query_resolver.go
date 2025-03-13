package graphql

import (
	"context"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
)

type queryResolver struct {
	server *Server
}

func (r *queryResolver) GetAuthByID(ctx context.Context, in GetAuthByIDInput) (*Auth, error) {
	auth, err := r.server.authenticationClient.GetAuthById(ctx, &pb.GetAuthByIdReq{Id: in.ID})
	if err != nil {
		return nil, err
	}

	return ToAuth(auth), nil
}

func (r *queryResolver) GetAuth(ctx context.Context, in GetAuthInput) (*Auth, error) {
	auth, err := r.server.authenticationClient.GetAuth(ctx, &pb.GetAuthReq{Email: in.Email, Phone: in.Phone})
	if err != nil {
		return nil, err
	}

	return ToAuth(auth), nil
}

func (r *queryResolver) GetProfile(ctx context.Context, in GetProfileInput) (*Profile, error) {
	return nil, nil
}

func (r *queryResolver) GetDeliveryAddress(ctx context.Context, id string) (*DeliveryAddress, error) {
	deliveryAddress, err := r.server.userClient.GetDeliveryAddress(ctx, &pb.GetDeliveryAddressReq{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	return ToGQDeliveryAddress(deliveryAddress), nil
}

func (r *queryResolver) GetDeliveryAddressDetail(ctx context.Context, addressId string) (*AddressDetail, error) {
	deliveryAddress, err := r.server.userClient.GetDeliveryAddress(ctx, &pb.GetDeliveryAddressReq{Id: addressId})
	if err != nil {
		return nil, err
	}

	res, err := r.server.geolocationClient.ReverseGeocode(ctx, &pb.ReverseGeocodeReq{
		Latitude:  deliveryAddress.Latitude,
		Longitude: deliveryAddress.Longitude,
		AddressId: addressId,
	})

	if err != nil {
		return nil, err
	}

	return ToGQAddressDetail(res), nil
}

func (r *queryResolver) ListDeliveryAddress(ctx context.Context, authId string) (*ListDeliveryAddressOutput, error) {

	res, err := r.server.userClient.ListDeliveryAddress(ctx, &pb.ListDeliveryAddressReq{
		AuthId: authId,
	})

	if err != nil {
		return nil, err
	}

	return &ListDeliveryAddressOutput{
		DeliveryAddress: ToDeliveryAddressList(res.DeliveryAddresses),
	}, nil
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

func (r *queryResolver) GetItemVariant(ctx context.Context, id string) (*ItemVariant, error) {
	res, err := r.server.productClient.GetItemVariant(ctx, &pb.GetItemVariantReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQItemVariant(res), err
}

func (r *queryResolver) GetItemAddon(ctx context.Context, id string) (*ItemAddon, error) {
	res, err := r.server.productClient.GetItemAddon(ctx, &pb.GetItemAddonReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQItemAddon(res), err
}

func (r *queryResolver) GetItemVariants(ctx context.Context, itemId string) (*GetItemVariantsOutput, error) {
	res, err := r.server.productClient.GetItemVariants(ctx, &pb.GetItemVariantsReq{ItemId: itemId})
	if err != nil {
		return nil, err
	}

	variants := ToGQItemVariants(res.Variants)

	return &GetItemVariantsOutput{
		Variants:      variants,
		TotalVariants: int32(len(variants)),
	}, err
}

func (r *queryResolver) GetItemAddons(ctx context.Context, itemId string) (*GetItemAddonsOutput, error) {
	res, err := r.server.productClient.GetItemAddons(ctx, &pb.GetItemAddonsReq{ItemId: itemId})
	if err != nil {
		return nil, err
	}

	addons := ToGQItemAddons(res.Addons)

	return &GetItemAddonsOutput{
		Addons:      addons,
		TotalAddons: int32(len(addons)),
	}, err
}

func (r *queryResolver) GetRestaurantMenu(ctx context.Context, id string) (*RestaurantMenu, error) {
	res, err := r.server.productClient.GetRestaurantMenu(ctx, &pb.GetRestaurantMenuReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQRestaurantMenu(res), nil
}

func (r *queryResolver) ListRestaurantMenu(ctx context.Context, shopId string) (*ListRestaurantMenuOutput, error) {
	res, err := r.server.productClient.ListRestaurantMenu(ctx, &pb.ListRestaurantMenuReq{ShopId: shopId})
	if err != nil {
		return nil, err
	}

	restaurantMenuList := ToGQRestaurantMenuList(res.RestaurantMenuList)
	return &ListRestaurantMenuOutput{
		RestaurantMenuList: restaurantMenuList,
		TotalMenu:          int32(len(restaurantMenuList)),
	}, nil
}

func (r *queryResolver) GetRetailCategory(ctx context.Context, id string) (*RetailCategory, error) {
	res, err := r.server.productClient.GetRetailCategory(ctx, &pb.GetRetailCategoryReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQRetailCategory(res), nil
}

func (r *queryResolver) ListRetailCategory(ctx context.Context, shopId string) (*ListRetailCategoryOutput, error) {
	res, err := r.server.productClient.ListRetailCategory(ctx, &pb.ListRetailCategoryReq{ShopId: shopId})
	if err != nil {
		return nil, err
	}

	retailCategoryList := ToGQRetailCategoryList(res.RetailCategoryList)

	return &ListRetailCategoryOutput{
		RetailCategoryList: retailCategoryList,
		TotalCategory:      int32(len(retailCategoryList)),
	}, nil
}

func (r *queryResolver) GetMedicineCategory(ctx context.Context, id string) (*MedicineCategory, error) {
	res, err := r.server.productClient.GetMedicineCategory(ctx, &pb.GetMedicineCategoryReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQMedicineCategory(res), nil
}

func (r *queryResolver) ListMedicineCategory(ctx context.Context, shopId string) (*ListMedicineCategoryOutput, error) {
	res, err := r.server.productClient.ListMedicineCategory(ctx, &pb.ListMedicineCategoryReq{ShopId: shopId})
	if err != nil {
		return nil, err
	}

	medicineCategoryList := ToGQMedicineCategoryList(res.MedicineCategoryList)

	return &ListMedicineCategoryOutput{
		MedicineCategoryList: medicineCategoryList,
		TotalCategory:        int32(len(medicineCategoryList)),
	}, nil
}
