package inventory

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/yash91989201/superfast-delivery-api/common/types"
	"github.com/yash91989201/superfast-delivery-api/services/inventory/db/queries"
)

type Repository interface {
	Close() error
	InsertItemStock(ctx context.Context, stock *types.ItemStock) error
	InsertVariantStock(ctx context.Context, stock *types.VariantStock) error
	InsertAddonStock(ctx context.Context, stock *types.AddonStock) error
}

type mysqlRepository struct {
	db *sqlx.DB
}

func NewMysqlRepository(dbUrl string) (Repository, error) {
	db, err := sqlx.Open("mysql", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %w", err)
	}

	return &mysqlRepository{
		db: db,
	}, nil
}

func (r *mysqlRepository) Close() error {
	return r.db.Close()
}

func (r *mysqlRepository) InsertItemStock(ctx context.Context, stock *types.ItemStock) error {
	queryRes, err := r.db.NamedExecContext(ctx, queries.INSERT_ITEM_STOCK, stock)
	if err != nil {
		return err
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to create item stock, 0 rows affected: %w", err)
	}

	return nil
}

func (r *mysqlRepository) InsertVariantStock(ctx context.Context, stock *types.VariantStock) error {
	queryRes, err := r.db.NamedExecContext(ctx, queries.INSERT_VARIANT_STOCK, stock)
	if err != nil {
		return err
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to create variant stock, 0 rows affected: %w", err)
	}

	return nil
}

func (r *mysqlRepository) InsertAddonStock(ctx context.Context, stock *types.AddonStock) error {
	queryRes, err := r.db.NamedExecContext(ctx, queries.INSERT_ADDON_STOCK, stock)
	if err != nil {
		return err
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to create addon stock, 0 rows affected: %w", err)
	}

	return nil
}
