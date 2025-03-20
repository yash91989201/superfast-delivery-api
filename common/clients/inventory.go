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

func (c *InventoryClient) GetItemStock(ctx context.Context, req *pb.GetItemStockReq) (*pb.ItemStock, error) {
	res, err := c.service.GetItemStock(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *InventoryClient) GetVariantStock(ctx context.Context, req *pb.GetVariantStockReq) (*pb.VariantStock, error) {
	res, err := c.service.GetVariantStock(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *InventoryClient) GetAddonStock(ctx context.Context, req *pb.GetAddonStockReq) (*pb.AddonStock, error) {
	res, err := c.service.GetAddonStock(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Update operations
func (c *InventoryClient) UpdateItemStock(ctx context.Context, req *pb.UpdateItemStockReq) (*pb.ItemStock, error) {
	res, err := c.service.UpdateItemStock(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *InventoryClient) UpdateVariantStock(ctx context.Context, req *pb.UpdateVariantStockReq) (*pb.VariantStock, error) {
	res, err := c.service.UpdateVariantStock(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *InventoryClient) UpdateAddonStock(ctx context.Context, req *pb.UpdateAddonStockReq) (*pb.AddonStock, error) {
	res, err := c.service.UpdateAddonStock(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Delete operations
func (c *InventoryClient) DeleteItemStock(ctx context.Context, req *pb.DeleteItemStockReq) (*pb.EmptyRes, error) {
	res, err := c.service.DeleteItemStock(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *InventoryClient) DeleteVariantStock(ctx context.Context, req *pb.DeleteVariantStockReq) (*pb.EmptyRes, error) {
	res, err := c.service.DeleteVariantStock(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *InventoryClient) DeleteAddonStock(ctx context.Context, req *pb.DeleteAddonStockReq) (*pb.EmptyRes, error) {
	res, err := c.service.DeleteAddonStock(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
