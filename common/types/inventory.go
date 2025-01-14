package types

import (
	"time"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ItemStock struct {
	ID         string    `json:"id" db:"id"`
	ItemID     string    `json:"item_id" db:"item_id"`
	Quantity   int32     `json:"quantity" db:"quantity"`
	RestockQty int32     `json:"restock_qty" db:"restock_qty"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type VariantStock struct {
	ID         string    `json:"id" db:"id"`
	VariantID  string    `json:"variant_id" db:"variant_id"`
	Quantity   int32     `json:"quantity" db:"quantity"`
	RestockQty int32     `json:"restock_qty" db:"restock_qty"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type AddonStock struct {
	ID         string    `json:"id" db:"id"`
	AddonID    string    `json:"addon_id" db:"addon_id"`
	Quantity   int32     `json:"quantity" db:"quantity"`
	RestockQty int32     `json:"restock_qty" db:"restock_qty"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type CreateItemStock struct {
	ItemID   string `json:"item_id" db:"item_id"`
	Quantity int32  `json:"quantity" db:"quantity"`
}

type CreateVariantStock struct {
	VariantID string `json:"variant_id" db:"variant_id"`
	Quantity  int32  `json:"quantity" db:"quantity"`
}

type CreateAddonStock struct {
	AddonID  string `json:"addon_id" db:"addon_id"`
	Quantity int32  `json:"quantity" db:"quantity"`
}

func ToPbItemStock(s *ItemStock) *pb.ItemStock {
	return &pb.ItemStock{
		Id:         s.ID,
		ItemId:     s.ItemID,
		Quantity:   s.Quantity,
		RestockQty: s.RestockQty,
		UpdatedAt:  timestamppb.New(s.UpdatedAt),
	}
}

func ToPbVariantStock(s *VariantStock) *pb.VariantStock {
	return &pb.VariantStock{
		Id:         s.ID,
		VariantId:  s.VariantID,
		Quantity:   s.Quantity,
		RestockQty: s.RestockQty,
		UpdatedAt:  timestamppb.New(s.UpdatedAt),
	}
}

func ToPbAddonStock(s *AddonStock) *pb.AddonStock {
	return &pb.AddonStock{
		Id:         s.ID,
		AddonId:    s.AddonID,
		Quantity:   s.Quantity,
		RestockQty: s.RestockQty,
		UpdatedAt:  timestamppb.New(s.UpdatedAt),
	}
}
