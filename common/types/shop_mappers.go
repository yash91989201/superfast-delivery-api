package types

import (
	"fmt"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
)

func ToPbShopType(t ShopType) pb.ShopType {
	switch t {

	case Grocery:
		return pb.ShopType_GROCERY
	case Pharmaceutical:
		return pb.ShopType_PHARMACEUTICAL
	case Restaurant:
		return pb.ShopType_RESTAURANT
	default:
		panic(fmt.Sprintf("unexpected types.ShopType: %#v", t))
	}
}

func ToPbShopStatus(s ShopStatus) pb.ShopStatus {
	switch s {

	case Closed:
		return pb.ShopStatus_CLOSED
	case Open:
		return pb.ShopStatus_OPEN
	default:
		panic(fmt.Sprintf("unexpected types.ShopStatus: %#v", s))
	}
}

func ToPbDayOfWeek(d DayOfWeek) pb.DayOfWeek {
	switch d {

	case Friday:
		return pb.DayOfWeek_FRIDAY
	case Monday:
		return pb.DayOfWeek_MONDAY
	case Saturday:
		return pb.DayOfWeek_SATURDAY
	case Sunday:
		return pb.DayOfWeek_SUNDAY
	case Thursday:
		return pb.DayOfWeek_THURSDAY
	case Tuesday:
		return pb.DayOfWeek_TUESDAY
	case Wednesday:
		return pb.DayOfWeek_WEDNESDAY
	default:
		panic(fmt.Sprintf("unexpected types.DayOfWeek: %#v", d))
	}
}

func ToDayOfWeek(d pb.DayOfWeek) DayOfWeek {
	switch d {

	case pb.DayOfWeek_MONDAY:
		return Monday
	case pb.DayOfWeek_TUESDAY:
		return Tuesday
	case pb.DayOfWeek_WEDNESDAY:
		return Wednesday
	case pb.DayOfWeek_THURSDAY:
		return Thursday
	case pb.DayOfWeek_FRIDAY:
		return Friday
	case pb.DayOfWeek_SATURDAY:
		return Saturday
	case pb.DayOfWeek_SUNDAY:
		return Sunday
	default:
		panic(fmt.Sprintf("unexpected pb.DayOfWeek: %#v", d))
	}
}

func ToShopType(t pb.ShopType) ShopType {
	switch t {

	case pb.ShopType_RESTAURANT:
		return Restaurant
	case pb.ShopType_GROCERY:
		return Grocery
	case pb.ShopType_PHARMACEUTICAL:
		return Pharmaceutical
	default:
		panic(fmt.Sprintf("unexpected pb.ShopType: %#v", t))
	}
}

func ToShopStatus(s pb.ShopStatus) ShopStatus {
	switch s {

	case pb.ShopStatus_OPEN:
		return Open
	case pb.ShopStatus_CLOSED:
		return Closed
	default:
		panic(fmt.Sprintf("unexpected pb.ShopStatus: %#v", s))
	}
}

func ToCreateShopImage(images []*pb.CreateShopImage) []CreateShopImage {
	if images == nil {
		return make([]CreateShopImage, 0)
	}

	imgs := make([]CreateShopImage, len(images))
	for i, img := range images {
		imgs[i] = CreateShopImage{
			ImageUrl:    img.ImageUrl,
			Description: img.Description,
		}
	}

	return imgs
}

func ToCreateShopTiming(t []*pb.CreateShopTiming) []CreateShopTiming {
	if t == nil {
		return make([]CreateShopTiming, 0)
	}

	timings := make([]CreateShopTiming, len(t))
	for i, time := range t {
		timings[i] = CreateShopTiming{
			Day:      ToDayOfWeek(time.Day),
			OpensAt:  time.OpensAt.AsTime(),
			ClosesAt: time.ClosesAt.AsTime(),
		}
	}

	return timings
}

func ToCreateShop(req *pb.CreateShopReq) *CreateShop {
	if req == nil {
		return nil
	}

	var address CreateShopAddress
	if req.Address != nil {
		address = CreateShopAddress{
			Latitude:       req.Address.Latitude,
			Longitude:      req.Address.Longitude,
			Address:        req.Address.Address,
			NearbyLandmark: req.Address.NearbyLandmark,
		}
	}

	// Convert CreateShopContact
	var contact CreateShopContact
	if req.Contact != nil {
		contact = CreateShopContact{
			Name:        req.Contact.Name,
			PhoneNumber: req.Contact.PhoneNumber,
			Email:       req.Contact.Email,
		}
	}

	return &CreateShop{
		Name:        req.Name,
		ShopType:    ToShopType(req.ShopType),
		ShopStatus:  ToShopStatus(req.ShopStatus),
		OwnerAuthId: req.OwnerAuthId,
		Address:     address,
		Contact:     contact,
		Image:       ToCreateShopImage(req.Images),
		Timing:      ToCreateShopTiming(req.Timings),
	}
}

