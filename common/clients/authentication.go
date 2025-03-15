package clients

import (
	"context"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthenticationClient struct {
	conn    *grpc.ClientConn
	service pb.AuthenticationServiceClient
}

func NewAuthenticationClient(serviceUrl string) (*AuthenticationClient, error) {
	conn, err := grpc.NewClient(serviceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	s := pb.NewAuthenticationServiceClient(conn)
	return &AuthenticationClient{conn, s}, nil
}

func (c *AuthenticationClient) GetConn() *grpc.ClientConn {
	return c.conn
}

func (c *AuthenticationClient) Close() {
	c.conn.Close()
}

func (c *AuthenticationClient) SignInWithEmail(ctx context.Context, req *pb.SignInWithEmailReq) (*pb.SignInRes, error) {
	res, err := c.service.SignInWithEmail(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *AuthenticationClient) SignInWithPhone(ctx context.Context, req *pb.SignInWithPhoneReq) (*pb.SignInRes, error) {
	res, err := c.service.SignInWithPhone(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *AuthenticationClient) SignInWithGoogle(ctx context.Context, req *pb.SignInWithGoogleReq) (*pb.SignInRes, error) {
	res, err := c.service.SignInWithGoogle(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *AuthenticationClient) RefreshToken(ctx context.Context, req *pb.RefreshTokenReq) (*pb.SignInRes, error) {
	res, err := c.service.RefreshToken(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *AuthenticationClient) LogOut(ctx context.Context, req *pb.LogOutReq) (*pb.SignInRes, error) {
	res, err := c.service.LogOut(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *AuthenticationClient) GetAuthById(ctx context.Context, req *pb.GetAuthByIdReq) (*pb.Auth, error) {
	res, err := c.service.GetAuthById(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *AuthenticationClient) GetAuth(ctx context.Context, req *pb.GetAuthReq) (*pb.Auth, error) {
	res, err := c.service.GetAuth(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *AuthenticationClient) ValidateSession(ctx context.Context, req *pb.ValidateSessionReq) (*pb.ValidateSessionRes, error) {
	res, err := c.service.ValidateSession(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
