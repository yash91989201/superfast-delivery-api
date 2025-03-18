package clients

import (
	"context"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ShopClient struct {
	conn    *grpc.ClientConn
	service pb.ShopServiceClient
}

func NewShopClient(serviceUrl string) (*ShopClient, error) {
	conn, err := grpc.NewClient(serviceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	s := pb.NewShopServiceClient(conn)

	return &ShopClient{conn, s}, nil
}

func (c *ShopClient) GetConn() *grpc.ClientConn {
	return c.conn
}

func (c *ShopClient) Close() {
	c.conn.Close()
}

func (s *ShopClient) CreateShop(ctx context.Context, req *pb.CreateShopReq) (*pb.Shop, error) {
	return s.service.CreateShop(ctx, req)
}

func (s *ShopClient) GetShop(ctx context.Context, req *pb.GetShopReq) (*pb.Shop, error) {
	return s.service.GetShop(ctx, req)
}

func (s *ShopClient) ListShops(ctx context.Context, req *pb.ListShopsReq) (*pb.ListShopsRes, error) {
	return s.service.ListShops(ctx, req)
}

func (s *ShopClient) UpdateShopAddress(ctx context.Context, req *pb.UpdateShopAddressReq) (*pb.UpdateShopAddressRes, error) {
	return s.service.UpdateShopAddress(ctx, req)
}

func (s *ShopClient) UpdateShopContact(ctx context.Context, req *pb.UpdateShopContactReq) (*pb.UpdateShopContactRes, error) {
	return s.service.UpdateShopContact(ctx, req)
}

func (s *ShopClient) UpdateShopImages(ctx context.Context, req *pb.UpdateShopImagesReq) (*pb.UpdateShopImagesRes, error) {
	return s.service.UpdateShopImages(ctx, req)
}

func (s *ShopClient) UpdateShopTimings(ctx context.Context, req *pb.UpdateShopTimingsReq) (*pb.UpdateShopTimingsRes, error) {
	return s.service.UpdateShopTimings(ctx, req)
}

func (s *ShopClient) UpdateShop(ctx context.Context, req *pb.UpdateShopReq) (*pb.UpdateShopRes, error) {
	return s.service.UpdateShop(ctx, req)
}

func (s *ShopClient) DeleteShop(ctx context.Context, req *pb.DeleteShopReq) (*pb.DeleteShopRes, error) {
	return s.service.DeleteShop(ctx, req)
}
