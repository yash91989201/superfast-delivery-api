package types

import (
	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToAuthRole(t pb.AuthRole) AuthRole {
	switch t {
	case pb.AuthRole_CUSTOMER:
		return Customer
	case pb.AuthRole_DELIVERY_PARTNER:
		return DeliveryPartner
	case pb.AuthRole_VENDOR:
		return Vendor
	case pb.AuthRole_ADMIN:
		return Admin
	default:
		return Customer // Default fallback
	}
}

func ToPbAuthRole(t AuthRole) pb.AuthRole {
	switch t {
	case Admin:
		return pb.AuthRole_ADMIN
	case Customer:
		return pb.AuthRole_CUSTOMER
	case DeliveryPartner:
		return pb.AuthRole_DELIVERY_PARTNER
	case Vendor:
		return pb.AuthRole_VENDOR
	default:
		return pb.AuthRole_CUSTOMER
	}
}

func ToAuth(a *pb.Auth) *Auth {
	return &Auth{
		ID:            a.Id,
		Email:         a.Email,
		EmailVerified: a.EmailVerified,
		Phone:         a.Phone,
		AuthRole:      ToAuthRole(a.AuthRole),
		CreatedAt:     ToTime(a.CreatedAt),
		UpdatedAt:     ToTime(a.CreatedAt),
		DeletedAt:     ToTimePtr(a.DeletedAt),
	}
}

func ToPbAuth(p *Auth) *pb.Auth {
	var deletedAt *timestamppb.Timestamp
	if p.DeletedAt != nil {
		deletedAt = ToPbTimestamp(*p.DeletedAt)
	}

	return &pb.Auth{
		Id:            p.ID,
		Email:         p.Email,
		EmailVerified: p.EmailVerified,
		Phone:         p.Phone,
		AuthRole:      ToPbAuthRole(p.AuthRole),
		CreatedAt:     ToPbTimestamp(p.CreatedAt),
		UpdatedAt:     ToPbTimestamp(p.UpdatedAt),
		DeletedAt:     deletedAt,
	}
}

func ToPbSession(s *ClientSession) *pb.Session {
	return &pb.Session{
		Id:                   s.ID,
		AccessToken:          s.AccessToken,
		AccessTokenExpiresAt: ToPbTimestamp(s.AccessTokenExpiresAt),
	}
}

func ToPbSignInRes(r *SignInRes) *pb.SignInRes {
	return &pb.SignInRes{
		Auth:    ToPbAuth(r.Auth),
		Session: ToPbSession(r.Session),
	}
}
