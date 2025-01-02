package product

import (
	"context"
	"fmt"
	"log"

	"github.com/yash91989201/superfast-delivery-api/common/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Repository interface {
	Close(ctx context.Context) error
	Ping(ctx context.Context) error

	InsertItemVariant(ctx context.Context, v *types.ItemVariant) error
	InsertItemAddon(ctx context.Context, a *types.ItemAddon) error
	InsertRestaurantMenu(ctx context.Context, m *types.RestaurantMenu) error
	InsertMenuItem(ctx context.Context, i *types.MenuItem) error
	InsertRetailCategory(ctx context.Context, c *types.RetailCategory) error
	InsertRetailItem(ctx context.Context, c *types.RetailItem) error
	InsertMedicineCategory(ctx context.Context, c *types.MedicineCategory) error
	InsertMedicineItem(ctx context.Context, i *types.MedicineItem) error

	GetRestaurantMenu(ctx context.Context, id string) (*types.RestaurantMenu, error)
	ListRestaurantMenu(ctx context.Context, shopId string) ([]*types.RestaurantMenu, error)
}

type mongoRepository struct {
	client           *mongo.Client
	db               *mongo.Database
	itemVariant      *mongo.Collection
	itemAddon        *mongo.Collection
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
		itemVariant:      mongoDb.Collection("item_variant"),
		itemAddon:        mongoDb.Collection("item_addon"),
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

func (r *mongoRepository) InsertItemVariant(ctx context.Context, v *types.ItemVariant) error {
	res, err := r.itemVariant.InsertOne(ctx, v)
	if err != nil || !res.Acknowledged {
		return err
	}

	return nil
}

func (r *mongoRepository) InsertItemAddon(ctx context.Context, a *types.ItemAddon) error {
	res, err := r.itemAddon.InsertOne(ctx, a)
	if err != nil || !res.Acknowledged {
		return err
	}

	return nil
}

func (r *mongoRepository) InsertRestaurantMenu(ctx context.Context, m *types.RestaurantMenu) error {
	res, err := r.restaurantMenu.InsertOne(ctx, m)
	log.Print("res,err")
	log.Println(res)
	log.Println(err)
	if err != nil || !res.Acknowledged {
		return err
	}

	return nil
}

func (r *mongoRepository) InsertMenuItem(ctx context.Context, i *types.MenuItem) error {
	res, err := r.menuItem.InsertOne(ctx, i)
	if err != nil || !res.Acknowledged {
		return err
	}

	return nil
}

func (r *mongoRepository) InsertMedicineCategory(ctx context.Context, c *types.MedicineCategory) error {
	res, err := r.medicineCategory.InsertOne(ctx, c)
	if err != nil || !res.Acknowledged {
		return err
	}

	return nil
}

func (r *mongoRepository) InsertMedicineItem(ctx context.Context, i *types.MedicineItem) error {
	res, err := r.medicineItem.InsertOne(ctx, i)
	if err != nil || !res.Acknowledged {
		return err
	}

	return nil
}

func (r *mongoRepository) InsertRetailCategory(ctx context.Context, c *types.RetailCategory) error {
	res, err := r.retailCategory.InsertOne(ctx, c)
	if err != nil || !res.Acknowledged {
		return err
	}

	return nil
}

func (r *mongoRepository) InsertRetailItem(ctx context.Context, i *types.RetailItem) error {
	res, err := r.retailItem.InsertOne(ctx, i)
	if err != nil || !res.Acknowledged {
		return err
	}

	return nil
}

func (r *mongoRepository) GetRestaurantMenu(ctx context.Context, id string) (*types.RestaurantMenu, error) {
	matchRestaurantMenu := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: types.HexToObjectId(id)}}}}

	joinMenuItemsByMenuId := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: "menu_item"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "menu_id"},
			{Key: "as", Value: "menu_items"},
		},
	}}

	projectToModel := bson.D{{
		Key: "$project",
		Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "menu_name", Value: 1},
			{Key: "shop_id", Value: 1},
			{Key: "created_at", Value: 1},
			{Key: "updated_at", Value: 1},
			{Key: "deleted_at", Value: 1},
			{Key: "menu_items", Value: 1},
		},
	}}

	cursor, err := r.restaurantMenu.Aggregate(
		ctx,
		mongo.Pipeline{
			matchRestaurantMenu,
			joinMenuItemsByMenuId,
			projectToModel,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to execute aggregation: %v", err)
	}

	defer cursor.Close(ctx)

	var restaurantMenu []*types.RestaurantMenu
	if cursor.Next(ctx) {
		if err := cursor.Decode(&restaurantMenu); err != nil {
			return nil, fmt.Errorf("failed to decode result: %v", err)
		}
	}

	if len(restaurantMenu) == 0 {
		return nil, fmt.Errorf("restaurant menu not found")
	}

	return restaurantMenu[0], nil
}

func (r *mongoRepository) ListRestaurantMenu(ctx context.Context, shopId string) ([]*types.RestaurantMenu, error) {
	filter := bson.D{{Key: "shop_id", Value: shopId}}
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
