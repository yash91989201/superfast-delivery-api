package authentication

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/nrednav/cuid2"
	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"github.com/yash91989201/superfast-delivery-api/common/types"
	"github.com/yash91989201/superfast-delivery-api/common/utils/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type grpcServer struct {
	service  Service
	JwtMaker *token.JWTMaker
	pb.UnimplementedAuthenticationServiceServer
}

func StartGRPCServer(s Service, jwtMaker *token.JWTMaker, serviceUrl string) error {
	listener, err := net.Listen("tcp", serviceUrl)
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	pb.RegisterAuthenticationServiceServer(server, &grpcServer{service: s, JwtMaker: jwtMaker})

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
					AuthRole:      types.Customer,
				})

			if err != nil {
				return nil, fmt.Errorf("Error creating account: %w", err)
			}
		}

		newAuthClaim := token.NewAuthClaim{
			Email:         auth.Email,
			EmailVerified: auth.EmailVerified,
			Phone:         auth.Phone,
			Role:          auth.AuthRole,
			AuthId:        auth.ID,
			SessionId:     cuid2.Generate(),
			Duration:      15 * time.Minute,
		}

		accessToken, accessClaims, err := s.JwtMaker.CreateToken(newAuthClaim)
		if err != nil {
			return nil, fmt.Errorf("Failed to create access token: %w", err)
		}

		refreshToken, refreshClaims, err := s.JwtMaker.CreateToken(newAuthClaim)
		if err != nil {
			return nil, fmt.Errorf("Failed to create refresh token: %w", err)
		}

		session := &types.Session{
			ID:           refreshClaims.RegisteredClaims.ID,
			AuthID:       refreshClaims.RegisteredClaims.Subject,
			RefreshToken: refreshToken,
			IsRevoked:    false,
			ExpiresAt:    refreshClaims.ExpiresAt.Time,
		}

		if err = s.service.CreateSession(ctx, session); err != nil {
			return nil, fmt.Errorf("Failed to create session: %w", err)
		}

		signInRes := &types.SignInRes{
			Auth: auth,
			Session: &types.ClientSession{
				ID:                   session.ID,
				AccessToken:          accessToken,
				AccessTokenExpiresAt: accessClaims.ExpiresAt.Time,
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
					EmailVerified: true,
					Phone:         &pv.Phone,
					AuthRole:      types.Customer,
				})

			if err != nil {
				return nil, fmt.Errorf("Error creating account: %w", err)
			}
		}

		newAuthClaim := token.NewAuthClaim{
			Email:         auth.Email,
			EmailVerified: auth.EmailVerified,
			Phone:         auth.Phone,
			Role:          auth.AuthRole,
			AuthId:        auth.ID,
			SessionId:     cuid2.Generate(),
			Duration:      15 * time.Minute,
		}

		accessToken, accessClaims, err := s.JwtMaker.CreateToken(newAuthClaim)
		if err != nil {
			return nil, fmt.Errorf("Failed to create access token: %w", err)
		}

		refreshToken, refreshClaims, err := s.JwtMaker.CreateToken(newAuthClaim)
		if err != nil {
			return nil, fmt.Errorf("Failed to create refresh token: %w", err)
		}

		session := &types.Session{
			ID:           refreshClaims.RegisteredClaims.ID,
			AuthID:       refreshClaims.RegisteredClaims.Subject,
			RefreshToken: refreshToken,
			IsRevoked:    false,
			ExpiresAt:    refreshClaims.ExpiresAt.Time,
		}

		if err = s.service.CreateSession(ctx, session); err != nil {
			return nil, fmt.Errorf("Failed to create session: %w", err)
		}

		signInRes := &types.SignInRes{
			Auth: auth,
			Session: &types.ClientSession{
				ID:                   session.ID,
				AccessToken:          accessToken,
				AccessTokenExpiresAt: accessClaims.ExpiresAt.Time,
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

func (s *grpcServer) RefreshToken(ctx context.Context, req *pb.RefreshTokenReq) (*pb.SignInRes, error) {
	// get session from db
	session, err := s.service.GetSession(ctx, req.SessionId)
	if err != nil || session.IsRevoked || session.ExpiresAt.Before(time.Now()) {
		return nil, status.Errorf(codes.NotFound, "Session expired, sign in again")
	}

	auth, err := s.service.GetAuthById(ctx, session.AuthID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Account not found")
	}

	// delete and create new refresh refresh token
	err = s.service.DeleteSession(ctx, session.ID)
	if err != nil {
		return nil, err
	}

	newAuthClaim := token.NewAuthClaim{
		Email:         auth.Email,
		EmailVerified: auth.EmailVerified,
		Phone:         auth.Phone,
		Role:          auth.AuthRole,
		AuthId:        auth.ID,
		SessionId:     cuid2.Generate(),
		Duration:      15 * time.Minute,
	}

	accessToken, accessClaims, err := s.JwtMaker.CreateToken(newAuthClaim)
	if err != nil {
		return nil, fmt.Errorf("Failed to create access token: %w", err)
	}

	refreshToken, refreshClaims, err := s.JwtMaker.CreateToken(newAuthClaim)
	if err != nil {
		return nil, fmt.Errorf("Failed to create refresh token: %w", err)
	}

	newSession := &types.Session{
		ID:           refreshClaims.RegisteredClaims.ID,
		AuthID:       refreshClaims.RegisteredClaims.Subject,
		RefreshToken: refreshToken,
		IsRevoked:    false,
		ExpiresAt:    refreshClaims.ExpiresAt.Time,
	}

	if err = s.service.CreateSession(ctx, newSession); err != nil {
		return nil, fmt.Errorf("Failed to create session: %w", err)
	}

	signInRes := &types.SignInRes{
		Auth: auth,
		Session: &types.ClientSession{
			ID:                   session.ID,
			AccessToken:          accessToken,
			AccessTokenExpiresAt: accessClaims.ExpiresAt.Time,
		},
	}

	return types.ToPbSignInRes(signInRes), nil
}

func (s *grpcServer) LogOut(ctx context.Context, req *pb.LogOutReq) (*pb.SignInRes, error) {

	session, err := s.service.GetSession(ctx, req.SessionId)
	if err != nil || session.IsRevoked || session.ExpiresAt.Before(time.Now()) {
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

	token, err := s.JwtMaker.VerifyToken(req.AuthToken)
	if err != nil {
		return nil, err
	}

	auth, err := s.service.GetAuthById(ctx, token.RegisteredClaims.Subject)
	if err != nil {
		return nil, err
	}

	session, err := s.service.GetSession(ctx, token.RegisteredClaims.ID)
	if err != nil {
		return nil, err
	}

	if auth.ID != session.AuthID {
		return nil, fmt.Errorf("invalid auth, please login")
	}

	return &pb.ValidateSessionRes{
		Valid: true,
		Auth:  types.ToPbAuth(auth),
	}, nil
}
