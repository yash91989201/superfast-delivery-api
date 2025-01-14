package inventory

import (
	"context"
	"net"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"github.com/yash91989201/superfast-delivery-api/common/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedInventoryServiceServer
	service Service
}

func Start(s Service, serviceUrl string) error {
	listener, err := net.Listen("tcp", serviceUrl)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterInventoryServiceServer(server, &grpcServer{service: s})

	reflection.Register(server)

	return server.Serve(listener)
}

func (s *grpcServer) CreateItemStock(ctx context.Context, req *pb.CreateItemStockReq) (*pb.ItemStock, error) {
	res, err := s.service.InsertItemStock(ctx, &types.CreateItemStock{
		ItemID:   req.ItemId,
		Quantity: req.Quantity,
	})

	if err != nil {
		return nil, err
	}

	return types.ToPbItemStock(res), nil
}

func (s *grpcServer) CreateVariantStock(ctx context.Context, req *pb.CreateVariantStockReq) (*pb.VariantStock, error) {
	res, err := s.service.InsertVariantStock(ctx, &types.CreateVariantStock{
		VariantID: req.VariantId,
		Quantity:  req.Quantity,
	})

	if err != nil {
		return nil, err
	}

	return types.ToPbVariantStock(res), nil
}

func (s *grpcServer) CreateAddonStock(ctx context.Context, req *pb.CreateAddonStockReq) (*pb.AddonStock, error) {
	res, err := s.service.InsertAddonStock(ctx, &types.CreateAddonStock{
		AddonID:  req.AddonId,
		Quantity: req.Quantity,
	})

	if err != nil {
		return nil, err
	}

	return types.ToPbAddonStock(res), nil
}
