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

func (s *ShopClient) CreateShop(ctx context.Context, req *pb.CreateShopReq) (*pb.CreateShopRes, error) {
	res, err := s.service.CreateShop(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ShopClient) GetShop(ctx context.Context, req *pb.GetShopReq) (*pb.GetShopRes, error) {
	res, err := s.service.GetShop(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ShopClient) ListShops(ctx context.Context, req *pb.ListShopsReq) (*pb.ListShopsRes, error) {
	res, err := s.service.ListShops(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ShopClient) UpdateShopAddress(ctx context.Context, req *pb.UpdateShopAddressReq) (*pb.UpdateShopAddressRes, error) {
	res, err := s.service.UpdateShopAddress(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ShopClient) UpdateShopContact(ctx context.Context, req *pb.UpdateShopContactReq) (*pb.UpdateShopContactRes, error) {
	res, err := s.service.UpdateShopContact(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ShopClient) UpdateShopImages(ctx context.Context, req *pb.UpdateShopImagesReq) (*pb.UpdateShopImagesRes, error) {
	res, err := s.service.UpdateShopImages(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ShopClient) UpdateShopTimings(ctx context.Context, req *pb.UpdateShopTimingsReq) (*pb.UpdateShopTimingsRes, error) {
	res, err := s.service.UpdateShopTimings(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ShopClient) UpdateShop(ctx context.Context, req *pb.UpdateShopReq) (*pb.UpdateShopRes, error) {
	res, err := s.service.UpdateShop(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ShopClient) DeleteShop(ctx context.Context, req *pb.DeleteShopReq) (*pb.DeleteShopRes, error) {
	res, err := s.service.DeleteShop(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
