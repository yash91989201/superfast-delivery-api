// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Auth struct {
	ID            string   `json:"id"`
	Email         *string  `json:"email,omitempty"`
	EmailVerified bool     `json:"email_verified"`
	Phone         *string  `json:"phone,omitempty"`
	AuthRole      AuthRole `json:"auth_role"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
	DeletedAt     *string  `json:"deleted_at,omitempty"`
}

type CreateProfileInput struct {
	Name        string  `json:"name"`
	ImageURL    *string `json:"image_url,omitempty"`
	Dob         *string `json:"dob,omitempty"`
	Anniversary *string `json:"anniversary,omitempty"`
	Gender      *Gender `json:"gender,omitempty"`
	AuthID      string  `json:"auth_id"`
}

type CreateShopAddressInput struct {
	Address1       string  `json:"address1"`
	Address2       *string `json:"address2,omitempty"`
	Longitude      float64 `json:"longitude"`
	Latitude       float64 `json:"latitude"`
	NearbyLandmark string  `json:"nearby_landmark"`
	City           string  `json:"city"`
	State          string  `json:"state"`
	Pincode        string  `json:"pincode"`
	Country        string  `json:"country"`
}

type CreateShopContactInput struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type CreateShopImageInput struct {
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

type CreateShopInput struct {
	Name       string                   `json:"name"`
	ShopType   ShopType                 `json:"shop_type"`
	ShopStatus ShopStatus               `json:"shop_status"`
	OwnerID    string                   `json:"owner_id"`
	Address    *CreateShopAddressInput  `json:"address"`
	Contact    *CreateShopContactInput  `json:"contact"`
	Images     []*CreateShopImageInput  `json:"images,omitempty"`
	Timings    []*CreateShopTimingInput `json:"timings,omitempty"`
}

type CreateShopOutput struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type CreateShopTimingInput struct {
	Day      DayOfWeek `json:"day"`
	OpensAt  time.Time `json:"opens_at"`
	ClosesAt time.Time `json:"closes_at"`
}

type GetAuthByIDInput struct {
	ID string `json:"id"`
}

type GetAuthInput struct {
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

type GetProfileInput struct {
	AuthID string `json:"auth_id"`
}

type GetShopOutput struct {
	Shop    *Shop         `json:"shop"`
	Address *ShopAddress  `json:"address"`
	Contact *ShopContact  `json:"contact"`
	Images  []*ShopImage  `json:"images,omitempty"`
	Timings []*ShopTiming `json:"timings,omitempty"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type LatLngInput struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type ListShopsInput struct {
	ShopType   *ShopType   `json:"shop_type,omitempty"`
	ShopStatus *ShopStatus `json:"shop_status,omitempty"`
	Page       *int32      `json:"page,omitempty"`
	PageSize   *int32      `json:"page_size,omitempty"`
}

type ListShopsOutput struct {
	Shops []*Shop `json:"shops,omitempty"`
	Total int32   `json:"total"`
}

type Mutation struct {
}

type Profile struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	ImageURL    *string `json:"image_url,omitempty"`
	Dob         *string `json:"dob,omitempty"`
	Anniversary *string `json:"anniversary,omitempty"`
	Gender      *Gender `json:"gender,omitempty"`
	AuthID      string  `json:"auth_id"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type Query struct {
}

type Shop struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	ShopType   ShopType   `json:"shop_type"`
	ShopStatus ShopStatus `json:"shop_status"`
	OwnerID    string     `json:"owner_id"`
	CreatedAt  string     `json:"created_at"`
	UpdatedAt  string     `json:"updated_at"`
	DeletedAt  *string    `json:"deleted_at,omitempty"`
}

type ShopAddress struct {
	ID             string  `json:"id"`
	Address1       string  `json:"address1"`
	Address2       *string `json:"address2,omitempty"`
	Longitude      float64 `json:"longitude"`
	Latitude       float64 `json:"latitude"`
	NearbyLandmark *string `json:"nearby_landmark,omitempty"`
	City           string  `json:"city"`
	State          string  `json:"state"`
	Pincode        string  `json:"pincode"`
	Country        string  `json:"country"`
	ShopID         string  `json:"shop_id"`
	CreatedAt      string  `json:"created_at"`
}

type ShopContact struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	ShopID      string `json:"shop_id"`
	CreatedAt   string `json:"created_at"`
}

type ShopImage struct {
	ID          string  `json:"id"`
	ImageURL    string  `json:"image_url"`
	Description *string `json:"description,omitempty"`
	ShopID      string  `json:"shop_id"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type ShopTiming struct {
	ID        string    `json:"id"`
	Day       DayOfWeek `json:"day"`
	OpensAt   time.Time `json:"opens_at"`
	ClosesAt  time.Time `json:"closes_at"`
	ShopID    string    `json:"shop_id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type SignInOutput struct {
	Auth                 *Auth    `json:"auth,omitempty"`
	Profile              *Profile `json:"profile,omitempty"`
	SessionID            *string  `json:"session_id,omitempty"`
	AccessToken          *string  `json:"access_token,omitempty"`
	AccessTokenExpiresAt *string  `json:"access_token_expires_at,omitempty"`
	CreateProfile        bool     `json:"create_profile"`
	VerifyOtp            bool     `json:"verify_otp"`
}

type SignInWithEmailInput struct {
	Email string  `json:"email"`
	Otp   *string `json:"otp,omitempty"`
}

type SignInWithGoogleInput struct {
	IDToken string `json:"id_token"`
}

type SignInWithPhoneInput struct {
	Phone string  `json:"phone"`
	Otp   *string `json:"otp,omitempty"`
}

type UpdateProfileInput struct {
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	ImageURL    *string `json:"image_url,omitempty"`
	Dob         *string `json:"dob,omitempty"`
	Anniversary *string `json:"anniversary,omitempty"`
	Gender      *Gender `json:"gender,omitempty"`
	AuthID      string  `json:"auth_id"`
}

type UpdateShopAddressInput struct {
	ID             string       `json:"id"`
	Address1       *string      `json:"address1,omitempty"`
	Address2       *string      `json:"address2,omitempty"`
	Location       *LatLngInput `json:"location,omitempty"`
	NearbyLandmark *string      `json:"nearby_landmark,omitempty"`
	City           *string      `json:"city,omitempty"`
	State          *string      `json:"state,omitempty"`
	Pincode        *string      `json:"pincode,omitempty"`
	Country        *string      `json:"country,omitempty"`
}

type UpdateShopContactInput struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Email       *string `json:"email,omitempty"`
}

type UpdateShopImageInput struct {
	ID          string  `json:"id"`
	ImageURL    *string `json:"image_url,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdateShopInput struct {
	ID         string      `json:"id"`
	Name       *string     `json:"name,omitempty"`
	ShopType   *ShopType   `json:"shop_type,omitempty"`
	ShopStatus *ShopStatus `json:"shop_status,omitempty"`
}

