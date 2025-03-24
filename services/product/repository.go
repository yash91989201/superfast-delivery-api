package product

import (
	"context"
	"fmt"
	"time"

	"github.com/yash91989201/superfast-delivery-api/common/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Repository interface {
	Close(ctx context.Context) error
	Ping(ctx context.Context) error

	CreateRestaurantMenu(ctx context.Context, menu *types.RestaurantMenu) error
	CreateMenuItem(ctx context.Context, item *types.MenuItem) error
	CreateMenuItemVariant(ctx context.Context, variant *types.ItemVariant) error
	CreateMenuItemAddon(ctx context.Context, addon *types.ItemAddon) error
	CreateRetailCategory(ctx context.Context, category *types.RetailCategory) error
	CreateRetailItem(ctx context.Context, item *types.RetailItem) error
	CreateRetailItemVariant(ctx context.Context, variant *types.ItemVariant) error
	CreateMedicineCategory(ctx context.Context, category *types.MedicineCategory) error
	CreateMedicineItem(ctx context.Context, item *types.MedicineItem) error

	GetRestaurantMenu(ctx context.Context, menuID bson.ObjectID) (*types.RestaurantMenu, error)
	GetMenuItem(ctx context.Context, itemID bson.ObjectID) (*types.MenuItem, error)
	GetMenuItemVariant(ctx context.Context, itemID bson.ObjectID, variantID bson.ObjectID) (*types.ItemVariant, error)
	GetMenuItemAddon(ctx context.Context, itemID bson.ObjectID, addonID bson.ObjectID) (*types.ItemAddon, error)
	GetRetailCategory(ctx context.Context, categoryID bson.ObjectID) (*types.RetailCategory, error)
	GetRetailItem(ctx context.Context, itemID bson.ObjectID) (*types.RetailItem, error)
	GetRetailItemVariant(ctx context.Context, itemID bson.ObjectID, variantID bson.ObjectID) (*types.ItemVariant, error)
	GetMedicineCategory(ctx context.Context, categoryID bson.ObjectID) (*types.MedicineCategory, error)
	GetMedicineItem(ctx context.Context, itemID bson.ObjectID) (*types.MedicineItem, error)

	ListRestaurantMenu(ctx context.Context, shopID string) ([]*types.RestaurantMenu, error)
	ListMenuItem(ctx context.Context, menuID bson.ObjectID) ([]*types.MenuItem, error)
	ListMenuItemVariant(ctx context.Context, itemID bson.ObjectID) ([]*types.ItemVariant, error)
	ListMenuItemAddon(ctx context.Context, itemID bson.ObjectID) ([]*types.ItemAddon, error)
	ListRetailCategory(ctx context.Context, shopID string) ([]*types.RetailCategory, error)
	ListRetailItem(ctx context.Context, categoryID bson.ObjectID) ([]*types.RetailItem, error)
	ListRetailItemVariant(ctx context.Context, itemID bson.ObjectID) ([]*types.ItemVariant, error)
	ListMedicineCategory(ctx context.Context, shopID string) ([]*types.MedicineCategory, error)
	ListMedicineItem(ctx context.Context, categoryID bson.ObjectID) ([]*types.MedicineItem, error)

	UpdateRestaurantMenu(ctx context.Context, menu *types.UpdateRestaurantMenu) error
	UpdateMenuItem(ctx context.Context, item *types.UpdateMenuItem) error
	UpdateMenuItemVariant(ctx context.Context, variant *types.UpdateItemVariant) error
	UpdateMenuItemAddon(ctx context.Context, addon *types.UpdateItemAddon) error
	UpdateRetailCategory(ctx context.Context, category *types.UpdateRetailCategory) error
	UpdateRetailItem(ctx context.Context, item *types.UpdateRetailItem) error
	UpdateRetailItemVariant(ctx context.Context, variant *types.UpdateItemVariant) error
	UpdateMedicineCategory(ctx context.Context, category *types.UpdateMedicineCategory) error
	UpdateMedicineItem(ctx context.Context, item *types.UpdateMedicineItem) error

	DeleteRestaurantMenu(ctx context.Context, menuID bson.ObjectID) error
	DeleteMenuItem(ctx context.Context, itemID bson.ObjectID) error
	DeleteMenuItemVariant(ctx context.Context, itemID bson.ObjectID, variantID bson.ObjectID) error
	DeleteMenuItemAddon(ctx context.Context, itemID bson.ObjectID, addonID bson.ObjectID) error
	DeleteRetailCategory(ctx context.Context, categoryID bson.ObjectID) error
	DeleteRetailItem(ctx context.Context, itemID bson.ObjectID) error
	DeleteRetailItemVariant(ctx context.Context, itemID bson.ObjectID, variantID bson.ObjectID) error
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

func (r *mongoRepository) CreateMenuItemVariant(ctx context.Context, variant *types.ItemVariant) error {
	filter := bson.M{"_id": variant.ItemID}
	update := bson.M{
		"$push": bson.M{"variants": variant},
		"$set":  bson.M{"updated_at": time.Now()},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updated types.MenuItem
	err := r.menuItem.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updated)
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoRepository) CreateMenuItemAddon(ctx context.Context, addon *types.ItemAddon) error {
	filter := bson.M{"_id": addon.ItemID}
	update := bson.M{
		"$push": bson.M{"addons": addon},
		"$set":  bson.M{"updated_at": time.Now()},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updated types.MenuItem
	err := r.menuItem.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updated)
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

func (r *mongoRepository) CreateRetailItemVariant(ctx context.Context, variant *types.ItemVariant) error {
	filter := bson.M{"_id": variant.ItemID}
	update := bson.M{
		"$push": bson.M{
			"variants": variant,
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	_, err := r.retailItem.UpdateOne(ctx, filter, update)
	return err
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

func (r *mongoRepository) GetMenuItemVariant(ctx context.Context, itemID bson.ObjectID, variantID bson.ObjectID) (*types.ItemVariant, error) {
	var item types.MenuItem
	err := r.menuItem.FindOne(
		ctx,
		bson.M{"_id": itemID, "variants._id": variantID},
		options.FindOne().SetProjection(bson.M{"variants.$": 1}),
	).Decode(&item)
	if err != nil {
		return nil, err
	}

	if len(item.Variants) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return item.Variants[0], nil
}

func (r *mongoRepository) GetMenuItemAddon(ctx context.Context, itemID bson.ObjectID, addonID bson.ObjectID) (*types.ItemAddon, error) {
	var item types.MenuItem
	err := r.menuItem.FindOne(
		ctx,
		bson.M{"_id": itemID, "addons._id": addonID},
		options.FindOne().SetProjection(bson.M{"addons.$": 1}),
	).Decode(&item)
	if err != nil {
		return nil, err
	}

	if len(item.Addons) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return item.Addons[0], nil
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

func (r *mongoRepository) GetRetailItemVariant(ctx context.Context, itemID bson.ObjectID, variantID bson.ObjectID) (*types.ItemVariant, error) {
	var item types.RetailItem
	err := r.retailItem.FindOne(
		ctx,
		bson.M{"_id": itemID, "variants._id": variantID},
		options.FindOne().SetProjection(bson.M{"variants.$": 1}),
	).Decode(&item)
	if err != nil {
		return nil, err
	}

	if len(item.Variants) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return item.Variants[0], nil
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
	var item types.MenuItem
	err := r.menuItem.FindOne(ctx, bson.M{"_id": itemID}).Decode(&item)
	if err != nil {
		return nil, err
	}

	return item.Variants, nil
}

func (r *mongoRepository) ListItemAddon(ctx context.Context, itemID bson.ObjectID) ([]*types.ItemAddon, error) {
	var item types.MenuItem
	err := r.menuItem.FindOne(ctx, bson.M{"_id": itemID}).Decode(&item)
	if err != nil {
		return nil, err
	}

	return item.Addons, nil
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

func (r *mongoRepository) ListMenuItemVariant(ctx context.Context, itemID bson.ObjectID) ([]*types.ItemVariant, error) {
	var item types.MenuItem
	err := r.menuItem.FindOne(ctx, bson.M{"_id": itemID}).Decode(&item)
	if err != nil {
		return nil, err
	}
	return item.Variants, nil
}

func (r *mongoRepository) ListMenuItemAddon(ctx context.Context, itemID bson.ObjectID) ([]*types.ItemAddon, error) {
	var item types.MenuItem
	err := r.menuItem.FindOne(ctx, bson.M{"_id": itemID}).Decode(&item)
	if err != nil {
		return nil, err
	}
	return item.Addons, nil
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

func (r *mongoRepository) ListRetailItemVariant(ctx context.Context, itemID bson.ObjectID) ([]*types.ItemVariant, error) {
	var item types.RetailItem
	err := r.retailItem.FindOne(ctx, bson.M{"_id": itemID}).Decode(&item)
	if err != nil {
		return nil, err
	}
	return item.Variants, nil
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

func (r *mongoRepository) UpdateRestaurantMenu(ctx context.Context, menu *types.UpdateRestaurantMenu) error {
	filter := bson.M{"_id": menu.ID}
	setFields := bson.M{}

	if menu.MenuName != nil {
		setFields["menu_name"] = *menu.MenuName
	}
	if menu.ImageURL != nil {
		setFields["image_url"] = *menu.ImageURL
	}

	if len(setFields) == 0 {
		return fmt.Errorf("no fields to update")
	}

	update := bson.M{
		"$set":         setFields,
		"$currentDate": bson.M{"updated_at": true},
	}

	_, err := r.restaurantMenu.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) UpdateMenuItem(ctx context.Context, item *types.UpdateMenuItem) error {
	filter := bson.M{"_id": item.ID}
	setFields := bson.M{}

	if item.Name != nil {
		setFields["name"] = *item.Name
	}
	if item.Price != nil {
		setFields["price"] = *item.Price
	}
	if item.ImageURL != nil {
		setFields["image_url"] = *item.ImageURL
	}
	if item.Description != nil {
		setFields["description"] = *item.Description
	}

	if len(setFields) == 0 {
		return fmt.Errorf("no fields to update")
	}

	update := bson.M{
		"$set":         setFields,
		"$currentDate": bson.M{"updated_at": true},
	}
	_, err := r.menuItem.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) UpdateMenuItemVariant(ctx context.Context, variant *types.UpdateItemVariant) error {
	filter := bson.M{
		"_id":          variant.ItemID,
		"variants._id": variant.ID,
	}

	setFields := bson.M{}
	if variant.VariantName != nil {
		setFields["variants.$.variant_name"] = *variant.VariantName
	}
	if variant.RelativePricing != nil {
		setFields["variants.$.relative_pricing"] = *variant.RelativePricing
	}
	if variant.RelativePrice != nil {
		setFields["variants.$.relative_price"] = *variant.RelativePrice
	}
	if variant.Price != nil {
		setFields["variants.$.price"] = *variant.Price
	}
	if variant.ImageURL != nil {
		setFields["variants.$.image_url"] = *variant.ImageURL
	}
	if variant.Description != nil {
		setFields["variants.$.description"] = *variant.Description
	}

	if len(setFields) == 0 {
		return fmt.Errorf("No fields to update")
	}

	update := bson.M{
		"$set":         setFields,
		"$currentDate": bson.M{"updated_at": true},
	}

	_, err := r.menuItem.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) UpdateMenuItemAddon(ctx context.Context, addon *types.UpdateItemAddon) error {
	filter := bson.M{
		"_id":        addon.ItemID,
		"addons._id": addon.ID,
	}

	setFields := bson.M{}
	if addon.AddonName != nil {
		setFields["addons.$.addon_name"] = *addon.AddonName
	}
	if addon.AddonPrice != nil {
		setFields["addons.$.addon_price"] = *addon.AddonPrice
	}
	if addon.ImageURL != nil {
		setFields["addons.$.image_url"] = *addon.ImageURL
	}
	if addon.Description != nil {
		setFields["addons.$.description"] = *addon.Description
	}

	if len(setFields) == 0 {
		return fmt.Errorf("no fields to update")
	}

	update := bson.M{
		"$set":         setFields,
		"$currentDate": bson.M{"updated_at": true},
	}

	_, err := r.menuItem.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) UpdateRetailCategory(ctx context.Context, category *types.UpdateRetailCategory) error {
	filter := bson.M{"_id": category.ID}
	setFields := bson.M{}

	if category.CategoryName != nil {
		setFields["category_name"] = *category.CategoryName
	}
	if category.ImageURL != nil {
		setFields["image_url"] = *category.ImageURL
	}

	if len(setFields) == 0 {
		return fmt.Errorf("No fields to update")
	}

	update := bson.M{
		"$set":         setFields,
		"$currentDate": bson.M{"updated_at": true},
	}
	_, err := r.retailCategory.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) UpdateRetailItem(ctx context.Context, item *types.UpdateRetailItem) error {
	filter := bson.M{"_id": item.ID}
	setFields := bson.M{}

	if item.Name != nil {
		setFields["name"] = *item.Name
	}
	if item.Price != nil {
		setFields["price"] = *item.Price
	}
	if item.ImageURL != nil {
		setFields["image_url"] = *item.ImageURL
	}
	if item.Description != nil {
		setFields["description"] = *item.Description
	}

	if len(setFields) == 0 {
		return fmt.Errorf("no fields to update")
	}

	update := bson.M{
		"$set":         setFields,
		"$currentDate": bson.M{"updated_at": true},
	}

	_, err := r.retailItem.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) UpdateRetailItemVariant(ctx context.Context, variant *types.UpdateItemVariant) error {
	filter := bson.M{
		"_id":          variant.ItemID,
		"variants._id": variant.ID,
	}

	setFields := bson.M{}

	if variant.VariantName != nil {
		setFields["variants.$.variant_name"] = *variant.VariantName
	}
	if variant.RelativePricing != nil {
		setFields["variants.$.relative_pricing"] = *variant.RelativePricing
	}
	if variant.RelativePrice != nil {
		setFields["variants.$.relative_price"] = *variant.RelativePrice
	}
	if variant.Price != nil {
		setFields["variants.$.price"] = *variant.Price
	}
	if variant.ImageURL != nil {
		setFields["variants.$.image_url"] = *variant.ImageURL
	}
	if variant.Description != nil {
		setFields["variants.$.description"] = *variant.Description
	}

	if len(setFields) == 0 {
		return fmt.Errorf("No fields to update")
	}

	update := bson.M{
		"$set":         setFields,
		"$currentDate": bson.M{"updated_at": true},
	}

	_, err := r.retailItem.UpdateOne(ctx, filter, update)

	return err
}

func (r *mongoRepository) UpdateMedicineCategory(ctx context.Context, category *types.UpdateMedicineCategory) error {
	filter := bson.M{"_id": category.ID}
	setFields := bson.M{}

	if category.CategoryName != nil {
		setFields["category_name"] = *category.CategoryName
	}
	if category.ImageURL != nil {
		setFields["image_url"] = *category.ImageURL
	}

	if len(setFields) == 0 {
		return fmt.Errorf("No fields to update")
	}

	update := bson.M{
		"$set":         setFields,
		"$currentDate": bson.M{"updated_at": true},
	}

	_, err := r.medicineCategory.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) UpdateMedicineItem(ctx context.Context, item *types.UpdateMedicineItem) error {
	filter := bson.M{"_id": item.ID}
	setFields := bson.M{}

	if item.Name != nil {
		setFields["name"] = *item.Name
	}
	if item.Price != nil {
		setFields["price"] = *item.Price
	}
	if item.ImageURL != nil {
		setFields["image_url"] = *item.ImageURL
	}
	if item.Description != nil {
		setFields["description"] = *item.Description
	}

	if len(setFields) == 0 {
		return fmt.Errorf("No fields to update")
	}

	update := bson.M{
		"$set":         setFields,
		"$currentDate": bson.M{"updated_at": true},
	}

	_, err := r.medicineItem.UpdateOne(ctx, filter, update)
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

func (r *mongoRepository) DeleteMenuItemVariant(ctx context.Context, itemID bson.ObjectID, variantID bson.ObjectID) error {
	filter := bson.M{"_id": itemID}
	update := bson.M{
		"$pull": bson.M{
			"variants": bson.M{"_id": variantID},
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	_, err := r.menuItem.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) DeleteMenuItemAddon(ctx context.Context, itemID bson.ObjectID, addonID bson.ObjectID) error {
	filter := bson.M{"_id": itemID}
	update := bson.M{
		"$pull": bson.M{
			"addons": bson.M{"_id": addonID},
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	_, err := r.menuItem.UpdateOne(ctx, filter, update)
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

func (r *mongoRepository) DeleteRetailItemVariant(ctx context.Context, itemID bson.ObjectID, variantID bson.ObjectID) error {
	filter := bson.M{"_id": itemID}
	update := bson.M{
		"$pull": bson.M{
			"variants": bson.M{"_id": variantID},
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	_, err := r.retailItem.UpdateOne(ctx, filter, update)
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
