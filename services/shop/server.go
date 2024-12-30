package shop

import (
	"context"
	"net"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"github.com/yash91989201/superfast-delivery-api/common/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service Service
	pb.UnimplementedShopServiceServer
}

func StartGRPCServer(s Service, serviceUrl string) error {

	listener, err := net.Listen("tcp", serviceUrl)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterShopServiceServer(server, &grpcServer{service: s})

	reflection.Register(server)

	return server.Serve(listener)
}

func (s *grpcServer) CreateShop(ctx context.Context, req *pb.CreateShopReq) (*pb.CreateShopRes, error) {
	newShop, err := s.service.InsertShop(ctx, types.ToCreateShop(req))
	if err != nil {
		return nil, err
	}

	return &pb.CreateShopRes{
		Id:      newShop.ID,
		Message: "Shop created successfully",
	}, nil
}

func (s *grpcServer) GetShop(ctx context.Context, req *pb.GetShopReq) (*pb.Shop, error) {
	res, err := s.service.GetShop(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return types.ToPbShop(res), nil
}

func (s *grpcServer) ListShops(context.Context, *pb.ListShopsReq) (*pb.ListShopsRes, error) {
	return nil, nil
}

func (s *grpcServer) UpdateShopAddress(context.Context, *pb.UpdateShopAddressReq) (*pb.UpdateShopAddressRes, error) {
	return nil, nil
}

func (s *grpcServer) UpdateShopContact(context.Context, *pb.UpdateShopContactReq) (*pb.UpdateShopContactRes, error) {
	return nil, nil
}

func (s *grpcServer) UpdateShopImages(context.Context, *pb.UpdateShopImagesReq) (*pb.UpdateShopImagesRes, error) {
	return nil, nil
}

func (s *grpcServer) UpdateShopTimings(context.Context, *pb.UpdateShopTimingsReq) (*pb.UpdateShopTimingsRes, error) {
	return nil, nil
}

func (s *grpcServer) UpdateShop(context.Context, *pb.UpdateShopReq) (*pb.UpdateShopRes, error) {
	return nil, nil
}

func (s *grpcServer) DeleteShop(context.Context, *pb.DeleteShopReq) (*pb.DeleteShopRes, error) {
	return nil, nil
}
