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

	InsertItemVariant(ctx context.Context, v *types.ItemVariant) error
	InsertItemAddon(ctx context.Context, a *types.ItemAddon) error
	InsertRestaurantMenu(ctx context.Context, m *types.RestaurantMenu) error
	InsertMenuItem(ctx context.Context, i *types.MenuItem) error
	InsertRetailCategory(ctx context.Context, c *types.RetailCategory) error
	InsertRetailItem(ctx context.Context, c *types.RetailItem) error
	InsertMedicineCategory(ctx context.Context, c *types.MedicineCategory) error
	InsertMedicineItem(ctx context.Context, i *types.MedicineItem) error

	GetItemVariant(ctx context.Context, id string) (*types.ItemVariant, error)
	GetItemAddon(ctx context.Context, id string) (*types.ItemAddon, error)
	GetItemVariants(ctx context.Context, itemId string) ([]*types.ItemVariant, error)
	GetItemAddons(ctx context.Context, itemId string) ([]*types.ItemAddon, error)
	GetRestaurantMenu(ctx context.Context, id string) (*types.RestaurantMenu, error)
	ListRestaurantMenu(ctx context.Context, shopId string) ([]*types.RestaurantMenu, error)
	GetRetailCategory(ctx context.Context, id string) (*types.RetailCategory, error)
	ListRetailCategory(ctx context.Context, shopId string) ([]*types.RetailCategory, error)
	GetMedicineCategory(ctx context.Context, id string) (*types.MedicineCategory, error)
	ListMedicineCategory(ctx context.Context, shopId string) ([]*types.MedicineCategory, error)
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

func (r *mongoRepository) GetItemVariant(ctx context.Context, id string) (*types.ItemVariant, error) {
	itemVariant := &types.ItemVariant{}
	err := r.itemVariant.FindOne(ctx, bson.D{{Key: "_id", Value: types.HexToObjectId(id)}}).Decode(&itemVariant)
	if errors.Is(err, mongo.ErrNoDocuments) || err != nil {
		return nil, fmt.Errorf("Item variant not found: %v", err)
	}

	return itemVariant, nil
}

func (r *mongoRepository) GetItemAddon(ctx context.Context, id string) (*types.ItemAddon, error) {
	addon := &types.ItemAddon{}
	err := r.itemAddon.FindOne(ctx, bson.M{"_id": types.HexToObjectId(id)}).Decode(addon)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("item addon not found")
		}
		return nil, fmt.Errorf("failed to fetch item addon: %v", err)
	}

	return addon, nil
}

func (r *mongoRepository) GetItemVariants(ctx context.Context, itemId string) ([]*types.ItemVariant, error) {
	cursor, err := r.itemVariant.Find(ctx, bson.M{"item_id": types.HexToObjectId(itemId)})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch item variants: %v", err)
	}

	defer cursor.Close(ctx)

	var variants []*types.ItemVariant
	if err = cursor.All(ctx, &variants); err != nil {
		return nil, fmt.Errorf("failed to decode item variants: %v", err)
	}

	return variants, nil
}

func (r *mongoRepository) GetItemAddons(ctx context.Context, itemId string) ([]*types.ItemAddon, error) {
	cursor, err := r.itemAddon.Find(ctx, bson.M{"item_id": types.HexToObjectId(itemId)})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch item addons: %v", err)
	}

	defer cursor.Close(ctx)

	var addons []*types.ItemAddon
	if err = cursor.All(ctx, &addons); err != nil {
		return nil, fmt.Errorf("failed to decode item addons: %v", err)
	}

	return addons, nil
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

	var restaurantMenu *types.RestaurantMenu
	if cursor.Next(ctx) {
		if err := cursor.Decode(&restaurantMenu); err != nil {
			return nil, fmt.Errorf("failed to decode result: %v", err)
		}
	}

	if restaurantMenu == nil {
		return nil, fmt.Errorf("restaurant menu not found")
	}

	return restaurantMenu, nil
}

func (r *mongoRepository) ListRestaurantMenu(ctx context.Context, shopId string) ([]*types.RestaurantMenu, error) {
	restaurantMenuByShopId := bson.D{{Key: "$match", Value: bson.D{{Key: "shop_id", Value: shopId}}}}

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
			restaurantMenuByShopId,
			joinMenuItemsByMenuId,
			projectToModel,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to execute aggregation: %v", err)
	}

	defer cursor.Close(ctx)

	var restaurantMenuList []*types.RestaurantMenu
	for cursor.Next(ctx) {
		var restaurantMenu types.RestaurantMenu
		if err := cursor.Decode(&restaurantMenu); err != nil {
			return nil, fmt.Errorf("failed to decode result: %v", err)
		}
		restaurantMenuList = append(restaurantMenuList, &restaurantMenu)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor iteration error: %v", err)
	}

	if len(restaurantMenuList) == 0 {
		return nil, fmt.Errorf("restaurant menu not found")
	}

	return restaurantMenuList, nil
}

