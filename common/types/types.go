package types

import (
	"time"

	geo "github.com/kellydunn/golang-geo"
)

type AuthType string
type Gender string
type AddressAlias string

const (
	Customer        AuthType = "customer"
	DeliveryPartner AuthType = "delivery_partner"
	Vendor          AuthType = "vendor"
	Admin           AuthType = "auth"
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

type Date struct {
	Year  int32
	Month int32
	Day   int32
}

func (d *Date) ToTime() *time.Time {
	t := time.Date(int(d.Year), time.Month(d.Month), int(d.Day), 0, 0, 0, 0, time.UTC)
	return &t
}

type Auth struct {
	Id            string     `json:"id" db:"id"`
	Email         *string    `json:"email" db:"email"`
	EmailVerified bool       `json:"email_verified" db:"email_verified"`
	Phone         *string    `json:"phone" db:"phone"`
	Type          AuthType   `json:"type" db:"type"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at" db:"deleted_at"`
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
	Id          string     `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	ImageUrl    *string    `json:"image_url" db:"image_url"`
	Dob         *time.Time `json:"dob" db:"dob"`
	Anniversary *time.Time `json:"anniversary" db:"anniversary"`
	Gender      *Gender    `json:"gender" db:"gender"`
	AuthId      string     `json:"auth_id" db:"auth_id"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

type DeliveryAddress struct {
	Id                  string       `json:"id" db:"id"`
	ReceiverName        string       `json:"receiver_name" db:"receiver_name"`
	ReceiverPhone       string       `json:"receiver_phone" db:"receiver_phone"`
	AddressAlias        AddressAlias `json:"address_alias" db:"address_alias"`
	OtherAlias          *string      `json:"other_alias" db:"other_alias"`
	Location            geo.Point    `json:"location" db:"location"`
	Address             string       `json:"address" db:"address"`
	NearbyLandmark      *string      `json:"nearby_landmark" db:"nearby_landmark"`
	DeliveryInstruction *string      `json:"delivery_instruction" db:"delivery_instruction"`
	AuthId              string       `json:"auth_id" db:"auth_id"`
	CreatedAt           time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time    `json:"updated_at" db:"updated_at"`
}
