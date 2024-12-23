package user

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/yash91989201/superfast-delivery-api/common/types"
	"github.com/yash91989201/superfast-delivery-api/services/authentication/db/queries"
)

type Repository interface {
	Close() error
	CreateProfile(ctx context.Context, p *types.Profile) error
	GetProfileById(ctx context.Context, id string) (*types.Profile, error)
	GetProfileByAuthId(ctx context.Context, auth_id string) (*types.Profile, error)
	UpdateProfile(ctx context.Context, p *types.Profile) error
	DeleteProfile(ctx context.Context, id string) error
}

type mysqlRepository struct {
	db *sqlx.DB
}

func NewMysqlRepository(dbUrl string) (Repository, error) {
	db, err := sqlx.Open("mysql", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("Database connection failed: %w", err)
	}

	return &mysqlRepository{
		db: db,
	}, nil
}

func (r *mysqlRepository) Close() error {
	return r.db.Close()
}

func (r *mysqlRepository) CreateProfile(ctx context.Context, p *types.Profile) error {
	queryRes, err := r.db.NamedExecContext(ctx, queries.CREATE_PROFILE, p)
	if err != nil {
		return fmt.Errorf("Failed to create profile: %w", err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to create profile, 0 rows affected: %w", err)
	}

	return nil
}

func (r *mysqlRepository) GetProfileById(ctx context.Context, id string) (*types.Profile, error) {
	p := &types.Profile{}
	if err := r.db.GetContext(ctx, p, queries.GET_PROFILE_BY_ID, id); err != nil {
		return nil, fmt.Errorf("Profile not found: %w", err)
	}

	return p, nil
}

func (r *mysqlRepository) GetProfileByAuthId(ctx context.Context, auth_id string) (*types.Profile, error) {
	p := &types.Profile{}
	if err := r.db.GetContext(ctx, p, queries.GET_PROFILE_BY_AUTH_ID, auth_id); err != nil {
		return nil, fmt.Errorf("Profile not found: %w", err)
	}

	return p, nil
}

func (r *mysqlRepository) UpdateProfile(ctx context.Context, p *types.Profile) error {
	queryRes, err := r.db.NamedExecContext(ctx, queries.UPDATE_PROFILE, p)
	if err != nil {
		return fmt.Errorf("Failed to update profile id %s : %w", p.Id, err)
	}
	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to update profile, 0 rows affected: %w", err)
	}

	return nil
}

func (r *mysqlRepository) DeleteProfile(ctx context.Context, id string) error {
	queryRes, err := r.db.ExecContext(ctx, queries.DELETE_PROFILE, id)
	if err != nil {
		return fmt.Errorf("Failed to delete profile id %s : %w", id, err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to delete profile, 0 rows affected: %w", err)
	}

	return nil
}
