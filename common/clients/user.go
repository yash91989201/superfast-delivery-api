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
	res, err := c.service.CreateProfile(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *UserClient) GetProfile(ctx context.Context, req *pb.GetProfileReq) (*pb.Profile, error) {
	res, err := c.service.GetProfile(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (c *UserClient) UpdateProfile(ctx context.Context, req *pb.UpdateProfileReq) (*pb.Profile, error) {
	res, err := c.service.UpdateProfile(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (c *UserClient) DeleteProfile(ctx context.Context, req *pb.DeleteProfileReq) (*pb.EmptyRes, error) {
	res, err := c.service.DeleteProfile(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (c *UserClient) CreateDeliveryAddress(ctx context.Context, req *pb.CreateDeliveryAddressReq) (*pb.DeliveryAddress, error) {
	res, err := c.service.CreateDeliveryAddress(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (c *UserClient) GetDeliveryAddress(ctx context.Context, req *pb.GetDeliveryAddressReq) (*pb.DeliveryAddress, error) {
	res, err := c.service.GetDeliveryAddress(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (c *UserClient) ListDeliveryAddress(ctx context.Context, req *pb.ListDeliveryAddressReq) (*pb.ListDeliveryAddressRes, error) {
	res, err := c.service.ListDeliveryAddress(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (c *UserClient) UpdateDeliveryAddress(ctx context.Context, req *pb.UpdateDeliveryAddressReq) (*pb.DeliveryAddress, error) {
	res, err := c.service.UpdateDeliveryAddress(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (c *UserClient) DeleteDeliveryAddress(ctx context.Context, req *pb.DeleteDeliveryAddressReq) (*pb.EmptyRes, error) {
	res, err := c.service.DeleteDeliveryAddress(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}
