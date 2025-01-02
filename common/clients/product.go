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
	res, err := c.service.CreateItemVariant(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ProductClient) CreateItemAddon(ctx context.Context, req *pb.CreateItemAddonReq) (*pb.ItemAddon, error) {
	res, err := c.service.CreateItemAddon(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ProductClient) CreateRestaurantMenu(ctx context.Context, req *pb.CreateRestaurantMenuReq) (*pb.RestaurantMenu, error) {
	res, err := c.service.CreateRestaurantMenu(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ProductClient) CreateMenuItem(ctx context.Context, req *pb.CreateMenuItemReq) (*pb.MenuItem, error) {
	res, err := c.service.CreateMenuItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ProductClient) CreateRetailCategory(ctx context.Context, req *pb.CreateRetailCategoryReq) (*pb.RetailCategory, error) {
	res, err := c.service.CreateRetailCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ProductClient) CreateMedicineCategory(ctx context.Context, req *pb.CreateMedicineCategoryReq) (*pb.MedicineCategory, error) {
	res, err := c.service.CreateMedicineCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
