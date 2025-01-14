package inventory

import (
	"context"

	"github.com/nrednav/cuid2"
	"github.com/yash91989201/superfast-delivery-api/common/types"
)

type Service interface {
	InsertItemStock(ctx context.Context, stock *types.CreateItemStock) (*types.ItemStock, error)
	InsertVariantStock(ctx context.Context, stock *types.CreateVariantStock) (*types.VariantStock, error)
	InsertAddonStock(ctx context.Context, stock *types.CreateAddonStock) (*types.AddonStock, error)
}

type inventoryService struct {
	r Repository
}

func New(r Repository) Service {
	return &inventoryService{
		r: r,
	}
}

func (s *inventoryService) InsertItemStock(ctx context.Context, stock *types.CreateItemStock) (*types.ItemStock, error) {
	iStock := &types.ItemStock{
		ID:       cuid2.Generate(),
		ItemID:   stock.ItemID,
		Quantity: stock.Quantity,
	}

	if err := s.r.InsertItemStock(ctx, iStock); err != nil {
		return nil, err
	}

	return iStock, nil
}

func (s *inventoryService) InsertVariantStock(ctx context.Context, stock *types.CreateVariantStock) (*types.VariantStock, error) {
	vStock := &types.VariantStock{
		ID:        cuid2.Generate(),
		VariantID: stock.VariantID,
		Quantity:  stock.Quantity,
	}

	if err := s.r.InsertVariantStock(ctx, vStock); err != nil {
		return nil, err
	}

	return vStock, nil
}

func (s *inventoryService) InsertAddonStock(ctx context.Context, stock *types.CreateAddonStock) (*types.AddonStock, error) {
	aStock := &types.AddonStock{
		ID:       cuid2.Generate(),
		AddonID:  stock.AddonID,
		Quantity: stock.Quantity,
	}

	if err := s.r.InsertAddonStock(ctx, aStock); err != nil {
		return nil, err
	}

	return aStock, nil
}
