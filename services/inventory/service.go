package inventory

import (
	"context"

	"github.com/nrednav/cuid2"
	"github.com/yash91989201/superfast-delivery-api/common/types"
)

type Service interface {
	CreateItemStock(ctx context.Context, stock *types.CreateItemStock) (*types.ItemStock, error)
	CreateVariantStock(ctx context.Context, stock *types.CreateVariantStock) (*types.VariantStock, error)
	CreateAddonStock(ctx context.Context, stock *types.CreateAddonStock) (*types.AddonStock, error)
	GetItemStockByID(ctx context.Context, id string) (*types.ItemStock, error)
	GetVariantStockByID(ctx context.Context, id string) (*types.VariantStock, error)
	GetAddonStockByID(ctx context.Context, id string) (*types.AddonStock, error)
	UpdateItemStock(ctx context.Context, stock *types.ItemStock) (*types.ItemStock, error)
	UpdateVariantStock(ctx context.Context, stock *types.VariantStock) (*types.VariantStock, error)
	UpdateAddonStock(ctx context.Context, stock *types.AddonStock) (*types.AddonStock, error)
	DeleteItemStock(ctx context.Context, id string) error
	DeleteVariantStock(ctx context.Context, id string) error
	DeleteAddonStock(ctx context.Context, id string) error
}

type inventoryService struct {
	r Repository
}

func New(r Repository) Service {
	return &inventoryService{
		r: r,
	}
}

func (s *inventoryService) CreateItemStock(ctx context.Context, stock *types.CreateItemStock) (*types.ItemStock, error) {
	iStock := &types.ItemStock{
		ID:       cuid2.Generate(),
		ItemID:   stock.ItemID,
		Quantity: stock.Quantity,
	}

	if err := s.r.CreateItemStock(ctx, iStock); err != nil {
		return nil, err
	}

	return iStock, nil
}

func (s *inventoryService) CreateVariantStock(ctx context.Context, stock *types.CreateVariantStock) (*types.VariantStock, error) {
	vStock := &types.VariantStock{
		ID:        cuid2.Generate(),
		VariantID: stock.VariantID,
		Quantity:  stock.Quantity,
	}

	if err := s.r.CreateVariantStock(ctx, vStock); err != nil {
		return nil, err
	}

	return vStock, nil
}

func (s *inventoryService) CreateAddonStock(ctx context.Context, stock *types.CreateAddonStock) (*types.AddonStock, error) {
	aStock := &types.AddonStock{
		ID:       cuid2.Generate(),
		AddonID:  stock.AddonID,
		Quantity: stock.Quantity,
	}

	if err := s.r.CreateAddonStock(ctx, aStock); err != nil {
		return nil, err
	}

	return aStock, nil
}

func (s *inventoryService) GetItemStockByID(ctx context.Context, id string) (*types.ItemStock, error) {
	return s.r.GetItemStockByID(ctx, id)
}

func (s *inventoryService) GetVariantStockByID(ctx context.Context, id string) (*types.VariantStock, error) {
	return s.r.GetVariantStockByID(ctx, id)
}

func (s *inventoryService) GetAddonStockByID(ctx context.Context, id string) (*types.AddonStock, error) {
	return s.r.GetAddonStockByID(ctx, id)
}

func (s *inventoryService) UpdateItemStock(ctx context.Context, stock *types.ItemStock) (*types.ItemStock, error) {
	if err := s.r.UpdateItemStock(ctx, stock); err != nil {
		return nil, err
	}

	return s.r.GetItemStockByID(ctx, stock.ID)
}

func (s *inventoryService) UpdateVariantStock(ctx context.Context, stock *types.VariantStock) (*types.VariantStock, error) {
	if err := s.r.UpdateVariantStock(ctx, stock); err != nil {
		return nil, err
	}
	return s.r.GetVariantStockByID(ctx, stock.ID)
}

func (s *inventoryService) UpdateAddonStock(ctx context.Context, stock *types.AddonStock) (*types.AddonStock, error) {
	if err := s.r.UpdateAddonStock(ctx, stock); err != nil {
		return nil, err
	}

	return s.r.GetAddonStockByID(ctx, stock.ID)
}

func (s *inventoryService) DeleteItemStock(ctx context.Context, id string) error {
	return s.r.DeleteItemStock(ctx, id)
}

func (s *inventoryService) DeleteVariantStock(ctx context.Context, id string) error {
	return s.r.DeleteVariantStock(ctx, id)
}

func (s *inventoryService) DeleteAddonStock(ctx context.Context, id string) error {
	return s.r.DeleteAddonStock(ctx, id)
}
