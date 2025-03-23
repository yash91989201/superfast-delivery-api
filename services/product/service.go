package product

import (
	"context"
	"time"

	"github.com/yash91989201/superfast-delivery-api/common/types"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Service interface {
	CreateRestaurantMenu(ctx context.Context, rm *types.CreateRestaurantMenu) (*types.RestaurantMenu, error)
	CreateMenuItem(ctx context.Context, rm *types.CreateMenuItem) (*types.MenuItem, error)
	CreateMenuItemVariant(ctx context.Context, iv *types.CreateItemVariant) (*types.ItemVariant, error)
	CreateMenuItemAddon(ctx context.Context, ia *types.CreateItemAddon) (*types.ItemAddon, error)
	CreateRetailCategory(ctx context.Context, rc *types.CreateRetailCategory) (*types.RetailCategory, error)
	CreateRetailItem(ctx context.Context, ri *types.CreateRetailItem) (*types.RetailItem, error)
	CreateRetailItemVariant(ctx context.Context, iv *types.CreateItemVariant) (*types.ItemVariant, error)
	CreateMedicineCategory(ctx context.Context, mc *types.CreateMedicineCategory) (*types.MedicineCategory, error)
	CreateMedicineItem(ctx context.Context, mi *types.CreateMedicineItem) (*types.MedicineItem, error)

	GetRestaurantMenu(ctx context.Context, id string) (*types.RestaurantMenu, error)
	GetMenuItem(ctx context.Context, id string) (*types.MenuItem, error)
	GetMenuItemVariant(ctx context.Context, itemID string, variantID string) (*types.ItemVariant, error)
	GetMenuItemAddon(ctx context.Context, itemId string, addonID string) (*types.ItemAddon, error)
	GetRetailCategory(ctx context.Context, id string) (*types.RetailCategory, error)
	GetRetailItem(ctx context.Context, id string) (*types.RetailItem, error)
	GetRetailItemVariant(ctx context.Context, itemID string, variantID string) (*types.ItemVariant, error)
	GetMedicineCategory(ctx context.Context, id string) (*types.MedicineCategory, error)
	GetMedicineItem(ctx context.Context, id string) (*types.MedicineItem, error)

	ListRestaurantMenu(ctx context.Context, shopID string) ([]*types.RestaurantMenu, error)
	ListMenuItem(ctx context.Context, menuID string) ([]*types.MenuItem, error)
	ListMenuItemVariant(ctx context.Context, itemID string) ([]*types.ItemVariant, error)
	ListMenuItemAddon(ctx context.Context, itemID string) ([]*types.ItemAddon, error)
	ListRetailCategory(ctx context.Context, shopID string) ([]*types.RetailCategory, error)
	ListRetailItem(ctx context.Context, categoryID string) ([]*types.RetailItem, error)
	ListRetailItemVariant(ctx context.Context, itemID string) ([]*types.ItemVariant, error)
	ListMedicineCategory(ctx context.Context, shopID string) ([]*types.MedicineCategory, error)
	ListMedicineItem(ctx context.Context, categoryID string) ([]*types.MedicineItem, error)

	UpdateRestaurantMenu(ctx context.Context, menu *types.UpdateRestaurantMenu) error
	UpdateMenuItem(ctx context.Context, item *types.UpdateMenuItem) error
	UpdateMenuItemVariant(ctx context.Context, variant *types.UpdateItemVariant) error
	UpdateMenuItemAddon(ctx context.Context, addon *types.UpdateItemAddon) error
	UpdateRetailCategory(ctx context.Context, category *types.UpdateRetailCategory) error
	UpdateRetailItem(ctx context.Context, item *types.UpdateRetailItem) error
	UpdateRetailItemVariant(ctx context.Context, variant *types.UpdateItemVariant) error
	UpdateMedicineCategory(ctx context.Context, category *types.UpdateMedicineCategory) error
	UpdateMedicineItem(ctx context.Context, item *types.UpdateMedicineItem) error

	DeleteRestaurantMenu(ctx context.Context, id string) error
	DeleteMenuItem(ctx context.Context, id string) error
	DeleteMenuItemVariant(ctx context.Context, itemID string, variantID string) error
	DeleteMenuItemAddon(ctx context.Context, itemID string, addonID string) error
	DeleteRetailCategory(ctx context.Context, id string) error
	DeleteRetailItem(ctx context.Context, id string) error
	DeleteRetailItemVariant(ctx context.Context, itemID string, variantID string) error
	DeleteMedicineCategory(ctx context.Context, id string) error
	DeleteMedicineItem(ctx context.Context, id string) error
}

type productService struct {
	r Repository
}

func New(r Repository) Service {
	return &productService{r: r}
}

func (s *productService) CreateRestaurantMenu(ctx context.Context, rm *types.CreateRestaurantMenu) (*types.RestaurantMenu, error) {
	newRestaurantMenu := &types.RestaurantMenu{
		ID:        bson.NewObjectID(),
		MenuName:  rm.MenuName,
		ImageURL:  rm.ImageURL,
		ShopID:    rm.ShopID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.r.CreateRestaurantMenu(ctx, newRestaurantMenu); err != nil {
		return nil, err
	}

	return newRestaurantMenu, nil
}

func (s *productService) CreateMenuItem(ctx context.Context, mi *types.CreateMenuItem) (*types.MenuItem, error) {
	newMenuItem := &types.MenuItem{
		ID:          bson.NewObjectID(),
		Name:        mi.Name,
		ImageURL:    mi.ImageUrl,
		Description: mi.Description,
		Price:       mi.Price,
		MenuID:      mi.MenuID,
		Variants:    []*types.ItemVariant{},
		Addons:      []*types.ItemAddon{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.r.CreateMenuItem(ctx, newMenuItem); err != nil {
		return nil, err
	}

	return newMenuItem, nil
}

func (s *productService) CreateMenuItemVariant(ctx context.Context, iv *types.CreateItemVariant) (*types.ItemVariant, error) {
	itemVariant := &types.ItemVariant{
		ID:              bson.NewObjectID(),
		VariantName:     iv.VariantName,
		RelativePricing: iv.RelativePricing,
		RelativePrice:   iv.RelativePrice,
		Price:           iv.Price,
		ImageURL:        iv.ImageURL,
		Description:     iv.Description,
		ItemID:          iv.ItemID,
	}

	if err := s.r.CreateMenuItemVariant(ctx, itemVariant); err != nil {
		return nil, err
	}

	return itemVariant, nil
}

func (s *productService) CreateMenuItemAddon(ctx context.Context, ia *types.CreateItemAddon) (*types.ItemAddon, error) {
	itemAddon := &types.ItemAddon{
		ID:          bson.NewObjectID(),
		AddonName:   ia.AddonName,
		AddonPrice:  ia.AddonPrice,
		ImageURL:    ia.ImageURL,
		Description: ia.Description,
		ItemID:      ia.ItemID,
	}

	if err := s.r.CreateMenuItemAddon(ctx, itemAddon); err != nil {
		return nil, err
	}

	return itemAddon, nil
}

func (s *productService) CreateRetailCategory(ctx context.Context, rc *types.CreateRetailCategory) (*types.RetailCategory, error) {
	retailCategory := &types.RetailCategory{
		ID:           bson.NewObjectID(),
		CategoryName: rc.CategoryName,
		ImageURL:     rc.ImageURL,
		ShopID:       rc.ShopID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := s.r.CreateRetailCategory(ctx, retailCategory)
	if err != nil {
		return nil, err
	}

	return retailCategory, nil
}

func (s *productService) CreateRetailItem(ctx context.Context, ri *types.CreateRetailItem) (*types.RetailItem, error) {
	retailItem := &types.RetailItem{
		ID:          bson.NewObjectID(),
		Name:        ri.Name,
		ImageURL:    ri.ImageURL,
		Description: ri.Description,
		Price:       ri.Price,
		CategoryID:  ri.CategoryID,
		Variants:    []*types.ItemVariant{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := s.r.CreateRetailItem(ctx, retailItem)
	if err != nil {
		return nil, err
	}

	return retailItem, nil
}

func (s *productService) CreateRetailItemVariant(ctx context.Context, iv *types.CreateItemVariant) (*types.ItemVariant, error) {
	itemVariant := &types.ItemVariant{
		ID:              bson.NewObjectID(),
		VariantName:     iv.VariantName,
		RelativePricing: iv.RelativePricing,
		RelativePrice:   iv.RelativePrice,
		Price:           iv.Price,
		ImageURL:        iv.ImageURL,
		Description:     iv.Description,
		ItemID:          iv.ItemID,
	}

	if err := s.r.CreateRetailItemVariant(ctx, itemVariant); err != nil {
		return nil, err
	}

	return itemVariant, nil
}

func (s *productService) CreateMedicineCategory(ctx context.Context, mc *types.CreateMedicineCategory) (*types.MedicineCategory, error) {
	medicineCategory := &types.MedicineCategory{
		ID:           bson.NewObjectID(),
		CategoryName: mc.CategoryName,
		ImageURL:     mc.ImageURL,
		ShopID:       mc.ShopID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := s.r.CreateMedicineCategory(ctx, medicineCategory)
	if err != nil {
		return nil, err
	}

	return medicineCategory, nil
}

func (s *productService) CreateMedicineItem(ctx context.Context, mi *types.CreateMedicineItem) (*types.MedicineItem, error) {
	medicineItem := &types.MedicineItem{
		ID:          bson.NewObjectID(),
		Name:        mi.Name,
		Price:       mi.Price,
		Description: mi.Description,
		ImageURL:    mi.ImageURL,
		CategoryID:  mi.CategoryID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := s.r.CreateMedicineItem(ctx, medicineItem)
	if err != nil {
		return nil, err
	}

	return medicineItem, nil
}

func (s *productService) GetRestaurantMenu(ctx context.Context, id string) (*types.RestaurantMenu, error) {
	return s.r.GetRestaurantMenu(ctx, types.HexToObjectID(id))
}

func (s *productService) GetMenuItem(ctx context.Context, id string) (*types.MenuItem, error) {
	return s.r.GetMenuItem(ctx, types.HexToObjectID(id))
}

func (s *productService) GetMenuItemVariant(ctx context.Context, itemID string, variantID string) (*types.ItemVariant, error) {
	return s.r.GetMenuItemVariant(ctx, types.HexToObjectID(itemID), types.HexToObjectID(variantID))
}

func (s *productService) GetMenuItemAddon(ctx context.Context, itemID string, addonID string) (*types.ItemAddon, error) {
	return s.r.GetMenuItemAddon(ctx, types.HexToObjectID(itemID), types.HexToObjectID(addonID))
}

func (s *productService) GetRetailCategory(ctx context.Context, id string) (*types.RetailCategory, error) {
	return s.r.GetRetailCategory(ctx, types.HexToObjectID(id))
}

func (s *productService) GetRetailItem(ctx context.Context, id string) (*types.RetailItem, error) {
	return s.r.GetRetailItem(ctx, types.HexToObjectID(id))
}

func (s *productService) GetRetailItemVariant(ctx context.Context, itemID string, variantID string) (*types.ItemVariant, error) {
	return s.r.GetRetailItemVariant(ctx, types.HexToObjectID(itemID), types.HexToObjectID(variantID))
}

func (s *productService) GetMedicineCategory(ctx context.Context, id string) (*types.MedicineCategory, error) {
	return s.r.GetMedicineCategory(ctx, types.HexToObjectID(id))
}

func (s *productService) GetMedicineItem(ctx context.Context, id string) (*types.MedicineItem, error) {
	return s.r.GetMedicineItem(ctx, types.HexToObjectID(id))
}

func (s *productService) ListRestaurantMenu(ctx context.Context, shopID string) ([]*types.RestaurantMenu, error) {
	return s.r.ListRestaurantMenu(ctx, shopID)
}

func (s *productService) ListMenuItem(ctx context.Context, menuID string) ([]*types.MenuItem, error) {
	return s.r.ListMenuItem(ctx, types.HexToObjectID(menuID))
}

func (s *productService) ListMenuItemVariant(ctx context.Context, itemID string) ([]*types.ItemVariant, error) {
	return s.r.ListMenuItemVariant(ctx, types.HexToObjectID(itemID))
}

func (s *productService) ListMenuItemAddon(ctx context.Context, itemID string) ([]*types.ItemAddon, error) {
	return s.r.ListMenuItemAddon(ctx, types.HexToObjectID(itemID))
}

func (s *productService) ListRetailCategory(ctx context.Context, shopID string) ([]*types.RetailCategory, error) {
	return s.r.ListRetailCategory(ctx, shopID)
}

func (s *productService) ListRetailItem(ctx context.Context, categoryID string) ([]*types.RetailItem, error) {
	return s.r.ListRetailItem(ctx, types.HexToObjectID(categoryID))
}

func (s *productService) ListRetailItemVariant(ctx context.Context, itemID string) ([]*types.ItemVariant, error) {
	return s.r.ListRetailItemVariant(ctx, types.HexToObjectID(itemID))
}

func (s *productService) ListMedicineCategory(ctx context.Context, shopID string) ([]*types.MedicineCategory, error) {
	return s.r.ListMedicineCategory(ctx, shopID)
}

func (s *productService) ListMedicineItem(ctx context.Context, categoryID string) ([]*types.MedicineItem, error) {
	return s.r.ListMedicineItem(ctx, types.HexToObjectID(categoryID))
}

func (s *productService) UpdateRestaurantMenu(ctx context.Context, menu *types.UpdateRestaurantMenu) error {
	return s.r.UpdateRestaurantMenu(ctx, menu)
}

func (s *productService) UpdateMenuItem(ctx context.Context, item *types.UpdateMenuItem) error {
	return s.r.UpdateMenuItem(ctx, item)
}

func (s *productService) UpdateMenuItemVariant(ctx context.Context, variant *types.UpdateItemVariant) error {
	return s.r.UpdateMenuItemVariant(ctx, variant)
}

func (s *productService) UpdateMenuItemAddon(ctx context.Context, addon *types.UpdateItemAddon) error {
	return s.r.UpdateMenuItemAddon(ctx, addon)
}

func (s *productService) UpdateRetailCategory(ctx context.Context, category *types.UpdateRetailCategory) error {
	return s.r.UpdateRetailCategory(ctx, category)
}

func (s *productService) UpdateRetailItem(ctx context.Context, item *types.UpdateRetailItem) error {
	return s.r.UpdateRetailItem(ctx, item)
}

func (s *productService) UpdateRetailItemVariant(ctx context.Context, variant *types.UpdateItemVariant) error {
	return s.r.UpdateRetailItemVariant(ctx, variant)
}

func (s *productService) UpdateMedicineCategory(ctx context.Context, category *types.UpdateMedicineCategory) error {
	return s.r.UpdateMedicineCategory(ctx, category)
}

func (s *productService) UpdateMedicineItem(ctx context.Context, item *types.UpdateMedicineItem) error {
	return s.r.UpdateMedicineItem(ctx, item)
}

func (s *productService) DeleteRestaurantMenu(ctx context.Context, id string) error {
	return s.r.DeleteRestaurantMenu(ctx, types.HexToObjectID(id))
}

func (s *productService) DeleteMenuItem(ctx context.Context, id string) error {
	return s.r.DeleteMenuItem(ctx, types.HexToObjectID(id))
}

func (s *productService) DeleteMenuItemVariant(ctx context.Context, itemID string, variantID string) error {
	return s.r.DeleteMenuItemVariant(ctx, types.HexToObjectID(itemID), types.HexToObjectID(variantID))
}

func (s *productService) DeleteMenuItemAddon(ctx context.Context, itemID string, addonID string) error {
	return s.r.DeleteMenuItemAddon(ctx, types.HexToObjectID(itemID), types.HexToObjectID(addonID))
}

func (s *productService) DeleteRetailCategory(ctx context.Context, id string) error {
	return s.r.DeleteRetailCategory(ctx, types.HexToObjectID(id))
}

func (s *productService) DeleteRetailItem(ctx context.Context, id string) error {
	return s.r.DeleteRetailItem(ctx, types.HexToObjectID(id))
}

func (s *productService) DeleteRetailItemVariant(ctx context.Context, itemID string, variantID string) error {
	return s.r.DeleteRetailItemVariant(ctx, types.HexToObjectID(itemID), types.HexToObjectID(variantID))
}

func (s *productService) DeleteMedicineCategory(ctx context.Context, id string) error {
	return s.r.DeleteMedicineCategory(ctx, types.HexToObjectID(id))
}

func (s *productService) DeleteMedicineItem(ctx context.Context, id string) error {
	return s.r.DeleteMedicineItem(ctx, types.HexToObjectID(id))
}
