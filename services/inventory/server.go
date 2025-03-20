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
	res, err := s.service.CreateItemStock(ctx, &types.CreateItemStock{
		ItemID:   req.ItemId,
		Quantity: req.Quantity,
	})

	if err != nil {
		return nil, err
	}

	return types.ToPbItemStock(res), nil
}

func (s *grpcServer) CreateVariantStock(ctx context.Context, req *pb.CreateVariantStockReq) (*pb.VariantStock, error) {
	res, err := s.service.CreateVariantStock(ctx, &types.CreateVariantStock{
		VariantID: req.VariantId,
		Quantity:  req.Quantity,
	})

	if err != nil {
		return nil, err
	}

	return types.ToPbVariantStock(res), nil
}

func (s *grpcServer) CreateAddonStock(ctx context.Context, req *pb.CreateAddonStockReq) (*pb.AddonStock, error) {
	res, err := s.service.CreateAddonStock(ctx, &types.CreateAddonStock{
		AddonID:  req.AddonId,
		Quantity: req.Quantity,
	})

	if err != nil {
		return nil, err
	}

	return types.ToPbAddonStock(res), nil
}

func (s *grpcServer) GetItemStock(ctx context.Context, req *pb.GetItemStockReq) (*pb.ItemStock, error) {
	stock, err := s.service.GetItemStockByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return types.ToPbItemStock(stock), nil
}

func (s *grpcServer) GetVariantStock(ctx context.Context, req *pb.GetVariantStockReq) (*pb.VariantStock, error) {
	stock, err := s.service.GetVariantStockByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return types.ToPbVariantStock(stock), nil
}

func (s *grpcServer) GetAddonStock(ctx context.Context, req *pb.GetAddonStockReq) (*pb.AddonStock, error) {
	stock, err := s.service.GetAddonStockByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return types.ToPbAddonStock(stock), nil
}

func (s *grpcServer) UpdateItemStock(ctx context.Context, req *pb.UpdateItemStockReq) (*pb.ItemStock, error) {
	stock, err := s.service.UpdateItemStock(ctx, &types.ItemStock{
		ID:         req.Id,
		Quantity:   req.Quantity,
		RestockQty: req.RestockQty,
	})
	if err != nil {
		return nil, err
	}
	return types.ToPbItemStock(stock), nil
}

func (s *grpcServer) UpdateVariantStock(ctx context.Context, req *pb.UpdateVariantStockReq) (*pb.VariantStock, error) {
	stock, err := s.service.UpdateVariantStock(ctx, &types.VariantStock{
		ID:         req.Id,
		Quantity:   req.Quantity,
		RestockQty: req.RestockQty,
	})
	if err != nil {
		return nil, err
	}
	return types.ToPbVariantStock(stock), nil
}

func (s *grpcServer) UpdateAddonStock(ctx context.Context, req *pb.UpdateAddonStockReq) (*pb.AddonStock, error) {
	stock, err := s.service.UpdateAddonStock(ctx, &types.AddonStock{
		ID:         req.Id,
		Quantity:   req.Quantity,
		RestockQty: req.RestockQty,
	})
	if err != nil {
		return nil, err
	}
	return types.ToPbAddonStock(stock), nil
}

func (s *grpcServer) DeleteItemStock(ctx context.Context, req *pb.DeleteItemStockReq) (*pb.EmptyRes, error) {
	if err := s.service.DeleteItemStock(ctx, req.Id); err != nil {
		return nil, err
	}
	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) DeleteVariantStock(ctx context.Context, req *pb.DeleteVariantStockReq) (*pb.EmptyRes, error) {
	if err := s.service.DeleteVariantStock(ctx, req.Id); err != nil {
		return nil, err
	}
	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) DeleteAddonStock(ctx context.Context, req *pb.DeleteAddonStockReq) (*pb.EmptyRes, error) {
	if err := s.service.DeleteAddonStock(ctx, req.Id); err != nil {
		return nil, err
	}
	return &pb.EmptyRes{}, nil
}
