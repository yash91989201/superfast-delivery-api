package authentication

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
	pb.UnimplementedAuthenticationServiceServer
	service Service
}

func Start(s Service, serviceUrl string) error {

	listener, err := net.Listen("tcp", serviceUrl)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterAuthenticationServiceServer(server, &grpcServer{service: s})

	reflection.Register(server)

	return server.Serve(listener)
}

func (s *grpcServer) SignInWithEmail(ctx context.Context, req *pb.SignInWithEmailReq) (*pb.Auth, error) {
	if req.Otp != nil {
		ev, err := s.service.GetEmailVerification(ctx, req.Email)
		if err != nil {
			return nil, err
		}

		if !isTokenValid(req.Otp, ev.Token) {
			return nil, fmt.Errorf("Invalid OTP")
		}

		if isTokenExpired(ev.ExpiresAt) {
			return nil, fmt.Errorf("OTP expired! Try again.")
		}

		auth, err := s.service.CreateAuth(
			ctx,
			&req.Email,
			true,
			nil,
			types.Customer,
		)

		if err != nil {
			return nil, err
		}

		return types.ToPbAuth(auth), nil
	}

	// create email verification
	ev := &types.EmailVerification{
		Token:     generateToken(),
		Email:     req.Email,
		ExpiresAt: getTokenExpiresAt(),
	}

	if err := s.service.CreateEmailVerification(ctx, ev); err != nil {
		return nil, err
	}

	// send otp to email
	return nil, nil
}

func (s *grpcServer) SignInWithPhone(ctx context.Context, req *pb.SignInWithPhoneReq) (*pb.Auth, error) {
	return nil, nil
}

func (s *grpcServer) SignInWithGoogle(ctx context.Context, req *pb.SignInWithGoogleReq) (*pb.Auth, error) {
	return nil, nil
}

func (s *grpcServer) GetAuthById(ctx context.Context, req *pb.GetAuthByIdReq) (*pb.Auth, error) {
	return nil, nil
}

func (s *grpcServer) GetAuth(ctx context.Context, req *pb.GetAuthReq) (*pb.Auth, error) {
	return nil, nil
}
