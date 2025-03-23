package product

import (
	"context"
	"log"
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

func (s *grpcServer) CreateMenuItemVariant(ctx context.Context, req *pb.CreateItemVariantReq) (*pb.ItemVariant, error) {
	res, err := s.service.CreateMenuItemVariant(ctx, types.ToCreateItemVariant(req))
	if err != nil {
		return nil, err
	}

	return types.ToPbItemVariant(res), nil
}

func (s *grpcServer) CreateMenuItemAddon(ctx context.Context, req *pb.CreateItemAddonReq) (*pb.ItemAddon, error) {
	res, err := s.service.CreateMenuItemAddon(ctx, types.ToCreateItemAddon(req))
	if err != nil {
		return nil, err
	}

	return types.ToPbItemAddon(res), nil
}

func (s *grpcServer) CreateRetailCategory(ctx context.Context, req *pb.CreateRetailCategoryReq) (*pb.RetailCategory, error) {
	res, err := s.service.CreateRetailCategory(ctx, &types.CreateRetailCategory{
		CategoryName: req.CategoryName,
		ImageURL:     req.ImageUrl,
		ShopID:       req.ShopId,
	})
	if err != nil {
		return nil, err
	}

	return types.ToPbRetailCategory(res), nil
}

func (s *grpcServer) CreateRetailItem(ctx context.Context, req *pb.CreateRetailItemReq) (*pb.RetailItem, error) {
	res, err := s.service.CreateRetailItem(ctx, &types.CreateRetailItem{
		Name:        req.Name,
		ImageURL:    req.ImageUrl,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  types.HexToObjectID(req.CategoryId),
	})

	if err != nil {
		return nil, err
	}
	return types.ToPbRetailItem(res), nil
}

func (s *grpcServer) CreateRetailItemVariant(ctx context.Context, req *pb.CreateItemVariantReq) (*pb.ItemVariant, error) {
	res, err := s.service.CreateRetailItemVariant(ctx, types.ToCreateItemVariant(req))
	if err != nil {
		return nil, err
	}

	return types.ToPbItemVariant(res), nil
}

func (s *grpcServer) CreateMedicineCategory(ctx context.Context, req *pb.CreateMedicineCategoryReq) (*pb.MedicineCategory, error) {
	res, err := s.service.CreateMedicineCategory(ctx, &types.CreateMedicineCategory{
		CategoryName: req.CategoryName,
		ImageURL:     req.ImageUrl,
		ShopID:       req.ShopId,
	})
	if err != nil {
		return nil, err
	}
	return types.ToPbMedicineCategory(res), nil
}

func (s *grpcServer) CreateMedicineItem(ctx context.Context, req *pb.CreateMedicineItemReq) (*pb.MedicineItem, error) {
	res, err := s.service.CreateMedicineItem(ctx, &types.CreateMedicineItem{
		Name:        req.Name,
		Price:       req.Price,
		ImageURL:    req.ImageUrl,
		Description: req.Description,
		CategoryID:  types.HexToObjectID(req.CategoryId),
	})
	if err != nil {
		return nil, err
	}
	return types.ToPbMedicineItem(res), nil
}

func (s *grpcServer) GetRestaurantMenu(ctx context.Context, req *pb.GetRestaurantMenuReq) (*pb.RestaurantMenu, error) {
	res, err := s.service.GetRestaurantMenu(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return types.ToPbRestaurantMenu(res), nil
}

func (s *grpcServer) GetMenuItem(ctx context.Context, req *pb.GetMenuItemReq) (*pb.MenuItem, error) {
	res, err := s.service.GetMenuItem(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return types.ToPbMenuItem(res), nil
}

func (s *grpcServer) GetMenuItemVariant(ctx context.Context, req *pb.GetItemVariantReq) (*pb.ItemVariant, error) {
	res, err := s.service.GetMenuItemVariant(ctx, req.ItemId, req.VariantId)
	if err != nil {
		return nil, err
	}

	return types.ToPbItemVariant(res), nil
}

func (s *grpcServer) GetMenuItemAddon(ctx context.Context, req *pb.GetItemAddonReq) (*pb.ItemAddon, error) {
	res, err := s.service.GetMenuItemAddon(ctx, req.ItemId, req.AddonId)
	if err != nil {
		return nil, err
	}

	return types.ToPbItemAddon(res), nil
}

func (s *grpcServer) GetRetailCategory(ctx context.Context, req *pb.GetRetailCategoryReq) (*pb.RetailCategory, error) {
	res, err := s.service.GetRetailCategory(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return types.ToPbRetailCategory(res), nil
}

func (s *grpcServer) GetRetailItem(ctx context.Context, req *pb.GetRetailItemReq) (*pb.RetailItem, error) {
	res, err := s.service.GetRetailItem(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return types.ToPbRetailItem(res), nil
}

func (s *grpcServer) GetRetailItemVariant(ctx context.Context, req *pb.GetItemVariantReq) (*pb.ItemVariant, error) {
	res, err := s.service.GetRetailItemVariant(ctx, req.ItemId, req.VariantId)
	if err != nil {
		return nil, err
	}

	return types.ToPbItemVariant(res), nil
}

func (s *grpcServer) GetMedicineCategory(ctx context.Context, req *pb.GetMedicineCategoryReq) (*pb.MedicineCategory, error) {
	res, err := s.service.GetMedicineCategory(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return types.ToPbMedicineCategory(res), nil
}

func (s *grpcServer) GetMedicineItem(ctx context.Context, req *pb.GetMedicineItemReq) (*pb.MedicineItem, error) {
	res, err := s.service.GetMedicineItem(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return types.ToPbMedicineItem(res), nil
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

func (s *grpcServer) ListMenuItem(ctx context.Context, req *pb.ListMenuItemReq) (*pb.ListMenuItemRes, error) {
	res, err := s.service.ListMenuItem(ctx, req.MenuId)
	if err != nil {
		return nil, err
	}
	return &pb.ListMenuItemRes{
		MenuItemList: types.ToPbMenuItemList(res),
	}, nil
}

func (s *grpcServer) ListMenuItemVariant(ctx context.Context, req *pb.ListItemVariantReq) (*pb.ListItemVariantRes, error) {
	res, err := s.service.ListMenuItemVariant(ctx, req.ItemId)
	if err != nil {
		return nil, err
	}

	return &pb.ListItemVariantRes{
		Variants: types.ToPbItemVariants(res),
	}, nil
}

func (s *grpcServer) ListMenuItemAddon(ctx context.Context, req *pb.ListItemAddonReq) (*pb.ListItemAddonRes, error) {
	res, err := s.service.ListMenuItemAddon(ctx, req.ItemId)
	if err != nil {
		return nil, err
	}

	return &pb.ListItemAddonRes{
		Addons: types.ToPbItemAddons(res),
	}, nil
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

func (s *grpcServer) ListRetailItem(ctx context.Context, req *pb.ListRetailItemReq) (*pb.ListRetailItemRes, error) {
	res, err := s.service.ListRetailItem(ctx, req.CategoryId)
	if err != nil {
		return nil, err
	}
	return &pb.ListRetailItemRes{
		RetailItemList: types.ToPbRetailItemList(res),
	}, nil
}

func (s *grpcServer) ListRetailItemVariant(ctx context.Context, req *pb.ListItemVariantReq) (*pb.ListItemVariantRes, error) {
	res, err := s.service.ListRetailItemVariant(ctx, req.ItemId)
	if err != nil {
		return nil, err
	}

	return &pb.ListItemVariantRes{
		Variants: types.ToPbItemVariants(res),
	}, nil
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

func (s *grpcServer) ListMedicineItem(ctx context.Context, req *pb.ListMedicineItemReq) (*pb.ListMedicineItemRes, error) {
	res, err := s.service.ListMedicineItem(ctx, req.CategoryId)
	if err != nil {
		return nil, err
	}
	return &pb.ListMedicineItemRes{
		MedicineItemList: types.ToPbMedicineItemList(res),
	}, nil
}

func (s *grpcServer) UpdateRestaurantMenu(ctx context.Context, req *pb.UpdateRestaurantMenuReq) (*pb.EmptyRes, error) {
	err := s.service.UpdateRestaurantMenu(ctx, types.ToUpdateRestaurantMenu(req))
	if err != nil {
		return nil, err
	}
	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) UpdateMenuItem(ctx context.Context, req *pb.UpdateMenuItemReq) (*pb.EmptyRes, error) {
	err := s.service.UpdateMenuItem(ctx, types.ToUpdateMenuItem(req))
	if err != nil {
		return nil, err
	}
	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) UpdateMenuItemVariant(ctx context.Context, req *pb.UpdateItemVariantReq) (*pb.EmptyRes, error) {
	if err := s.service.UpdateMenuItemVariant(ctx, types.ToUpdateItemVariant(req)); err != nil {
		return nil, err
	}

	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) UpdateMenuItemAddon(ctx context.Context, req *pb.UpdateItemAddonReq) (*pb.EmptyRes, error) {
	if err := s.service.UpdateMenuItemAddon(ctx, types.ToUpdateItemAddon(req)); err != nil {
		return nil, err
	}

	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) UpdateRetailCategory(ctx context.Context, req *pb.UpdateRetailCategoryReq) (*pb.EmptyRes, error) {
	log.Printf("server: %+v", req)
	err := s.service.UpdateRetailCategory(ctx, types.ToUpdateRetailCategory(req))
	if err != nil {
		return nil, err
	}
	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) UpdateRetailItem(ctx context.Context, req *pb.UpdateRetailItemReq) (*pb.EmptyRes, error) {
	err := s.service.UpdateRetailItem(ctx, types.ToUpdateRetailItem(req))
	if err != nil {
		return nil, err
	}
	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) UpdateRetailItemVariant(ctx context.Context, req *pb.UpdateItemVariantReq) (*pb.EmptyRes, error) {
	if err := s.service.UpdateRetailItemVariant(ctx, types.ToUpdateItemVariant(req)); err != nil {
		return nil, err
	}

	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) UpdateMedicineCategory(ctx context.Context, req *pb.UpdateMedicineCategoryReq) (*pb.EmptyRes, error) {
	err := s.service.UpdateMedicineCategory(ctx, types.ToUpdateMedicineCategory(req))
	if err != nil {
		return nil, err
	}
	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) UpdateMedicineItem(ctx context.Context, req *pb.UpdateMedicineItemReq) (*pb.EmptyRes, error) {
	err := s.service.UpdateMedicineItem(ctx, types.ToUpdateMedicineItem(req))
	if err != nil {
		return nil, err
	}
	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) DeleteRestaurantMenu(ctx context.Context, req *pb.DeleteRestaurantMenuReq) (*pb.EmptyRes, error) {
	err := s.service.DeleteRestaurantMenu(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) DeleteMenuItem(ctx context.Context, req *pb.DeleteMenuItemReq) (*pb.EmptyRes, error) {
	err := s.service.DeleteMenuItem(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) DeleteMenuItemVariant(ctx context.Context, req *pb.DeleteItemVariantReq) (*pb.EmptyRes, error) {
	if err := s.service.DeleteMenuItemVariant(ctx, req.ItemId, req.VariantId); err != nil {
		return nil, err
	}

	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) DeleteMenuItemAddon(ctx context.Context, req *pb.DeleteItemAddonReq) (*pb.EmptyRes, error) {
	if err := s.service.DeleteMenuItemAddon(ctx, req.ItemId, req.AddonId); err != nil {
		return nil, err
	}

	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) DeleteRetailCategory(ctx context.Context, req *pb.DeleteRetailCategoryReq) (*pb.EmptyRes, error) {
	err := s.service.DeleteRetailCategory(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) DeleteRetailItem(ctx context.Context, req *pb.DeleteRetailItemReq) (*pb.EmptyRes, error) {
	err := s.service.DeleteRetailItem(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) DeleteRetailItemVariant(ctx context.Context, req *pb.DeleteItemVariantReq) (*pb.EmptyRes, error) {
	if err := s.service.DeleteRetailItemVariant(ctx, req.ItemId, req.VariantId); err != nil {
		return nil, err
	}

	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) DeleteMedicineCategory(ctx context.Context, req *pb.DeleteMedicineCategoryReq) (*pb.EmptyRes, error) {
	err := s.service.DeleteMedicineCategory(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) DeleteMedicineItem(ctx context.Context, req *pb.DeleteMedicineItemReq) (*pb.EmptyRes, error) {
	err := s.service.DeleteMedicineItem(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.EmptyRes{}, nil
}
