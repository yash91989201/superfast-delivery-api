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
		InsertShop(ctx context.Context, shop *types.Shop) error

		GetShopInfo(ctx context.Context, id string) (*types.ShopInfo, error)
		GetShopInfoByOwnerId(ctx context.Context, ownerId string) (*types.ShopInfo, error)
		GetShop(ctx context.Context, id string) (*types.Shop, error)
		GetShopByOwnerId(ctx context.Context, ownerId string) (*types.Shop, error)
		GetShopAddress(ctx context.Context, id string) (*types.ShopAddress, error)
		GetShopAddressByShopId(ctx context.Context, shopId string) (*types.ShopAddress, error)
		GetShopContact(ctx context.Context, id string) (*types.ShopContact, error)
		GetShopContactByShopId(ctx context.Context, shopId string) (*types.ShopContact, error)
		GetShopTiming(ctx context.Context, id string) (*types.ShopTiming, error)
		GetShopTimings(ctx context.Context, shopId string) ([]*types.ShopTiming, error)
		GetShopImage(ctx context.Context, id string) (*types.ShopImage, error)
		GetShopImages(ctx context.Context, shopId string) ([]*types.ShopImage, error)
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

		if err := insertShopAddress(ctx, tx, shop.Address); err != nil {
			return err
		}

		if err := insertShopContact(ctx, tx, shop.Contact); err != nil {
			return err
		}

		for _, t := range shop.Timing {
			if err := insertShopTiming(ctx, tx, t); err != nil {
				return err
			}
		}

		for _, i := range shop.Image {
			if err := insertShopImage(ctx, tx, i); err != nil {
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

func (r *pgRepository) GetShopInfo(ctx context.Context, id string) (*types.ShopInfo, error) {
	shopInfo := &types.ShopInfo{}
	if err := r.db.GetContext(ctx, shopInfo, queries.GET_SHOP, id); err != nil {
		return nil, fmt.Errorf("Failed to get shop info: %w", err)
	}

	return shopInfo, nil
}

func (r *pgRepository) GetShopInfoByOwnerId(ctx context.Context, ownerId string) (*types.ShopInfo, error) {
	shopInfo := &types.ShopInfo{}
	if err := r.db.GetContext(ctx, shopInfo, queries.GET_SHOP_BY_OWNER_ID, ownerId); err != nil {
		return nil, fmt.Errorf("Failed to get shop info by owner_id: %w", err)
	}

	return shopInfo, nil
}

func (r *pgRepository) GetShop(ctx context.Context, id string) (*types.Shop, error) {
	return nil, nil
}

func (r *pgRepository) GetShopByOwnerId(ctx context.Context, ownerId string) (*types.Shop, error) {
	return nil, nil
}

func (r *pgRepository) GetShopAddress(ctx context.Context, id string) (*types.ShopAddress, error) {
	shopAddress := &types.ShopAddress{}
	if err := r.db.GetContext(ctx, shopAddress, queries.GET_SHOP_ADDRESS, id); err != nil {
		return nil, fmt.Errorf("Failed to get shop address: %w", err)
	}

	return shopAddress, nil
}

func (r *pgRepository) GetShopAddressByShopId(ctx context.Context, shopId string) (*types.ShopAddress, error) {
	shopAddress := &types.ShopAddress{}
	if err := r.db.GetContext(ctx, shopAddress, queries.GET_SHOP_ADDRESS_BY_SHOP_ID, shopId); err != nil {
		return nil, fmt.Errorf("Failed to get shop address by shop_id: %w", err)
	}

	return shopAddress, nil
}

func (r *pgRepository) GetShopContact(ctx context.Context, id string) (*types.ShopContact, error) {
	shopContact := &types.ShopContact{}
	if err := r.db.GetContext(ctx, shopContact, queries.GET_SHOP_CONTACT, id); err != nil {
		return nil, fmt.Errorf("Failed to get shop contact: %w", err)
	}

	return shopContact, nil
}

func (r *pgRepository) GetShopContactByShopId(ctx context.Context, shopId string) (*types.ShopContact, error) {
	shopContact := &types.ShopContact{}
	if err := r.db.GetContext(ctx, shopContact, queries.GET_SHOP_CONTACT_BY_SHOP_ID, shopId); err != nil {
		return nil, fmt.Errorf("Failed to get shop contact by shop_id: %w", err)
	}

	return shopContact, nil
}

func (r *pgRepository) GetShopTiming(ctx context.Context, id string) (*types.ShopTiming, error) {
	shopTiming := &types.ShopTiming{}
	if err := r.db.GetContext(ctx, shopTiming, queries.GET_SHOP_TIMING, id); err != nil {
		return nil, fmt.Errorf("Failed to get shop timing: %w", err)
	}

	return shopTiming, nil
}

func (r *pgRepository) GetShopTimings(ctx context.Context, shopId string) ([]*types.ShopTiming, error) {
	var shopTimings []*types.ShopTiming
	if err := r.db.SelectContext(ctx, &shopTimings, queries.GET_SHOP_TIMINGS, shopId); err != nil {
		return nil, fmt.Errorf("Failed to get shop timings: %w", err)
	}

	return shopTimings, nil
}

func (r *pgRepository) GetShopImage(ctx context.Context, id string) (*types.ShopImage, error) {
	shopImage := &types.ShopImage{}
	if err := r.db.GetContext(ctx, shopImage, queries.GET_SHOP_IMAGE, id); err != nil {
		return nil, fmt.Errorf("Failed to get shop image: %w", err)
	}

	return shopImage, nil
}

func (r *pgRepository) GetShopImages(ctx context.Context, shopId string) ([]*types.ShopImage, error) {
	var shopImages []*types.ShopImage
	if err := r.db.SelectContext(ctx, &shopImages, queries.GET_SHOP_IMAGES, shopId); err != nil {
		return nil, fmt.Errorf("Failed to get shop images: %w", err)
	}

	return shopImages, nil
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
