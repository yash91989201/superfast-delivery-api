package clients

import (
	"context"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type InventoryClient struct {
	conn    *grpc.ClientConn
	service pb.InventoryServiceClient
}

func NewInventoryClient(serviceUrl string) (*InventoryClient, error) {
	conn, err := grpc.NewClient(serviceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	s := pb.NewInventoryServiceClient(conn)

	return &InventoryClient{conn, s}, nil
}

func (c *InventoryClient) GetConn() *grpc.ClientConn {
	return c.conn
}

func (c *InventoryClient) Close() {
	c.conn.Close()
}

func (c *InventoryClient) CreateItemStock(ctx context.Context, req *pb.CreateItemStockReq) (*pb.ItemStock, error) {
	res, err := c.service.CreateItemStock(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *InventoryClient) CreateVariantStock(ctx context.Context, req *pb.CreateVariantStockReq) (*pb.VariantStock, error) {
	res, err := c.service.CreateVariantStock(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *InventoryClient) CreateAddonStock(ctx context.Context, req *pb.CreateAddonStockReq) (*pb.AddonStock, error) {
	res, err := c.service.CreateAddonStock(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
