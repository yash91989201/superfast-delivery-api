package types

import (
	"time"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToAuthType(t pb.AuthType) AuthType {
	switch t {
	case pb.AuthType_CUSTOMER:
		return Customer
	case pb.AuthType_DELIVERY_PARTNER:
		return DeliveryPartner
	case pb.AuthType_VENDOR:
		return Vendor
	case pb.AuthType_ADMIN:
		return Admin
	default:
		return Customer // Default fallback
	}
}

func ToPbAuthType(t AuthType) pb.AuthType {
	switch t {
	case Admin:
		return pb.AuthType_ADMIN
	case Customer:
		return pb.AuthType_CUSTOMER
	case DeliveryPartner:
		return pb.AuthType_DELIVERY_PARTNER
	case Vendor:
		return pb.AuthType_VENDOR
	default:
		return pb.AuthType_CUSTOMER
	}
}

func ToTime(ts *timestamppb.Timestamp) time.Time {
	if ts != nil {
		return ts.AsTime()
	}
	return time.Time{}
}

func ToTimePtr(ts *timestamppb.Timestamp) *time.Time {
	if ts != nil {
		t := ts.AsTime()
		return &t
	}
	return nil
}

func ToDate(d *pb.Date) *Date {
	return &Date{
		Year:  d.Year,
		Month: d.Month,
		Day:   d.Day,
	}
}

func ToPbDate(d *Date) *pb.Date {
	return &pb.Date{
		Year:  d.Year,
		Month: d.Month,
		Day:   d.Day,
	}
}

func TimeToPbDate(t *time.Time) *pb.Date {
	return &pb.Date{
		Year:  int32(t.Year()),
		Month: int32(t.Month()),
		Day:   int32(t.Day()),
	}
}

func ToGender(g pb.Gender) Gender {
	switch g {
	case pb.Gender_MALE:
		return Male
	case pb.Gender_FEMALE:
		return Female
	case pb.Gender_OTHERS:
		return Others
	case pb.Gender_UNDISCLOSED:
		return Undisclosed
	default:
		return Undisclosed
	}
}

func ToGenderPtr(g *pb.Gender) *Gender {
	if g == nil {
		return nil
	}

	switch *g {
	case pb.Gender_MALE:
		return genderPtr(Male)
	case pb.Gender_FEMALE:
		return genderPtr(Female)
	case pb.Gender_OTHERS:
		return genderPtr(Others)
	case pb.Gender_UNDISCLOSED:
		return genderPtr(Undisclosed)
	default:
		return nil
	}
}

func genderPtr(g Gender) *Gender {
	return &g
}

func pbGenderPtr(g pb.Gender) *pb.Gender {
	return &g
}

func ToPbGenderPtr(g *Gender) *pb.Gender {
	if g == nil {
		return nil
	}

	switch *g {
	case Male:
		return pbGenderPtr(pb.Gender_MALE)
	case Female:
		return pbGenderPtr(pb.Gender_FEMALE)
	case Others:
		return pbGenderPtr(pb.Gender_OTHERS)
	case Undisclosed:
		return pbGenderPtr(pb.Gender_UNDISCLOSED)
	default:
		return nil
	}
}

func ToAuth(a *pb.Auth) *Auth {
	return &Auth{
		Id:            a.Id,
		Email:         a.Email,
		EmailVerified: a.EmailVerified,
		Phone:         a.Phone,
		Type:          ToAuthType(a.Type),
		CreatedAt:     ToTime(a.CreatedAt),
		UpdatedAt:     ToTime(a.CreatedAt),
		DeletedAt:     ToTimePtr(a.DeletedAt),
	}
}

func ToProfile(p *pb.Profile) *Profile {
	return &Profile{
		Id:          "",
		Name:        "",
		ImageUrl:    p.ImageUrl,
		Dob:         ToDate(p.Dob).ToTime(),
		Anniversary: ToDate(p.Anniversary).ToTime(),
		Gender:      ToGenderPtr(p.Gender),
		AuthId:      p.AuthId,
		CreatedAt:   ToTime(p.CreatedAt),
		UpdatedAt:   ToTime(p.UpdatedAt),
	}
}

func ToPbProfile(p *Profile) *pb.Profile {
	return &pb.Profile{
		Id:          p.Id,
		Name:        p.Name,
		ImageUrl:    p.ImageUrl,
		Dob:         TimeToPbDate(p.Dob),
		Anniversary: &pb.Date{},
		Gender:      ToPbGenderPtr(p.Gender),
		AuthId:      p.AuthId,
		CreatedAt:   timestamppb.New(p.CreatedAt),
		UpdatedAt:   timestamppb.New(p.UpdatedAt),
	}
}

func ToPbAuth(p *Auth) *pb.Auth {
	return &pb.Auth{
		Id:            p.Id,
		Email:         p.Email,
		EmailVerified: p.EmailVerified,
		Phone:         p.Phone,
		Type:          ToPbAuthType(p.Type),
		CreatedAt:     timestamppb.New(p.CreatedAt),
		UpdatedAt:     timestamppb.New(p.UpdatedAt),
		DeletedAt:     timestamppb.New(*p.DeletedAt),
	}
}
