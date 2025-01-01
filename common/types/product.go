package types

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ItemVariant struct {
	ID              bson.ObjectID `json:"id" bson:"_id"`
	VariantName     string        `json:"variant_name" bson:"variant_name"`
	RelativePrice   float32       `json:"relative_price" bson:"relative_price"`
	RelativePricing bool          `json:"relative_pricing" bson:"relative_pricing"`
	Price           float32       `json:"price" bson:"price"`
	Description     *string       `json:"description" bson:"description"`
	ItemId          bson.ObjectID `json:"item_id" bson:"item_id"`
}

type ItemAddon struct {
	ID          bson.ObjectID `json:"id" bson:"_id"`
	AddonName   string        `json:"addon_name" bson:"addon_name"`
	AddonPrice  float32       `json:"addon_price" bson:"addon_price"`
	Description *string       `json:"description" bson:"description"`
	ItemId      string        `json:"item_id" bson:"item_id"`
}

type RestaurantMenu struct {
	ID        bson.ObjectID `json:"id" bson:"_id"`
	MenuName  string        `json:"menu_name" bson:"menu_name"`
	ShopID    string        `json:"shop_id" bson:"shop_id"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time    `json:"deleted_at" bson:"deleted_at"`
}

type MenuItem struct {
	ID          bson.ObjectID  `json:"id" bson:"_id"`
	Name        string         `json:"name" bson:"name"`
	Description string         `json:"description" bson:"description"`
	Price       float32        `json:"price" bson:"price"`
	CreatedAt   time.Time      `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" bson:"updated_at"`
	DeletedAt   *time.Time     `json:"deleted_at" bson:"deleted_at"`
	MenuID      bson.ObjectID  `json:"menu_id" bson:"menu_id"`
	Variants    []*ItemVariant `json:"variants" bson:"variants"`
	AddOns      []*ItemAddon   `json:"addons" bson:"addons"`
}

type RetailCategory struct {
	ID           bson.ObjectID `json:"id" bson:"_id"`
	CategoryName string        `json:"category_name" bson:"category_name"`
	ShopID       string        `json:"shop_id" bson:"shop_id"`
	CreatedAt    time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at" bson:"updated_at"`
	DeletedAt    *time.Time    `json:"deleted_at" bson:"deleted_at"`
}

type RetailItem struct {
	ID          bson.ObjectID  `json:"id" bson:"_id"`
	Name        string         `json:"name" bson:"name"`
	Description string         `json:"description" bson:"description"`
	Price       float32        `json:"price" bson:"price"`
	CreatedAt   time.Time      `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" bson:"updated_at"`
	DeletedAt   *time.Time     `json:"deleted_at" bson:"deleted_at"`
	Variants    []*ItemVariant `json:"variants" bson:"variants"`
	AddOns      []*ItemAddon   `json:"addons" bson:"addons"`
	CategoryID  bson.ObjectID  `json:"category_id" bson:"category_id"`
}

type MedicineCategory struct {
	ID           bson.ObjectID `json:"id" bson:"_id"`
	CategoryName string        `json:"category_name" bson:"category_name"`
	ShopID       string        `json:"shop_id" bson:"shop_id"`
	CreatedAt    time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at" bson:"updated_at"`
	DeletedAt    *time.Time    `json:"deleted_at" bson:"deleted_at"`
}

type MedicineItem struct {
	ID          bson.ObjectID `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Price       float32       `json:"price" bson:"price"`
	Description *string       `json:"description" bson:"description"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
	DeletedAt   *time.Time    `json:"deleted_at" bson:"deleted_at"`
	CategoryID  bson.ObjectID `json:"category_id" bson:"category_id"`
}
