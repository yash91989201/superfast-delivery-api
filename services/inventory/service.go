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
