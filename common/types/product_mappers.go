package types

import (
	"github.com/yash91989201/superfast-delivery-api/common/pb"
)

func ToItemAddon(a *pb.ItemAddon) *ItemAddon {
	if a == nil {
		return nil
	}

	return &ItemAddon{
		ID:          HexToObjectID(a.Id),
		AddonName:   a.AddonName,
		AddonPrice:  a.AddonPrice,
		ImageURL:    a.ImageUrl,
		Description: a.Description,
		ItemID:      HexToObjectID(a.ItemId),
	}
}

func ToRestaurantMenu(rm *pb.RestaurantMenu) *RestaurantMenu {
	if rm == nil {
		return nil
	}

	return &RestaurantMenu{
		ID:        HexToObjectID(rm.Id),
		MenuName:  rm.MenuName,
		ImageURL:  rm.ImageUrl,
		ShopID:    rm.ShopId,
		CreatedAt: ToTime(rm.CreatedAt),
		UpdatedAt: ToTime(rm.UpdatedAt),
	}
}

func ToMenuItem(mi *pb.MenuItem) *MenuItem {
	if mi == nil {
		return nil
	}

	return &MenuItem{
		ID:          HexToObjectID(mi.Id),
		Name:        mi.Name,
		Price:       mi.Price,
		ImageURL:    mi.ImageUrl,
		Description: mi.Description,
		Variants:    ToItemVariants(mi.Variants),
		Addons:      ToItemAddons(mi.Addons),
		MenuID:      HexToObjectID(mi.MenuId),
		CreatedAt:   ToTime(mi.CreatedAt),
		UpdatedAt:   ToTime(mi.UpdatedAt),
	}
}

func ToRetailCategory(rc *pb.RetailCategory) *RetailCategory {
	if rc == nil {
		return nil
	}

	return &RetailCategory{
		ID:           HexToObjectID(rc.Id),
		CategoryName: rc.CategoryName,
		ImageURL:     rc.ImageUrl,
		ShopID:       rc.ShopId,
		CreatedAt:    ToTime(rc.CreatedAt),
		UpdatedAt:    ToTime(rc.UpdatedAt),
	}
}

func ToRetailItem(ri *pb.RetailItem) *RetailItem {
	if ri == nil {
		return nil
	}

	return &RetailItem{
		ID:          HexToObjectID(ri.Id),
		Name:        ri.Name,
		Price:       ri.Price,
		ImageURL:    ri.ImageUrl,
		Description: ri.Description,
		CategoryID:  HexToObjectID(ri.CategoryId),
		Variants:    ToItemVariants(ri.Variants),
		CreatedAt:   ToTime(ri.CreatedAt),
		UpdatedAt:   ToTime(ri.UpdatedAt),
	}
}

func ToMedicineCategory(mc *pb.MedicineCategory) *MedicineCategory {
	if mc == nil {
		return nil
	}

	return &MedicineCategory{
		ID:           HexToObjectID(mc.Id),
		CategoryName: mc.CategoryName,
		ImageURL:     mc.ImageUrl,
		ShopID:       mc.ShopId,
		CreatedAt:    ToTime(mc.CreatedAt),
		UpdatedAt:    ToTime(mc.UpdatedAt),
	}
}

func ToMedicineItem(mi *pb.MedicineItem) *MedicineItem {
	if mi == nil {
		return nil
	}

	return &MedicineItem{
		ID:          HexToObjectID(mi.Id),
		Name:        mi.Name,
		Price:       mi.Price,
		ImageURL:    mi.ImageUrl,
		Description: mi.Description,
		CategoryID:  HexToObjectID(mi.CategoryId),
		CreatedAt:   ToTime(mi.CreatedAt),
		UpdatedAt:   ToTime(mi.UpdatedAt),
	}
}

func ToItemVariants(variants []*pb.ItemVariant) []*ItemVariant {
	itemVariants := make([]*ItemVariant, len(variants))
	for i, v := range variants {
		if v != nil {
			itemVariants[i] = &ItemVariant{
				ID:              HexToObjectID(v.Id),
				VariantName:     v.VariantName,
				RelativePrice:   v.RelativePrice,
				RelativePricing: v.RelativePricing,
				Price:           v.Price,
				ImageURL:        v.ImageUrl,
				Description:     v.Description,
				ItemID:          HexToObjectID(v.ItemId),
			}
		}
	}

	return itemVariants
}

