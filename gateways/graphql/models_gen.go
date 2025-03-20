// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type AddonStock struct {
	ID         string `json:"id"`
	AddonID    string `json:"addon_id"`
	Quantity   int32  `json:"quantity"`
	RestockQty int32  `json:"restock_qty"`
	UpdatedAt  string `json:"updated_at"`
}

type AddressDetail struct {
	ID               string  `json:"id"`
	Route            string  `json:"route"`
	Town             string  `json:"town"`
	PostalCode       string  `json:"postal_code"`
	District         string  `json:"district"`
	State            string  `json:"state"`
	Country          string  `json:"country"`
	PlusCode         string  `json:"plus_code"`
	PlaceID          string  `json:"place_id"`
	FormattedAddress string  `json:"formatted_address"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	AddressID        string  `json:"address_id"`
}

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

type CreateAddonStockInput struct {
	AddonID  string `json:"addon_id"`
	Quantity int32  `json:"quantity"`
}

type CreateDeliveryAddressInput struct {
	ReceiverName        string       `json:"receiver_name"`
	ReceiverPhone       string       `json:"receiver_phone"`
	AddressAlias        AddressAlias `json:"address_alias"`
	OtherAlias          *string      `json:"other_alias,omitempty"`
	Latitude            float64      `json:"latitude"`
	Longitude           float64      `json:"longitude"`
	Address             string       `json:"address"`
	NearbyLandmark      *string      `json:"nearby_landmark,omitempty"`
	DeliveryInstruction *string      `json:"delivery_instruction,omitempty"`
	IsDefault           bool         `json:"is_default"`
	AuthID              string       `json:"auth_id"`
}

type CreateItemAddonInput struct {
	AddonName   string  `json:"addon_name"`
	AddonPrice  float64 `json:"addon_price"`
	Description *string `json:"description,omitempty"`
	ItemID      string  `json:"item_id"`
}

type CreateItemStockInput struct {
	ItemID   string `json:"item_id"`
	Quantity int32  `json:"quantity"`
}

type CreateItemVariantInput struct {
	VariantName     string  `json:"variant_name"`
	RelativePrice   float64 `json:"relative_price"`
	RelativePricing bool    `json:"relative_pricing"`
	Price           float64 `json:"price"`
	Description     *string `json:"description,omitempty"`
	ItemID          string  `json:"item_id"`
}

type CreateMedicineCategoryInput struct {
	CategoryName string `json:"category_name"`
	ShopID       string `json:"shop_id"`
}

type CreateMedicineItemInput struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CategoryID  string  `json:"category_id"`
}

type CreateMenuItemInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	MenuID      string  `json:"menu_id"`
}

type CreateProfileInput struct {
	Name        string  `json:"name"`
	ImageURL    *string `json:"image_url,omitempty"`
	Dob         *string `json:"dob,omitempty"`
	Anniversary *string `json:"anniversary,omitempty"`
	Gender      *Gender `json:"gender,omitempty"`
	AuthID      string  `json:"auth_id"`
}

type CreateRestaurantMenuInput struct {
	MenuName string `json:"menu_name"`
	ShopID   string `json:"shop_id"`
}

type CreateRetailCategoryInput struct {
	CategoryName string `json:"category_name"`
	ShopID       string `json:"shop_id"`
}

type CreateRetailItemInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  string  `json:"category_id"`
}

type CreateShopAddressInput struct {
	Longitude      float64 `json:"longitude"`
	Latitude       float64 `json:"latitude"`
	Address        string  `json:"address"`
	NearbyLandmark string  `json:"nearby_landmark"`
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
	Name        string                   `json:"name"`
	ShopType    ShopType                 `json:"shop_type"`
	ShopStatus  ShopStatus               `json:"shop_status"`
	OwnerAuthID string                   `json:"owner_auth_id"`
	Address     *CreateShopAddressInput  `json:"address"`
	Contact     *CreateShopContactInput  `json:"contact"`
	Images      []*CreateShopImageInput  `json:"images,omitempty"`
	Timings     []*CreateShopTimingInput `json:"timings,omitempty"`
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

type CreateVariantStockInput struct {
	VariantID string `json:"variant_id"`
	Quantity  int32  `json:"quantity"`
}

type DeleteOutput struct {
	Message string `json:"message"`
}

type DeliveryAddress struct {
	ID                  string       `json:"id"`
	ReceiverName        string       `json:"receiver_name"`
	ReceiverPhone       string       `json:"receiver_phone"`
	AddressAlias        AddressAlias `json:"address_alias"`
	OtherAlias          *string      `json:"other_alias,omitempty"`
	Latitude            float64      `json:"latitude"`
	Longitude           float64      `json:"longitude"`
	Address             string       `json:"address"`
	NearbyLandmark      *string      `json:"nearby_landmark,omitempty"`
	DeliveryInstruction *string      `json:"delivery_instruction,omitempty"`
	IsDefault           bool         `json:"is_default"`
	AuthID              string       `json:"auth_id"`
	CreatedAt           string       `json:"created_at"`
	UpdatedAt           string       `json:"updated_at"`
}

type GetAuthByIDInput struct {
	ID string `json:"id"`
}

type GetAuthInput struct {
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

type GetItemAddonsOutput struct {
	Addons      []*ItemAddon `json:"addons"`
	TotalAddons int32        `json:"total_addons"`
}

type GetItemVariantsOutput struct {
	Variants      []*ItemVariant `json:"variants"`
	TotalVariants int32          `json:"total_variants"`
}

type GetProfileInput struct {
	AuthID string `json:"auth_id"`
}

type ItemAddon struct {
	ID          string  `json:"id"`
	AddonName   string  `json:"addon_name"`
	AddonPrice  float64 `json:"addon_price"`
	Description *string `json:"description,omitempty"`
	ItemID      string  `json:"item_id"`
}

type ItemStock struct {
	ID         string `json:"id"`
	ItemID     string `json:"item_id"`
	Quantity   int32  `json:"quantity"`
	RestockQty int32  `json:"restock_qty"`
	UpdatedAt  string `json:"updated_at"`
}

type ItemVariant struct {
	ID              string  `json:"id"`
	VariantName     string  `json:"variant_name"`
	RelativePrice   float64 `json:"relative_price"`
	RelativePricing bool    `json:"relative_pricing"`
	Price           float64 `json:"price"`
	Description     *string `json:"description,omitempty"`
	ItemID          string  `json:"item_id"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type LatLngInput struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type ListDeliveryAddressOutput struct {
	DeliveryAddress []*DeliveryAddress `json:"delivery_address"`
}

type ListMedicineCategoryOutput struct {
	MedicineCategoryList []*MedicineCategory `json:"medicine_category_list"`
	TotalCategory        int32               `json:"total_category"`
}

type ListRestaurantMenuOutput struct {
	RestaurantMenuList []*RestaurantMenu `json:"restaurant_menu_list"`
	TotalMenu          int32             `json:"total_menu"`
}

type ListRetailCategoryOutput struct {
	RetailCategoryList []*RetailCategory `json:"retail_category_list"`
	TotalCategory      int32             `json:"total_category"`
}

type ListShopsInput struct {
	Name       *string     `json:"name,omitempty"`
	ShopType   *ShopType   `json:"shop_type,omitempty"`
	ShopStatus *ShopStatus `json:"shop_status,omitempty"`
	OrderBy    *OrderBy    `json:"order_by,omitempty"`
	Limit      *int32      `json:"limit,omitempty"`
	Offset     *int32      `json:"offset,omitempty"`
}

type ListShopsOutput struct {
	Shops []*Shop `json:"shops,omitempty"`
	Total int32   `json:"total"`
}

type MedicineCategory struct {
	ID            string          `json:"id"`
	CategoryName  string          `json:"category_name"`
	ShopID        string          `json:"shop_id"`
	MedicineItems []*MedicineItem `json:"medicine_items"`
	CreatedAt     string          `json:"created_at"`
	UpdatedAt     string          `json:"updated_at"`
	DeletedAt     *string         `json:"deleted_at,omitempty"`
}

type MedicineItem struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CategoryID  string  `json:"category_id"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   *string `json:"deleted_at,omitempty"`
}

type MenuItem struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description *string        `json:"description,omitempty"`
	Price       float64        `json:"price"`
	MenuID      string         `json:"menu_id"`
	Variants    []*ItemVariant `json:"variants"`
	Addons      []*ItemAddon   `json:"addons"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
	DeletedAt   *string        `json:"deleted_at,omitempty"`
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

type RestaurantMenu struct {
	ID        string      `json:"id"`
	MenuName  string      `json:"menu_name"`
	ShopID    string      `json:"shop_id"`
	MenuItems []*MenuItem `json:"menu_items"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	DeletedAt *string     `json:"deleted_at,omitempty"`
}

type RetailCategory struct {
	ID           string        `json:"id"`
	CategoryName string        `json:"category_name"`
	ShopID       string        `json:"shop_id"`
	RetailItems  []*RetailItem `json:"retail_items"`
	CreatedAt    string        `json:"created_at"`
	UpdatedAt    string        `json:"updated_at"`
	DeletedAt    *string       `json:"deleted_at,omitempty"`
}

type RetailItem struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	CategoryID  string         `json:"category_id"`
	Variants    []*ItemVariant `json:"variants"`
	Addons      []*ItemAddon   `json:"addons"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
	DeletedAt   *string        `json:"deleted_at,omitempty"`
}

type Session struct {
	ID                   string `json:"id"`
	AccessToken          string `json:"access_token"`
	AccessTokenExpiresAt string `json:"access_token_expires_at"`
}

type Shop struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	ShopType    ShopType      `json:"shop_type"`
	ShopStatus  ShopStatus    `json:"shop_status"`
	OwnerAuthID string        `json:"owner_auth_id"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
	DeletedAt   *string       `json:"deleted_at,omitempty"`
	Contact     *ShopContact  `json:"contact"`
	Address     *ShopAddress  `json:"address"`
	Timings     []*ShopTiming `json:"timings,omitempty"`
	Images      []*ShopImage  `json:"images,omitempty"`
}

type ShopAddress struct {
	ID             string  `json:"id"`
	Longitude      float64 `json:"longitude"`
	Latitude       float64 `json:"latitude"`
	Address        string  `json:"address"`
	NearbyLandmark string  `json:"nearby_landmark"`
	ShopID         string  `json:"shop_id"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
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
	ID          string `json:"id"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	ShopID      string `json:"shop_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
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
	Auth          *Auth    `json:"auth,omitempty"`
	Profile       *Profile `json:"profile,omitempty"`
	Session       *Session `json:"session,omitempty"`
	CreateProfile bool     `json:"create_profile"`
	VerifyOtp     bool     `json:"verify_otp"`
}

