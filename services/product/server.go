package product

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
	pb.UnimplementedProductServiceServer
}

func StartGRPCServer(s Service, serviceUrl string) error {
	listener, err := net.Listen("tcp", serviceUrl)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterProductServiceServer(server, &grpcServer{service: s})

	reflection.Register(server)

	return server.Serve(listener)
}

func (s *grpcServer) CreateRestaurantMenu(ctx context.Context, req *pb.CreateRestaurantMenuReq) (*pb.RestaurantMenu, error) {
	res, err := s.service.InsertRestaurantMenu(ctx, types.ToCreateRestaurantMenu(req))
	if err != nil {
		return nil, err
	}

	return types.ToPbRestaurantMenu(res), nil
}

func (s *grpcServer) CreateMenuItem(ctx context.Context, req *pb.CreateMenuItemReq) (*pb.MenuItem, error) {
	res, err := s.service.InsertMenuItem(ctx, types.ToCreateMenuItem(req))
	if err != nil {
		return nil, err
	}

	return types.ToPbMenuItem(res), nil
}

func (s *grpcServer) GetRestaurantMenu(ctx context.Context, req *pb.GetRestaurantMenuReq) (*pb.RestaurantMenu, error) {
	res, err := s.service.GetRestaurantMenu(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return types.ToPbRestaurantMenu(res), err
}

func (s *grpcServer) ListRestaurantMenu(ctx context.Context, req *pb.ListRestaurantMenuReq) (*pb.ListRestaurantMenuRes, error) {
	res, err := s.service.ListRestaurantMenu(ctx, req.ShopId)
	if err != nil {
		return nil, err
	}

	return &pb.ListRestaurantMenuRes{
		RestaurantMenuList: types.ToPbRestaurantMenuList(res),
	}, nil
}