func ToItemAddons(addons []*pb.ItemAddon) []*ItemAddon {
	itemAddons := make([]*ItemAddon, len(addons))
	for i, a := range addons {
		if a != nil {
			itemAddons[i] = &ItemAddon{
				ID:          HexToObjectID(a.Id),
				AddonName:   a.AddonName,
				AddonPrice:  a.AddonPrice,
				ImageURL:    a.ImageUrl,
				Description: a.Description,
				ItemID:      HexToObjectID(a.ItemId),
			}
		}
	}

	return itemAddons
}

func ToCreateItemVariant(iv *pb.CreateItemVariantReq) *CreateItemVariant {
	if iv == nil {
		return nil
	}

	return &CreateItemVariant{
		VariantName:     iv.VariantName,
		RelativePrice:   iv.RelativePrice,
		RelativePricing: iv.RelativePricing,
		Price:           iv.Price,
		ImageURL:        iv.ImageUrl,
		Description:     iv.Description,
		ItemID:          HexToObjectID(iv.ItemId),
	}
}

func ToCreateItemAddon(ia *pb.CreateItemAddonReq) *CreateItemAddon {
	if ia == nil {
		return nil
	}

	return &CreateItemAddon{
		AddonName:   ia.AddonName,
		AddonPrice:  ia.AddonPrice,
		ImageURL:    ia.ImageUrl,
		Description: ia.Description,
		ItemID:      HexToObjectID(ia.ItemId),
	}
}

func ToCreateRestaurantMenu(rm *pb.CreateRestaurantMenuReq) *CreateRestaurantMenu {
	return &CreateRestaurantMenu{
		MenuName: rm.MenuName,
		ImageURL: rm.ImageUrl,
		ShopID:   rm.ShopId,
	}
}

func ToCreateMenuItem(mi *pb.CreateMenuItemReq) *CreateMenuItem {
	return &CreateMenuItem{
		Name:        mi.Name,
		ImageUrl:    mi.ImageUrl,
		Description: mi.Description,
		Price:       mi.Price,
		MenuID:      HexToObjectID(mi.MenuId),
	}
}

func ToPbMenuItems(mi []*MenuItem) []*pb.MenuItem {
	menuItems := make([]*pb.MenuItem, len(mi))
	for i, m := range mi {
		if m != nil {
			menuItems[i] = ToPbMenuItem(m)
		}
	}

	return menuItems
}

func ToPbRetailItems(il []*RetailItem) []*pb.RetailItem {
	list := make([]*pb.RetailItem, len(il))
	for i, item := range il {
		if item != nil {
			list[i] = ToPbRetailItem(item)
		}
	}

	return list
}

func ToPbMedicineItem(mi *MedicineItem) *pb.MedicineItem {
	return &pb.MedicineItem{
		Id:          mi.ID.Hex(),
		Name:        mi.Name,
		ImageUrl:    mi.ImageURL,
		Description: mi.Description,
		Price:       mi.Price,
		CategoryId:  mi.CategoryID.Hex(),
		CreatedAt:   ToPbTimestamp(mi.CreatedAt),
		UpdatedAt:   ToPbTimestamp(mi.CreatedAt),
	}
}

func ToPbMedicineItems(mil []*MedicineItem) []*pb.MedicineItem {
	list := make([]*pb.MedicineItem, len(mil))
	for i, item := range mil {
		if item != nil {
			list[i] = ToPbMedicineItem(item)
		}
	}

	return list
}

func ToPbItemVariant(v *ItemVariant) *pb.ItemVariant {
	if v == nil {
		return nil
	}

	return &pb.ItemVariant{
		Id:              v.ID.Hex(),
		VariantName:     v.VariantName,
		RelativePrice:   v.RelativePrice,
		RelativePricing: v.RelativePricing,
		Price:           v.Price,
		ImageUrl:        v.ImageURL,
		Description:     v.Description,
		ItemId:          v.ItemID.Hex(),
	}
}

func ToPbItemVariants(variants []*ItemVariant) []*pb.ItemVariant {
	pbVariants := make([]*pb.ItemVariant, len(variants))
	for i, v := range variants {
		pbVariants[i] = ToPbItemVariant(v)
	}

	return pbVariants
}