type SignInWithEmailInput struct {
	Email    string   `json:"email"`
	AuthRole AuthRole `json:"auth_role"`
	Otp      *string  `json:"otp,omitempty"`
}

type SignInWithGoogleInput struct {
	IDToken  string   `json:"id_token"`
	AuthRole AuthRole `json:"auth_role"`
}

type SignInWithPhoneInput struct {
	Phone    string   `json:"phone"`
	AuthRole AuthRole `json:"auth_role"`
	Otp      *string  `json:"otp,omitempty"`
}

type UpdateDefaultDeliveryAddressInput struct {
	DeliveryAddressID string `json:"delivery_address_id"`
	AuthID            string `json:"auth_id"`
}

type UpdateDeliveryAddressInput struct {
	ReceiverName        *string       `json:"receiver_name,omitempty"`
	ReceiverPhone       *string       `json:"receiver_phone,omitempty"`
	AddressAlias        *AddressAlias `json:"address_alias,omitempty"`
	OtherAlias          *string       `json:"other_alias,omitempty"`
	Latitude            *float64      `json:"latitude,omitempty"`
	Longitude           *float64      `json:"longitude,omitempty"`
	NearbyLandmark      *string       `json:"nearby_landmark,omitempty"`
	DeliveryInstruction *string       `json:"delivery_instruction,omitempty"`
	IsDefault           bool          `json:"is_default"`
	AuthID              string        `json:"auth_id"`
}

