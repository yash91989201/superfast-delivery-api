package types

import (
	"database/sql/driver"
	"time"
)

// helper types
type ShopType string
type ShopStatus string
type DayOfWeek string
type OrderBy string

const (
	Restaurant     ShopType = "restaurant"
	Grocery        ShopType = "grocery"
	Pharmaceutical ShopType = "pharmaceutical"
)

const (
	Open   ShopStatus = "open"
	Closed ShopStatus = "closed"
)

const (
	Monday    DayOfWeek = "monday"
	Tuesday   DayOfWeek = "tuesday"
	Wednesday DayOfWeek = "wednesday"
	Thursday  DayOfWeek = "thursday"
	Friday    DayOfWeek = "friday"
	Saturday  DayOfWeek = "saturday"
	Sunday    DayOfWeek = "sunday"
)

const (
	Asc  OrderBy = "ASC"
	Desc OrderBy = "DESC"
)

type Shop struct {
	ID          string        `json:"id" db:"id"`
	Name        string        `json:"name" db:"name"`
	ShopType    ShopType      `json:"shop_type" db:"shop_type"`
	ShopStatus  ShopStatus    `json:"shop_status" db:"shop_status"`
	OwnerAuthID string        `json:"owner_auth_id" db:"owner_auth_id"`
	CreatedAt   time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time    `json:"deleted_at" db:"deleted_at"`
	Address     *ShopAddress  `json:"shop_address"`
	Contact     *ShopContact  `json:"shop_contact"`
	Timing      []*ShopTiming `json:"shop_timing"`
	Image       []*ShopImage  `json:"shop_image"`
}

type ShopInfo struct {
	ID          string     `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	ShopType    ShopType   `json:"shop_type" db:"shop_type"`
	ShopStatus  ShopStatus `json:"shop_status" db:"shop_status"`
	OwnerAuthID string     `json:"owner_auth_id" db:"owner_auth_id"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}

type ShopAddress struct {
	ID             string    `json:"id" db:"id"`
	Address1       string    `json:"address1" db:"address1"`
	Address2       string    `json:"address2" db:"address2"`
	Longitude      float64   `json:"longitude" db:"longitude"`
	Latitude       float64   `json:"latitude" db:"latitude"`
	NearbyLandmark string    `json:"nearby_landmark" db:"nearby_landmark"`
	City           string    `json:"city" db:"city"`
	State          string    `json:"state" db:"state"`
	Pincode        string    `json:"pincode" db:"pincode"`
	Country        string    `json:"country" db:"country"`
	ShopID         string    `json:"shop_id" db:"shop_id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type ShopContact struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	Email       string    `json:"email" db:"email"`
	ShopID      string    `json:"shop_id" db:"shop_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type ShopImage struct {
	ID          string    `json:"id" db:"id"`
	ImageUrl    string    `json:"image_url" db:"image_url"`
	Description string    `json:"description" db:"description"`
	ShopID      string    `json:"shop_id" db:"shop_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type ShopTiming struct {
	ID        string    `json:"id" db:"id"`
	Day       DayOfWeek `json:"day" db:"day"`
	OpensAt   time.Time `json:"opens_at" db:"opens_at"`
	ClosesAt  time.Time `json:"closes_at" db:"closes_at"`
	ShopID    string    `json:"shop_id" db:"shop_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// crud input/output types
type CreateShopAddress struct {
	Address1       string
	Address2       string
	Longitude      float64
	Latitude       float64
	NearbyLandmark string
	City           string
	State          string
	Pincode        string
	Country        string
}

type CreateShopContact struct {
	Name        string
	PhoneNumber string
	Email       string
	ShopID      string
}

type CreateShopImage struct {
	ImageUrl    string
	Description string
}

type CreateShopTiming struct {
	Day      DayOfWeek
	OpensAt  time.Time
	ClosesAt time.Time
	ShopID   string
}

type CreateShop struct {
	Name        string
	ShopType    ShopType
	ShopStatus  ShopStatus
	OwnerAuthId string
	Address     CreateShopAddress
	Contact     CreateShopContact
	Image       []CreateShopImage
	Timing      []CreateShopTiming
}

type ListShopFilters struct {
	Name       *string     `json:"name"`
	ShopType   *ShopType   `json:"shop_type"`
	ShopStatus *ShopStatus `json:"shop_status"`
	OrderBy    *OrderBy    `json:"order_by"`
	Limit      *int        `json:"limit"`
	Offset     *int        `json:"offset"`
}

// helper methods for sqlx to scan and insert
func (s ShopType) Value() (driver.Value, error) {
	return string(s), nil
}

func (s *ShopType) Scan(value interface{}) error {
	*s = ShopType(value.([]uint8))
	return nil
}

func (s ShopStatus) Value() (driver.Value, error) {
	return string(s), nil
}

func (s *ShopStatus) Scan(value interface{}) error {
	*s = ShopStatus(value.([]uint8))
	return nil
}

func (s DayOfWeek) Value() (driver.Value, error) {
	return string(s), nil
}

func (s *DayOfWeek) Scan(value interface{}) error {
	*s = DayOfWeek(value.([]uint8))
	return nil
}

func (s OrderBy) Value() (driver.Value, error) {
	return string(s), nil
}

func (s *OrderBy) Scan(value interface{}) error {
	*s = OrderBy(value.([]uint8))
	return nil
}
