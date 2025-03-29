package types

import "time"

type AuthRole string

const (
	Customer        AuthRole = "customer"
	DeliveryPartner AuthRole = "delivery_partner"
	Vendor          AuthRole = "vendor"
	Admin           AuthRole = "auth"
)

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

// crud input/output types
type CreateAuth struct {
	Email         *string
	EmailVerified bool
	Phone         *string
	AuthRole      AuthRole
}

type ClientSession struct {
	AccessToken  string
	RefreshToken string
}

type SignInRes struct {
	Auth    *Auth
	Session *ClientSession
}