func (r *mongoRepository) GetRetailCategory(ctx context.Context, id string) (*types.RetailCategory, error) {
	matchRetailCategory := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: types.HexToObjectId(id)}}}}

	joinRetailItemsByCategoryId := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: "retail_item"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "category_id"},
			{Key: "as", Value: "retail_items"},
		},
	}}

	projectToModel := bson.D{{
		Key: "$project",
		Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "category_name", Value: 1},
			{Key: "shop_id", Value: 1},
			{Key: "created_at", Value: 1},
			{Key: "updated_at", Value: 1},
			{Key: "deleted_at", Value: 1},
			{Key: "retail_items", Value: 1},
		},
	}}

	cursor, err := r.retailCategory.Aggregate(
		ctx,
		mongo.Pipeline{
			matchRetailCategory,
			joinRetailItemsByCategoryId,
			projectToModel,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to execute aggregation: %v", err)
	}

	defer cursor.Close(ctx)

	var retailCategory *types.RetailCategory
	if cursor.Next(ctx) {
		if err := cursor.Decode(&retailCategory); err != nil {
			return nil, fmt.Errorf("failed to decode result: %v", err)
		}
	}

	if retailCategory == nil {
		return nil, fmt.Errorf("retail category not found")
	}

	return retailCategory, nil
}

func (r *mongoRepository) ListRetailCategory(ctx context.Context, shopId string) ([]*types.RetailCategory, error) {
	retailCategoryByShopId := bson.D{{Key: "$match", Value: bson.D{{Key: "shop_id", Value: shopId}}}}

	joinRetailItemsByCategoryId := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: "retail_item"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "category_id"},
			{Key: "as", Value: "retail_items"},
		},
	}}

	projectToModel := bson.D{{
		Key: "$project",
		Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "category_name", Value: 1},
			{Key: "shop_id", Value: 1},
			{Key: "created_at", Value: 1},
			{Key: "updated_at", Value: 1},
			{Key: "deleted_at", Value: 1},
			{Key: "retail_items", Value: 1},
		},
	}}

	cursor, err := r.retailCategory.Aggregate(
		ctx,
		mongo.Pipeline{
			retailCategoryByShopId,
			joinRetailItemsByCategoryId,
			projectToModel,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to execute aggregation: %v", err)
	}

	defer cursor.Close(ctx)

	var retailCategoryList []*types.RetailCategory
	for cursor.Next(ctx) {
		var retailCategory types.RetailCategory
		if err := cursor.Decode(&retailCategory); err != nil {
			return nil, fmt.Errorf("failed to decode result: %v", err)
		}

		retailCategoryList = append(retailCategoryList, &retailCategory)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor iteration error: %v", err)
	}

	if len(retailCategoryList) == 0 {
		return nil, fmt.Errorf("restaurant menu not found")
	}

	return retailCategoryList, nil
}

func (r *mongoRepository) GetMedicineCategory(ctx context.Context, id string) (*types.MedicineCategory, error) {
	matchMedicineCategory := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: types.HexToObjectId(id)}}}}

	joinMedicineItemsByCategoryId := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: "medicine_item"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "category_id"},
			{Key: "as", Value: "medicine_items"},
		},
	}}

	projectToModel := bson.D{{
		Key: "$project",
		Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "category_name", Value: 1},
			{Key: "shop_id", Value: 1},
			{Key: "created_at", Value: 1},
			{Key: "updated_at", Value: 1},
			{Key: "deleted_at", Value: 1},
			{Key: "medicine_items", Value: 1},
		},
	}}

	cursor, err := r.medicineCategory.Aggregate(
		ctx,
		mongo.Pipeline{
			matchMedicineCategory,
			joinMedicineItemsByCategoryId,
			projectToModel,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to execute aggregation: %v", err)
	}

	defer cursor.Close(ctx)

	var medicineCategory *types.MedicineCategory
	if cursor.Next(ctx) {
		if err := cursor.Decode(&medicineCategory); err != nil {
			return nil, fmt.Errorf("failed to decode result: %v", err)
		}
	}

	if medicineCategory == nil {
		return nil, fmt.Errorf("retail category not found")
	}

	return medicineCategory, nil
}

func (r *mongoRepository) ListMedicineCategory(ctx context.Context, shopId string) ([]*types.MedicineCategory, error) {
	medicineCategoryByShopId := bson.D{{Key: "$match", Value: bson.D{{Key: "shop_id", Value: shopId}}}}

	joinMedicineItemsByCategoryId := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: "medicine_item"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "category_id"},
			{Key: "as", Value: "medicine_items"},
		},
	}}

	projectToModel := bson.D{{
		Key: "$project",
		Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "category_name", Value: 1},
			{Key: "shop_id", Value: 1},
			{Key: "created_at", Value: 1},
			{Key: "updated_at", Value: 1},
			{Key: "deleted_at", Value: 1},
			{Key: "medicine_items", Value: 1},
		},
	}}

	cursor, err := r.medicineCategory.Aggregate(
		ctx,
		mongo.Pipeline{
			medicineCategoryByShopId,
			joinMedicineItemsByCategoryId,
			projectToModel,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to execute aggregation: %v", err)
	}

	defer cursor.Close(ctx)

	var medicineCategoryList []*types.MedicineCategory
	for cursor.Next(ctx) {
		var medicineCategory types.MedicineCategory
		if err := cursor.Decode(&medicineCategory); err != nil {
			return nil, fmt.Errorf("failed to decode result: %v", err)
		}

		medicineCategoryList = append(medicineCategoryList, &medicineCategory)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor iteration error: %v", err)
	}

	if len(medicineCategoryList) == 0 {
		return nil, fmt.Errorf("restaurant menu not found")
	}

	return medicineCategoryList, nil
}
