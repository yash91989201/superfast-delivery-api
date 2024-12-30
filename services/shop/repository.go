package shop

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/yash91989201/superfast-delivery-api/common/types"
	"github.com/yash91989201/superfast-delivery-api/services/shop/db/queries"
)

type (
	Repository interface {
		Close() error
		InsertShop(context.Context, *types.Shop) error
	}

	pgRepository struct {
		db *sqlx.DB
	}
)

func NewPgRepository(dbUrl string) (Repository, error) {
	db, err := sqlx.Open("postgres", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("Database connection failed: %w", err)
	}

	return &pgRepository{
		db: db,
	}, nil
}

func (r *pgRepository) Close() error {
	return r.db.Close()
}

func (r *pgRepository) InsertShop(ctx context.Context, shop *types.Shop) error {
	err := r.execTx(ctx, func(tx *sqlx.Tx) error {
		if err := insertShop(ctx, tx, shop); err != nil {
			return err
		}

		if err := insertShopAddress(ctx, tx, &shop.Address); err != nil {
			return err
		}

		if err := insertShopContact(ctx, tx, &shop.Contact); err != nil {
			return err
		}

		for _, t := range shop.Timing {
			if err := insertShopTiming(ctx, tx, &t); err != nil {
				return err
			}
		}

		for _, i := range shop.Image {
			if err := insertShopImage(ctx, tx, &i); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func insertShop(ctx context.Context, tx *sqlx.Tx, s *types.Shop) error {
	queryRes, err := tx.NamedExecContext(ctx, queries.CREATE_SHOP, s)
	if err != nil {
		return fmt.Errorf("Failed to insert shop: %w", err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to insert shop, 0 rows affected: %w", err)
	}

	return nil
}

func insertShopAddress(ctx context.Context, tx *sqlx.Tx, a *types.ShopAddress) error {
	queryRes, err := tx.NamedExecContext(ctx, queries.CREATE_SHOP_ADDRESS, a)
	if err != nil {
		return fmt.Errorf("Failed to insert address: %w", err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to insert addres, 0 rows affected: %w", err)
	}

	return nil
}

func insertShopContact(ctx context.Context, tx *sqlx.Tx, c *types.ShopContact) error {
	queryRes, err := tx.NamedExecContext(ctx, queries.CREATE_SHOP_CONTACT, c)
	if err != nil {
		return fmt.Errorf("Failed to insert contact: %w", err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to insert contact, 0 rows affected: %w", err)
	}

	return nil
}

func insertShopImage(ctx context.Context, tx *sqlx.Tx, i *types.ShopImage) error {
	queryRes, err := tx.NamedExecContext(ctx, queries.CREATE_SHOP_IMAGE, i)
	if err != nil {
		return fmt.Errorf("Failed to insert image: %w", err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to insert image, 0 rows affected: %w", err)
	}

	return nil
}

func insertShopTiming(ctx context.Context, tx *sqlx.Tx, t *types.ShopTiming) error {
	queryRes, err := tx.NamedExecContext(ctx, queries.CREATE_SHOP_TIMING, t)
	if err != nil {
		return fmt.Errorf("Failed to insert timing: %w", err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to insert timing, 0 rows affected: %w", err)
	}

	return nil
}

func (r *pgRepository) execTx(ctx context.Context, fn func(*sqlx.Tx) error) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("Failed to start transaction: %w", err)
	}

	if err = fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("Failed to rollback transaction: %w", rbErr)
		}
		return fmt.Errorf("Failed to execute transaction: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("Failed to commit transaction: %w", err)
	}

	return nil
}
