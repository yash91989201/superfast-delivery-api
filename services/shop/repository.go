package shop

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type (
	Repository interface {
		Close() error
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
