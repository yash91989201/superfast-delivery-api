package types

import (
	"time"
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
