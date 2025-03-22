package product

import (
	"context"
	"fmt"
	"reflect"

	"github.com/yash91989201/superfast-delivery-api/common/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Repository interface {
	Close(ctx context.Context) error
	Ping(ctx context.Context) error

	CreateItemVariant(ctx context.Context, variant *types.ItemVariant) error
	CreateItemAddon(ctx context.Context, addon *types.ItemAddon) error
	CreateRestaurantMenu(ctx context.Context, menu *types.RestaurantMenu) error
	CreateMenuItem(ctx context.Context, item *types.MenuItem) error
	CreateRetailCategory(ctx context.Context, category *types.RetailCategory) error
	CreateRetailItem(ctx context.Context, item *types.RetailItem) error
	CreateMedicineCategory(ctx context.Context, category *types.MedicineCategory) error
	CreateMedicineItem(ctx context.Context, item *types.MedicineItem) error

	GetItemVariant(ctx context.Context, variantID bson.ObjectID) (*types.ItemVariant, error)
	GetItemAddon(ctx context.Context, addonID bson.ObjectID) (*types.ItemAddon, error)
	GetRestaurantMenu(ctx context.Context, menuID bson.ObjectID) (*types.RestaurantMenu, error)
	GetMenuItem(ctx context.Context, itemID bson.ObjectID) (*types.MenuItem, error)
	GetRetailCategory(ctx context.Context, categoryID bson.ObjectID) (*types.RetailCategory, error)
	GetRetailItem(ctx context.Context, itemID bson.ObjectID) (*types.RetailItem, error)
	GetMedicineCategory(ctx context.Context, categoryID bson.ObjectID) (*types.MedicineCategory, error)
	GetMedicineItem(ctx context.Context, itemID bson.ObjectID) (*types.MedicineItem, error)

	ListItemVariant(ctx context.Context, itemID bson.ObjectID) ([]*types.ItemVariant, error)
	ListItemAddon(ctx context.Context, itemID bson.ObjectID) ([]*types.ItemAddon, error)
	ListRestaurantMenu(ctx context.Context, shopID string) ([]*types.RestaurantMenu, error)
	ListMenuItem(ctx context.Context, menuID bson.ObjectID) ([]*types.MenuItem, error)
	ListRetailCategory(ctx context.Context, shopID string) ([]*types.RetailCategory, error)
	ListRetailItem(ctx context.Context, categoryID bson.ObjectID) ([]*types.RetailItem, error)
	ListMedicineCategory(ctx context.Context, shopID string) ([]*types.MedicineCategory, error)
	ListMedicineItem(ctx context.Context, categoryID bson.ObjectID) ([]*types.MedicineItem, error)

	UpdateItemVariant(ctx context.Context, variant *types.UpdateItemVariant) error
	UpdateItemAddon(ctx context.Context, addon *types.UpdateItemAddon) error
	UpdateRestaurantMenu(ctx context.Context, menu *types.UpdateRestaurantMenu) error
	UpdateMenuItem(ctx context.Context, item *types.UpdateMenuItem) error
	UpdateRetailCategory(ctx context.Context, category *types.UpdateRetailCategory) error
	UpdateRetailItem(ctx context.Context, item *types.UpdateRetailItem) error
	UpdateMedicineCategory(ctx context.Context, category *types.UpdateMedicineCategory) error
	UpdateMedicineItem(ctx context.Context, item *types.UpdateMedicineItem) error

	DeleteItemVariant(ctx context.Context, variantID bson.ObjectID) error
	DeleteItemAddon(ctx context.Context, addonID bson.ObjectID) error
	DeleteRestaurantMenu(ctx context.Context, menuID bson.ObjectID) error
	DeleteMenuItem(ctx context.Context, itemID bson.ObjectID) error
	DeleteRetailCategory(ctx context.Context, categoryID bson.ObjectID) error
	DeleteRetailItem(ctx context.Context, itemID bson.ObjectID) error
	DeleteMedicineCategory(ctx context.Context, categoryID bson.ObjectID) error
	DeleteMedicineItem(ctx context.Context, itemID bson.ObjectID) error
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

func (r *mongoRepository) CreateItemVariant(ctx context.Context, variant *types.ItemVariant) error {
	_, err := r.itemVariant.InsertOne(ctx, variant)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) CreateItemAddon(ctx context.Context, addon *types.ItemAddon) error {
	_, err := r.itemAddon.InsertOne(ctx, addon)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) CreateRestaurantMenu(ctx context.Context, menu *types.RestaurantMenu) error {
	_, err := r.restaurantMenu.InsertOne(ctx, menu)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) CreateMenuItem(ctx context.Context, item *types.MenuItem) error {
	_, err := r.menuItem.InsertOne(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) CreateRetailCategory(ctx context.Context, category *types.RetailCategory) error {
	_, err := r.retailCategory.InsertOne(ctx, category)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) CreateRetailItem(ctx context.Context, item *types.RetailItem) error {
	_, err := r.retailItem.InsertOne(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) CreateMedicineCategory(ctx context.Context, category *types.MedicineCategory) error {
	_, err := r.medicineCategory.InsertOne(ctx, category)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) CreateMedicineItem(ctx context.Context, item *types.MedicineItem) error {
	_, err := r.medicineItem.InsertOne(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) GetItemVariant(ctx context.Context, variantID bson.ObjectID) (*types.ItemVariant, error) {
	var variant types.ItemVariant
	err := r.itemVariant.FindOne(ctx, bson.M{"_id": variantID}).Decode(&variant)
	if err != nil {
		return nil, err
	}

	return &variant, nil
}

func (r *mongoRepository) GetItemAddon(ctx context.Context, addonID bson.ObjectID) (*types.ItemAddon, error) {
	var addon types.ItemAddon
	err := r.itemAddon.FindOne(ctx, bson.M{"_id": addonID}).Decode(&addon)
	if err != nil {
		return nil, err
	}

	return &addon, nil
}

func (r *mongoRepository) GetRestaurantMenu(ctx context.Context, menuID bson.ObjectID) (*types.RestaurantMenu, error) {
	var menu types.RestaurantMenu
	err := r.restaurantMenu.FindOne(ctx, bson.M{"_id": menuID}).Decode(&menu)
	if err != nil {
		return nil, err
	}

	return &menu, nil
}

func (r *mongoRepository) GetMenuItem(ctx context.Context, itemID bson.ObjectID) (*types.MenuItem, error) {
	var item types.MenuItem
	err := r.menuItem.FindOne(ctx, bson.M{"_id": itemID}).Decode(&item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *mongoRepository) GetRetailCategory(ctx context.Context, categoryID bson.ObjectID) (*types.RetailCategory, error) {
	var category types.RetailCategory
	err := r.retailCategory.FindOne(ctx, bson.M{"_id": categoryID}).Decode(&category)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *mongoRepository) GetRetailItem(ctx context.Context, itemID bson.ObjectID) (*types.RetailItem, error) {
	var item types.RetailItem
	err := r.retailItem.FindOne(ctx, bson.M{"_id": itemID}).Decode(&item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *mongoRepository) GetMedicineCategory(ctx context.Context, categoryID bson.ObjectID) (*types.MedicineCategory, error) {
	var category types.MedicineCategory
	err := r.medicineCategory.FindOne(ctx, bson.M{"_id": categoryID}).Decode(&category)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *mongoRepository) GetMedicineItem(ctx context.Context, itemID bson.ObjectID) (*types.MedicineItem, error) {
	var item types.MedicineItem
	err := r.medicineItem.FindOne(ctx, bson.M{"_id": itemID}).Decode(&item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *mongoRepository) ListItemVariant(ctx context.Context, itemID bson.ObjectID) ([]*types.ItemVariant, error) {
	var variants []*types.ItemVariant
	cursor, err := r.itemVariant.Find(ctx, bson.M{"item_id": itemID})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &variants); err != nil {
		return nil, err
	}

	return variants, nil
}

func (r *mongoRepository) ListItemAddon(ctx context.Context, itemID bson.ObjectID) ([]*types.ItemAddon, error) {
	var addons []*types.ItemAddon
	cursor, err := r.itemAddon.Find(ctx, bson.M{"item_id": itemID})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &addons); err != nil {
		return nil, err
	}

	return addons, nil
}

func (r *mongoRepository) ListRestaurantMenu(ctx context.Context, shopID string) ([]*types.RestaurantMenu, error) {
	var menus []*types.RestaurantMenu
	cursor, err := r.restaurantMenu.Find(ctx, bson.M{"shop_id": shopID})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &menus); err != nil {
		return nil, err
	}

	return menus, nil
}

func (r *mongoRepository) ListMenuItem(ctx context.Context, menuID bson.ObjectID) ([]*types.MenuItem, error) {
	var items []*types.MenuItem
	cursor, err := r.menuItem.Find(ctx, bson.M{"menu_id": menuID})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *mongoRepository) ListRetailCategory(ctx context.Context, shopID string) ([]*types.RetailCategory, error) {
	var categories []*types.RetailCategory
	cursor, err := r.retailCategory.Find(ctx, bson.M{"shop_id": shopID})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *mongoRepository) ListRetailItem(ctx context.Context, categoryID bson.ObjectID) ([]*types.RetailItem, error) {
	var items []*types.RetailItem
	cursor, err := r.retailItem.Find(ctx, bson.M{"category_id": categoryID})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *mongoRepository) ListMedicineCategory(ctx context.Context, shopID string) ([]*types.MedicineCategory, error) {
	var categories []*types.MedicineCategory
	cursor, err := r.medicineCategory.Find(ctx, bson.M{"shop_id": shopID})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *mongoRepository) ListMedicineItem(ctx context.Context, categoryID bson.ObjectID) ([]*types.MedicineItem, error) {
	var items []*types.MedicineItem
	cursor, err := r.medicineItem.Find(ctx, bson.M{"category_id": categoryID})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *mongoRepository) UpdateItemVariant(ctx context.Context, variant *types.UpdateItemVariant) error {
	filter, update, err := BuildUpdateDocument(variant)
	if err != nil {
		return err
	}

	if len(update) == 0 {
		return nil
	}

	_, err = r.itemVariant.UpdateOne(ctx, filter, update)

	return err
}

func (r *mongoRepository) UpdateItemAddon(ctx context.Context, addon *types.UpdateItemAddon) error {
	filter, update, err := BuildUpdateDocument(addon)
	if err != nil {
		return err
	}

	if len(update) == 0 {
		return nil
	}

	_, err = r.itemVariant.UpdateOne(ctx, filter, update)

	return err
}

func (r *mongoRepository) UpdateRestaurantMenu(ctx context.Context, menu *types.UpdateRestaurantMenu) error {
	filter, update, err := BuildUpdateDocument(menu)
	if err != nil {
		return err
	}

	if len(update) == 0 {
		return nil
	}

	_, err = r.restaurantMenu.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) UpdateMenuItem(ctx context.Context, item *types.UpdateMenuItem) error {
	filter, update, err := BuildUpdateDocument(item)
	if err != nil {
		return err
	}

	if len(update) == 0 {
		return nil
	}

	_, err = r.menuItem.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) UpdateRetailCategory(ctx context.Context, category *types.UpdateRetailCategory) error {
	filter, update, err := BuildUpdateDocument(category)
	if err != nil {
		return err
	}

	if len(update) == 0 {
		return nil
	}

	_, err = r.retailCategory.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) UpdateRetailItem(ctx context.Context, item *types.UpdateRetailItem) error {
	filter, update, err := BuildUpdateDocument(item)
	if err != nil {
		return err
	}

	if len(update) == 0 {
		return nil
	}

	_, err = r.retailItem.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) UpdateMedicineCategory(ctx context.Context, category *types.UpdateMedicineCategory) error {
	filter, update, err := BuildUpdateDocument(category)
	if err != nil {
		return err
	}

	if len(update) == 0 {
		return nil
	}

	_, err = r.medicineCategory.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) UpdateMedicineItem(ctx context.Context, item *types.UpdateMedicineItem) error {
	filter, update, err := BuildUpdateDocument(item)
	if err != nil {
		return err
	}

	if len(update) == 0 {
		return nil
	}

	_, err = r.medicineItem.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) DeleteItemVariant(ctx context.Context, variantID bson.ObjectID) error {
	_, err := r.itemVariant.DeleteOne(ctx, bson.M{"_id": variantID})
	return err
}

func (r *mongoRepository) DeleteItemAddon(ctx context.Context, addonID bson.ObjectID) error {
	_, err := r.itemAddon.DeleteOne(ctx, bson.M{"_id": addonID})
	return err
}

func (r *mongoRepository) DeleteRestaurantMenu(ctx context.Context, menuID bson.ObjectID) error {
	_, err := r.restaurantMenu.DeleteOne(ctx, bson.M{"_id": menuID})
	return err
}

func (r *mongoRepository) DeleteMenuItem(ctx context.Context, itemID bson.ObjectID) error {
	_, err := r.menuItem.DeleteOne(ctx, bson.M{"_id": itemID})
	return err
}

func (r *mongoRepository) DeleteRetailCategory(ctx context.Context, categoryID bson.ObjectID) error {
	_, err := r.retailCategory.DeleteOne(ctx, bson.M{"_id": categoryID})
	return err
}

func (r *mongoRepository) DeleteRetailItem(ctx context.Context, itemID bson.ObjectID) error {
	_, err := r.retailItem.DeleteOne(ctx, bson.M{"_id": itemID})
	return err
}

func (r *mongoRepository) DeleteMedicineCategory(ctx context.Context, categoryID bson.ObjectID) error {
	_, err := r.medicineCategory.DeleteOne(ctx, bson.M{"_id": categoryID})
	return err
}

func (r *mongoRepository) DeleteMedicineItem(ctx context.Context, itemID bson.ObjectID) error {
	_, err := r.medicineItem.DeleteOne(ctx, bson.M{"_id": itemID})
	return err
}

func BuildUpdateDocument(updateStruct any) (filter bson.M, update bson.M, err error) {
	v := reflect.ValueOf(updateStruct).Elem()
	t := v.Type()

	set := make(bson.M)
	unset := make(bson.M)
	filter = make(bson.M)
	update = make(bson.M)

	// Find ID field
	for i := range t.NumField() {
		field := t.Field(i)
		if bsonTag := field.Tag.Get("bson"); bsonTag == "_id" {
			filter["_id"] = v.Field(i).Interface()
			break
		}
	}

	if filter["_id"] == nil {
		return nil, nil, fmt.Errorf("missing _id field")
	}

	// Process fields
	for i := range t.NumField() {
		field := t.Field(i)
		bsonTag := field.Tag.Get("bson")
		fieldValue := v.Field(i)

		if bsonTag == "_id" || bsonTag == "" {
			continue
		}

		if fieldValue.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				unset[bsonTag] = ""
			} else {
				set[bsonTag] = fieldValue.Elem().Interface()
			}
		} else {
			set[bsonTag] = fieldValue.Interface()
		}
	}

	// Build final update command
	if len(set) > 0 {
		update["$set"] = set
	}
	if len(unset) > 0 {
		update["$unset"] = unset
	}

	return filter, update, nil
}