func ToPbItemAddon(a *ItemAddon) *pb.ItemAddon {
	if a == nil {
		return nil
	}

	return &pb.ItemAddon{
		Id:          a.ID.Hex(),
		AddonName:   a.AddonName,
		AddonPrice:  a.AddonPrice,
		ImageUrl:    a.ImageURL,
		Description: a.Description,
		ItemId:      a.ItemID.Hex(),
	}
}

func ToPbItemAddons(addons []*ItemAddon) []*pb.ItemAddon {
	pbAddons := make([]*pb.ItemAddon, len(addons))
	for i, a := range addons {
		pbAddons[i] = ToPbItemAddon(a)
	}

	return pbAddons
}

func ToPbRestaurantMenu(rm *RestaurantMenu) *pb.RestaurantMenu {
	if rm == nil {
		return nil
	}

	return &pb.RestaurantMenu{
		Id:        rm.ID.Hex(),
		MenuName:  rm.MenuName,
		ImageUrl:  rm.ImageURL,
		ShopId:    rm.ShopID,
		CreatedAt: ToPbTimestamp(rm.CreatedAt),
		UpdatedAt: ToPbTimestamp(rm.UpdatedAt),
	}
}

func ToPbRestaurantMenuList(rmList []*RestaurantMenu) []*pb.RestaurantMenu {
	pbRmList := make([]*pb.RestaurantMenu, len(rmList))
	for i, rm := range rmList {
		pbRmList[i] = ToPbRestaurantMenu(rm)
	}

	return pbRmList
}

func ToPbMenuItem(mi *MenuItem) *pb.MenuItem {
	if mi == nil {
		return nil
	}

	return &pb.MenuItem{
		Id:          mi.ID.Hex(),
		Name:        mi.Name,
		Price:       mi.Price,
		ImageUrl:    mi.ImageURL,
		Description: mi.Description,
		Variants:    ToPbItemVariants(mi.Variants),
		Addons:      ToPbItemAddons(mi.Addons),
		MenuId:      mi.MenuID.Hex(),
		CreatedAt:   ToPbTimestamp(mi.CreatedAt),
		UpdatedAt:   ToPbTimestamp(mi.UpdatedAt),
	}
}

func ToPbMenuItemList(miList []*MenuItem) []*pb.MenuItem {
	pbMiList := make([]*pb.MenuItem, len(miList))
	for i, mi := range miList {
		pbMiList[i] = ToPbMenuItem(mi)
	}

	return pbMiList
}

func ToPbRetailCategory(rc *RetailCategory) *pb.RetailCategory {
	if rc == nil {
		return nil
	}

	return &pb.RetailCategory{
		Id:           rc.ID.Hex(),
		CategoryName: rc.CategoryName,
		ImageUrl:     rc.ImageURL,
		ShopId:       rc.ShopID,
		CreatedAt:    ToPbTimestamp(rc.CreatedAt),
		UpdatedAt:    ToPbTimestamp(rc.UpdatedAt),
	}
}

func ToPbRetailCategoryList(rcList []*RetailCategory) []*pb.RetailCategory {
	pbRcList := make([]*pb.RetailCategory, len(rcList))
	for i, rc := range rcList {
		pbRcList[i] = ToPbRetailCategory(rc)
	}

	return pbRcList
}

func ToPbRetailItem(ri *RetailItem) *pb.RetailItem {
	if ri == nil {
		return nil
	}

	return &pb.RetailItem{
		Id:          ri.ID.Hex(),
		Name:        ri.Name,
		Price:       ri.Price,
		ImageUrl:    ri.ImageURL,
		Description: ri.Description,
		CategoryId:  ri.CategoryID.Hex(),
		Variants:    ToPbItemVariants(ri.Variants),
		CreatedAt:   ToPbTimestamp(ri.CreatedAt),
		UpdatedAt:   ToPbTimestamp(ri.UpdatedAt),
	}
}

func ToPbRetailItemList(riList []*RetailItem) []*pb.RetailItem {
	pbRiList := make([]*pb.RetailItem, len(riList))
	for i, ri := range riList {
		pbRiList[i] = ToPbRetailItem(ri)
	}

	return pbRiList
}

