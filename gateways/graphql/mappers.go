package graphql

import (
	"fmt"
	"time"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"github.com/yash91989201/superfast-delivery-api/common/types"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPbOrderBy(o *OrderBy) *pb.OrderBy {
	if o == nil {
		return nil
	}
	switch *o {
	case OrderByAsc:
		orderByAsc := pb.OrderBy_ASC
		return &orderByAsc
	case OrderByDesc:
		orderByDesc := pb.OrderBy_DESC
		return &orderByDesc
	default:
		panic("unexpected graphql.OrderBy")
	}
}

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
	parsedTime, err := time.Parse("2006-01-02", *dateStr)
	if err != nil {
		return nil
	}
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

func ToGQAuthRole(t pb.AuthRole) AuthRole {
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

func TypesToGQAuthRole(role types.AuthRole) AuthRole {
	switch role {
	case types.Customer:
		return AuthRoleCustomer
	case types.DeliveryPartner:
		return AuthRoleDeliveryPartner
	case types.Vendor:
		return AuthRoleVendor
	case types.Admin:
		return AuthRoleAdmin
	default:
		return ""
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
	if pbDate == nil || (pbDate.Year == 0 && pbDate.Month == 0 && pbDate.Day == 0) {
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
	return pbTime.AsTime().Format(time.RFC3339)
}

func ToGQTimePtr(pbTime *timestamppb.Timestamp) *string {
	if pbTime == nil || pbTime.AsTime().IsZero() {
		return nil
	}
	timeStr := pbTime.AsTime().Format(time.RFC3339)
	return &timeStr
}

func ToGQAuth(a *pb.Auth) *Auth {
	if a == nil {
		return nil
	}
	return &Auth{
		ID:            a.Id,
		Email:         a.Email,
		EmailVerified: a.EmailVerified,
		Phone:         a.Phone,
		AuthRole:      ToGQAuthRole(a.AuthRole),
		CreatedAt:     ToGQTime(a.CreatedAt),
		UpdatedAt:     ToGQTime(a.UpdatedAt),
		DeletedAt:     ToGQTimePtr(a.DeletedAt),
	}
}

func ToGQSession(s *pb.Session) *Session {
	return &Session{
		AccessToken:  s.AccessToken,
		RefreshToken: s.RefreshToken,
	}
}

func ToGQProfile(p *pb.Profile) *Profile {
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

func ToGQAddressAlias(a pb.AddressAlias) AddressAlias {
	switch a {
	case pb.AddressAlias_HOME:
		return AddressAliasHome
	case pb.AddressAlias_HOTEL:
		return AddressAliasHotel
	case pb.AddressAlias_OTHER:
		return AddressAliasOther
	case pb.AddressAlias_WORK:
		return AddressAliasWork
	default:
		return AddressAliasOther
	}
}

func ToGQDeliveryAddress(d *pb.DeliveryAddress) *DeliveryAddress {
	return &DeliveryAddress{
		ID:                  d.Id,
		ReceiverName:        d.ReceiverName,
		ReceiverPhone:       d.ReceiverPhone,
		AddressAlias:        ToGQAddressAlias(d.AddressAlias),
		OtherAlias:          d.OtherAlias,
		Latitude:            d.Latitude,
		Longitude:           d.Longitude,
		Address:             d.Address,
		NearbyLandmark:      d.NearbyLandmark,
		DeliveryInstruction: d.DeliveryInstruction,
		IsDefault:           d.IsDefault,
		AuthID:              d.AuthId,
		CreatedAt:           ToGQTime(d.CreatedAt),
		UpdatedAt:           ToGQTime(d.UpdatedAt),
	}
}

func ToPbAddressAlias(a AddressAlias) pb.AddressAlias {
	switch a {
	case AddressAliasHome:
		return pb.AddressAlias_HOME
	case AddressAliasHotel:
		return pb.AddressAlias_HOTEL
	case AddressAliasWork:
		return pb.AddressAlias_WORK
	case AddressAliasOther:
		return pb.AddressAlias_OTHER
	default:
		return pb.AddressAlias_OTHER
	}
}

func ToPbDeliveryAddress(d *DeliveryAddress) *pb.DeliveryAddress {
	if d == nil {
		return nil
	}
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
		AuthId:              d.AuthID,
		CreatedAt:           ToPbTime(&d.CreatedAt),
		UpdatedAt:           ToPbTime(&d.UpdatedAt),
	}
}

func ToPbCreateDeliveryAddress(in CreateDeliveryAddressInput) *pb.CreateDeliveryAddressReq {
	return &pb.CreateDeliveryAddressReq{
		ReceiverName:        in.ReceiverName,
		ReceiverPhone:       in.ReceiverPhone,
		AddressAlias:        ToPbAddressAlias(in.AddressAlias),
		OtherAlias:          in.OtherAlias,
		Latitude:            in.Latitude,
		Longitude:           in.Longitude,
		Address:             in.Address,
		NearbyLandmark:      in.NearbyLandmark,
		DeliveryInstruction: in.DeliveryInstruction,
		IsDefault:           in.IsDefault,
		AuthId:              in.AuthID,
	}
}

func ToGQDeliveryAddressList(pbAddresses []*pb.DeliveryAddress) []*DeliveryAddress {
	addresses := make([]*DeliveryAddress, len(pbAddresses))
	for i, d := range pbAddresses {
		addresses[i] = ToGQDeliveryAddress(d)
	}
	return addresses
}

func ToGQAddressDetail(d *pb.AddressDetail) *AddressDetail {
	if d == nil {
		return nil
	}
	return &AddressDetail{
		ID:               d.Id,
		Route:            d.Route,
		Town:             d.Town,
		PostalCode:       d.PostalCode,
		District:         d.District,
		State:            d.State,
		Country:          d.Country,
		PlusCode:         d.PlusCode,
		PlaceID:          d.PlaceId,
		FormattedAddress: d.FormattedAddress,
		Latitude:         d.Latitude,
		Longitude:        d.Longitude,
		AddressID:        d.AddressId,
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
		Name:        cs.Name,
		ShopType:    ToPbShopType(cs.ShopType),
		ShopStatus:  ToPbShopStatus(cs.ShopStatus),
		OwnerAuthId: cs.OwnerAuthID,
		Address: &pb.CreateShopAddress{
			Longitude:      cs.Address.Longitude,
			Latitude:       cs.Address.Latitude,
			Address:        cs.Address.Address,
			NearbyLandmark: cs.Address.NearbyLandmark,
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
		Longitude:      a.Longitude,
		Latitude:       a.Latitude,
		Address:        a.Address,
		NearbyLandmark: a.NearbyLandmark,
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
		ID:          s.Id,
		Name:        s.Name,
		ShopType:    ToGQShopType(s.ShopType),
		ShopStatus:  ToGQShopStatus(s.ShopStatus),
		OwnerAuthID: s.OwnerAuthId,
		CreatedAt:   ToGQTime(s.CreatedAt),
		UpdatedAt:   ToGQTime(s.UpdatedAt),
		DeletedAt:   ToGQTimePtr(s.DeletedAt),
		Contact:     ToGQContact(s.Contact),
		Address:     ToGQAddress(s.Address),
		Timings:     ToGQTimings(s.Timings),
		Images:      ToGQImages(s.Images),
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

	if r.OrderBy != nil {
		req.OrderBy = ToPbOrderBy(r.OrderBy)
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

func ToPbCreateRestaurantMenuReq(r *CreateRestaurantMenuInput) *pb.CreateRestaurantMenuReq {
	if r == nil {
		return nil
	}

	return &pb.CreateRestaurantMenuReq{
		MenuName: r.MenuName,
		ImageUrl: r.ImageURL,
		ShopId:   r.ShopID,
	}
}

func ToPbCreateMenuItemReq(r *CreateMenuItemInput) *pb.CreateMenuItemReq {
	if r == nil {
		return nil
	}

	return &pb.CreateMenuItemReq{
		Name:        r.Name,
		ImageUrl:    r.ImageURL,
		Description: r.Description,
		Price:       r.Price,
		MenuId:      r.MenuID,
	}
}

func ToPbCreateItemVariantReq(r *CreateItemVariantInput) *pb.CreateItemVariantReq {
	if r == nil {
		return nil
	}

	return &pb.CreateItemVariantReq{
		VariantName:     r.VariantName,
		RelativePrice:   r.RelativePrice,
		RelativePricing: r.RelativePricing,
		Price:           r.Price,
		ImageUrl:        r.ImageURL,
		Description:     r.Description,
		ItemId:          r.ItemID,
	}
}

func ToPbCreateItemAddonReq(r *CreateItemAddonInput) *pb.CreateItemAddonReq {
	if r == nil {
		return nil
	}

	return &pb.CreateItemAddonReq{
		AddonName:   r.AddonName,
		AddonPrice:  r.AddonPrice,
		ImageUrl:    r.ImageURL,
		Description: r.Description,
		ItemId:      r.ItemID,
	}
}

func ToGQRestaurantMenuList(m []*pb.RestaurantMenu) []*RestaurantMenu {
	list := make([]*RestaurantMenu, len(m))
	for i, rm := range m {
		list[i] = ToGQRestaurantMenu(rm)
	}
	return list
}

func ToGQRestaurantMenu(m *pb.RestaurantMenu) *RestaurantMenu {
	return &RestaurantMenu{
		ID:        m.Id,
		MenuName:  m.MenuName,
		ImageURL:  m.ImageUrl,
		ShopID:    m.ShopId,
		CreatedAt: ToGQTime(m.CreatedAt),
		UpdatedAt: ToGQTime(m.UpdatedAt),
	}
}

func ToGQMenuItems(m []*pb.MenuItem) []*MenuItem {
	menuItems := make([]*MenuItem, len(m))
	for i, mi := range m {
		menuItems[i] = ToGQMenuItem(mi)
	}
	return menuItems
}

func ToGQItemVariantList(v []*pb.ItemVariant) []*ItemVariant {
	variants := make([]*ItemVariant, len(v))
	for i, vars := range v {
		variants[i] = ToGQItemVariant(vars)
	}
	return variants
}

func ToGQItemAddonList(a []*pb.ItemAddon) []*ItemAddon {
	addons := make([]*ItemAddon, len(a))
	for i, addon := range a {
		addons[i] = ToGQItemAddon(addon)
	}
	return addons
}

func ToGQItemVariant(v *pb.ItemVariant) *ItemVariant {
	return &ItemVariant{
		ID:              v.Id,
		VariantName:     v.VariantName,
		RelativePrice:   float64(v.RelativePrice),
		RelativePricing: v.RelativePricing,
		Price:           float64(v.Price),
		ImageURL:        v.ImageUrl,
		Description:     v.Description,
		ItemID:          v.ItemId,
	}
}

func ToGQItemAddon(a *pb.ItemAddon) *ItemAddon {
	return &ItemAddon{
		ID:          a.Id,
		AddonName:   a.AddonName,
		AddonPrice:  float64(a.AddonPrice),
		ImageURL:    a.ImageUrl,
		Description: a.Description,
		ItemID:      a.ItemId,
	}
}

func ToGQRetailCategory(c *pb.RetailCategory) *RetailCategory {
	return &RetailCategory{
		ID:           c.Id,
		CategoryName: c.CategoryName,
		ImageURL:     c.ImageUrl,
		ShopID:       c.ShopId,
		CreatedAt:    ToGQTime(c.CreatedAt),
		UpdatedAt:    ToGQTime(c.UpdatedAt),
	}
}

func ToGQRetailItem(i *pb.RetailItem) *RetailItem {
	return &RetailItem{
		ID:          i.Id,
		Name:        i.Name,
		ImageURL:    i.ImageUrl,
		Description: i.Description,
		Price:       i.Price,
		CategoryID:  i.CategoryId,
		Variants:    ToGQItemVariantList(i.Variants),
		CreatedAt:   ToGQTime(i.CreatedAt),
		UpdatedAt:   ToGQTime(i.UpdatedAt),
	}
}

func ToGQRetailItems(il []*pb.RetailItem) []*RetailItem {
	list := make([]*RetailItem, len(il))
	for i, item := range il {
		if item != nil {
			list[i] = ToGQRetailItem(item)
		}
	}
	return list
}

func ToGQRetailCategoryList(cl []*pb.RetailCategory) []*RetailCategory {
	list := make([]*RetailCategory, len(cl))
	for i, item := range cl {
		if item != nil {
			list[i] = ToGQRetailCategory(item)
		}
	}
	return list
}

func ToGQMedicineCategory(c *pb.MedicineCategory) *MedicineCategory {
	return &MedicineCategory{
		ID:           c.Id,
		CategoryName: c.CategoryName,
		ImageURL:     c.ImageUrl,
		ShopID:       c.ShopId,
		CreatedAt:    ToGQTime(c.CreatedAt),
		UpdatedAt:    ToGQTime(c.UpdatedAt),
	}
}

func ToGQMedicineItem(i *pb.MedicineItem) *MedicineItem {
	return &MedicineItem{
		ID:          i.Id,
		Name:        i.Name,
		Price:       i.Price,
		ImageURL:    i.ImageUrl,
		Description: i.Description,
		CategoryID:  i.CategoryId,
		CreatedAt:   ToGQTime(i.CreatedAt),
		UpdatedAt:   ToGQTime(i.UpdatedAt),
	}
}

func ToGQMenuItem(item *pb.MenuItem) *MenuItem {
	return &MenuItem{
		ID:          item.Id,
		Name:        item.Name,
		Price:       item.Price,
		ImageURL:    item.ImageUrl,
		Description: item.Description,
		Variants:    ToGQItemVariantList(item.Variants),
		Addons:      ToGQItemAddonList(item.Addons),
		MenuID:      item.MenuId,
		CreatedAt:   item.CreatedAt.String(),
		UpdatedAt:   item.UpdatedAt.String(),
	}
}

func ToGQMenuItemList(items []*pb.MenuItem) []*MenuItem {
	var result []*MenuItem
	for _, item := range items {
		result = append(result, ToGQMenuItem(item))
	}
	return result
}

func ToGQRetailItemList(items []*pb.RetailItem) []*RetailItem {
	var result []*RetailItem
	for _, item := range items {
		result = append(result, ToGQRetailItem(item))
	}
	return result
}

func ToGQMedicineItemList(items []*pb.MedicineItem) []*MedicineItem {
	var result []*MedicineItem
	for _, item := range items {
		result = append(result, ToGQMedicineItem(item))
	}
	return result
}

func ToGQMedicineItems(il []*pb.MedicineItem) []*MedicineItem {
	list := make([]*MedicineItem, len(il))
	for i, item := range il {
		if item != nil {
			list[i] = ToGQMedicineItem(item)
		}
	}
	return list
}

func ToGQMedicineCategoryList(cl []*pb.MedicineCategory) []*MedicineCategory {
	list := make([]*MedicineCategory, len(cl))
	for i, item := range cl {
		if item != nil {
			list[i] = ToGQMedicineCategory(item)
		}
	}
	return list
}

func ToPbCreateItemStockReq(r *CreateItemStockInput) *pb.CreateItemStockReq {
	return &pb.CreateItemStockReq{
		ItemId:   r.ItemID,
		Quantity: r.Quantity,
	}
}

func ToPbCreateVariantStockReq(r *CreateVariantStockInput) *pb.CreateVariantStockReq {
	return &pb.CreateVariantStockReq{
		VariantId: r.VariantID,
		Quantity:  r.Quantity,
	}
}

func ToPbCreateAddonStockReq(r *CreateAddonStockInput) *pb.CreateAddonStockReq {
	return &pb.CreateAddonStockReq{
		AddonId:  r.AddonID,
		Quantity: r.Quantity,
	}
}

func ToGQItemStock(r *pb.ItemStock) *ItemStock {
	return &ItemStock{
		ID:         r.Id,
		ItemID:     r.ItemId,
		Quantity:   r.Quantity,
		RestockQty: r.RestockQty,
		UpdatedAt:  ToGQTime(r.UpdatedAt),
	}
}

func ToGQVariantStock(r *pb.VariantStock) *VariantStock {
	return &VariantStock{
		ID:         r.Id,
		VariantID:  r.VariantId,
		Quantity:   r.Quantity,
		RestockQty: r.RestockQty,
		UpdatedAt:  ToGQTime(r.UpdatedAt),
	}
}

func ToGQAddonStock(r *pb.AddonStock) *AddonStock {
	return &AddonStock{
		ID:         r.Id,
		AddonID:    r.AddonId,
		Quantity:   r.Quantity,
		RestockQty: r.RestockQty,
		UpdatedAt:  ToGQTime(r.UpdatedAt),
	}
}

func ToPbUpdateItemStockReq(in *UpdateItemStockInput) *pb.UpdateItemStockReq {
	return &pb.UpdateItemStockReq{
		Id:         in.ID,
		Quantity:   in.Quantity,
		RestockQty: in.RestockQty,
	}
}

func ToPbUpdateVariantStockReq(in *UpdateVariantStockInput) *pb.UpdateVariantStockReq {
	return &pb.UpdateVariantStockReq{
		Id:         in.ID,
		Quantity:   in.Quantity,
		RestockQty: in.RestockQty,
	}
}

func ToPbUpdateAddonStockReq(in *UpdateAddonStockInput) *pb.UpdateAddonStockReq {
	return &pb.UpdateAddonStockReq{
		Id:         in.ID,
		Quantity:   in.Quantity,
		RestockQty: in.RestockQty,
	}
}

func ToPbCreateRetailCategoryReq(in *CreateRetailCategoryInput) *pb.CreateRetailCategoryReq {
	return &pb.CreateRetailCategoryReq{
		CategoryName: in.CategoryName,
		ImageUrl:     in.ImageURL,
		ShopId:       in.ShopID,
	}
}

func ToPbCreateRetailItemReq(in *CreateRetailItemInput) *pb.CreateRetailItemReq {
	return &pb.CreateRetailItemReq{
		Name:        in.Name,
		Price:       in.Price,
		ImageUrl:    in.ImageURL,
		Description: in.Description,
		CategoryId:  in.CategoryID,
	}
}

func ToPbCreateMedicineCategoryReq(in *CreateMedicineCategoryInput) *pb.CreateMedicineCategoryReq {
	return &pb.CreateMedicineCategoryReq{
		CategoryName: in.CategoryName,
		ImageUrl:     in.ImageURL,
		ShopId:       in.ShopID,
	}
}

func ToPbCreateMedicineItemReq(in *CreateMedicineItemInput) *pb.CreateMedicineItemReq {
	return &pb.CreateMedicineItemReq{
		Name:        in.Name,
		Price:       in.Price,
		ImageUrl:    in.ImageURL,
		Description: in.Description,
		CategoryId:  in.CategoryID,
	}
}