type UpdateOutput struct {
	Message string `json:"message"`
}

type UpdateProfileInput struct {
	ID          string  `json:"id"`
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

type VariantStock struct {
	ID         string `json:"id"`
	VariantID  string `json:"variant_id"`
	Quantity   int32  `json:"quantity"`
	RestockQty int32  `json:"restock_qty"`
	UpdatedAt  string `json:"updated_at"`
}

type AddressAlias string

const (
	AddressAliasHome  AddressAlias = "HOME"
	AddressAliasWork  AddressAlias = "WORK"
	AddressAliasHotel AddressAlias = "HOTEL"
	AddressAliasOther AddressAlias = "OTHER"
)

var AllAddressAlias = []AddressAlias{
	AddressAliasHome,
	AddressAliasWork,
	AddressAliasHotel,
	AddressAliasOther,
}

func (e AddressAlias) IsValid() bool {
	switch e {
	case AddressAliasHome, AddressAliasWork, AddressAliasHotel, AddressAliasOther:
		return true
	}
	return false
}

func (e AddressAlias) String() string {
	return string(e)
}

func (e *AddressAlias) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AddressAlias(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AddressAlias", str)
	}
	return nil
}

func (e AddressAlias) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
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

type OrderBy string

const (
	OrderByAsc  OrderBy = "ASC"
	OrderByDesc OrderBy = "DESC"
)

var AllOrderBy = []OrderBy{
	OrderByAsc,
	OrderByDesc,
}

func (e OrderBy) IsValid() bool {
	switch e {
	case OrderByAsc, OrderByDesc:
		return true
	}
	return false
}

func (e OrderBy) String() string {
	return string(e)
}

func (e *OrderBy) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderBy", str)
	}
	return nil
}

func (e OrderBy) MarshalGQL(w io.Writer) {
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
