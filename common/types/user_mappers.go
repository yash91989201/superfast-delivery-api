package types

import (
	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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

func ToProfile(p *pb.Profile) *Profile {
	return &Profile{
		ID:          p.Id,
		Name:        p.Name,
		ImageUrl:    p.ImageUrl,
		Dob:         ToDate(p.Dob).ToTime(),
		Anniversary: ToDate(p.Anniversary).ToTime(),
		Gender:      ToGenderPtr(p.Gender),
		AuthID:      p.AuthId,
		CreatedAt:   ToTime(p.CreatedAt),
		UpdatedAt:   ToTime(p.UpdatedAt),
	}
}

func ToPbProfile(p *Profile) *pb.Profile {
	return &pb.Profile{
		Id:          p.ID,
		Name:        p.Name,
		ImageUrl:    p.ImageUrl,
		Dob:         TimeToPbDate(p.Dob),
		Anniversary: &pb.Date{},
		Gender:      ToPbGenderPtr(p.Gender),
		AuthId:      p.AuthID,
		CreatedAt:   timestamppb.New(p.CreatedAt),
		UpdatedAt:   timestamppb.New(p.UpdatedAt),
	}
}

func ToAddressAlias(alias pb.AddressAlias) AddressAlias {
	switch alias {
	case pb.AddressAlias_HOME:
		return Home
	case pb.AddressAlias_WORK:
		return Work
	case pb.AddressAlias_HOTEL:
		return Hotel
	case pb.AddressAlias_OTHER:
		return Other
	default:
		return Other
	}
}

func ToPbAddressAlias(alias AddressAlias) pb.AddressAlias {
	switch alias {
	case Home:
		return pb.AddressAlias_HOME
	case Work:
		return pb.AddressAlias_WORK
	case Hotel:
		return pb.AddressAlias_HOTEL
	case Other:
		return pb.AddressAlias_OTHER
	default:
		return pb.AddressAlias_OTHER
	}
}

func ToDeliveryAddress(d *pb.DeliveryAddress) *DeliveryAddress {
	return &DeliveryAddress{
		ID:                  d.Id,
		ReceiverName:        d.ReceiverName,
		ReceiverPhone:       d.ReceiverPhone,
		AddressAlias:        ToAddressAlias(d.AddressAlias),
		OtherAlias:          d.OtherAlias,
		Latitude:            d.Latitude,
		Longitude:           d.Longitude,
		Address:             d.Address,
		NearbyLandmark:      d.NearbyLandmark,
		DeliveryInstruction: d.DeliveryInstruction,
		IsDefault:           d.IsDefault,
		AuthId:              d.AuthId,
		CreatedAt:           d.CreatedAt.AsTime(),
		UpdatedAt:           d.UpdatedAt.AsTime(),
	}
}

func ToPbDeliveryAddress(d *DeliveryAddress) *pb.DeliveryAddress {
	return &pb.DeliveryAddress{
		Id:                  d.ID,
		ReceiverName:        d.ReceiverName,
		ReceiverPhone:       d.ReceiverPhone,
		AddressAlias:        ToPbAddressAlias(d.AddressAlias),
		OtherAlias:          d.OtherAlias,
		Latitude:            d.Latitude,
		Longitude:           d.Longitude,
		Address:             d.Address,
		NearbyLandmark:      d.NearbyLandmark,
		DeliveryInstruction: d.DeliveryInstruction,
		IsDefault:           d.IsDefault,
		AuthId:              d.AuthId,
		CreatedAt:           timestamppb.New(d.CreatedAt),
		UpdatedAt:           timestamppb.New(d.UpdatedAt),
	}
}

func ToCreateDeliveryAddress(req *pb.CreateDeliveryAddressReq) *CreateDeliveryAddress {
	return &CreateDeliveryAddress{
		ReceiverName:        req.ReceiverName,
		ReceiverPhone:       req.ReceiverPhone,
		AddressAlias:        ToAddressAlias(req.AddressAlias),
		OtherAlias:          req.OtherAlias,
		Latitude:            req.Latitude,
		Longitude:           req.Longitude,
		Address:             req.Address,
		NearbyLandmark:      req.NearbyLandmark,
		DeliveryInstruction: req.DeliveryInstruction,
		IsDefault:           req.IsDefault,
		AuthId:              req.AuthId,
	}
}

func ToPbDeliveryAddressList(addresses []*DeliveryAddress) []*pb.DeliveryAddress {
	pbAddresses := make([]*pb.DeliveryAddress, len(addresses))
	for i, addr := range addresses {
		pbAddresses[i] = ToPbDeliveryAddress(addr)
	}
	return pbAddresses
}

func ToPbAddressDetail(d *AddressDetail) *pb.AddressDetail {
	if d == nil {
		return nil
	}

	return &pb.AddressDetail{
		Id:               d.Id,
		Route:            d.Route,
		Town:             d.Town,
		PostalCode:       d.PostalCode,
		District:         d.District,
		State:            d.State,
		Country:          d.Country,
		PlusCode:         d.PlusCode,
		PlaceId:          d.PlaceId,
		FormattedAddress: d.FormattedAddress,
		Latitude:         d.Latitude,
		Longitude:        d.Longitude,
		AddressId:        d.AddressId,
	}
}

func PbUpdateProfileReqToProfile(p *pb.UpdateProfileReq) *Profile {
	if p == nil {
		return nil
	}

	profile := &Profile{
		ID:     p.Id,
		AuthID: p.AuthId,
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