type UpdateShopOutput struct {
	Message string `json:"message"`
}

type UpdateShopTimingInput struct {
	ID       string     `json:"id"`
	Day      *DayOfWeek `json:"day,omitempty"`
	OpensAt  *time.Time `json:"opens_at,omitempty"`
	ClosesAt *time.Time `json:"closes_at,omitempty"`
}

type AuthRole string

const (
	AuthRoleCustomer        AuthRole = "CUSTOMER"
	AuthRoleDeliveryPartner AuthRole = "DELIVERY_PARTNER"
	AuthRoleVendor          AuthRole = "VENDOR"
	AuthRoleAdmin           AuthRole = "ADMIN"
)

var AllAuthRole = []AuthRole{
	AuthRoleCustomer,
	AuthRoleDeliveryPartner,
	AuthRoleVendor,
	AuthRoleAdmin,
}

func (e AuthRole) IsValid() bool {
	switch e {
	case AuthRoleCustomer, AuthRoleDeliveryPartner, AuthRoleVendor, AuthRoleAdmin:
		return true
	}
	return false
}

func (e AuthRole) String() string {
	return string(e)
}

func (e *AuthRole) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AuthRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AuthRole", str)
	}
	return nil
}

func (e AuthRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DayOfWeek string

const (
	DayOfWeekMonday    DayOfWeek = "MONDAY"
	DayOfWeekTuesday   DayOfWeek = "TUESDAY"
	DayOfWeekWednesday DayOfWeek = "WEDNESDAY"
	DayOfWeekThursday  DayOfWeek = "THURSDAY"
	DayOfWeekFriday    DayOfWeek = "FRIDAY"
	DayOfWeekSaturday  DayOfWeek = "SATURDAY"
	DayOfWeekSunday    DayOfWeek = "SUNDAY"
)

var AllDayOfWeek = []DayOfWeek{
	DayOfWeekMonday,
	DayOfWeekTuesday,
	DayOfWeekWednesday,
	DayOfWeekThursday,
	DayOfWeekFriday,
	DayOfWeekSaturday,
	DayOfWeekSunday,
}

func (e DayOfWeek) IsValid() bool {
	switch e {
	case DayOfWeekMonday, DayOfWeekTuesday, DayOfWeekWednesday, DayOfWeekThursday, DayOfWeekFriday, DayOfWeekSaturday, DayOfWeekSunday:
		return true
	}
	return false
}

func (e DayOfWeek) String() string {
	return string(e)
}

func (e *DayOfWeek) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DayOfWeek(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DayOfWeek", str)
	}
	return nil
}

func (e DayOfWeek) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Gender string

const (
	GenderMale        Gender = "MALE"
	GenderFemale      Gender = "FEMALE"
	GenderOthers      Gender = "OTHERS"
	GenderUndisclosed Gender = "UNDISCLOSED"
)

var AllGender = []Gender{
	GenderMale,
	GenderFemale,
	GenderOthers,
	GenderUndisclosed,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderMale, GenderFemale, GenderOthers, GenderUndisclosed:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ShopStatus string

const (
	ShopStatusOpen   ShopStatus = "OPEN"
	ShopStatusClosed ShopStatus = "CLOSED"
)

var AllShopStatus = []ShopStatus{
	ShopStatusOpen,
	ShopStatusClosed,
}

func (e ShopStatus) IsValid() bool {
	switch e {
	case ShopStatusOpen, ShopStatusClosed:
		return true
	}
	return false
}

func (e ShopStatus) String() string {
	return string(e)
}

func (e *ShopStatus) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ShopStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ShopStatus", str)
	}
	return nil
}

func (e ShopStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ShopType string

const (
	ShopTypeRestaurant     ShopType = "RESTAURANT"
	ShopTypeGrocery        ShopType = "GROCERY"
	ShopTypePharmaceutical ShopType = "PHARMACEUTICAL"
)

var AllShopType = []ShopType{
	ShopTypeRestaurant,
	ShopTypeGrocery,
	ShopTypePharmaceutical,
}

func (e ShopType) IsValid() bool {
	switch e {
	case ShopTypeRestaurant, ShopTypeGrocery, ShopTypePharmaceutical:
		return true
	}
	return false
}

func (e ShopType) String() string {
	return string(e)
}

func (e *ShopType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ShopType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ShopType", str)
	}
	return nil
}

func (e ShopType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
