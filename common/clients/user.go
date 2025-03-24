package clients

import (
	"context"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	conn    *grpc.ClientConn
	service pb.UserServiceClient
}

func NewUserClient(serviceUrl string) (*UserClient, error) {
	conn, err := grpc.NewClient(serviceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	s := pb.NewUserServiceClient(conn)

	return &UserClient{conn, s}, nil
}

func (c *UserClient) GetConn() *grpc.ClientConn {
	return c.conn
}

func (c *UserClient) Close() {
	c.conn.Close()
}

func (c *UserClient) CreateProfile(ctx context.Context, req *pb.CreateProfileReq) (*pb.Profile, error) {
	return c.service.CreateProfile(ctx, req)
}

func (c *UserClient) GetProfile(ctx context.Context, req *pb.GetProfileReq) (*pb.Profile, error) {
	return c.service.GetProfile(ctx, req)
}

func (c *UserClient) UpdateProfile(ctx context.Context, req *pb.UpdateProfileReq) (*pb.Profile, error) {
	return c.service.UpdateProfile(ctx, req)
}

func (c *UserClient) DeleteProfile(ctx context.Context, req *pb.DeleteProfileReq) (*pb.EmptyRes, error) {
	return c.service.DeleteProfile(ctx, req)
}

func (c *UserClient) CreateDeliveryAddress(ctx context.Context, req *pb.CreateDeliveryAddressReq) (*pb.DeliveryAddress, error) {
	return c.service.CreateDeliveryAddress(ctx, req)
}

func (c *UserClient) GetDeliveryAddress(ctx context.Context, req *pb.GetDeliveryAddressReq) (*pb.DeliveryAddress, error) {
	return c.service.GetDeliveryAddress(ctx, req)
}

func (c *UserClient) GetDefaultDeliveryAddress(ctx context.Context, req *pb.GetDefaultDeliveryAddressReq) (*pb.DeliveryAddress, error) {
	return c.service.GetDefaultDeliveryAddress(ctx, req)
}

func (c *UserClient) ListDeliveryAddress(ctx context.Context, req *pb.ListDeliveryAddressReq) (*pb.ListDeliveryAddressRes, error) {
	return c.service.ListDeliveryAddress(ctx, req)
}

func (c *UserClient) UpdateDeliveryAddress(ctx context.Context, req *pb.DeliveryAddress) (*pb.DeliveryAddress, error) {
	return c.service.UpdateDeliveryAddress(ctx, req)
}

func (c *UserClient) UpdateDefaultDeliveryAddress(ctx context.Context, req *pb.UpdateDefaultDeliveryAddressReq) (*pb.EmptyRes, error) {
	return c.service.UpdateDefaultDeliveryAddress(ctx, req)
}

func (c *UserClient) DeleteDeliveryAddress(ctx context.Context, req *pb.DeleteDeliveryAddressReq) (*pb.EmptyRes, error) {
	return c.service.DeleteDeliveryAddress(ctx, req)
}
