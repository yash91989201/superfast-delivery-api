package user

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/yash91989201/superfast-delivery-api/common/types"
	"github.com/yash91989201/superfast-delivery-api/services/user/db/queries"
)

type Repository interface {
	Close() error
	CreateProfile(ctx context.Context, p *types.Profile) error
	GetProfileById(ctx context.Context, id string) (*types.Profile, error)
	GetProfileByAuthId(ctx context.Context, authID string) (*types.Profile, error)
	UpdateProfile(ctx context.Context, p *types.Profile) error
	DeleteProfile(ctx context.Context, id string) error

	CreateDeliveryAddress(ctx context.Context, d *types.DeliveryAddress) error
	GetDeliveryAddressById(ctx context.Context, id string) (*types.DeliveryAddress, error)
	GetDefaultDeliveryAddress(ctx context.Context, authID string) (*types.DeliveryAddress, error)
	GetDeliveryAddresses(ctx context.Context, authID string) ([]*types.DeliveryAddress, error)
	UpdateDeliveryAddress(ctx context.Context, d *types.DeliveryAddress) error
	UpdateDefaultDeliveryAddress(ctx context.Context, deliveryAddressId string, authId string) error
	DeleteDeliveryAddress(ctx context.Context, id string) error
}

type mysqlRepository struct {
	db *sqlx.DB
}

func NewMysqlRepository(dbUrl string) (Repository, error) {
	db, err := sqlx.Open("mysql", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
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
		return fmt.Errorf("failed to create profile: %w", err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("failed to create profile, 0 rows affected: %w", err)
	}

	return nil
}

func (r *mysqlRepository) GetProfileById(ctx context.Context, id string) (*types.Profile, error) {
	p := &types.Profile{}
	if err := r.db.GetContext(ctx, p, queries.GET_PROFILE_BY_ID, id); err != nil {
		return nil, fmt.Errorf("profile not found: %w", err)
	}

	return p, nil
}

func (r *mysqlRepository) GetProfileByAuthId(ctx context.Context, authID string) (*types.Profile, error) {
	p := &types.Profile{}
	if err := r.db.GetContext(ctx, p, queries.GET_PROFILE_BY_AUTH_ID, authID); err != nil {
		return nil, fmt.Errorf("profile not found: %w", err)
	}

	return p, nil
}

func (r *mysqlRepository) UpdateProfile(ctx context.Context, p *types.Profile) error {
	queryRes, err := r.db.NamedExecContext(ctx, queries.UPDATE_PROFILE, p)
	if err != nil {
		return fmt.Errorf("failed to update profile id %s : %w", p.ID, err)
	}
	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("failed to update profile, 0 rows affected: %w", err)
	}

	return nil
}

func (r *mysqlRepository) DeleteProfile(ctx context.Context, id string) error {
	queryRes, err := r.db.ExecContext(ctx, queries.DELETE_PROFILE, id)
	if err != nil {
		return fmt.Errorf("failed to delete profile id %s : %w", id, err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("failed to delete profile, 0 rows affected: %w", err)
	}

	return nil
}

func (r *mysqlRepository) CreateDeliveryAddress(ctx context.Context, d *types.DeliveryAddress) error {
	queryRes, err := r.db.NamedExecContext(ctx, queries.CREATE_DELIVERY_ADDRESS, d)
	if err != nil {
		return fmt.Errorf("failed to create delivery address: %w", err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("failed to create profile, 0 rows affected: %w", err)
	}

	return nil
}

func (r *mysqlRepository) DeleteDeliveryAddress(ctx context.Context, id string) error {
	queryRes, err := r.db.ExecContext(ctx, queries.DELETE_DELIVERY_ADDRESS, id)
	if err != nil {
		return fmt.Errorf("failed to delete delivery address: %w", err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("failed to delete delivery address, 0 rows affected: %w", err)
	}

	return nil
}

func (r *mysqlRepository) GetDeliveryAddressById(ctx context.Context, id string) (*types.DeliveryAddress, error) {
	var d types.DeliveryAddress
	err := r.db.GetContext(ctx, &d, queries.GET_DELIVERY_ADDRESS_BY_ID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("failed to get delivery address, no rows found")
		}

		return nil, fmt.Errorf("failed to get delivery address by id: %w", err)
	}
	return &d, nil
}

func (r *mysqlRepository) GetDefaultDeliveryAddress(ctx context.Context, authID string) (*types.DeliveryAddress, error) {
	var d types.DeliveryAddress
	err := r.db.GetContext(ctx, &d, queries.GET_DEFAULT_DELIVERY_ADDRESS, authID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("failed to get delivery address, no rows found")
		}

		return nil, fmt.Errorf("failed to get delivery address by auth id: %w", err)
	}

	return &d, nil
}

func (r *mysqlRepository) GetDeliveryAddresses(ctx context.Context, authID string) ([]*types.DeliveryAddress, error) {
	var addresses []*types.DeliveryAddress
	err := r.db.SelectContext(ctx, &addresses, queries.GET_DELIVERY_ADDRESSES_BY_AUTH_ID, authID)
	if err != nil {
		return nil, fmt.Errorf("failed to get delivery addresses: %w", err)
	}
	return addresses, nil
}

func (r *mysqlRepository) UpdateDeliveryAddress(ctx context.Context, d *types.DeliveryAddress) error {
	_, err := r.db.NamedExecContext(ctx, queries.UPDATE_DELIVERY_ADDRESS, d)
	if err != nil {
		return fmt.Errorf("failed to update delivery address: %w", err)
	}
	return nil
}

func (r *mysqlRepository) UpdateDefaultDeliveryAddress(ctx context.Context, deliveryAddressId string, authId string) error {
	queryRes, err := r.db.ExecContext(ctx, queries.UPDATE_DEFAULT_DELIVERY_ADDRESS, deliveryAddressId, authId)
	if err != nil {
		return err
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected != 1 || err != nil {
		return err
	}

	return nil
}
