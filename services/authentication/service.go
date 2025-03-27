package authentication

import (
	"context"
	"fmt"
	"time"

	"github.com/nrednav/cuid2"
	"github.com/yash91989201/superfast-delivery-api/common/types"
)

type Service interface {
	CreateAuth(ctx context.Context, ca *types.CreateAuth) (*types.Auth, error)
	GetAuthById(ctx context.Context, id string) (*types.Auth, error)
	GetAuth(ctx context.Context, email *string, phone *string) (*types.Auth, error)
	DeleteAuth(ctx context.Context, id string) error

	CreateEmailVerification(ctx context.Context, ev *types.EmailVerification) error
	CreatePhoneVerification(ctx context.Context, pb *types.PhoneVerification) error
	GetEmailVerification(ctx context.Context, email string) (*types.EmailVerification, error)
	GetPhoneVerification(ctx context.Context, phone string) (*types.PhoneVerification, error)
	DeleteEmailVerification(ctx context.Context, email string) error
	DeletePhoneVerification(ctx context.Context, phone string) error

	CreateSession(ctx context.Context, session *types.Session) error
	GetSession(ctx context.Context, refreshToken string) (*types.Session, error)
	GetSessionById(ctx context.Context, id string) (*types.Session, error)
	GetSessionByAuthId(ctx context.Context, authID string) (*types.Session, error)
	RevokeSession(ctx context.Context, id string) error
	DeleteSession(ctx context.Context, id string) error
}

type authenticationService struct {
	r Repository
}

func New(r Repository) Service {
	return &authenticationService{r: r}
}

func (s *authenticationService) CreateAuth(ctx context.Context, ca *types.CreateAuth) (*types.Auth, error) {
	if ca.Email == nil && ca.Phone == nil {
		return nil, fmt.Errorf("One of email or phone is required")
	}

	auth := &types.Auth{
		ID:            cuid2.Generate(),
		Email:         ca.Email,
		EmailVerified: ca.EmailVerified,
		Phone:         ca.Phone,
		AuthRole:      ca.AuthRole,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := s.r.CreateAuth(ctx, auth); err != nil {
		return nil, err
	}

	return auth, nil
}

func (s *authenticationService) GetAuthById(ctx context.Context, id string) (*types.Auth, error) {
	return s.r.GetAuthById(ctx, id)
}

func (s *authenticationService) GetAuth(ctx context.Context, email *string, phone *string) (*types.Auth, error) {
	if email != nil && phone != nil {
		return nil, fmt.Errorf("Either email or phone is needed")
	}

	if email == nil && phone == nil {
		return nil, fmt.Errorf("One of email or phone is required")
	}

	if phone != nil {
		return s.r.GetAuthByPhone(ctx, *phone)
	}

	if email != nil {
		return s.r.GetAuthByEmail(ctx, *email)
	}

	return nil, fmt.Errorf("unexpected error")
}

func (s *authenticationService) DeleteAuth(ctx context.Context, id string) error {
	return s.r.DeleteAuth(ctx, id)
}

func (s *authenticationService) CreateEmailVerification(ctx context.Context, ev *types.EmailVerification) error {
	return s.r.CreateEmailVerification(ctx, ev)
}

func (s *authenticationService) CreatePhoneVerification(ctx context.Context, pv *types.PhoneVerification) error {
	return s.r.CreatePhoneVerification(ctx, pv)
}

func (s *authenticationService) GetEmailVerification(ctx context.Context, email string) (*types.EmailVerification, error) {
	return s.r.GetEmailVerification(ctx, email)
}

func (s *authenticationService) GetPhoneVerification(ctx context.Context, phone string) (*types.PhoneVerification, error) {
	return s.r.GetPhoneVerification(ctx, phone)
}

func (s *authenticationService) DeleteEmailVerification(ctx context.Context, email string) error {
	return s.r.DeleteEmailVerification(ctx, email)
}

func (s *authenticationService) DeletePhoneVerification(ctx context.Context, phone string) error {
	return s.r.DeletePhoneVerification(ctx, phone)
}

func (s *authenticationService) CreateSession(ctx context.Context, session *types.Session) error {
	return s.r.CreateSession(ctx, session)
}

func (s *authenticationService) GetSession(ctx context.Context, refreshToken string) (*types.Session, error) {
	return s.r.GetSession(ctx, refreshToken)
}

func (s *authenticationService) GetSessionById(ctx context.Context, id string) (*types.Session, error) {
	return s.r.GetSessionById(ctx, id)
}

func (s *authenticationService) GetSessionByAuthId(ctx context.Context, authID string) (*types.Session, error) {
	return s.r.GetSessionByAuthId(ctx, authID)
}

func (s *authenticationService) RevokeSession(ctx context.Context, id string) error {
	return s.r.RevokeSession(ctx, id)
}

func (s *authenticationService) DeleteSession(ctx context.Context, id string) error {
	return s.r.DeleteSession(ctx, id)
}
