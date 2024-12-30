package types

import (
	"database/sql/driver"
	"time"
)

// helper types
type AuthRole string
type Gender string
type AddressAlias string
type ShopType string
type ShopStatus string
type DayOfWeek string

// helper methods for sqlx to scan and insert
func (s ShopType) Value() (driver.Value, error) {
	return string(s), nil
}

func (s *ShopType) Scan(value interface{}) error {
	*s = ShopType(value.(string))
	return nil
}

func (s ShopStatus) Value() (driver.Value, error) {
	return string(s), nil
}

func (s *ShopStatus) Scan(value interface{}) error {
	*s = ShopStatus(value.(string))
	return nil
}

func (s DayOfWeek) Value() (driver.Value, error) {
	return string(s), nil
}

func (s *DayOfWeek) Scan(value interface{}) error {
	*s = DayOfWeek(value.(string))
	return nil
}

const (
	Customer        AuthRole = "customer"
	DeliveryPartner AuthRole = "delivery_partner"
	Vendor          AuthRole = "vendor"
	Admin           AuthRole = "auth"
)

const (
	Male        Gender = "male"
	Female      Gender = "female"
	Others      Gender = "others"
	Undisclosed Gender = "undisclosed"
)

const (
	Home  AddressAlias = "home"
	Work  AddressAlias = "work"
	Hotel AddressAlias = "hotel"
	Other AddressAlias = "other"
)

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

type Date struct {
	Year  int32
	Month int32
	Day   int32
}

func (d *Date) ToTime() *time.Time {
	t := time.Date(int(d.Year), time.Month(d.Month), int(d.Day), 0, 0, 0, 0, time.UTC)
	return &t
}

// main types
type Auth struct {
	ID            string     `json:"id" db:"id"`
	Email         *string    `json:"email" db:"email"`
	EmailVerified bool       `json:"email_verified" db:"email_verified"`
	Phone         *string    `json:"phone" db:"phone"`
	AuthRole      AuthRole   `json:"auth_role" db:"auth_role"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at" db:"deleted_at"`
}

type Session struct {
	ID           string    `json:"id" db:"id"`
	AuthID       string    `json:"auth_id" db:"auth_id"`
	RefreshToken string    `json:"refresh_token" db:"refresh_token"`
	IsRevoked    bool      `json:"is_revoked" db:"is_revoked"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	ExpiresAt    time.Time `json:"expires_at" db:"expires_at"`
}

type EmailVerification struct {
	Token     string    `json:"token" db:"token"`
	Email     string    `json:"email" db:"email"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}

type PhoneVerification struct {
	Token     string    `json:"token" db:"token"`
	Phone     string    `json:"phone" db:"phone"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}

type Profile struct {
	ID          string     `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	ImageUrl    *string    `json:"image_url" db:"image_url"`
	Dob         *time.Time `json:"dob" db:"dob"`
	Anniversary *time.Time `json:"anniversary" db:"anniversary"`
	Gender      *Gender    `json:"gender" db:"gender"`
	AuthID      string     `json:"auth_id" db:"auth_id"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

type DeliveryAddress struct {
	ID                  string       `json:"id" db:"id"`
	ReceiverName        string       `json:"receiver_name" db:"receiver_name"`
	ReceiverPhone       string       `json:"receiver_phone" db:"receiver_phone"`
	AddressAlias        AddressAlias `json:"address_alias" db:"address_alias"`
	OtherAlias          *string      `json:"other_alias" db:"other_alias"`
	Location            Point        `json:"location" db:"location"`
	Address             string       `json:"address" db:"address"`
	NearbyLandmark      *string      `json:"nearby_landmark" db:"nearby_landmark"`
	DeliveryInstruction *string      `json:"delivery_instruction" db:"delivery_instruction"`
	AuthId              string       `json:"auth_id" db:"auth_id"`
	CreatedAt           time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time    `json:"updated_at" db:"updated_at"`
}

type Shop struct {
	ID         string       `json:"id" db:"id"`
	Name       string       `json:"name" db:"name"`
	ShopType   ShopType     `json:"shop_type" db:"shop_type"`
	ShopStatus ShopStatus   `json:"shop_status" db:"shop_status"`
	OwnerID    string       `json:"owner_id" db:"owner_id"`
	CreatedAt  string       `json:"created_at" db:"created_at"`
	UpdatedAt  string       `json:"updated_at" db:"updated_at"`
	DeletedAt  *string      `json:"deleted_at" db:"deleted_at"`
	Address    ShopAddress  `json:"shop_address"`
	Contact    ShopContact  `json:"shop_contact"`
	Timing     []ShopTiming `json:"shop_timing"`
	Image      []ShopImage  `json:"shop_image"`
}

type ShopAddress struct {
	ID             string  `json:"id" db:"id"`
	Address1       string  `json:"address1" db:"address1"`
	Address2       string  `json:"address2" db:"address2"`
	Longitude      float64 `json:"longitude" db:"longitude"`
	Latitude       float64 `json:"latitude" db:"latitude"`
	NearbyLandmark string  `json:"nearby_landmark" db:"nearby_landmark"`
	City           string  `json:"city" db:"city"`
	State          string  `json:"state" db:"state"`
	Pincode        string  `json:"pincode" db:"pincode"`
	Country        string  `json:"country" db:"country"`
	ShopID         string  `json:"shop_id" db:"shop_id"`
}

type ShopContact struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Email       string `json:"email" db:"email"`
	ShopID      string `json:"shop_id" db:"shop_id"`
}

type ShopImage struct {
	ID          string `json:"id" db:"id"`
	ImageUrl    string `json:"image_url" db:"image_url"`
	Description string `json:"description" db:"description"`
	ShopID      string `json:"shop_id" db:"shop_id"`
}

type ShopTiming struct {
	ID       string    `json:"id" db:"id"`
	Day      DayOfWeek `json:"day" db:"day"`
	OpensAt  time.Time `json:"opens_at" db:"opens_at"`
	ClosesAt time.Time `json:"closes_at" db:"closes_at"`
	ShopID   string    `json:"shop_id" db:"shop_id"`
}

// operation input types
type CreateAuth struct {
	Email         *string
	EmailVerified bool
	Phone         *string
	AuthRole      AuthRole
}

type SignInRes struct {
	Auth                 *Auth
	SessionId            *string
	AccessToken          *string
	AccessTokenExpiresAt *time.Time
}

type CreateProfile struct {
	Name        string
	ImageUrl    *string
	Dob         *time.Time
	Anniversary *time.Time
	Gender      *Gender
	AuthID      string
}

type CreateDeliveryAddress struct {
	ReceiverName        string
	ReceiverPhone       string
	AddressAlias        AddressAlias
	OtherAlias          *string
	Location            Point
	Address             string
	NearbyLandmark      *string
	DeliveryInstruction *string
	AuthId              string
}

type CreateShopAddress struct {
	Address1       string
	Address2       string
	Longitude      float64 `json:"longitude" db:"longitude"`
	Latitude       float64 `json:"latitude" db:"latitude"`
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
	Name       string
	ShopType   ShopType
	ShopStatus ShopStatus
	OwnerId    string
	Address    CreateShopAddress
	Contact    CreateShopContact
	Image      []CreateShopImage
	Timing     []CreateShopTiming
}
