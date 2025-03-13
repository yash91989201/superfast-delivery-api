package types

type AddressDetail struct {
	Id               string  `json:"id" db:"id"`
	Route            string  `json:"route" db:"route"`
	Town             string  `json:"locality" db:"locality"`
	PostalCode       string  `json:"postal_code" db:"postal_code"`
	District         string  `json:"district" db:"district"`
	State            string  `json:"state" db:"state"`
	Country          string  `json:"country" db:"country"`
	PlusCode         string  `json:"plus_code" db:"plus_code"`
	PlaceId          string  `json:"place_id" db:"place_id"`
	FormattedAddress string  `json:"formatted_address" db:"formatted_address"`
	Latitude         float64 `json:"latitude" db:"latitude"`
	Longitude        float64 `json:"longitude" db:"longitude"`
	AddressId        string  `json:"address_id" db:"address_id"`
}
