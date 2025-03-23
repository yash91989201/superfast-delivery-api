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

	return ToGQAuth(auth), nil
}

func (r *queryResolver) GetAuth(ctx context.Context, in GetAuthInput) (*Auth, error) {
	auth, err := r.server.authenticationClient.GetAuth(ctx, &pb.GetAuthReq{Email: in.Email, Phone: in.Phone})
	if err != nil {
		return nil, err
	}

	return ToGQAuth(auth), nil
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
		DeliveryAddress: ToGQDeliveryAddressList(res.DeliveryAddresses),
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

func (r *queryResolver) GetRestaurantMenu(ctx context.Context, id string) (*RestaurantMenu, error) {
	res, err := r.server.productClient.GetRestaurantMenu(ctx, &pb.GetRestaurantMenuReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQRestaurantMenu(res), nil
}

func (r *queryResolver) GetMenuItem(ctx context.Context, id string) (*MenuItem, error) {
	res, err := r.server.productClient.GetMenuItem(ctx, &pb.GetMenuItemReq{Id: id})
	if err != nil {
		return nil, err
	}
	return ToGQMenuItem(res), nil
}

func (r *queryResolver) GetMenuItemVariant(ctx context.Context, in GetItemVariantInput) (*ItemVariant, error) {
	res, err := r.server.productClient.GetMenuItemVariant(ctx, &pb.GetItemVariantReq{
		ItemId:    in.ItemID,
		VariantId: in.VariantID,
	})

	if err != nil {
		return nil, err
	}

	return ToGQItemVariant(res), err
}

func (r *queryResolver) GetMenuItemAddon(ctx context.Context, in GetItemAddonInput) (*ItemAddon, error) {
	res, err := r.server.productClient.GetMenuItemAddon(ctx, &pb.GetItemAddonReq{
		ItemId:  in.ItemID,
		AddonId: in.AddonID,
	})

	if err != nil {
		return nil, err
	}

	return ToGQItemAddon(res), err
}

func (r *queryResolver) GetRetailCategory(ctx context.Context, id string) (*RetailCategory, error) {
	res, err := r.server.productClient.GetRetailCategory(ctx, &pb.GetRetailCategoryReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQRetailCategory(res), nil
}

func (r *queryResolver) GetRetailItem(ctx context.Context, id string) (*RetailItem, error) {
	res, err := r.server.productClient.GetRetailItem(ctx, &pb.GetRetailItemReq{Id: id})
	if err != nil {
		return nil, err
	}
	return ToGQRetailItem(res), nil
}

func (r *queryResolver) GetRetailItemVariant(ctx context.Context, in GetItemVariantInput) (*ItemVariant, error) {
	res, err := r.server.productClient.GetRetailItemVariant(ctx, &pb.GetItemVariantReq{
		ItemId:    in.ItemID,
		VariantId: in.VariantID,
	})

	if err != nil {
		return nil, err
	}

	return ToGQItemVariant(res), nil
}

func (r *queryResolver) GetMedicineCategory(ctx context.Context, id string) (*MedicineCategory, error) {
	res, err := r.server.productClient.GetMedicineCategory(ctx, &pb.GetMedicineCategoryReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQMedicineCategory(res), nil
}

func (r *queryResolver) GetMedicineItem(ctx context.Context, id string) (*MedicineItem, error) {
	res, err := r.server.productClient.GetMedicineItem(ctx, &pb.GetMedicineItemReq{Id: id})
	if err != nil {
		return nil, err
	}
	return ToGQMedicineItem(res), nil
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

func (r *queryResolver) ListMenuItem(ctx context.Context, menuId string) (*ListMenuItemOutput, error) {
	res, err := r.server.productClient.ListMenuItem(ctx, &pb.ListMenuItemReq{MenuId: menuId})
	if err != nil {
		return nil, err
	}

	menuItems := ToGQMenuItemList(res.MenuItemList)
	return &ListMenuItemOutput{
		MenuItems:  menuItems,
		TotalItems: int32(len(menuItems)),
	}, nil
}

func (r *queryResolver) ListMenuItemVariant(ctx context.Context, itemID string) (*ListItemVariantOutput, error) {
	res, err := r.server.productClient.ListMenuItemVariant(ctx, &pb.ListItemVariantReq{
		ItemId: itemID,
	})

	if err != nil {
		return nil, err
	}

	variants := ToGQItemVariantList(res.Variants)

	return &ListItemVariantOutput{
		Variants:      variants,
		TotalVariants: int32(len(variants)),
	}, nil
}

func (r *queryResolver) ListMenuItemAddon(ctx context.Context, itemID string) (*ListItemAddonOutput, error) {
	res, err := r.server.productClient.ListMenuItemAddon(ctx, &pb.ListItemAddonReq{
		ItemId: itemID,
	})

	if err != nil {
		return nil, err
	}

	addons := ToGQItemAddonList(res.Addons)

	return &ListItemAddonOutput{
		Addons:      addons,
		TotalAddons: int32(len(addons)),
	}, nil
}

func (r *queryResolver) ListRetailItem(ctx context.Context, categoryId string) (*ListRetailItemOutput, error) {
	res, err := r.server.productClient.ListRetailItem(ctx, &pb.ListRetailItemReq{CategoryId: categoryId})
	if err != nil {
		return nil, err
	}

	retailItems := ToGQRetailItemList(res.RetailItemList)
	return &ListRetailItemOutput{
		RetailItems: retailItems,
		TotalItems:  int32(len(retailItems)),
	}, nil
}

func (r *queryResolver) ListRetailItemVariant(ctx context.Context, itemID string) (*ListItemVariantOutput, error) {
	res, err := r.server.productClient.ListRetailItemVariant(ctx, &pb.ListItemVariantReq{
		ItemId: itemID,
	})

	if err != nil {
		return nil, err
	}

	variants := ToGQItemVariantList(res.Variants)

	return &ListItemVariantOutput{
		Variants:      variants,
		TotalVariants: int32(len(variants)),
	}, nil
}

func (r *queryResolver) ListMedicineItem(ctx context.Context, categoryId string) (*ListMedicineItemOutput, error) {
	res, err := r.server.productClient.ListMedicineItem(ctx, &pb.ListMedicineItemReq{CategoryId: categoryId})
	if err != nil {
		return nil, err
	}

	medicineItems := ToGQMedicineItemList(res.MedicineItemList)
	return &ListMedicineItemOutput{
		MedicineItems: medicineItems,
		TotalItems:    int32(len(medicineItems)),
	}, nil
}

func (r *queryResolver) GetItemStock(ctx context.Context, id string) (*ItemStock, error) {
	res, err := r.server.inventoryClient.GetItemStock(ctx, &pb.GetItemStockReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQItemStock(res), nil
}

func (r *queryResolver) GetVariantStock(ctx context.Context, id string) (*VariantStock, error) {
	res, err := r.server.inventoryClient.GetVariantStock(ctx, &pb.GetVariantStockReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQVariantStock(res), nil
}

func (r *queryResolver) GetAddonStock(ctx context.Context, id string) (*AddonStock, error) {
	res, err := r.server.inventoryClient.GetAddonStock(ctx, &pb.GetAddonStockReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQAddonStock(res), nil
}