func ToPbShopContact(c *ShopContact) *pb.ShopContact {
	if c == nil {
		return nil
	}

	return &pb.ShopContact{
		Id:          c.ID,
		Name:        c.Name,
		PhoneNumber: c.PhoneNumber,
		Email:       c.Email,
		ShopId:      c.ShopID,
		CreatedAt:   ToPbTimestamp(c.CreatedAt),
	}
}

func ToPbShopAddress(a *ShopAddress) *pb.ShopAddress {
	if a == nil {
		return nil
	}

	return &pb.ShopAddress{
		Id:             a.ID,
		Longitude:      a.Longitude,
		Latitude:       a.Latitude,
		Address:        a.Address,
		NearbyLandmark: a.NearbyLandmark,
		ShopId:         a.ShopID,
		CreatedAt:      ToPbTimestamp(a.CreatedAt),
	}
}

func ToPbShopTimmings(t []*ShopTiming) []*pb.ShopTiming {
	timings := make([]*pb.ShopTiming, len(t))
	for i, time := range t {
		timings[i] = &pb.ShopTiming{
			Id:        time.ID,
			Day:       ToPbDayOfWeek(time.Day),
			OpensAt:   ToPbTimestamp(time.OpensAt),
			ClosesAt:  ToPbTimestamp(time.ClosesAt),
			ShopId:    time.ShopID,
			CreatedAt: ToPbTimestamp(time.CreatedAt),
			UpdatedAt: ToPbTimestamp(time.UpdatedAt),
		}
	}

	return timings
}

func ToPbShopImage(t []*ShopImage) []*pb.ShopImage {
	images := make([]*pb.ShopImage, len(t))
	for i, image := range t {
		images[i] = &pb.ShopImage{
			Id:          image.ID,
			ImageUrl:    image.ImageUrl,
			Description: image.Description,
			ShopId:      image.ShopID,
			CreatedAt:   ToPbTimestamp(image.CreatedAt),
			UpdatedAt:   ToPbTimestamp(image.UpdatedAt),
		}
	}

	return images
}

func ToPbShop(shop *Shop) *pb.Shop {
	if shop == nil {
		return nil
	}

	return &pb.Shop{
		Id:          shop.ID,
		Name:        shop.Name,
		ShopType:    ToPbShopType(shop.ShopType),
		ShopStatus:  ToPbShopStatus(shop.ShopStatus),
		OwnerAuthId: shop.OwnerAuthID,
		CreatedAt:   ToPbTimestamp(shop.CreatedAt),
		UpdatedAt:   ToPbTimestamp(shop.UpdatedAt),
		DeletedAt:   TimePtrToPbTime(shop.DeletedAt),
		Contact:     ToPbShopContact(shop.Contact),
		Address:     ToPbShopAddress(shop.Address),
		Timings:     ToPbShopTimmings(shop.Timing),
		Images:      ToPbShopImage(shop.Image),
	}
}

func ToPbShops(shopList []*Shop) []*pb.Shop {
	shops := make([]*pb.Shop, len(shopList))
	for i, shop := range shopList {
		shops[i] = ToPbShop(shop)
	}

	return shops
}

func ToPbListShopRes(shopList []*Shop) *pb.ListShopsRes {

	return &pb.ListShopsRes{
		Shops: ToPbShops(shopList),
		Total: int32(len(shopList)),
	}
}

func ToOrderBy(o *pb.OrderBy) *OrderBy {
	if o == nil {
		return nil
	}

	switch *o {

	case pb.OrderBy_ASC:
		orderByAsc := Asc
		return &orderByAsc
	case pb.OrderBy_DESC:
		orderByDesc := Desc
		return &orderByDesc

	default:
		panic("unexpected pb.OrderBy")
	}
}

func ToListShopFilters(f *pb.ListShopsReq) *ListShopFilters {
	if f == nil {
		return nil
	}

	filters := &ListShopFilters{}
	if f.Name != nil {
		filters.Name = f.Name
	}

	if f.ShopStatus != nil {
		status := ToShopStatus(*f.ShopStatus)
		filters.ShopStatus = &status
	}

	if f.ShopType != nil {
		shopType := ToShopType(*f.ShopType)
		filters.ShopType = &shopType
	}

	if f.OrderBy != nil {
		filters.OrderBy = ToOrderBy(f.OrderBy)
	}

	if f.Limit != nil {
		limit := int(*f.Limit)
		filters.Limit = &limit
	}

	if f.Offset != nil {
		offset := int(*f.Limit)
		filters.Offset = &offset
	}

	return filters
}