func ToPbMedicineCategory(mc *MedicineCategory) *pb.MedicineCategory {
	if mc == nil {
		return nil
	}

	return &pb.MedicineCategory{
		Id:           mc.ID.Hex(),
		CategoryName: mc.CategoryName,
		ImageUrl:     mc.ImageURL,
		ShopId:       mc.ShopID,
		CreatedAt:    ToPbTimestamp(mc.CreatedAt),
		UpdatedAt:    ToPbTimestamp(mc.UpdatedAt),
	}
}

func ToPbMedicineCategoryList(mcList []*MedicineCategory) []*pb.MedicineCategory {
	pbMcList := make([]*pb.MedicineCategory, len(mcList))
	for i, mc := range mcList {
		pbMcList[i] = ToPbMedicineCategory(mc)
	}

	return pbMcList
}

func ToPbMedicineItemList(miList []*MedicineItem) []*pb.MedicineItem {
	pbMiList := make([]*pb.MedicineItem, len(miList))
	for i, mi := range miList {
		pbMiList[i] = ToPbMedicineItem(mi)
	}

	return pbMiList
}

func ToUpdateItemVariant(req *pb.UpdateItemVariantReq) *UpdateItemVariant {
	if req == nil {
		return nil
	}

	return &UpdateItemVariant{
		ID:              HexToObjectID(req.Id),
		VariantName:     req.VariantName,
		RelativePricing: req.RelativePricing,
		RelativePrice:   req.RelativePrice,
		Price:           req.Price,
		ImageURL:        req.ImageUrl,
		Description:     req.Description,
		ItemID:          HexToObjectID(req.ItemId),
	}
}

func ToUpdateItemAddon(req *pb.UpdateItemAddonReq) *UpdateItemAddon {
	if req == nil {
		return nil
	}

	return &UpdateItemAddon{
		ID:          HexToObjectID(req.Id),
		AddonName:   req.AddonName,
		AddonPrice:  req.AddonPrice,
		ImageURL:    req.ImageUrl,
		Description: req.Description,
		ItemID:      HexToObjectID(req.ItemId),
	}
}

func ToUpdateRestaurantMenu(req *pb.UpdateRestaurantMenuReq) *UpdateRestaurantMenu {
	if req == nil {
		return nil
	}

	return &UpdateRestaurantMenu{
		ID:       HexToObjectID(req.Id),
		MenuName: req.MenuName,
		ImageURL: req.ImageUrl,
	}
}

func ToUpdateMenuItem(req *pb.UpdateMenuItemReq) *UpdateMenuItem {
	if req == nil {
		return nil
	}

	return &UpdateMenuItem{
		ID:          HexToObjectID(req.Id),
		Name:        req.Name,
		Price:       req.Price,
		ImageURL:    req.ImageUrl,
		Description: req.Description,
	}
}

func ToUpdateRetailCategory(req *pb.UpdateRetailCategoryReq) *UpdateRetailCategory {
	if req == nil {
		return nil
	}

	return &UpdateRetailCategory{
		ID:           HexToObjectID(req.Id),
		CategoryName: req.CategoryName,
		ImageURL:     req.ImageUrl,
	}
}

func ToUpdateRetailItem(req *pb.UpdateRetailItemReq) *UpdateRetailItem {
	if req == nil {
		return nil
	}

	return &UpdateRetailItem{
		ID:          HexToObjectID(req.Id),
		Name:        req.Name,
		Price:       req.Price,
		ImageURL:    req.ImageUrl,
		Description: req.Description,
	}
}

func ToUpdateMedicineCategory(req *pb.UpdateMedicineCategoryReq) *UpdateMedicineCategory {
	if req == nil {
		return nil
	}

	return &UpdateMedicineCategory{
		ID:           HexToObjectID(req.Id),
		CategoryName: req.CategoryName,
		ImageURL:     req.ImageUrl,
	}

}

func ToUpdateMedicineItem(req *pb.UpdateMedicineItemReq) *UpdateMedicineItem {
	if req == nil {
		return nil
	}

	return &UpdateMedicineItem{
		ID:          HexToObjectID(req.Id),
		Name:        req.Name,
		Price:       req.Price,
		ImageURL:    req.ImageUrl,
		Description: req.Description,
	}
}
