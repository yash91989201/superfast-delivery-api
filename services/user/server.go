package user

import (
	"context"
	"net"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"github.com/yash91989201/superfast-delivery-api/common/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedUserServiceServer
	service Service
}

func Start(s Service, serviceUrl string) error {

	listener, err := net.Listen("tcp", serviceUrl)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &grpcServer{service: s})

	reflection.Register(server)

	return server.Serve(listener)
}

func (s *grpcServer) CreateProfile(ctx context.Context, req *pb.CreateProfileReq) (*pb.Profile, error) {
	res, err := s.service.CreateProfile(ctx, &types.CreateProfile{
		Name:        req.Name,
		ImageUrl:    req.ImageUrl,
		Dob:         types.PbDateToTime(req.Dob),
		Anniversary: types.PbDateToTime(req.Anniversary),
		Gender:      types.ToGenderPtr(req.Gender),
		AuthID:      req.AuthId,
	})

	if err != nil {
		return nil, err
	}

	return types.ToPbProfile(res), nil
}

func (s *grpcServer) GetProfile(ctx context.Context, req *pb.GetProfileReq) (*pb.Profile, error) {
	res, err := s.service.GetProfile(ctx, req.AuthId)
	if err != nil {
		return nil, err
	}

	return types.ToPbProfile(res), nil
}

func (s *grpcServer) UpdateProfile(ctx context.Context, req *pb.UpdateProfileReq) (*pb.Profile, error) {
	updatedProfile := types.PbUpdateProfileReqToProfile(req)
	err := s.service.UpdateProfile(ctx, updatedProfile)
	if err != nil {
		return nil, err
	}

	return types.ToPbProfile(updatedProfile), nil
}

func (s *grpcServer) DeleteProfile(ctx context.Context, req *pb.DeleteProfileReq) (*pb.EmptyRes, error) {
	if err := s.service.DeleteProfile(ctx, req.AuthId); err != nil {
		return nil, err
	}

	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) CreateDeliveryAddress(ctx context.Context, req *pb.CreateDeliveryAddressReq) (*pb.DeliveryAddress, error) {
	res, err := s.service.CreateDeliveryAddress(ctx, types.ToCreateDeliveryAddress(req))

	if err != nil {
		return nil, err
	}

	return types.ToPbDeliveryAddress(res), nil
}

func (s *grpcServer) GetDeliveryAddress(ctx context.Context, req *pb.GetDeliveryAddressReq) (*pb.DeliveryAddress, error) {
	res, err := s.service.GetDeliveryAddress(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return types.ToPbDeliveryAddress(res), nil
}

func (s *grpcServer) GetDefaultDeliveryAddress(ctx context.Context, req *pb.GetDefaultDeliveryAddressReq) (*pb.DeliveryAddress, error) {
	res, err := s.service.GetDefaultDeliveryAddress(ctx, req.AuthId)
	if err != nil {
		return nil, err
	}

	return types.ToPbDeliveryAddress(res), err
}

func (s *grpcServer) ListDeliveryAddress(ctx context.Context, req *pb.ListDeliveryAddressReq) (*pb.ListDeliveryAddressRes, error) {
	res, err := s.service.GetDeliveryAddresses(ctx, req.AuthId)
	if err != nil {
		return nil, err
	}

	return &pb.ListDeliveryAddressRes{
		DeliveryAddresses: types.ToPbDeliveryAddressList(res),
	}, nil
}

func (s *grpcServer) UpdateDeliveryAddress(ctx context.Context, req *pb.DeliveryAddress) (*pb.DeliveryAddress, error) {
	updatedDeliveryAddress := types.ToDeliveryAddress(req)
	if err := s.service.UpdateDeliveryAddress(ctx, updatedDeliveryAddress); err != nil {
		return nil, err
	}

	return types.ToPbDeliveryAddress(updatedDeliveryAddress), nil
}

func (s *grpcServer) UpdateDefaultDeliveryAddress(ctx context.Context, req *pb.UpdateDefaultDeliveryAddressReq) (*pb.EmptyRes, error) {
	if err := s.service.UpdateDefaultDeliveryAddress(ctx, req.DeliveryAddressId, req.AuthId); err != nil {
		return nil, err
	}

	return &pb.EmptyRes{}, nil
}

func (s *grpcServer) DeleteDeliveryAddress(ctx context.Context, req *pb.DeleteDeliveryAddressReq) (*pb.EmptyRes, error) {
	if err := s.service.DeleteDeliveryAddress(ctx, req.Id); err != nil {
		return nil, err
	}

	return &pb.EmptyRes{}, nil
}
