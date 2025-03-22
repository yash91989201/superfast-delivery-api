package product

import (
	"context"
	"time"

	"github.com/yash91989201/superfast-delivery-api/common/types"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Service interface {
	CreateItemVariant(ctx context.Context, iv *types.CreateItemVariant) (*types.ItemVariant, error)
	CreateItemAddon(ctx context.Context, ia *types.CreateItemAddon) (*types.ItemAddon, error)
	CreateRestaurantMenu(ctx context.Context, rm *types.CreateRestaurantMenu) (*types.RestaurantMenu, error)
	CreateMenuItem(ctx context.Context, rm *types.CreateMenuItem) (*types.MenuItem, error)
	CreateRetailCategory(ctx context.Context, rc *types.CreateRetailCategory) (*types.RetailCategory, error)
	CreateRetailItem(ctx context.Context, ri *types.CreateRetailItem) (*types.RetailItem, error)
	CreateMedicineCategory(ctx context.Context, mc *types.CreateMedicineCategory) (*types.MedicineCategory, error)
	CreateMedicineItem(ctx context.Context, mi *types.CreateMedicineItem) (*types.MedicineItem, error)

	GetItemVariant(ctx context.Context, id string) (*types.ItemVariant, error)
	GetItemAddon(ctx context.Context, id string) (*types.ItemAddon, error)
	GetRestaurantMenu(ctx context.Context, id string) (*types.RestaurantMenu, error)
	GetMenuItem(ctx context.Context, id string) (*types.MenuItem, error)
	GetRetailCategory(ctx context.Context, id string) (*types.RetailCategory, error)
	GetRetailItem(ctx context.Context, id string) (*types.RetailItem, error)
	GetMedicineCategory(ctx context.Context, id string) (*types.MedicineCategory, error)
	GetMedicineItem(ctx context.Context, id string) (*types.MedicineItem, error)

	ListItemVariant(ctx context.Context, itemID string) ([]*types.ItemVariant, error)
	ListItemAddon(ctx context.Context, itemID string) ([]*types.ItemAddon, error)
	ListRestaurantMenu(ctx context.Context, shopID string) ([]*types.RestaurantMenu, error)
	ListMenuItem(ctx context.Context, menuID string) ([]*types.MenuItem, error)
	ListRetailCategory(ctx context.Context, shopID string) ([]*types.RetailCategory, error)
	ListRetailItem(ctx context.Context, categoryID string) ([]*types.RetailItem, error)
	ListMedicineCategory(ctx context.Context, shopID string) ([]*types.MedicineCategory, error)
	ListMedicineItem(ctx context.Context, categoryID string) ([]*types.MedicineItem, error)

	UpdateItemVariant(ctx context.Context, variant *types.UpdateItemVariant) error
	UpdateItemAddon(ctx context.Context, addon *types.UpdateItemAddon) error
	UpdateRestaurantMenu(ctx context.Context, menu *types.UpdateRestaurantMenu) error
	UpdateMenuItem(ctx context.Context, item *types.UpdateMenuItem) error
	UpdateRetailCategory(ctx context.Context, category *types.UpdateRetailCategory) error
	UpdateRetailItem(ctx context.Context, item *types.UpdateRetailItem) error
	UpdateMedicineCategory(ctx context.Context, category *types.UpdateMedicineCategory) error
	UpdateMedicineItem(ctx context.Context, item *types.UpdateMedicineItem) error

	DeleteItemVariant(ctx context.Context, id string) error
	DeleteItemAddon(ctx context.Context, id string) error
	DeleteRestaurantMenu(ctx context.Context, id string) error
	DeleteMenuItem(ctx context.Context, id string) error
	DeleteRetailCategory(ctx context.Context, id string) error
	DeleteRetailItem(ctx context.Context, id string) error
	DeleteMedicineCategory(ctx context.Context, id string) error
	DeleteMedicineItem(ctx context.Context, id string) error
}

type productService struct {
	r Repository
}

func New(r Repository) Service {
	return &productService{r: r}
}

func (s *productService) CreateItemVariant(ctx context.Context, v *types.CreateItemVariant) (*types.ItemVariant, error) {
	newVariant := &types.ItemVariant{
		ID:              bson.NewObjectID(),
		VariantName:     v.VariantName,
		RelativePrice:   v.RelativePrice,
		RelativePricing: v.RelativePricing,
		Price:           v.Price,
		Description:     v.Description,
		ItemID:          v.ItemID,
	}

	if err := s.r.CreateItemVariant(ctx, newVariant); err != nil {
		return nil, err
	}

	return newVariant, nil
}

func (s *productService) CreateItemAddon(ctx context.Context, a *types.CreateItemAddon) (*types.ItemAddon, error) {
	newAddon := &types.ItemAddon{
		ID:          bson.NewObjectID(),
		AddonName:   a.AddonName,
		AddonPrice:  a.AddonPrice,
		Description: a.Description,
		ItemID:      a.ItemID,
	}

	if err := s.r.CreateItemAddon(ctx, newAddon); err != nil {
		return nil, err
	}

	return newAddon, nil
}

func (s *productService) CreateRestaurantMenu(ctx context.Context, rm *types.CreateRestaurantMenu) (*types.RestaurantMenu, error) {
	newRestaurantMenu := &types.RestaurantMenu{
		ID:        bson.NewObjectID(),
		MenuName:  rm.MenuName,
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
		Description: mi.Description,
		Price:       mi.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		MenuID:      mi.MenuID,
	}

	if err := s.r.CreateMenuItem(ctx, newMenuItem); err != nil {
		return nil, err
	}

	return newMenuItem, nil
}

func (s *productService) CreateMedicineCategory(ctx context.Context, mc *types.CreateMedicineCategory) (*types.MedicineCategory, error) {
	medicineCategory := &types.MedicineCategory{
		ID:           bson.NewObjectID(),
		CategoryName: mc.CategoryName,
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

func (s *productService) CreateRetailCategory(ctx context.Context, rc *types.CreateRetailCategory) (*types.RetailCategory, error) {
	retailCategory := &types.RetailCategory{
		ID:           bson.NewObjectID(),
		CategoryName: rc.CategoryName,
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
		Description: ri.Description,
		Price:       ri.Price,
		CategoryID:  ri.CategoryID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := s.r.CreateRetailItem(ctx, retailItem)
	if err != nil {
		return nil, err
	}

	return retailItem, nil
}

func (s *productService) GetItemVariant(ctx context.Context, id string) (*types.ItemVariant, error) {
	return s.r.GetItemVariant(ctx, types.HexToObjectID(id))
}

func (s *productService) GetItemAddon(ctx context.Context, id string) (*types.ItemAddon, error) {
	return s.r.GetItemAddon(ctx, types.HexToObjectID(id))
}

func (s *productService) GetRestaurantMenu(ctx context.Context, id string) (*types.RestaurantMenu, error) {
	return s.r.GetRestaurantMenu(ctx, types.HexToObjectID(id))
}

func (s *productService) GetMenuItem(ctx context.Context, id string) (*types.MenuItem, error) {
	return s.r.GetMenuItem(ctx, types.HexToObjectID(id))
}

func (s *productService) GetRetailCategory(ctx context.Context, id string) (*types.RetailCategory, error) {
	return s.r.GetRetailCategory(ctx, types.HexToObjectID(id))
}

func (s *productService) GetRetailItem(ctx context.Context, id string) (*types.RetailItem, error) {
	return s.r.GetRetailItem(ctx, types.HexToObjectID(id))
}

func (s *productService) GetMedicineCategory(ctx context.Context, id string) (*types.MedicineCategory, error) {
	return s.r.GetMedicineCategory(ctx, types.HexToObjectID(id))
}

func (s *productService) GetMedicineItem(ctx context.Context, id string) (*types.MedicineItem, error) {
	return s.r.GetMedicineItem(ctx, types.HexToObjectID(id))
}

func (s *productService) ListItemVariant(ctx context.Context, itemID string) ([]*types.ItemVariant, error) {
	return s.r.ListItemVariant(ctx, types.HexToObjectID(itemID))
}

func (s *productService) ListItemAddon(ctx context.Context, itemID string) ([]*types.ItemAddon, error) {
	return s.r.ListItemAddon(ctx, types.HexToObjectID(itemID))
}

func (s *productService) ListRestaurantMenu(ctx context.Context, shopID string) ([]*types.RestaurantMenu, error) {
	return s.r.ListRestaurantMenu(ctx, shopID)
}

func (s *productService) ListMenuItem(ctx context.Context, menuID string) ([]*types.MenuItem, error) {
	return s.r.ListMenuItem(ctx, types.HexToObjectID(menuID))
}

func (s *productService) ListRetailCategory(ctx context.Context, shopID string) ([]*types.RetailCategory, error) {
	return s.r.ListRetailCategory(ctx, shopID)
}

func (s *productService) ListRetailItem(ctx context.Context, categoryID string) ([]*types.RetailItem, error) {
	return s.r.ListRetailItem(ctx, types.HexToObjectID(categoryID))
}

func (s *productService) ListMedicineCategory(ctx context.Context, shopID string) ([]*types.MedicineCategory, error) {
	return s.r.ListMedicineCategory(ctx, shopID)
}

func (s *productService) ListMedicineItem(ctx context.Context, categoryID string) ([]*types.MedicineItem, error) {
	return s.r.ListMedicineItem(ctx, types.HexToObjectID(categoryID))
}

func (s *productService) UpdateItemVariant(ctx context.Context, variant *types.UpdateItemVariant) error {
	return s.r.UpdateItemVariant(ctx, variant)
}

func (s *productService) UpdateItemAddon(ctx context.Context, addon *types.UpdateItemAddon) error {
	return s.r.UpdateItemAddon(ctx, addon)
}

func (s *productService) UpdateRestaurantMenu(ctx context.Context, menu *types.UpdateRestaurantMenu) error {
	return s.r.UpdateRestaurantMenu(ctx, menu)
}

func (s *productService) UpdateMenuItem(ctx context.Context, item *types.UpdateMenuItem) error {
	return s.r.UpdateMenuItem(ctx, item)
}

func (s *productService) UpdateRetailCategory(ctx context.Context, category *types.UpdateRetailCategory) error {
	return s.r.UpdateRetailCategory(ctx, category)
}

func (s *productService) UpdateRetailItem(ctx context.Context, item *types.UpdateRetailItem) error {
	return s.r.UpdateRetailItem(ctx, item)
}

func (s *productService) UpdateMedicineCategory(ctx context.Context, category *types.UpdateMedicineCategory) error {
	return s.r.UpdateMedicineCategory(ctx, category)
}

func (s *productService) UpdateMedicineItem(ctx context.Context, item *types.UpdateMedicineItem) error {
	return s.r.UpdateMedicineItem(ctx, item)
}

func (s *productService) DeleteItemVariant(ctx context.Context, id string) error {
	return s.r.DeleteItemVariant(ctx, types.HexToObjectID(id))
}

func (s *productService) DeleteItemAddon(ctx context.Context, id string) error {
	return s.r.DeleteItemAddon(ctx, types.HexToObjectID(id))
}

func (s *productService) DeleteRestaurantMenu(ctx context.Context, id string) error {
	return s.r.DeleteRestaurantMenu(ctx, types.HexToObjectID(id))
}

func (s *productService) DeleteMenuItem(ctx context.Context, id string) error {
	return s.r.DeleteMenuItem(ctx, types.HexToObjectID(id))
}

func (s *productService) DeleteRetailCategory(ctx context.Context, id string) error {
	return s.r.DeleteRetailCategory(ctx, types.HexToObjectID(id))
}

func (s *productService) DeleteRetailItem(ctx context.Context, id string) error {
	return s.r.DeleteRetailItem(ctx, types.HexToObjectID(id))
}

func (s *productService) DeleteMedicineCategory(ctx context.Context, id string) error {
	return s.r.DeleteMedicineCategory(ctx, types.HexToObjectID(id))
}

func (s *productService) DeleteMedicineItem(ctx context.Context, id string) error {
	return s.r.DeleteMedicineItem(ctx, types.HexToObjectID(id))
}
