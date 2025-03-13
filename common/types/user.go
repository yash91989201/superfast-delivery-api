package types

import "time"

type Gender string
type AddressAlias string

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
	Latitude            float64      `json:"latitude" db:"latitude"`
	Longitude           float64      `json:"longitude" db:"longitude"`
	Address             string       `json:"address" db:"address"`
	NearbyLandmark      *string      `json:"nearby_landmark" db:"nearby_landmark"`
	DeliveryInstruction *string      `json:"delivery_instruction" db:"delivery_instruction"`
	IsDefault           bool         `json:"is_default" db:"is_default"`
	AuthId              string       `json:"auth_id" db:"auth_id"`
	CreatedAt           time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time    `json:"updated_at" db:"updated_at"`
}

// crud input/output types
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
	Latitude            float64
	Longitude           float64
	Address             string
	NearbyLandmark      *string
	DeliveryInstruction *string
	IsDefault           bool
	AuthId              string
}
