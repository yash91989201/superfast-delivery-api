package product

import (
	"context"
	"errors"
	"fmt"

	"github.com/yash91989201/superfast-delivery-api/common/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Repository interface {
	Close(ctx context.Context) error
	Ping(ctx context.Context) error

	InsertRestaurantMenu(ctx context.Context, m *types.RestaurantMenu) error
	GetRestaurantMenu(ctx context.Context, id string) (*types.RestaurantMenu, error)
	GetRestaurantMenuByShopId(ctx context.Context, shop_id string) ([]*types.RestaurantMenu, error)
}

type mongoRepository struct {
	client           *mongo.Client
	db               *mongo.Database
	restaurantMenu   *mongo.Collection
	menuItem         *mongo.Collection
	retailCategory   *mongo.Collection
	retailItem       *mongo.Collection
	medicineCategory *mongo.Collection
	medicineItem     *mongo.Collection
}

func NewMongoRepository(dbUrl string, dbName string) (Repository, error) {
	mongoClient, err := mongo.Connect(options.Client().ApplyURI(dbUrl))
	if err != nil {
		return nil, err
	}

	mongoDb := mongoClient.Database(dbName)
	return &mongoRepository{
		client:           mongoClient,
		db:               mongoDb,
		restaurantMenu:   mongoDb.Collection("restaurant_menu"),
		menuItem:         mongoDb.Collection("menu_item"),
		retailCategory:   mongoDb.Collection("retail_category"),
		retailItem:       mongoDb.Collection("retail_item"),
		medicineCategory: mongoDb.Collection("medicine_category"),
		medicineItem:     mongoDb.Collection("medicine_item"),
	}, nil
}

func (r *mongoRepository) Close(ctx context.Context) error {
	return r.client.Disconnect(ctx)
}

func (r *mongoRepository) Ping(ctx context.Context) error {
	return r.client.Ping(ctx, readpref.Primary())
}

func (r *mongoRepository) InsertRestaurantMenu(ctx context.Context, m *types.RestaurantMenu) error {
	res, err := r.restaurantMenu.InsertOne(ctx, m)
	if err != nil || !res.Acknowledged {
		return err
	}

	return nil
}

func (r *mongoRepository) GetRestaurantMenu(ctx context.Context, id string) (*types.RestaurantMenu, error) {
	m := &types.RestaurantMenu{}
	res := r.restaurantMenu.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&m)
	if errors.Is(res, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("Restaurant menu not found")
	}

	return m, nil
}

func (r *mongoRepository) GetRestaurantMenuByShopId(ctx context.Context, shop_id string) ([]*types.RestaurantMenu, error) {
	filter := bson.D{{Key: "shop_id", Value: shop_id}}
	cursor, err := r.restaurantMenu.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("Restaurant menu not found: %w", err)
	}

	m := make([]*types.RestaurantMenu, 0)
	if err := cursor.All(ctx, &m); err != nil {
		return nil, fmt.Errorf("Restaurant menu not found: %w", err)
	}

	return m, nil
}
