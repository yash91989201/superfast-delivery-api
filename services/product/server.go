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

func (s *grpcServer) GetItemVariant(ctx context.Context, req *pb.GetItemVariantReq) (*pb.ItemVariant, error) {
	res, err := s.service.GetItemVariant(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return types.ToPbItemVariant(res), nil
}
func (s *grpcServer) GetItemAddon(ctx context.Context, req *pb.GetItemAddonReq) (*pb.ItemAddon, error) {
	res, err := s.service.GetItemAddon(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return types.ToPbItemAddon(res), nil
}
func (s *grpcServer) GetItemVariants(ctx context.Context, req *pb.GetItemVariantsReq) (*pb.GetItemVariantsRes, error) {
	res, err := s.service.GetItemVariants(ctx, req.ItemId)
	if err != nil {
		return nil, err
	}

	return &pb.GetItemVariantsRes{
		Variants: types.ToPbItemVariants(res),
	}, nil
}
func (s *grpcServer) GetItemAddons(ctx context.Context, req *pb.GetItemAddonsReq) (*pb.GetItemAddonsRes, error) {
	res, err := s.service.GetItemAddons(ctx, req.ItemId)
	if err != nil {
		return nil, err
	}

	return &pb.GetItemAddonsRes{
		Addons: types.ToPbItemAddons(res),
	}, nil
}

func (s *grpcServer) CreateItemVariant(ctx context.Context, req *pb.CreateItemVariantReq) (*pb.ItemVariant, error) {
	res, err := s.service.CreateItemVariant(ctx, types.ToCreateItemVariant(req))
	if err != nil {
		return nil, err
	}

	return types.ToPbItemVariant(res), err
}

func (s *grpcServer) CreateItemAddon(ctx context.Context, req *pb.CreateItemAddonReq) (*pb.ItemAddon, error) {
	res, err := s.service.CreateItemAddon(ctx, types.ToCreateItemAddon(req))
	if err != nil {
		return nil, err
	}

	return types.ToPbItemAddon(res), err
}

func (s *grpcServer) CreateRestaurantMenu(ctx context.Context, req *pb.CreateRestaurantMenuReq) (*pb.RestaurantMenu, error) {
	res, err := s.service.CreateRestaurantMenu(ctx, types.ToCreateRestaurantMenu(req))
	if err != nil {
		return nil, err
	}

	return types.ToPbRestaurantMenu(res), nil
}

func (s *grpcServer) CreateMenuItem(ctx context.Context, req *pb.CreateMenuItemReq) (*pb.MenuItem, error) {
	res, err := s.service.CreateMenuItem(ctx, types.ToCreateMenuItem(req))
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

func (s *grpcServer) GetRetailCategory(ctx context.Context, req *pb.GetRetailCategoryReq) (*pb.RetailCategory, error) {
	res, err := s.service.GetRetailCategory(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return types.ToPbRetailCategory(res), err
}
func (s *grpcServer) ListRetailCategory(ctx context.Context, req *pb.ListRetailCategoryReq) (*pb.ListRetailCategoryRes, error) {
	res, err := s.service.ListRetailCategory(ctx, req.ShopId)
	if err != nil {
		return nil, err
	}

	return &pb.ListRetailCategoryRes{
		RetailCategoryList: types.ToPbRetailCategoryList(res),
	}, nil
}
func (s *grpcServer) GetMedicineCategory(ctx context.Context, req *pb.GetMedicineCategoryReq) (*pb.MedicineCategory, error) {
	res, err := s.service.GetMedicineCategory(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return types.ToPbMedicineCategory(res), err
}
func (s *grpcServer) ListMedicineCategory(ctx context.Context, req *pb.ListMedicineCategoryReq) (*pb.ListMedicineCategoryRes, error) {
	res, err := s.service.ListMedicineCategory(ctx, req.ShopId)
	if err != nil {
		return nil, err
	}

	return &pb.ListMedicineCategoryRes{
		MedicineCategoryList: types.ToPbMedicineCategoryList(res),
	}, nil
}
