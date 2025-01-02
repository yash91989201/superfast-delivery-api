package graphql

import (
	"fmt"
	"time"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
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

func ToPbCreateRestaurantMenuReq(r *CreateRestaurantMenuReq) *pb.CreateRestaurantMenuReq {
	if r == nil {
		return nil
	}

	return &pb.CreateRestaurantMenuReq{
		MenuName: r.MenuName,
		ShopId:   r.ShopID,
	}
}

func ToPbCreateMenuItemReq(r *CreateMenuItemReq) *pb.CreateMenuItemReq {
	if r == nil {
		return nil
	}

	return &pb.CreateMenuItemReq{
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
		MenuId:      r.MenuID,
	}
}

func ToPbCreateItemVariantReq(r *CreateItemVariantReq) *pb.CreateItemVariantReq {
	if r == nil {
		return nil
	}

	return &pb.CreateItemVariantReq{
		VariantName:     r.VariantName,
		RelativePrice:   r.RelativePrice,
		RelativePricing: r.RelativePricing,
		Price:           r.Price,
		Description:     r.Description,
	}
}

func ToPbCreateItemAddonReq(r *CreateItemAddonReq) *pb.CreateItemAddonReq {
	if r == nil {
		return nil
	}

	return &pb.CreateItemAddonReq{
		AddonName:   r.AddonName,
		AddonPrice:  r.AddonPrice,
		Description: r.Description,
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
		ShopID:    m.ShopId,
		MenuItems: ToGQMenuItems(m.MenuItems),
		CreatedAt: ToGQTime(m.CreatedAt),
		UpdatedAt: ToGQTime(m.UpdatedAt),
		DeletedAt: ToGQTimePtr(m.DeletedAt),
	}
}

func ToGQMenuItems(m []*pb.MenuItem) []*MenuItem {
	menuItems := make([]*MenuItem, len(m))
	for i, mi := range m {
		menuItems[i] = ToGQMenuItem(mi)
	}

	return menuItems
}

func ToGQMenuItem(m *pb.MenuItem) *MenuItem {

	return &MenuItem{
		ID:          m.Id,
		Name:        m.Name,
		Description: &m.Description,
		Price:       float64(m.Price),
		MenuID:      m.MenuId,
		Variants:    ToGQItemVariants(m.Variants),
		Addons:      ToGQItemAddons(m.Addons),
		CreatedAt:   ToGQTime(m.CreatedAt),
		UpdatedAt:   ToGQTime(m.UpdatedAt),
		DeletedAt:   ToGQTimePtr(m.DeletedAt),
	}
}

func ToGQItemVariants(v []*pb.ItemVariant) []*ItemVariant {
	variants := make([]*ItemVariant, len(v))
	for i, vars := range v {
		variants[i] = ToGQItemVariant(vars)
	}

	return variants
}

func ToGQItemAddons(a []*pb.ItemAddon) []*ItemAddon {
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
		Description:     v.Description,
		ItemID:          v.ItemId,
	}
}

func ToGQItemAddon(a *pb.ItemAddon) *ItemAddon {

	return &ItemAddon{
		ID:          a.Id,
		AddonName:   a.AddonName,
		AddonPrice:  float64(a.AddonPrice),
		Description: a.Description,
		ItemID:      a.ItemId,
	}
}

func ToGQRetailCategory(c *pb.RetailCategory) *RetailCategory {
	return &RetailCategory{
		ID:           c.Id,
		CategoryName: c.CategoryName,
		ShopID:       c.ShopId,
		RetailItems:  ToGQRetailItems(c.RetailItems),
		CreatedAt:    ToGQTime(c.CreatedAt),
		UpdatedAt:    ToGQTime(c.UpdatedAt),
		DeletedAt:    ToGQTimePtr(c.DeletedAt),
	}
}

func ToGQRetailItem(i *pb.RetailItem) *RetailItem {
	return &RetailItem{
		ID:          i.Id,
		Name:        i.Name,
		Description: i.Description,
		Price:       i.Price,
		CategoryID:  i.CategoryId,
		Variants:    ToGQItemVariants(i.Variants),
		Addons:      ToGQItemAddons(i.Addons),
		CreatedAt:   ToGQTime(i.CreatedAt),
		UpdatedAt:   ToGQTime(i.UpdatedAt),
		DeletedAt:   ToGQTimePtr(i.DeletedAt),
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
		ID:            c.Id,
		CategoryName:  c.CategoryName,
		ShopID:        c.ShopId,
		MedicineItems: ToGQMedicineItems(c.MedicineItems),
		CreatedAt:     ToGQTime(c.CreatedAt),
		UpdatedAt:     ToGQTime(c.UpdatedAt),
		DeletedAt:     ToGQTimePtr(c.DeletedAt),
	}
}

func ToGQMedicineItem(i *pb.MedicineItem) *MedicineItem {
	return &MedicineItem{
		ID:          i.Id,
		Name:        i.Name,
		Price:       i.Price,
		Description: i.Description,
		CategoryID:  i.CategoryId,
		CreatedAt:   ToGQTime(i.CreatedAt),
		UpdatedAt:   ToGQTime(i.UpdatedAt),
		DeletedAt:   ToGQTimePtr(i.DeletedAt),
	}
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
