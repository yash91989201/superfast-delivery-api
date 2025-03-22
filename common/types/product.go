package types

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ItemVariant struct {
	ID              bson.ObjectID `bson:"_id" json:"id"`
	VariantName     string        `bson:"variant_name" json:"variant_name"`
	RelativePricing bool          `bson:"relative_pricing" json:"relative_pricing"`
	RelativePrice   float64       `bson:"relative_price" json:"relative_price"`
	Price           float64       `bson:"price" json:"price"`
	ImageURL        *string       `bson:"image_url" json:"imageUrl"`
	Description     *string       `bson:"description" json:"description"`
	ItemID          bson.ObjectID `bson:"item_id" json:"item_id"`
}

type ItemAddon struct {
	ID          bson.ObjectID `bson:"_id" json:"id"`
	AddonName   string        `bson:"addon_name" json:"addon_name"`
	AddonPrice  float64       `bson:"addon_price" json:"addon_price"`
	ImageURL    *string       `bson:"image_url" json:"imageUrl"`
	Description *string       `bson:"description" json:"description"`
	ItemID      bson.ObjectID `bson:"item_id" json:"item_id"`
}

type RestaurantMenu struct {
	ID          bson.ObjectID   `bson:"_id" json:"id"`
	MenuName    string          `bson:"menu_name" json:"menu_name"`
	ImageURL    *string         `bson:"image_url" json:"imageUrl"`
	ShopID      string          `bson:"shop_id" json:"shop_id"`
	MenuItemsID []bson.ObjectID `bson:"menu_items_id" json:"menu_items_id"`
	CreatedAt   time.Time       `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time       `bson:"updated_at" json:"updated_at"`
}

type MenuItem struct {
	ID          bson.ObjectID  `bson:"_id" json:"id"`
	Name        string         `bson:"name" json:"name"`
	Price       float64        `bson:"price" json:"price"`
	ImageURL    *string        `bson:"image_url" json:"imageUrl"`
	Description *string        `bson:"description" json:"description"`
	Variants    []*ItemVariant `bson:"variants" json:"variants"`
	Addons      []*ItemAddon   `bson:"addons" json:"addons"`
	MenuID      bson.ObjectID  `bson:"menu_id" json:"menu_id"`
	CreatedAt   time.Time      `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `bson:"updated_at" json:"updated_at"`
}

type RetailCategory struct {
	ID            bson.ObjectID   `bson:"_id" json:"id"`
	CategoryName  string          `bson:"category_name" json:"category_name"`
	ImageURL      *string         `bson:"image_url" json:"imageUrl"`
	ShopID        string          `bson:"shop_id" json:"shop_id"`
	RetailItemsID []bson.ObjectID `bson:"retail_items_id" json:"retail_items_id"`
	CreatedAt     time.Time       `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time       `bson:"updated_at" json:"updated_at"`
}

type RetailItem struct {
	ID          bson.ObjectID  `bson:"_id" json:"id"`
	Name        string         `bson:"name" json:"name"`
	Price       float64        `bson:"price" json:"price"`
	ImageURL    *string        `bson:"image_url" json:"imageUrl"`
	Description *string        `bson:"description" json:"description"`
	CategoryID  bson.ObjectID  `bson:"category_id" json:"category_id"`
	Variants    []*ItemVariant `bson:"variants" json:"variants"`
	CreatedAt   time.Time      `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `bson:"updated_at" json:"updated_at"`
}

type MedicineCategory struct {
	ID              bson.ObjectID   `bson:"_id" json:"id"`
	CategoryName    string          `bson:"category_name" json:"category_name"`
	ImageURL        *string         `bson:"image_url" json:"imageUrl"`
	ShopID          string          `bson:"shop_id" json:"shop_id"`
	MedicineItemsID []bson.ObjectID `bson:"medicine_items_id" json:"medicine_items_id"`
	CreatedAt       time.Time       `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time       `bson:"updated_at" json:"updated_at"`
}

type MedicineItem struct {
	ID          bson.ObjectID `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Price       float64       `bson:"price" json:"price"`
	ImageURL    *string       `bson:"image_url" json:"imageUrl"`
	Description *string       `bson:"description" json:"description"`
	CategoryID  bson.ObjectID `bson:"category_id" json:"category_id"`
	CreatedAt   time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at" json:"updated_at"`
}

type CreateItemVariant struct {
	VariantName     string
	RelativePricing bool
	RelativePrice   float64
	Price           float64
	ImageURL        *string
	Description     *string
	ItemID          bson.ObjectID
}

