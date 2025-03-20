package inventory

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/yash91989201/superfast-delivery-api/common/types"
	"github.com/yash91989201/superfast-delivery-api/services/inventory/db/queries"
)

type Repository interface {
	Close() error
	CreateItemStock(ctx context.Context, stock *types.ItemStock) error
	CreateVariantStock(ctx context.Context, stock *types.VariantStock) error
	CreateAddonStock(ctx context.Context, stock *types.AddonStock) error

	GetItemStockByID(ctx context.Context, id string) (*types.ItemStock, error)
	GetVariantStockByID(ctx context.Context, id string) (*types.VariantStock, error)
	GetAddonStockByID(ctx context.Context, id string) (*types.AddonStock, error)

	UpdateItemStock(ctx context.Context, stock *types.ItemStock) error
	UpdateVariantStock(ctx context.Context, stock *types.VariantStock) error
	UpdateAddonStock(ctx context.Context, stock *types.AddonStock) error

	DeleteItemStock(ctx context.Context, id string) error
	DeleteVariantStock(ctx context.Context, id string) error
	DeleteAddonStock(ctx context.Context, id string) error
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

func (r *mysqlRepository) CreateItemStock(ctx context.Context, stock *types.ItemStock) error {
	queryRes, err := r.db.NamedExecContext(ctx, queries.INSERT_ITEM_STOCK, stock)
	if err != nil {
		return err
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to create item stock, 0 rows affected: %w", err)
	}

	return nil
}

func (r *mysqlRepository) CreateVariantStock(ctx context.Context, stock *types.VariantStock) error {
	queryRes, err := r.db.NamedExecContext(ctx, queries.INSERT_VARIANT_STOCK, stock)
	if err != nil {
		return err
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to create variant stock, 0 rows affected: %w", err)
	}

	return nil
}

func (r *mysqlRepository) CreateAddonStock(ctx context.Context, stock *types.AddonStock) error {
	queryRes, err := r.db.NamedExecContext(ctx, queries.INSERT_ADDON_STOCK, stock)
	if err != nil {
		return err
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to create addon stock, 0 rows affected: %w", err)
	}

	return nil
}

func (r *mysqlRepository) GetItemStockByID(ctx context.Context, id string) (*types.ItemStock, error) {
	var stock types.ItemStock
	err := r.db.GetContext(ctx, &stock, queries.GET_ITEM_STOCK_BY_ID, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get item stock: %w", err)
	}
	return &stock, nil
}

func (r *mysqlRepository) GetVariantStockByID(ctx context.Context, id string) (*types.VariantStock, error) {
	var stock types.VariantStock
	err := r.db.GetContext(ctx, &stock, queries.GET_VARIANT_STOCK_BY_ID, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get variant stock: %w", err)
	}
	return &stock, nil
}

func (r *mysqlRepository) GetAddonStockByID(ctx context.Context, id string) (*types.AddonStock, error) {
	var stock types.AddonStock
	err := r.db.GetContext(ctx, &stock, queries.GET_ADDON_STOCK_BY_ID, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get addon stock: %w", err)
	}
	return &stock, nil
}

func (r *mysqlRepository) UpdateItemStock(ctx context.Context, stock *types.ItemStock) error {
	result, err := r.db.NamedExecContext(ctx, queries.UPDATE_ITEM_STOCK, stock)
	if err != nil {
		return fmt.Errorf("failed to update item stock: %w", err)
	}
	return checkRowsAffected(result)
}

func (r *mysqlRepository) UpdateVariantStock(ctx context.Context, stock *types.VariantStock) error {
	result, err := r.db.NamedExecContext(ctx, queries.UPDATE_VARIANT_STOCK, stock)
	if err != nil {
		return fmt.Errorf("failed to update variant stock: %w", err)
	}
	return checkRowsAffected(result)
}

func (r *mysqlRepository) UpdateAddonStock(ctx context.Context, stock *types.AddonStock) error {
	result, err := r.db.NamedExecContext(ctx, queries.UPDATE_ADDON_STOCK, stock)
	if err != nil {
		return fmt.Errorf("failed to update addon stock: %w", err)
	}
	return checkRowsAffected(result)
}

func (r *mysqlRepository) DeleteItemStock(ctx context.Context, id string) error {
	result, err := r.db.ExecContext(ctx, queries.DELETE_ITEM_STOCK, id)
	if err != nil {
		return fmt.Errorf("failed to delete item stock: %w", err)
	}
	return checkRowsAffected(result)
}

func (r *mysqlRepository) DeleteVariantStock(ctx context.Context, id string) error {
	result, err := r.db.ExecContext(ctx, queries.DELETE_VARIANT_STOCK, id)
	if err != nil {
		return fmt.Errorf("failed to delete variant stock: %w", err)
	}
	return checkRowsAffected(result)
}

func (r *mysqlRepository) DeleteAddonStock(ctx context.Context, id string) error {
	result, err := r.db.ExecContext(ctx, queries.DELETE_ADDON_STOCK, id)
	if err != nil {
		return fmt.Errorf("failed to delete addon stock: %w", err)
	}
	return checkRowsAffected(result)
}

func checkRowsAffected(result sql.Result) error {
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}
