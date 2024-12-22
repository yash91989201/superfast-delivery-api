package types

import (
	"time"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToBoolPtr(b bool) *bool {
	return &b
}

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

func ToPbTimestamp(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
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

func PbDateToTime(d *pb.Date) *time.Time {
	if d == nil {
		return nil
	}
	t := time.Date(
		int(d.Year),
		time.Month(d.Month),
		int(d.Day),
		0, 0, 0, 0,
		time.UTC,
	)
	return &t
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
		Role:          ToAuthRole(a.Role),
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

func PbUpdateProfileReqToProfile(p *pb.UpdateProfileReq) *Profile {
	if p == nil {
		return nil
	}

	profile := &Profile{
		AuthId: p.AuthId,
	}

	if p.Name != nil {
		profile.Name = *p.Name
	}

	if p.ImageUrl != nil {
		profile.ImageUrl = p.ImageUrl
	}

	if p.Dob != nil {
		profile.Dob = PbDateToTime(p.Dob)
	}

	if p.Anniversary != nil {
		profile.Anniversary = PbDateToTime(p.Anniversary)
	}

	if p.Gender != nil {
		profile.Gender = ToGenderPtr(p.Gender)
	}

	return profile
}

func ToPbAuth(p *Auth) *pb.Auth {
	var deletedAt *timestamppb.Timestamp
	if p.DeletedAt != nil {
		deletedAt = ToPbTimestamp(*p.DeletedAt)
	}

	return &pb.Auth{
		Id:            p.Id,
		Email:         p.Email,
		EmailVerified: p.EmailVerified,
		Phone:         p.Phone,
		Role:          ToPbAuthRole(p.Role),
		CreatedAt:     ToPbTimestamp(p.CreatedAt),
		UpdatedAt:     ToPbTimestamp(p.UpdatedAt),
		DeletedAt:     deletedAt,
	}
}
