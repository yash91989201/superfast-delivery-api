package user

import (
	"context"
	"time"

	"github.com/nrednav/cuid2"
	"github.com/yash91989201/superfast-delivery-api/common/types"
)

type Service interface {
	CreateProfile(ctx context.Context, p *types.CreateProfile) (*types.Profile, error)
	GetProfile(ctx context.Context, auth_id string) (*types.Profile, error)
	UpdateProfile(ctx context.Context, p *types.Profile) error
	DeleteProfile(ctx context.Context, id string) error

	CreateDeliveryAddress(ctx context.Context, d *types.CreateDeliveryAddress) (*types.DeliveryAddress, error)
	GetDeliveryAddress(ctx context.Context, id string) (*types.DeliveryAddress, error)
	GetDeliveryAddresses(ctx context.Context, authID string) ([]*types.DeliveryAddress, error)
	UpdateDeliveryAddress(ctx context.Context, d *types.DeliveryAddress) error
	UpdateDefaultDeliveryAddress(ctx context.Context, deliveryAddressId string, authId string) error
	DeleteDeliveryAddress(ctx context.Context, deliveryAddressId string) error
}

type userService struct {
	r Repository
}

func New(r Repository) Service {
	return &userService{r: r}
}

func (s *userService) CreateProfile(ctx context.Context, p *types.CreateProfile) (*types.Profile, error) {
	profile := &types.Profile{
		ID:          cuid2.Generate(),
		Name:        p.Name,
		ImageUrl:    p.ImageUrl,
		Dob:         p.Dob,
		Anniversary: p.Anniversary,
		Gender:      p.Gender,
		AuthID:      p.AuthID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.r.CreateProfile(ctx, profile); err != nil {
		return nil, err
	}

	return profile, nil
}

func (s *userService) GetProfile(ctx context.Context, auth_id string) (*types.Profile, error) {
	return s.r.GetProfileByAuthId(ctx, auth_id)
}

func (s *userService) UpdateProfile(ctx context.Context, p *types.Profile) error {
	return s.r.UpdateProfile(ctx, p)
}

func (s *userService) DeleteProfile(ctx context.Context, id string) error {
	return s.r.DeleteProfile(ctx, id)
}

func (s *userService) CreateDeliveryAddress(ctx context.Context, d *types.CreateDeliveryAddress) (*types.DeliveryAddress, error) {
	deliveryAddress := &types.DeliveryAddress{
		ID:                  cuid2.Generate(),
		ReceiverName:        d.ReceiverName,
		ReceiverPhone:       d.ReceiverPhone,
		AddressAlias:        d.AddressAlias,
		OtherAlias:          d.OtherAlias,
		Latitude:            d.Latitude,
		Longitude:           d.Longitude,
		Address:             d.Address,
		NearbyLandmark:      d.NearbyLandmark,
		DeliveryInstruction: d.DeliveryInstruction,
		AuthId:              d.AuthId,
	}

	if err := s.r.CreateDeliveryAddress(ctx, deliveryAddress); err != nil {
		return nil, err
	}

	return deliveryAddress, nil
}

func (s *userService) GetDeliveryAddress(ctx context.Context, id string) (*types.DeliveryAddress, error) {
	return s.r.GetDeliveryAddressById(ctx, id)
}

func (s *userService) GetDeliveryAddresses(ctx context.Context, authID string) ([]*types.DeliveryAddress, error) {
	return s.r.GetDeliveryAddresses(ctx, authID)
}

func (s *userService) UpdateDeliveryAddress(ctx context.Context, d *types.DeliveryAddress) error {
	return s.r.UpdateDeliveryAddress(ctx, d)
}

func (s *userService) UpdateDefaultDeliveryAddress(ctx context.Context, deliveryAddressId string, authId string) error {
	return s.r.UpdateDefaultDeliveryAddress(ctx, deliveryAddressId, authId)
}

func (s *userService) DeleteDeliveryAddress(ctx context.Context, deliveryAddressId string) error {
	return s.r.DeleteDeliveryAddress(ctx, deliveryAddressId)
}
