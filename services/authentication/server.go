package authentication

import (
	"context"
	"fmt"
	"log"
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
	// Check if the account already exists
	auth, err := s.service.GetAuth(ctx, &req.Email, nil)
	if err == nil {
		log.Println("1. account not found")
		// Account found, return the authentication details
		return types.ToPbAuth(auth), nil
	}

	if req.Otp != nil {
		ev, err := s.service.GetEmailVerification(ctx, req.Email)
		if err != nil {
			return nil, fmt.Errorf("Verification failed, try again: %w", err)
		}

		if isTokenExpired(ev.ExpiresAt) {
			return nil, fmt.Errorf("Otp expired, try again: %w", err)
		}

		if !isTokenValid(req.Otp, ev.Token) {
			return nil, fmt.Errorf("Otp incorrect, try again")
		}

		_ = s.service.DeleteEmailVerification(ctx, ev.Email)

		auth, err := s.service.CreateAuth(
			ctx,
			&types.CreateAuth{
				Email:         &ev.Email,
				EmailVerified: true,
				Phone:         nil,
				Role:          types.Customer,
			})

		if err != nil {
			return nil, fmt.Errorf("Error creating account, try again :%w", err)
		}

		log.Println("2. account created")
		return types.ToPbAuth(auth), nil
	}

	ev, err := s.service.GetEmailVerification(ctx, req.Email)
	if err != nil {
		err := s.service.CreateEmailVerification(
			ctx,
			&types.EmailVerification{
				Token:     generateToken(),
				Email:     req.Email,
				ExpiresAt: getTokenExpiresAt(),
			})

		if err != nil {
			return nil, fmt.Errorf("Verification failed, try again :%w", err)
		}

		// TODO: send otp using email service
		return &pb.Auth{}, nil
	}

	if isTokenExpired(ev.ExpiresAt) {
		err = s.service.DeleteEmailVerification(ctx, ev.Email)
		if err != nil {
			return nil, fmt.Errorf("Verification failed, try again :%w", err)
		}

		err := s.service.CreateEmailVerification(
			ctx,
			&types.EmailVerification{
				Token:     generateToken(),
				Email:     req.Email,
				ExpiresAt: getTokenExpiresAt(),
			})

		if err != nil {
			return nil, fmt.Errorf("Verification failed, try again :%w", err)
		}

		// TODO: send otp using email service
		return &pb.Auth{}, nil
	}

	// TODO: existing token not expired send email again
	return &pb.Auth{}, nil
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
