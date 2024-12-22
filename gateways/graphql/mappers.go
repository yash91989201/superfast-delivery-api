package graphql

import (
	"time"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPbAuthRole(a AuthRole) pb.AuthRole {
	switch a {
	case AuthRoleCustomer:
		return pb.AuthRole_CUSTOMER
	case AuthRoleDeliveryPartner:
		return pb.AuthRole_DELIVERY_PARTNER
	case AuthRoleVendor:
		return pb.AuthRole_VENDOR
	case AuthRoleAdmin:
		return pb.AuthRole_ADMIN
	default:
		return pb.AuthRole_CUSTOMER
	}
}

func ToPbGenderPtr(g *Gender) *pb.Gender {
	if g == nil {
		return nil
	}

	switch *g {
	case GenderMale:
		return pbGenderPtr(pb.Gender_MALE)
	case GenderFemale:
		return pbGenderPtr(pb.Gender_FEMALE)
	case GenderOthers:
		return pbGenderPtr(pb.Gender_OTHERS)
	case GenderUndisclosed:
		return pbGenderPtr(pb.Gender_UNDISCLOSED)
	default:
		return nil
	}

}

func pbGenderPtr(g pb.Gender) *pb.Gender {
	return &g
}

func ToPbDate(dateStr *string) *pb.Date {
	if dateStr == nil || *dateStr == "" {
		return nil
	}

	// Parse the date string (ISO 8601 format)
	parsedTime, err := time.Parse("2006-01-02", *dateStr)
	if err != nil {
		// Handle parse error
		return nil
	}

	// Extract year, month, and day
	return &pb.Date{
		Year:  int32(parsedTime.Year()),
		Month: int32(parsedTime.Month()),
		Day:   int32(parsedTime.Day()),
	}
}

func ToPbTime(dateTimeStr *string) *timestamppb.Timestamp {
	if dateTimeStr == nil || *dateTimeStr == "" {
		return nil
	}

	parsedTime, err := time.Parse(time.RFC3339, *dateTimeStr)
	if err != nil {
		return nil
	}

	return timestamppb.New(parsedTime)
}

func ToPbAuth(a *Auth) *pb.Auth {
	return &pb.Auth{
		Id:            a.ID,
		Email:         a.Email,
		EmailVerified: *a.EmailVerified,
		Phone:         a.Phone,
		Role:          ToPbAuthRole(a.Role),
		CreatedAt:     ToPbTime(&a.CreatedAt),
		UpdatedAt:     ToPbTime(&a.UpdatedAt),
		DeletedAt:     ToPbTime(a.DeletedAt),
	}
}

func ToPbProfile(p *Profile) *pb.Profile {
	return &pb.Profile{
		Id:          p.ID,
		Name:        p.Name,
		ImageUrl:    p.ImageURL,
		Dob:         ToPbDate(p.Dob),
		Anniversary: ToPbDate(p.Anniversary),
		Gender:      ToPbGenderPtr(p.Gender),
		AuthId:      p.AuthID,
		CreatedAt:   ToPbTime(&p.CreatedAt),
		UpdatedAt:   ToPbTime(&p.UpdatedAt),
	}
}

func ToAuthRole(t pb.AuthRole) AuthRole {
	switch t {
	case pb.AuthRole_CUSTOMER:
		return AuthRoleCustomer
	case pb.AuthRole_DELIVERY_PARTNER:
		return AuthRoleDeliveryPartner
	case pb.AuthRole_VENDOR:
		return AuthRoleVendor
	case pb.AuthRole_ADMIN:
		return AuthRoleAdmin
	default:
		return AuthRoleCustomer
	}
}

func ToGenderPtr(pbGender *pb.Gender) *Gender {
	if pbGender == nil {
		return nil
	}

	g := ToGender(*pbGender)
	return &g
}

func ToGender(pbGender pb.Gender) Gender {
	switch pbGender {
	case pb.Gender_MALE:
		return GenderMale
	case pb.Gender_FEMALE:
		return GenderFemale
	case pb.Gender_OTHERS:
		return GenderOthers
	case pb.Gender_UNDISCLOSED:
		return GenderUndisclosed
	default:
		return GenderUndisclosed
	}
}

func ToDate(pbDate *pb.Date) *string {
	if pbDate == nil {
		return nil
	}

	dateStr := time.Date(
		int(pbDate.Year),
		time.Month(pbDate.Month),
		int(pbDate.Day),
		0, 0, 0, 0, time.UTC,
	).Format("2006-01-02")

	return &dateStr
}

func ToTime(pbTime *timestamppb.Timestamp) *string {
	if pbTime == nil || pbTime.AsTime().IsZero() {
		return nil
	}
	timeStr := pbTime.AsTime().Format(time.RFC3339)
	return &timeStr
}

func ToAuth(a *pb.Auth) *Auth {
	if a == nil {
		return nil
	}

	return &Auth{
		ID:            a.Id,
		Email:         a.Email,
		EmailVerified: &a.EmailVerified,
		Phone:         a.Phone,
		Role:          ToAuthRole(a.Role),
		CreatedAt:     *ToTime(a.CreatedAt),
		UpdatedAt:     *ToTime(a.UpdatedAt),
		DeletedAt:     ToTime(a.DeletedAt),
	}
}

func ToProfile(p *pb.Profile) *Profile {
	if p == nil {
		return nil
	}

	return &Profile{
		ID:          p.Id,
		Name:        p.Name,
		ImageURL:    p.ImageUrl,
		Dob:         ToDate(p.Dob),
		Anniversary: ToDate(p.Anniversary),
		Gender:      ToGenderPtr(p.Gender),
		AuthID:      p.AuthId,
		CreatedAt:   *ToTime(p.CreatedAt),
		UpdatedAt:   *ToTime(p.UpdatedAt),
	}
}
