package product

import (
	"net"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
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