type CreateItemAddon struct {
	AddonName   string
	AddonPrice  float64
	ImageURL    *string
	Description *string
	ItemID      bson.ObjectID
}

type CreateRestaurantMenu struct {
	MenuName string
	ImageURL *string
	ShopID   string
}

type CreateMenuItem struct {
	Name        string
	Price       float64
	ImageUrl    *string
	Description *string
	MenuID      bson.ObjectID
}

type CreateRetailCategory struct {
	CategoryName string
	ImageURL     *string
	ShopID       string
}

type CreateRetailItem struct {
	Name        string
	Price       float64
	ImageURL    *string
	Description *string
	CategoryID  bson.ObjectID
}

type CreateMedicineCategory struct {
	CategoryName string
	ImageURL     *string
	ShopID       string
}

type CreateMedicineItem struct {
	Name        string
	Price       float64
	ImageURL    *string
	Description *string
	CategoryID  bson.ObjectID
}

type UpdateItemVariant struct {
	ID              bson.ObjectID `bson:"_id" json:"id"`
	VariantName     *string       `bson:"variant_name,omitempty" json:"variant_name,omitempty"`
	RelativePricing *bool         `bson:"relative_pricing,omitempty" json:"relative_pricing,omitempty"`
	RelativePrice   *float64      `bson:"relative_price,omitempty" json:"relative_price,omitempty"`
	Price           *float64      `bson:"price,omitempty" json:"price,omitempty"`
	ImageURL        *string       `bson:"image_url,omitempty" json:"imageUrl,omitempty"`
	Description     *string       `bson:"description,omitempty" json:"description,omitempty"`
}

type UpdateItemAddon struct {
	ID          bson.ObjectID `bson:"_id" json:"id"`
	AddonName   *string       `bson:"addon_name,omitempty" json:"addon_name,omitempty"`
	AddonPrice  *float64      `bson:"addon_price,omitempty" json:"addon_price,omitempty"`
	ImageURL    *string       `bson:"image_url,omitempty" json:"imageUrl,omitempty"`
	Description *string       `bson:"description,omitempty" json:"description,omitempty"`
}

type UpdateRestaurantMenu struct {
	ID       bson.ObjectID `bson:"_id" json:"id"`
	MenuName *string       `bson:"menu_name,omitempty" json:"menu_name,omitempty"`
	ImageURL *string       `bson:"image_url,omitempty" json:"imageUrl,omitempty"`
	ShopID   *string       `bson:"shop_id,omitempty" json:"shop_id,omitempty"`
}

type UpdateMenuItem struct {
	ID          bson.ObjectID `bson:"_id" json:"id"`
	Name        *string       `bson:"name,omitempty" json:"name,omitempty"`
	Price       *float64      `bson:"price,omitempty" json:"price,omitempty"`
	ImageURL    *string       `bson:"image_url,omitempty" json:"imageUrl,omitempty"`
	Description *string       `bson:"description,omitempty" json:"description,omitempty"`
}

type UpdateRetailCategory struct {
	ID           bson.ObjectID `bson:"_id" json:"id"`
	CategoryName *string       `bson:"category_name,omitempty" json:"category_name,omitempty"`
	ImageURL     *string       `bson:"image_url,omitempty" json:"imageUrl,omitempty"`
	ShopID       *string       `bson:"shop_id,omitempty" json:"shop_id,omitempty"`
}

type UpdateRetailItem struct {
	ID          bson.ObjectID `bson:"_id" json:"id"`
	Name        *string       `bson:"name,omitempty" json:"name,omitempty"`
	Price       *float64      `bson:"price,omitempty" json:"price,omitempty"`
	ImageURL    *string       `bson:"image_url,omitempty" json:"imageUrl,omitempty"`
	Description *string       `bson:"description,omitempty" json:"description,omitempty"`
}

type UpdateMedicineCategory struct {
	ID           bson.ObjectID `bson:"_id" json:"id"`
	CategoryName *string       `bson:"category_name,omitempty" json:"category_name,omitempty"`
	ImageURL     *string       `bson:"image_url,omitempty" json:"imageUrl,omitempty"`
	ShopID       *string       `bson:"shop_id,omitempty" json:"shop_id,omitempty"`
}

type UpdateMedicineItem struct {
	ID          bson.ObjectID `bson:"_id" json:"id"`
	Name        *string       `bson:"name,omitempty" json:"name,omitempty"`
	Price       *float64      `bson:"price,omitempty" json:"price,omitempty"`
	ImageURL    *string       `bson:"image_url,omitempty" json:"imageUrl,omitempty"`
	Description *string       `bson:"description,omitempty" json:"description,omitempty"`
}
