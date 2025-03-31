package authentication

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/nrednav/cuid2"
	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"github.com/yash91989201/superfast-delivery-api/common/types"
	"github.com/yash91989201/superfast-delivery-api/common/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type grpcServer struct {
	service      Service
	TokenManager *utils.TokenManager
	pb.UnimplementedAuthenticationServiceServer
}

func StartGRPCServer(s Service, tokenManager *utils.TokenManager, serviceUrl string) error {
	listener, err := net.Listen("tcp", serviceUrl)
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	pb.RegisterAuthenticationServiceServer(server, &grpcServer{service: s, TokenManager: tokenManager})

	reflection.Register(server)

	return server.Serve(listener)
}

func (s *grpcServer) SignInWithEmail(ctx context.Context, req *pb.SignInWithEmailReq) (*pb.SignInRes, error) {
	if req.Otp != nil {
		ev, err := s.service.GetEmailVerification(ctx, req.Email)
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "Verification not found")
		}

		if isTokenExpired(ev.ExpiresAt) {
			return nil, status.Errorf(codes.Unavailable, "Otp expired, try again")
		}

		if !isTokenValid(req.Otp, ev.Token) {
			return nil, fmt.Errorf("Otp incorrect, try again")
		}

		_ = s.service.DeleteEmailVerification(ctx, ev.Email)

		// user has now verified themselves using otp
		// Get auth, if not found then create
		var auth *types.Auth
		auth, err = s.service.GetAuth(ctx, &ev.Email, nil)
		if err != nil {
			auth, err = s.service.CreateAuth(
				ctx,
				&types.CreateAuth{
					Email:         &ev.Email,
					EmailVerified: true,
					Phone:         nil,
					AuthRole:      types.ToAuthRole(req.AuthRole),
				})

			if err != nil {
				return nil, fmt.Errorf("Error creating account: %w", err)
			}
		}

		sessionID := cuid2.Generate()
		accessToken, err := s.TokenManager.GenerateAccessToken(auth, sessionID)
		if err != nil {
			return nil, fmt.Errorf("failed to create access token: %w", err)
		}

		refreshToken, err := s.TokenManager.GenerateRefreshToken()
		if err != nil {
			return nil, fmt.Errorf("Failed to create refresh token: %w", err)
		}

		session := &types.Session{
			ID:           sessionID,
			AuthID:       auth.ID,
			RefreshToken: refreshToken,
			IsRevoked:    false,
			ExpiresAt:    time.Now().Add(30 * 24 * time.Hour),
		}

		if err = s.service.CreateSession(ctx, session); err != nil {
			return nil, fmt.Errorf("Failed to create session: %w", err)
		}

		signInRes := &types.SignInRes{
			Auth: auth,
			Session: &types.ClientSession{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
		}

		return types.ToPbSignInRes(signInRes), nil
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
		return &pb.SignInRes{Auth: nil}, nil
	}

	if isTokenExpired(ev.ExpiresAt) {
		_ = s.service.DeleteEmailVerification(ctx, ev.Email)

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
		return &pb.SignInRes{Auth: nil}, nil
	}

	// TODO: existing token not expired send email again
	return &pb.SignInRes{Auth: nil}, nil
}

