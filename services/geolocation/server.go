package geolocation

import (
	"context"
	"fmt"
	"net"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"github.com/yash91989201/superfast-delivery-api/common/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service Service
	pb.UnimplementedGeolocationServiceServer
}

func StartGRPCServer(s Service, serviceUrl string) error {
	listener, err := net.Listen("tcp", serviceUrl)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	server := grpc.NewServer()
	pb.RegisterGeolocationServiceServer(server, &grpcServer{service: s})
	reflection.Register(server)

	return server.Serve(listener)
}

func (s *grpcServer) ReverseGeocode(ctx context.Context, req *pb.ReverseGeocodeReq) (*pb.AddressDetail, error) {
	addressDetail, err := s.service.ReverseGeocode(ctx, req.Latitude, req.Longitude, req.AddressId)
	if err != nil {
		return nil, fmt.Errorf("failed to get address: %w", err)
	}

	return types.ToPbAddressDetail(addressDetail), nil
}
