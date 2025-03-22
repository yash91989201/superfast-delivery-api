package types

import (
	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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
