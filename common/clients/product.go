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

func (c *ProductClient) CreateItemVariant(ctx context.Context, req *pb.CreateItemVariantReq) (*pb.ItemVariant, error) {
	return c.service.CreateItemVariant(ctx, req)
}

func (c *ProductClient) CreateItemAddon(ctx context.Context, req *pb.CreateItemAddonReq) (*pb.ItemAddon, error) {
	return c.service.CreateItemAddon(ctx, req)
}

func (c *ProductClient) CreateRestaurantMenu(ctx context.Context, req *pb.CreateRestaurantMenuReq) (*pb.RestaurantMenu, error) {
	return c.service.CreateRestaurantMenu(ctx, req)
}

func (c *ProductClient) CreateMenuItem(ctx context.Context, req *pb.CreateMenuItemReq) (*pb.MenuItem, error) {
	return c.service.CreateMenuItem(ctx, req)
}

func (c *ProductClient) CreateRetailCategory(ctx context.Context, req *pb.CreateRetailCategoryReq) (*pb.RetailCategory, error) {
	return c.service.CreateRetailCategory(ctx, req)
}

func (c *ProductClient) CreateMedicineCategory(ctx context.Context, req *pb.CreateMedicineCategoryReq) (*pb.MedicineCategory, error) {
	return c.service.CreateMedicineCategory(ctx, req)
}

func (c *ProductClient) GetRestaurantMenu(ctx context.Context, req *pb.GetRestaurantMenuReq) (*pb.RestaurantMenu, error) {
	return c.service.GetRestaurantMenu(ctx, req)
}

func (c *ProductClient) ListRestaurantMenu(ctx context.Context, req *pb.ListRestaurantMenuReq) (*pb.ListRestaurantMenuRes, error) {
	return c.service.ListRestaurantMenu(ctx, req)
}
