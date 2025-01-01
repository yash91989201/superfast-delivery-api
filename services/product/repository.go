package product

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Repository interface {
	Close(ctx context.Context) error
	Ping(ctx context.Context) error
}

type mongoRepository struct {
	db *mongo.Client
}

func NewMongoRepository(dbUrl string) (Repository, error) {
	db, err := mongo.Connect(options.Client().ApplyURI(dbUrl))
	if err != nil {
		return nil, err
	}

	return &mongoRepository{
		db: db,
	}, nil
}

func (r *mongoRepository) Close(ctx context.Context) error {
	return r.db.Disconnect(ctx)
}

func (r *mongoRepository) Ping(ctx context.Context) error {
	return r.db.Ping(ctx, readpref.Primary())
}
