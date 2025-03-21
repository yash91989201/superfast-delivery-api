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
	GetItemVariants(ctx context.Context, itemId string) ([]*types.ItemVariant, error)
	GetItemAddon(ctx context.Context, id string) (*types.ItemAddon, error)
	GetItemAddons(ctx context.Context, itemId string) ([]*types.ItemAddon, error)
	GetRestaurantMenu(ctx context.Context, id string) (*types.RestaurantMenu, error)
	ListRestaurantMenu(ctx context.Context, shopId string) ([]*types.RestaurantMenu, error)
	GetRetailCategory(ctx context.Context, id string) (*types.RetailCategory, error)
	ListRetailCategory(ctx context.Context, shopId string) ([]*types.RetailCategory, error)
	GetMedicineCategory(ctx context.Context, id string) (*types.MedicineCategory, error)
	ListMedicineCategory(ctx context.Context, shopId string) ([]*types.MedicineCategory, error)

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
		ItemId:          v.ItemId,
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
		ItemId:      a.ItemId,
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
		DeletedAt: nil,
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
		DeletedAt:   nil,
		MenuID:      mi.MenuID,
	}

	if err := s.r.CreateMenuItem(ctx, newMenuItem); err != nil {
		return nil, err
	}

	return newMenuItem, nil
}

// CreateMedicineCategory creates a new medicine category
func (s *productService) CreateMedicineCategory(ctx context.Context, mc *types.CreateMedicineCategory) (*types.MedicineCategory, error) {
	medicineCategory := &types.MedicineCategory{
		ID:            bson.NewObjectID(),
		CategoryName:  mc.CategoryName,
		ShopID:        mc.ShopID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     nil,
		MedicineItems: []*types.MedicineItem{},
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
		CategoryID:  mi.CategoryId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
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
		DeletedAt:    nil,
		RetailItems:  []*types.RetailItem{},
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
		CategoryID:  ri.CategoryId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
		Variants:    []*types.ItemVariant{},
		AddOns:      []*types.ItemAddon{},
	}

	err := s.r.CreateRetailItem(ctx, retailItem)
	if err != nil {
		return nil, err
	}

	return retailItem, nil
}

func (s *productService) GetItemVariant(ctx context.Context, id string) (*types.ItemVariant, error) {
	return s.r.GetItemVariant(ctx, id)
}

func (s *productService) GetItemAddon(ctx context.Context, id string) (*types.ItemAddon, error) {
	return s.r.GetItemAddon(ctx, id)
}

func (s *productService) GetItemVariants(ctx context.Context, itemId string) ([]*types.ItemVariant, error) {
	return s.r.GetItemVariants(ctx, itemId)
}

func (s *productService) GetItemAddons(ctx context.Context, itemId string) ([]*types.ItemAddon, error) {
	return s.r.GetItemAddons(ctx, itemId)
}

func (s *productService) GetRestaurantMenu(ctx context.Context, id string) (*types.RestaurantMenu, error) {
	return s.r.GetRestaurantMenu(ctx, id)
}

func (s *productService) ListRestaurantMenu(ctx context.Context, shopId string) ([]*types.RestaurantMenu, error) {
	return s.r.ListRestaurantMenu(ctx, shopId)
}

func (s *productService) GetRetailCategory(ctx context.Context, id string) (*types.RetailCategory, error) {
	return s.r.GetRetailCategory(ctx, id)
}

func (s *productService) ListRetailCategory(ctx context.Context, shopId string) ([]*types.RetailCategory, error) {
	return s.r.ListRetailCategory(ctx, shopId)
}

func (s *productService) GetMedicineCategory(ctx context.Context, id string) (*types.MedicineCategory, error) {
	return s.r.GetMedicineCategory(ctx, id)
}

func (s *productService) ListMedicineCategory(ctx context.Context, shopId string) ([]*types.MedicineCategory, error) {
	return s.r.ListMedicineCategory(ctx, shopId)
}

func (s *productService) DeleteItemAddon(ctx context.Context, id string) error {
	return s.r.DeleteItemAddon(ctx, id)
}

func (s *productService) DeleteItemVariant(ctx context.Context, id string) error {
	return s.r.DeleteItemVariant(ctx, id)
}

func (s *productService) DeleteMedicineCategory(ctx context.Context, id string) error {
	return s.r.DeleteMedicineCategory(ctx, id)
}

func (s *productService) DeleteMedicineItem(ctx context.Context, id string) error {
	return s.r.DeleteMedicineItem(ctx, id)
}

func (s *productService) DeleteMenuItem(ctx context.Context, id string) error {
	return s.r.DeleteMenuItem(ctx, id)
}

func (s *productService) DeleteRestaurantMenu(ctx context.Context, id string) error {
	return s.r.DeleteMenuItem(ctx, id)
}

func (s *productService) DeleteRetailCategory(ctx context.Context, id string) error {
	return s.r.DeleteRetailCategory(ctx, id)
}

func (s *productService) DeleteRetailItem(ctx context.Context, id string) error {
	return s.r.DeleteRetailItem(ctx, id)
}
