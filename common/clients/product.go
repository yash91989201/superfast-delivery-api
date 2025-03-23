package clients

import (
	"context"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductClient struct {
	conn    *grpc.ClientConn
	service pb.ProductServiceClient
}

func NewProductClient(serviceUrl string) (*ProductClient, error) {
	conn, err := grpc.NewClient(serviceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	s := pb.NewProductServiceClient(conn)

	return &ProductClient{conn, s}, nil
}

func (c *ProductClient) GetConn() *grpc.ClientConn {
	return c.conn
}

func (c *ProductClient) Close() {
	c.conn.Close()
}

func (c *ProductClient) CreateRestaurantMenu(ctx context.Context, req *pb.CreateRestaurantMenuReq) (*pb.RestaurantMenu, error) {
	return c.service.CreateRestaurantMenu(ctx, req)
}

func (c *ProductClient) CreateMenuItem(ctx context.Context, req *pb.CreateMenuItemReq) (*pb.MenuItem, error) {
	return c.service.CreateMenuItem(ctx, req)
}

func (c *ProductClient) CreateMenuItemVariant(ctx context.Context, req *pb.CreateItemVariantReq) (*pb.ItemVariant, error) {
	return c.service.CreateMenuItemVariant(ctx, req)
}

func (c *ProductClient) CreateMenuItemAddon(ctx context.Context, req *pb.CreateItemAddonReq) (*pb.ItemAddon, error) {
	return c.service.CreateMenuItemAddon(ctx, req)
}

func (c *ProductClient) CreateRetailCategory(ctx context.Context, req *pb.CreateRetailCategoryReq) (*pb.RetailCategory, error) {
	return c.service.CreateRetailCategory(ctx, req)
}

func (c *ProductClient) CreateRetailItem(ctx context.Context, req *pb.CreateRetailItemReq) (*pb.RetailItem, error) {
	return c.service.CreateRetailItem(ctx, req)
}

func (c *ProductClient) CreateRetailItemVariant(ctx context.Context, req *pb.CreateItemVariantReq) (*pb.ItemVariant, error) {
	return c.service.CreateRetailItemVariant(ctx, req)
}

func (c *ProductClient) CreateMedicineCategory(ctx context.Context, req *pb.CreateMedicineCategoryReq) (*pb.MedicineCategory, error) {
	return c.service.CreateMedicineCategory(ctx, req)
}

func (c *ProductClient) CreateMedicineItem(ctx context.Context, req *pb.CreateMedicineItemReq) (*pb.MedicineItem, error) {
	return c.service.CreateMedicineItem(ctx, req)
}

func (c *ProductClient) GetRestaurantMenu(ctx context.Context, req *pb.GetRestaurantMenuReq) (*pb.RestaurantMenu, error) {
	return c.service.GetRestaurantMenu(ctx, req)
}

func (c *ProductClient) GetMenuItem(ctx context.Context, req *pb.GetMenuItemReq) (*pb.MenuItem, error) {
	return c.service.GetMenuItem(ctx, req)
}

func (c *ProductClient) GetMenuItemVariant(ctx context.Context, req *pb.GetItemVariantReq) (*pb.ItemVariant, error) {
	return c.service.GetMenuItemVariant(ctx, req)
}

func (c *ProductClient) GetMenuItemAddon(ctx context.Context, req *pb.GetItemAddonReq) (*pb.ItemAddon, error) {
	return c.service.GetMenuItemAddon(ctx, req)
}

func (c *ProductClient) GetRetailCategory(ctx context.Context, req *pb.GetRetailCategoryReq) (*pb.RetailCategory, error) {
	return c.service.GetRetailCategory(ctx, req)
}

func (c *ProductClient) GetRetailItem(ctx context.Context, req *pb.GetRetailItemReq) (*pb.RetailItem, error) {
	return c.service.GetRetailItem(ctx, req)
}

func (c *ProductClient) GetRetailItemVariant(ctx context.Context, req *pb.GetItemVariantReq) (*pb.ItemVariant, error) {
	return c.service.GetRetailItemVariant(ctx, req)
}

func (c *ProductClient) GetMedicineCategory(ctx context.Context, req *pb.GetMedicineCategoryReq) (*pb.MedicineCategory, error) {
	return c.service.GetMedicineCategory(ctx, req)
}

func (c *ProductClient) GetMedicineItem(ctx context.Context, req *pb.GetMedicineItemReq) (*pb.MedicineItem, error) {
	return c.service.GetMedicineItem(ctx, req)
}

func (c *ProductClient) ListRestaurantMenu(ctx context.Context, req *pb.ListRestaurantMenuReq) (*pb.ListRestaurantMenuRes, error) {
	return c.service.ListRestaurantMenu(ctx, req)
}

func (c *ProductClient) ListMenuItem(ctx context.Context, req *pb.ListMenuItemReq) (*pb.ListMenuItemRes, error) {
	return c.service.ListMenuItem(ctx, req)
}

func (c *ProductClient) ListMenuItemVariant(ctx context.Context, req *pb.ListItemVariantReq) (*pb.ListItemVariantRes, error) {
	return c.service.ListMenuItemVariant(ctx, req)
}

func (c *ProductClient) ListMenuItemAddon(ctx context.Context, req *pb.ListItemAddonReq) (*pb.ListItemAddonRes, error) {
	return c.service.ListMenuItemAddon(ctx, req)
}

func (c *ProductClient) ListRetailCategory(ctx context.Context, req *pb.ListRetailCategoryReq) (*pb.ListRetailCategoryRes, error) {
	return c.service.ListRetailCategory(ctx, req)
}

func (c *ProductClient) ListRetailItem(ctx context.Context, req *pb.ListRetailItemReq) (*pb.ListRetailItemRes, error) {
	return c.service.ListRetailItem(ctx, req)
}

func (c *ProductClient) ListRetailItemVariant(ctx context.Context, req *pb.ListItemVariantReq) (*pb.ListItemVariantRes, error) {
	return c.service.ListRetailItemVariant(ctx, req)
}

func (c *ProductClient) ListMedicineCategory(ctx context.Context, req *pb.ListMedicineCategoryReq) (*pb.ListMedicineCategoryRes, error) {
	return c.service.ListMedicineCategory(ctx, req)
}

func (c *ProductClient) ListMedicineItem(ctx context.Context, req *pb.ListMedicineItemReq) (*pb.ListMedicineItemRes, error) {
	return c.service.ListMedicineItem(ctx, req)
}

func (c *ProductClient) UpdateRestaurantMenu(ctx context.Context, req *pb.UpdateRestaurantMenuReq) (*pb.EmptyRes, error) {
	return c.service.UpdateRestaurantMenu(ctx, req)
}

func (c *ProductClient) UpdateMenuItem(ctx context.Context, req *pb.UpdateMenuItemReq) (*pb.EmptyRes, error) {
	return c.service.UpdateMenuItem(ctx, req)
}

func (c *ProductClient) UpdateMenuItemVariant(ctx context.Context, req *pb.UpdateItemVariantReq) (*pb.EmptyRes, error) {
	return c.service.UpdateMenuItemVariant(ctx, req)
}

func (c *ProductClient) UpdateMenuItemAddon(ctx context.Context, req *pb.UpdateItemAddonReq) (*pb.EmptyRes, error) {
	return c.service.UpdateMenuItemAddon(ctx, req)
}

func (c *ProductClient) UpdateRetailCategory(ctx context.Context, req *pb.UpdateRetailCategoryReq) (*pb.EmptyRes, error) {
	return c.service.UpdateRetailCategory(ctx, req)
}

func (c *ProductClient) UpdateRetailItem(ctx context.Context, req *pb.UpdateRetailItemReq) (*pb.EmptyRes, error) {
	return c.service.UpdateRetailItem(ctx, req)
}

func (c *ProductClient) UpdateRetailItemVariant(ctx context.Context, req *pb.UpdateItemVariantReq) (*pb.EmptyRes, error) {
	return c.service.UpdateRetailItemVariant(ctx, req)
}

func (c *ProductClient) UpdateMedicineCategory(ctx context.Context, req *pb.UpdateMedicineCategoryReq) (*pb.EmptyRes, error) {
	return c.service.UpdateMedicineCategory(ctx, req)
}

func (c *ProductClient) UpdateMedicineItem(ctx context.Context, req *pb.UpdateMedicineItemReq) (*pb.EmptyRes, error) {
	return c.service.UpdateMedicineItem(ctx, req)
}

func (c *ProductClient) DeleteRestaurantMenu(ctx context.Context, req *pb.DeleteRestaurantMenuReq) (*pb.EmptyRes, error) {
	return c.service.DeleteRestaurantMenu(ctx, req)
}

func (c *ProductClient) DeleteMenuItem(ctx context.Context, req *pb.DeleteMenuItemReq) (*pb.EmptyRes, error) {
	return c.service.DeleteMenuItem(ctx, req)
}

func (c *ProductClient) DeleteMenuItemVariant(ctx context.Context, req *pb.DeleteItemVariantReq) (*pb.EmptyRes, error) {
	return c.service.DeleteMenuItemVariant(ctx, req)
}

func (c *ProductClient) DeleteMenuItemAddon(ctx context.Context, req *pb.DeleteItemAddonReq) (*pb.EmptyRes, error) {
	return c.service.DeleteMenuItemAddon(ctx, req)
}

func (c *ProductClient) DeleteRetailCategory(ctx context.Context, req *pb.DeleteRetailCategoryReq) (*pb.EmptyRes, error) {
	return c.service.DeleteRetailCategory(ctx, req)
}

func (c *ProductClient) DeleteRetailItem(ctx context.Context, req *pb.DeleteRetailItemReq) (*pb.EmptyRes, error) {
	return c.service.DeleteRetailItem(ctx, req)
}

func (c *ProductClient) DeleteRetailItemVariant(ctx context.Context, req *pb.DeleteItemVariantReq) (*pb.EmptyRes, error) {
	return c.service.DeleteRetailItemVariant(ctx, req)
}

func (c *ProductClient) DeleteMedicineCategory(ctx context.Context, req *pb.DeleteMedicineCategoryReq) (*pb.EmptyRes, error) {
	return c.service.DeleteMedicineCategory(ctx, req)
}

func (c *ProductClient) DeleteMedicineItem(ctx context.Context, req *pb.DeleteMedicineItemReq) (*pb.EmptyRes, error) {
	return c.service.DeleteMedicineItem(ctx, req)
}
