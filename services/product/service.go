package product

import (
	"context"
	"time"

	"github.com/yash91989201/superfast-delivery-api/common/types"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Service interface {
	InsertRestaurantMenu(ctx context.Context, rm *types.CreateRestaurantMenu) (*types.RestaurantMenu, error)
	InsertMenuItem(ctx context.Context, rm *types.CreateMenuItem) (*types.MenuItem, error)
}

type productService struct {
	r Repository
}

func New(r Repository) Service {
	return &productService{r: r}
}

func (s *productService) InsertItemVariant(ctx context.Context, v *types.CreateItemVariant) (*types.ItemVariant, error) {
	newVariant := &types.ItemVariant{
		ID:              bson.NewObjectID(),
		VariantName:     v.VariantName,
		RelativePrice:   v.RelativePrice,
		RelativePricing: v.RelativePricing,
		Price:           v.Price,
		Description:     v.Description,
		ItemId:          v.ItemId,
	}

	if err := s.r.InsertItemVariant(ctx, newVariant); err != nil {
		return nil, err
	}

	return newVariant, nil
}

func (s *productService) InsertItemAddon(ctx context.Context, a *types.CreateItemAddon) (*types.ItemAddon, error) {
	newAddon := &types.ItemAddon{
		ID:          bson.NewObjectID(),
		AddonName:   a.AddonName,
		AddonPrice:  a.AddonPrice,
		Description: a.Description,
		ItemId:      a.ItemId,
	}

	if err := s.r.InsertItemAddon(ctx, newAddon); err != nil {
		return nil, err
	}

	return newAddon, nil
}

func (s *productService) InsertRestaurantMenu(ctx context.Context, rm *types.CreateRestaurantMenu) (*types.RestaurantMenu, error) {
	newRestaurantMenu := &types.RestaurantMenu{
		ID:        bson.NewObjectID(),
		MenuName:  rm.MenuName,
		ShopID:    rm.ShopID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	if err := s.r.InsertRestaurantMenu(ctx, newRestaurantMenu); err != nil {
		return nil, err
	}

	return newRestaurantMenu, nil
}

func (s *productService) InsertMenuItem(ctx context.Context, mi *types.CreateMenuItem) (*types.MenuItem, error) {
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

	if err := s.r.InsertMenuItem(ctx, newMenuItem); err != nil {
		return nil, err
	}

	return newMenuItem, nil
}
