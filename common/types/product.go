package types

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ItemVariant struct {
	ID              bson.ObjectID `json:"id" bson:"_id"`
	VariantName     string        `json:"variant_name" bson:"variant_name"`
	RelativePrice   float64       `json:"relative_price" bson:"relative_price"`
	RelativePricing bool          `json:"relative_pricing" bson:"relative_pricing"`
	Price           float64       `json:"price" bson:"price"`
	Description     *string       `json:"description" bson:"description"`
	ItemId          bson.ObjectID `json:"item_id" bson:"item_id"`
}

type ItemAddon struct {
	ID          bson.ObjectID `json:"id" bson:"_id"`
	AddonName   string        `json:"addon_name" bson:"addon_name"`
	AddonPrice  float64       `json:"addon_price" bson:"addon_price"`
	Description *string       `json:"description" bson:"description"`
	ItemId      bson.ObjectID `json:"item_id" bson:"item_id"`
}

type RestaurantMenu struct {
	ID        bson.ObjectID `json:"id" bson:"_id"`
	MenuName  string        `json:"menu_name" bson:"menu_name"`
	ShopID    string        `json:"shop_id" bson:"shop_id"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time    `json:"deleted_at" bson:"deleted_at"`
	MenuItems []*MenuItem   `json:"menu_items" bson:"menu_items"`
}

type MenuItem struct {
	ID          bson.ObjectID  `json:"id" bson:"_id"`
	Name        string         `json:"name" bson:"name"`
	Description string         `json:"description" bson:"description"`
	Price       float64        `json:"price" bson:"price"`
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
	RetailItems  []*RetailItem `json:"retail_items" bson:"retail_items"`
}

type RetailItem struct {
	ID          bson.ObjectID  `json:"id" bson:"_id"`
	Name        string         `json:"name" bson:"name"`
	Description string         `json:"description" bson:"description"`
	Price       float64        `json:"price" bson:"price"`
	CategoryID  bson.ObjectID  `json:"category_id" bson:"category_id"`
	CreatedAt   time.Time      `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" bson:"updated_at"`
	DeletedAt   *time.Time     `json:"deleted_at" bson:"deleted_at"`
	Variants    []*ItemVariant `json:"variants" bson:"variants"`
	AddOns      []*ItemAddon   `json:"addons" bson:"addons"`
}

type MedicineCategory struct {
	ID            bson.ObjectID   `json:"id" bson:"_id"`
	CategoryName  string          `json:"category_name" bson:"category_name"`
	ShopID        string          `json:"shop_id" bson:"shop_id"`
	CreatedAt     time.Time       `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at" bson:"updated_at"`
	DeletedAt     *time.Time      `json:"deleted_at" bson:"deleted_at"`
	MedicineItems []*MedicineItem `json:"medicine_items" bson:"medicine_items"`
}

type MedicineItem struct {
	ID          bson.ObjectID `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Price       float64       `json:"price" bson:"price"`
	Description *string       `json:"description" bson:"description"`
	CategoryID  bson.ObjectID `json:"category_id" bson:"category_id"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
	DeletedAt   *time.Time    `json:"deleted_at" bson:"deleted_at"`
}

type CreateItemVariant struct {
	VariantName     string        `json:"variant_name" bson:"variant_name"`
	RelativePrice   float64       `json:"relative_price" bson:"relative_price"`
	RelativePricing bool          `json:"relative_pricing" bson:"relative_pricing"`
	Price           float64       `json:"price" bson:"price"`
	Description     *string       `json:"description" bson:"description"`
	ItemId          bson.ObjectID `json:"item_id" bson:"item_id"`
}

type CreateItemAddon struct {
	AddonName   string        `json:"addon_name" bson:"addon_name"`
	AddonPrice  float64       `json:"addon_price" bson:"addon_price"`
	Description *string       `json:"description" bson:"description"`
	ItemId      bson.ObjectID `json:"item_id" bson:"item_id"`
}

type CreateRestaurantMenu struct {
	MenuName string `json:"menu_name" bson:"menu_name"`
	ShopID   string `json:"shop_id" bson:"shop_id"`
}

type CreateMenuItem struct {
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Price       float64       `json:"price" bson:"price"`
	MenuID      bson.ObjectID `json:"menu_id" bson:"menu_id"`
}

type CreateRetailCategory struct {
	CategoryName string `json:"category_name" bson:"category_name"`
	ShopID       string `json:"shop_id" bson:"shop_id"`
}

type CreateRetailItem struct {
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Price       float64       `json:"price" bson:"price"`
	CategoryId  bson.ObjectID `json:"category_id" bson:"category_id"`
}

type CreateMedicineCategory struct {
	CategoryName string `json:"category_name" bson:"category_name"`
	ShopID       string `json:"shop_id" bson:"shop_id"`
}

type CreateMedicineItem struct {
	Name        string        `json:"name" bson:"name"`
	Price       float64       `json:"price" bson:"price"`
	Description *string       `json:"description" bson:"description"`
	CategoryId  bson.ObjectID `json:"category_id" bson:"category_id"`
}
