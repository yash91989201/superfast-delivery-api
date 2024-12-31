package graphql

import (
	"fmt"
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
		EmailVerified: a.EmailVerified,
		Phone:         a.Phone,
		AuthRole:      ToPbAuthRole(a.AuthRole),
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

func ToGQShopType(t pb.ShopType) ShopType {
	switch t {

	case pb.ShopType_GROCERY:
		return ShopTypeGrocery
	case pb.ShopType_PHARMACEUTICAL:
		return ShopTypePharmaceutical
	case pb.ShopType_RESTAURANT:
		return ShopTypeRestaurant
	default:
		panic(fmt.Sprintf("unexpected pb.ShopType: %#v", t))
	}
}

func ToGQShopStatus(s pb.ShopStatus) ShopStatus {
	switch s {

	case pb.ShopStatus_CLOSED:
		return ShopStatusClosed
	case pb.ShopStatus_OPEN:
		return ShopStatusOpen
	default:
		panic(fmt.Sprintf("unexpected pb.ShopStatus: %#v", s))
	}
}

func ToGQDayOfWeek(d pb.DayOfWeek) DayOfWeek {
	switch d {

	case pb.DayOfWeek_FRIDAY:
		return DayOfWeekFriday
	case pb.DayOfWeek_MONDAY:
		return DayOfWeekMonday
	case pb.DayOfWeek_SATURDAY:
		return DayOfWeekSaturday
	case pb.DayOfWeek_SUNDAY:
		return DayOfWeekSunday
	case pb.DayOfWeek_THURSDAY:
		return DayOfWeekThursday
	case pb.DayOfWeek_TUESDAY:
		return DayOfWeekTuesday
	case pb.DayOfWeek_WEDNESDAY:
		return DayOfWeekWednesday
	default:
		panic(fmt.Sprintf("unexpected pb.DayOfWeek: %#v", d))
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

func ToGQTime(pbTime *timestamppb.Timestamp) string {
	if pbTime == nil || pbTime.AsTime().IsZero() {
		return time.Now().UTC().String()
	}
	timeStr := pbTime.AsTime().Format(time.RFC3339)
	return timeStr
}
func ToGQTimePtr(pbTime *timestamppb.Timestamp) *string {
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
		EmailVerified: a.EmailVerified,
		Phone:         a.Phone,
		AuthRole:      ToAuthRole(a.AuthRole),
		CreatedAt:     ToGQTime(a.CreatedAt),
		UpdatedAt:     ToGQTime(a.UpdatedAt),
		DeletedAt:     ToGQTimePtr(a.DeletedAt),
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
		CreatedAt:   ToGQTime(p.CreatedAt),
		UpdatedAt:   ToGQTime(p.UpdatedAt),
	}
}

func ToPbShopType(t ShopType) pb.ShopType {
	switch t {

	case ShopTypeGrocery:
		return pb.ShopType_GROCERY
	case ShopTypePharmaceutical:
		return pb.ShopType_PHARMACEUTICAL
	case ShopTypeRestaurant:
		return pb.ShopType_RESTAURANT
	default:
		panic(fmt.Sprintf("unexpected graphql.ShopType: %#v", t))
	}
}

func ToPbShopStatus(s ShopStatus) pb.ShopStatus {
	switch s {

	case ShopStatusClosed:
		return pb.ShopStatus_CLOSED
	case ShopStatusOpen:
		return pb.ShopStatus_OPEN
	default:
		panic(fmt.Sprintf("unexpected graphql.ShopStatus: %#v", s))
	}
}

func ToPbDayOfWeek(d DayOfWeek) pb.DayOfWeek {
	switch d {

	case DayOfWeekFriday:
		return pb.DayOfWeek_FRIDAY
	case DayOfWeekMonday:
		return pb.DayOfWeek_MONDAY
	case DayOfWeekSaturday:
		return pb.DayOfWeek_SATURDAY
	case DayOfWeekSunday:
		return pb.DayOfWeek_SUNDAY
	case DayOfWeekThursday:
		return pb.DayOfWeek_THURSDAY
	case DayOfWeekTuesday:
		return pb.DayOfWeek_TUESDAY
	case DayOfWeekWednesday:
		return pb.DayOfWeek_WEDNESDAY
	default:
		panic(fmt.Sprintf("unexpected graphql.DayOfWeek: %#v", d))
	}
}

func ToPbCreateShopReq(cs CreateShopInput) *pb.CreateShopReq {

	images := make([]*pb.CreateShopImage, len(cs.Images))
	for i, img := range cs.Images {
		images[i] = &pb.CreateShopImage{
			ImageUrl:    img.ImageURL,
			Description: img.Description,
		}
	}

	timings := make([]*pb.CreateShopTiming, len(cs.Timings))
	for i, t := range cs.Timings {
		timings[i] = &pb.CreateShopTiming{
			Day:      ToPbDayOfWeek(t.Day),
			OpensAt:  timestamppb.New(t.OpensAt),
			ClosesAt: timestamppb.New(t.ClosesAt),
		}
	}

	return &pb.CreateShopReq{
		Name:       cs.Name,
		ShopType:   ToPbShopType(cs.ShopType),
		ShopStatus: ToPbShopStatus(cs.ShopStatus),
		OwnerId:    cs.OwnerID,
		Address: &pb.CreateShopAddress{
			Address1:       cs.Address.Address1,
			Address2:       *cs.Address.Address2,
			Longitude:      cs.Address.Longitude,
			Latitude:       cs.Address.Latitude,
			NearbyLandmark: cs.Address.NearbyLandmark,
			City:           cs.Address.City,
			State:          cs.Address.State,
			Pincode:        cs.Address.Pincode,
			Country:        cs.Address.Country,
		},
		Contact: &pb.CreateShopContact{
			Name:        cs.Contact.Name,
			PhoneNumber: cs.Contact.PhoneNumber,
			Email:       cs.Contact.Email,
		},
		Images:  images,
		Timings: timings,
	}
}

func ToGQCreateShopOutput(cs *pb.CreateShopRes) *CreateShopOutput {
	return &CreateShopOutput{
		ID:      cs.Id,
		Message: cs.Message,
	}
}

func ToGQContact(c *pb.ShopContact) *ShopContact {
	return &ShopContact{
		ID:          c.Id,
		Name:        c.Name,
		PhoneNumber: c.PhoneNumber,
		Email:       c.Email,
		ShopID:      c.ShopId,
		CreatedAt:   ToGQTime(c.CreatedAt),
	}
}

func ToGQAddress(a *pb.ShopAddress) *ShopAddress {
	return &ShopAddress{
		ID:             a.Id,
		Address1:       a.Address1,
		Address2:       a.Address2,
		Longitude:      a.Longitude,
		Latitude:       a.Latitude,
		NearbyLandmark: a.NearbyLandmark,
		City:           a.City,
		State:          a.State,
		Pincode:        a.Pincode,
		Country:        a.Country,
		ShopID:         a.ShopId,
		CreatedAt:      ToGQTime(a.CreatedAt),
	}
}

func ToGQTiming(t *pb.ShopTiming) *ShopTiming {
	return &ShopTiming{
		ID:        t.Id,
		Day:       ToGQDayOfWeek(t.Day),
		OpensAt:   t.OpensAt.AsTime(),
		ClosesAt:  t.ClosesAt.AsTime(),
		ShopID:    t.ShopId,
		CreatedAt: ToGQTime(t.CreatedAt),
		UpdatedAt: ToGQTime(t.CreatedAt),
	}
}

func ToGQTimings(t []*pb.ShopTiming) []*ShopTiming {
	timings := make([]*ShopTiming, len(t))
	for i, timing := range t {
		timings[i] = ToGQTiming(timing)
	}

	return timings
}

func ToGQImage(t *pb.ShopImage) *ShopImage {
	return &ShopImage{
		ID:          t.Id,
		ImageURL:    t.ImageUrl,
		Description: t.Description,
		ShopID:      t.ShopId,
		CreatedAt:   ToGQTime(t.CreatedAt),
		UpdatedAt:   ToGQTime(t.CreatedAt),
	}
}

func ToGQImages(i []*pb.ShopImage) []*ShopImage {
	images := make([]*ShopImage, len(i))
	for idx, image := range i {
		images[idx] = ToGQImage(image)
	}

	return images
}

func ToGQShop(s *pb.Shop) *Shop {

	return &Shop{
		ID:         s.Id,
		Name:       s.Name,
		ShopType:   ToGQShopType(s.ShopType),
		ShopStatus: ToGQShopStatus(s.ShopStatus),
		OwnerID:    s.OwnerId,
		CreatedAt:  ToGQTime(s.CreatedAt),
		UpdatedAt:  ToGQTime(s.UpdatedAt),
		DeletedAt:  ToGQTimePtr(s.DeletedAt),
		Contact:    ToGQContact(s.Contact),
		Address:    ToGQAddress(s.Address),
		Timings:    ToGQTimings(s.Timings),
		Images:     ToGQImages(s.Images),
	}
}

func ToGQShops(shopList []*pb.Shop) []*Shop {
	shops := make([]*Shop, len(shopList))
	for i, shop := range shopList {
		shops[i] = ToGQShop(shop)
	}

	return shops
}

func ToPbListShopReq(r *ListShopsInput) *pb.ListShopsReq {
	if r == nil {
		return nil
	}

	req := &pb.ListShopsReq{}

	if r.Name != nil {
		req.Name = r.Name
	}

	if r.ShopType != nil {
		shopType := ToPbShopType(*r.ShopType)
		req.ShopType = &shopType
	}

	if r.ShopStatus != nil {
		shopStatus := ToPbShopStatus(*r.ShopStatus)
		req.ShopStatus = &shopStatus
	}

	if r.Limit != nil {
		limit := int32(*r.Limit)
		req.Limit = &limit
	}

	if r.Offset != nil {
		offset := int32(*r.Offset)
		req.Offset = &offset
	}

	return req
}