func (s *grpcServer) SignInWithPhone(ctx context.Context, req *pb.SignInWithPhoneReq) (*pb.SignInRes, error) {
	if req.Otp != nil {
		pv, err := s.service.GetPhoneVerification(ctx, req.Phone)
		if err != nil {
			return nil, fmt.Errorf("Verification failed, try again: %w", err)
		}

		if isTokenExpired(pv.ExpiresAt) {
			return nil, fmt.Errorf("Otp expired, try again")
		}

		if !isTokenValid(req.Otp, pv.Token) {
			return nil, fmt.Errorf("Otp incorrect, try again")
		}

		_ = s.service.DeletePhoneVerification(ctx, pv.Phone)

		// Get auth, if not found then create
		var auth *types.Auth
		auth, err = s.service.GetAuth(ctx, nil, &pv.Phone)
		if err != nil {
			auth, err = s.service.CreateAuth(
				ctx,
				&types.CreateAuth{
					Email:         nil,
					EmailVerified: false,
					Phone:         &pv.Phone,
					AuthRole:      types.ToAuthRole(req.AuthRole),
				})

			if err != nil {
				return nil, fmt.Errorf("Error creating account: %w", err)
			}
		}

		sessionID := cuid2.Generate()
		accessToken, err := s.TokenManager.GenerateAccessToken(auth, sessionID)
		if err != nil {
			return nil, err
		}

		refreshToken, err := s.TokenManager.GenerateRefreshToken()
		if err != nil {
			return nil, err
		}

		session := &types.Session{
			ID:           sessionID,
			AuthID:       auth.ID,
			RefreshToken: refreshToken,
			IsRevoked:    false,
			ExpiresAt:    time.Now().Add(30 * 24 * time.Hour),
		}

		if err = s.service.CreateSession(ctx, session); err != nil {
			return nil, fmt.Errorf("Failed to create session: %w", err)
		}

		signInRes := &types.SignInRes{
			Auth: auth,
			Session: &types.ClientSession{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
		}

		return types.ToPbSignInRes(signInRes), nil
	}

	pv, err := s.service.GetPhoneVerification(ctx, req.Phone)
	if err != nil {
		err := s.service.CreatePhoneVerification(
			ctx,
			&types.PhoneVerification{
				Token:     generateToken(),
				Phone:     req.Phone,
				ExpiresAt: getTokenExpiresAt(),
			})

		if err != nil {
			return nil, fmt.Errorf("Verification failed, try again :%w", err)
		}

		// TODO: send otp using email service
		return &pb.SignInRes{}, nil
	}

	if isTokenExpired(pv.ExpiresAt) {
		_ = s.service.DeletePhoneVerification(ctx, pv.Phone)

		err := s.service.CreatePhoneVerification(
			ctx,
			&types.PhoneVerification{
				Token:     generateToken(),
				Phone:     req.Phone,
				ExpiresAt: getTokenExpiresAt(),
			})

		if err != nil {
			return nil, fmt.Errorf("Verification failed, try again :%w", err)
		}

		// TODO: send otp using email service
		return &pb.SignInRes{}, nil
	}

	// TODO: existing token not expired send email again
	return &pb.SignInRes{}, nil
}

func (s *grpcServer) SignInWithGoogle(ctx context.Context, req *pb.SignInWithGoogleReq) (*pb.SignInRes, error) {
	return nil, nil
}

func (s *grpcServer) GetAuthById(ctx context.Context, req *pb.GetAuthByIdReq) (*pb.Auth, error) {
	auth, err := s.service.GetAuthById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return types.ToPbAuth(auth), nil
}

func (s *grpcServer) GetAuth(ctx context.Context, req *pb.GetAuthReq) (*pb.Auth, error) {
	auth, err := s.service.GetAuth(ctx, req.Email, req.Phone)
	if err != nil {
		return nil, err
	}

	return types.ToPbAuth(auth), nil
}

func (s *grpcServer) RefreshAccessToken(ctx context.Context, req *pb.RefreshAccessTokenReq) (*pb.SignInRes, error) {
	session, err := s.service.GetSession(ctx, req.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Session expired, sign in again")
	}

	if session.IsRevoked || session.ExpiresAt.Before(time.Now()) {
		return nil, status.Errorf(codes.NotFound, "Session expired, sign in again")
	}

	auth, err := s.service.GetAuthById(ctx, session.AuthID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Account not found")
	}

	err = s.service.DeleteSession(ctx, session.ID)
	if err != nil {
		return nil, err
	}

	sessionID := cuid2.Generate()
	accessToken, err := s.TokenManager.GenerateAccessToken(auth, sessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %w", err)
	}

	refreshToken, err := s.TokenManager.GenerateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("failed to create refresh token: %w", err)
	}

	session = &types.Session{
		ID:           sessionID,
		AuthID:       auth.ID,
		RefreshToken: refreshToken,
		IsRevoked:    false,
		ExpiresAt:    time.Now().Add(30 * 24 * time.Hour),
	}

	if err = s.service.CreateSession(ctx, session); err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	signInRes := &types.SignInRes{
		Auth: auth,
		Session: &types.ClientSession{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}

	return types.ToPbSignInRes(signInRes), nil
}

func (s *grpcServer) LogOut(ctx context.Context, req *pb.LogOutReq) (*pb.SignInRes, error) {
	session, err := s.service.GetSessionById(ctx, req.SessionId)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Session expired, sign in again")
	}

	if session.IsRevoked || session.ExpiresAt.Before(time.Now()) {
		return nil, status.Errorf(codes.Unauthenticated, "Session expired, sign in again")
	}

	_, err = s.service.GetAuthById(ctx, session.AuthID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Account not found")
	}

	err = s.service.DeleteSession(ctx, session.ID)
	if err != nil {
		return nil, err
	}

	return &pb.SignInRes{}, nil
}

func (s *grpcServer) ValidateSession(ctx context.Context, req *pb.ValidateSessionReq) (*pb.ValidateSessionRes, error) {
	token, err := s.TokenManager.VerifyAccessToken(req.AccessToken)
	if err != nil {
		return nil, err
	}

	session, err := s.service.GetSessionById(ctx, token.RegisteredClaims.ID)
	if err != nil {
		return nil, err
	}

	auth, err := s.service.GetAuthById(ctx, token.RegisteredClaims.Subject)
	if err != nil {
		return nil, err
	}

	if session.AuthID != auth.ID {
		return nil, fmt.Errorf("invalid auth, please login")
	}

	return &pb.ValidateSessionRes{
		Auth:      types.ToPbAuth(auth),
		SessionId: session.ID,
	}, nil
}
