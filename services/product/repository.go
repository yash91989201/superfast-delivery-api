package product

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Close() error
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
