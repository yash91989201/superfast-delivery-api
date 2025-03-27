package graphql

import (
	"context"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
)

type queryResolver struct {
	server *Server
}

func (r *queryResolver) GetAuthByID(ctx context.Context, in GetAuthByIDInput) (*Auth, error) {
	res, err := r.server.AuthenticationClient.GetAuthById(ctx, &pb.GetAuthByIdReq{Id: in.ID})
	if err != nil {
		return nil, err
	}

	return ToGQAuth(res), nil
}

func (r *queryResolver) GetAuth(ctx context.Context, in GetAuthInput) (*Auth, error) {
	res, err := r.server.AuthenticationClient.GetAuth(ctx, &pb.GetAuthReq{Email: in.Email, Phone: in.Phone})
	if err != nil {
		return nil, err
	}

	return ToGQAuth(res), nil
}

func (r *queryResolver) GetProfile(ctx context.Context, in GetProfileInput) (*Profile, error) {
	return nil, nil
}

func (r *queryResolver) GetDeliveryAddress(ctx context.Context, id string) (*DeliveryAddress, error) {
	res, err := r.server.UserClient.GetDeliveryAddress(ctx, &pb.GetDeliveryAddressReq{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	return ToGQDeliveryAddress(res), nil
}

func (r *queryResolver) GetDefaultDeliveryAddress(ctx context.Context, authID string) (*DeliveryAddress, error) {
	res, err := r.server.UserClient.GetDefaultDeliveryAddress(ctx, &pb.GetDefaultDeliveryAddressReq{
		AuthId: authID,
	})

	if err != nil {
		return nil, err
	}

	return ToGQDeliveryAddress(res), nil
}

func (r *queryResolver) GetDeliveryAddressDetail(ctx context.Context, addressId string) (*AddressDetail, error) {
	res, err := r.server.UserClient.GetDeliveryAddress(ctx, &pb.GetDeliveryAddressReq{Id: addressId})
	if err != nil {
		return nil, err
	}

	addressDetail, err := r.server.GeolocationClient.ReverseGeocode(ctx, &pb.ReverseGeocodeReq{
		Latitude:  res.Latitude,
		Longitude: res.Longitude,
		AddressId: addressId,
	})

	if err != nil {
		return nil, err
	}

	return ToGQAddressDetail(addressDetail), nil
}

func (r *queryResolver) ListDeliveryAddress(ctx context.Context, authId string) (*ListDeliveryAddressOutput, error) {

	res, err := r.server.UserClient.ListDeliveryAddress(ctx, &pb.ListDeliveryAddressReq{
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
	res, err := r.server.ShopClient.GetShop(ctx, &pb.GetShopReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQShop(res), nil
}

func (r *queryResolver) ListShops(ctx context.Context, in *ListShopsInput) (*ListShopsOutput, error) {
	res, err := r.server.ShopClient.ListShops(ctx, ToPbListShopReq(in))
	if err != nil {
		return nil, err
	}

	return &ListShopsOutput{
		Shops: ToGQShops(res.Shops),
		Total: int32(len(res.Shops)),
	}, nil
}

func (r *queryResolver) GetRestaurantMenu(ctx context.Context, id string) (*RestaurantMenu, error) {
	res, err := r.server.ProductClient.GetRestaurantMenu(ctx, &pb.GetRestaurantMenuReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQRestaurantMenu(res), nil
}

func (r *queryResolver) GetMenuItem(ctx context.Context, id string) (*MenuItem, error) {
	res, err := r.server.ProductClient.GetMenuItem(ctx, &pb.GetMenuItemReq{Id: id})
	if err != nil {
		return nil, err
	}
	return ToGQMenuItem(res), nil
}

func (r *queryResolver) GetMenuItemVariant(ctx context.Context, in GetItemVariantInput) (*ItemVariant, error) {
	res, err := r.server.ProductClient.GetMenuItemVariant(ctx, &pb.GetItemVariantReq{
		ItemId:    in.ItemID,
		VariantId: in.VariantID,
	})

	if err != nil {
		return nil, err
	}

	return ToGQItemVariant(res), err
}

func (r *queryResolver) GetMenuItemAddon(ctx context.Context, in GetItemAddonInput) (*ItemAddon, error) {
	res, err := r.server.ProductClient.GetMenuItemAddon(ctx, &pb.GetItemAddonReq{
		ItemId:  in.ItemID,
		AddonId: in.AddonID,
	})

	if err != nil {
		return nil, err
	}

	return ToGQItemAddon(res), err
}

func (r *queryResolver) GetRetailCategory(ctx context.Context, id string) (*RetailCategory, error) {
	res, err := r.server.ProductClient.GetRetailCategory(ctx, &pb.GetRetailCategoryReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQRetailCategory(res), nil
}

func (r *queryResolver) GetRetailItem(ctx context.Context, id string) (*RetailItem, error) {
	res, err := r.server.ProductClient.GetRetailItem(ctx, &pb.GetRetailItemReq{Id: id})
	if err != nil {
		return nil, err
	}
	return ToGQRetailItem(res), nil
}

func (r *queryResolver) GetRetailItemVariant(ctx context.Context, in GetItemVariantInput) (*ItemVariant, error) {
	res, err := r.server.ProductClient.GetRetailItemVariant(ctx, &pb.GetItemVariantReq{
		ItemId:    in.ItemID,
		VariantId: in.VariantID,
	})

	if err != nil {
		return nil, err
	}

	return ToGQItemVariant(res), nil
}

func (r *queryResolver) GetMedicineCategory(ctx context.Context, id string) (*MedicineCategory, error) {
	res, err := r.server.ProductClient.GetMedicineCategory(ctx, &pb.GetMedicineCategoryReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQMedicineCategory(res), nil
}

func (r *queryResolver) GetMedicineItem(ctx context.Context, id string) (*MedicineItem, error) {
	res, err := r.server.ProductClient.GetMedicineItem(ctx, &pb.GetMedicineItemReq{Id: id})
	if err != nil {
		return nil, err
	}
	return ToGQMedicineItem(res), nil
}

func (r *queryResolver) ListRestaurantMenu(ctx context.Context, shopId string) (*ListRestaurantMenuOutput, error) {
	res, err := r.server.ProductClient.ListRestaurantMenu(ctx, &pb.ListRestaurantMenuReq{ShopId: shopId})
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
	res, err := r.server.ProductClient.ListRetailCategory(ctx, &pb.ListRetailCategoryReq{ShopId: shopId})
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
	res, err := r.server.ProductClient.ListMedicineCategory(ctx, &pb.ListMedicineCategoryReq{ShopId: shopId})
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
	res, err := r.server.ProductClient.ListMenuItem(ctx, &pb.ListMenuItemReq{MenuId: menuId})
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
	res, err := r.server.ProductClient.ListMenuItemVariant(ctx, &pb.ListItemVariantReq{
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
	res, err := r.server.ProductClient.ListMenuItemAddon(ctx, &pb.ListItemAddonReq{
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
	res, err := r.server.ProductClient.ListRetailItem(ctx, &pb.ListRetailItemReq{CategoryId: categoryId})
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
	res, err := r.server.ProductClient.ListRetailItemVariant(ctx, &pb.ListItemVariantReq{
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
	res, err := r.server.ProductClient.ListMedicineItem(ctx, &pb.ListMedicineItemReq{CategoryId: categoryId})
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
	res, err := r.server.InventoryClient.GetItemStock(ctx, &pb.GetItemStockReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQItemStock(res), nil
}

func (r *queryResolver) GetVariantStock(ctx context.Context, id string) (*VariantStock, error) {
	res, err := r.server.InventoryClient.GetVariantStock(ctx, &pb.GetVariantStockReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQVariantStock(res), nil
}

func (r *queryResolver) GetAddonStock(ctx context.Context, id string) (*AddonStock, error) {
	res, err := r.server.InventoryClient.GetAddonStock(ctx, &pb.GetAddonStockReq{Id: id})
	if err != nil {
		return nil, err
	}

	return ToGQAddonStock(res), nil
}
