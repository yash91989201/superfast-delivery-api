// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"fmt"
	"io"
	"strconv"
)

type Auth struct {
	ID            string   `json:"id"`
	Email         *string  `json:"email,omitempty"`
	EmailVerified *bool    `json:"email_verified,omitempty"`
	Phone         *string  `json:"phone,omitempty"`
	Role          AuthRole `json:"role"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
	DeletedAt     *string  `json:"deleted_at,omitempty"`
}

type GetAuthByIDInput struct {
	ID string `json:"id"`
}

type GetAuthInput struct {
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
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

type SignInOutput struct {
	Auth      *Auth    `json:"auth,omitempty"`
	Profile   *Profile `json:"profile,omitempty"`
	VerifyOtp *bool    `json:"verify_otp,omitempty"`
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
