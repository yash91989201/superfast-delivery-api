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
}

type userService struct {
	r Repository
}

func New(r Repository) Service {
	return &userService{r: r}
}

func (s *userService) CreateProfile(ctx context.Context, p *types.CreateProfile) (*types.Profile, error) {
	profile := &types.Profile{
		Id:          cuid2.Generate(),
		Name:        p.Name,
		ImageUrl:    p.ImageUrl,
		Dob:         p.Dob,
		Anniversary: p.Anniversary,
		Gender:      p.Gender,
		AuthId:      p.AuthId,
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
