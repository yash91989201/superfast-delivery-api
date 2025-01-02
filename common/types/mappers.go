package types

import (
	"fmt"
	"log"
	"time"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Date struct {
	Year  int32
	Month int32
	Day   int32
}

func (d *Date) ToTime() *time.Time {
	t := time.Date(int(d.Year), time.Month(d.Month), int(d.Day), 0, 0, 0, 0, time.UTC)
	return &t
}

func ToBoolPtr(b bool) *bool {
	return &b
}

func ToStrPtr(s string) *string {
	return &s
}

func HexToObjectId(id string) bson.ObjectID {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Fatalf("Invalid object id: %v", err)
	}

	return objectID
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

func PbTimeStampToStrPtr(t *timestamppb.Timestamp) *string {
	if t == nil {
		return nil
	}

	timeStr := t.String()
	return &timeStr
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

func TimePtrToPbTime(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}

	return timestamppb.New(*t)
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

func PbUpdateProfileReqToProfile(p *pb.UpdateProfileReq) *Profile {
	if p == nil {
		return nil
	}

	profile := &Profile{
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

func ToPbSignInRes(r *SignInRes) *pb.SignInRes {
	return &pb.SignInRes{
		Auth:                 ToPbAuth(r.Auth),
		SessionId:            *r.SessionId,
		AccessToken:          *r.AccessToken,
		AccessTokenExpiresAt: ToPbTimestamp(*r.AccessTokenExpiresAt),
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
			Address1:       req.Address.Address1,
			Address2:       req.Address.Address2,
			Latitude:       req.Address.Latitude,
			Longitude:      req.Address.Longitude,
			NearbyLandmark: req.Address.NearbyLandmark,
			City:           req.Address.City,
			State:          req.Address.State,
			Pincode:        req.Address.Pincode,
			Country:        req.Address.Country,
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
		Address1:       a.Address1,
		Address2:       a.Address2,
		Longitude:      a.Longitude,
		Latitude:       a.Latitude,
		NearbyLandmark: a.NearbyLandmark,
		City:           a.City,
		State:          a.State,
		Pincode:        a.Pincode,
		Country:        a.Country,
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

func ToCreateRestaurantMenu(rm *pb.CreateRestaurantMenuReq) *CreateRestaurantMenu {
	return &CreateRestaurantMenu{
		MenuName: rm.MenuName,
		ShopID:   rm.ShopId,
	}
}

func ToCreateMenuItem(mi *pb.CreateMenuItemReq) *CreateMenuItem {
	return &CreateMenuItem{
		Name:        mi.Name,
		Description: mi.Description,
		Price:       mi.Price,
		MenuID:      HexToObjectId(mi.MenuId),
	}
}

func ToPbRestaurantMenuList(rm []*RestaurantMenu) []*pb.RestaurantMenu {
	list := make([]*pb.RestaurantMenu, len(rm))
	for i, l := range rm {
		if l != nil {
			list[i] = ToPbRestaurantMenu(l)
		}
	}

	return list
}

func ToPbRestaurantMenu(rm *RestaurantMenu) *pb.RestaurantMenu {
	if rm == nil {
		return nil
	}

	return &pb.RestaurantMenu{
		Id:        rm.ID.Hex(),
		MenuName:  rm.MenuName,
		ShopId:    rm.ShopID,
		MenuItems: ToPbMenuItems(rm.MenuItems),
		CreatedAt: ToPbTimestamp(rm.CreatedAt),
		UpdatedAt: ToPbTimestamp(rm.CreatedAt),
		DeletedAt: TimePtrToPbTime(rm.DeletedAt),
	}
}

func ToPbMenuItems(mi []*MenuItem) []*pb.MenuItem {
	menuItems := make([]*pb.MenuItem, len(mi))
	for i, m := range mi {
		if m != nil {
			menuItems[i] = ToPbMenuItem(m)
		}
	}
	return menuItems
}

func ToPbMenuItem(mi *MenuItem) *pb.MenuItem {
	return &pb.MenuItem{
		Id:          mi.ID.Hex(),
		Name:        mi.Name,
		Description: mi.Description,
		Price:       mi.Price,
		MenuId:      mi.MenuID.Hex(),
		Variants:    ToPbItemVariants(mi.Variants),
		Addons:      ToPbItemAddons(mi.AddOns),
		CreatedAt:   ToPbTimestamp(mi.CreatedAt),
		UpdatedAt:   ToPbTimestamp(mi.CreatedAt),
		DeletedAt:   TimePtrToPbTime(mi.DeletedAt),
	}
}

func ToPbItemVariants(v []*ItemVariant) []*pb.ItemVariant {
	variants := make([]*pb.ItemVariant, len(v))
	for i, vrs := range v {
		if vrs != nil {
			variants[i] = ToPbItemVariant(vrs)
		}
	}
	return variants
}

func ToPbItemVariant(v *ItemVariant) *pb.ItemVariant {
	return &pb.ItemVariant{
		Id:              v.ID.Hex(),
		VariantName:     v.VariantName,
		RelativePrice:   v.RelativePrice,
		RelativePricing: v.RelativePricing,
		Price:           v.Price,
		Description:     v.Description,
		ItemId:          v.ItemId.Hex(),
	}
}

func ToPbItemAddons(a []*ItemAddon) []*pb.ItemAddon {
	addons := make([]*pb.ItemAddon, len(a))
	for i, addon := range a {
		if addon != nil {
			addons[i] = ToPbItemAddon(addon)
		}
	}
	return addons
}

func ToPbItemAddon(a *ItemAddon) *pb.ItemAddon {
	return &pb.ItemAddon{
		Id:          a.ID.Hex(),
		AddonName:   a.AddonName,
		AddonPrice:  a.AddonPrice,
		Description: a.Description,
		ItemId:      a.ItemId.Hex(),
	}
}

func ToPbRetailItem(i *RetailItem) *pb.RetailItem {
	return &pb.RetailItem{
		Id:          i.ID.Hex(),
		Name:        i.Name,
		Description: i.Description,
		Price:       i.Price,
		CategoryId:  i.CategoryID.Hex(),
		Variants:    ToPbItemVariants(i.Variants),
		Addons:      ToPbItemAddons(i.AddOns),
		CreatedAt:   ToPbTimestamp(i.CreatedAt),
		UpdatedAt:   ToPbTimestamp(i.CreatedAt),
		DeletedAt:   TimePtrToPbTime(i.DeletedAt),
	}
}

func ToPbRetailItems(il []*RetailItem) []*pb.RetailItem {
	list := make([]*pb.RetailItem, len(il))
	for i, item := range il {
		if item != nil {
			list[i] = ToPbRetailItem(item)
		}
	}

	return list
}

func ToPbRetailCategory(c *RetailCategory) *pb.RetailCategory {
	return &pb.RetailCategory{
		Id:           c.ID.Hex(),
		CategoryName: c.CategoryName,
		ShopId:       c.ShopID,
		RetailItems:  ToPbRetailItems(c.RetailItems),
		CreatedAt:    ToPbTimestamp(c.CreatedAt),
		UpdatedAt:    ToPbTimestamp(c.CreatedAt),
		DeletedAt:    TimePtrToPbTime(c.DeletedAt),
	}
}
func ToPbRetailCategoryList(cl []*RetailCategory) []*pb.RetailCategory {
	list := make([]*pb.RetailCategory, len(cl))
	for i, item := range cl {
		if item != nil {
			list[i] = ToPbRetailCategory(item)
		}
	}

	return list
}

func ToPbMedicineItem(mi *MedicineItem) *pb.MedicineItem {
	return &pb.MedicineItem{
		Id:          mi.ID.Hex(),
		Name:        mi.Name,
		Price:       mi.Price,
		Description: mi.Description,
		CategoryId:  mi.CategoryID.Hex(),
		CreatedAt:   ToPbTimestamp(mi.CreatedAt),
		UpdatedAt:   ToPbTimestamp(mi.CreatedAt),
		DeletedAt:   TimePtrToPbTime(mi.DeletedAt),
	}
}

func ToPbMedicineItems(mil []*MedicineItem) []*pb.MedicineItem {
	list := make([]*pb.MedicineItem, len(mil))
	for i, item := range mil {
		if item != nil {
			list[i] = ToPbMedicineItem(item)
		}
	}

	return list
}

func ToPbMedicineCategory(c *MedicineCategory) *pb.MedicineCategory {
	return &pb.MedicineCategory{
		Id:            c.ID.Hex(),
		CategoryName:  c.CategoryName,
		ShopId:        c.ShopID,
		MedicineItems: ToPbMedicineItems(c.MedicineItems),
		CreatedAt:     ToPbTimestamp(c.CreatedAt),
		UpdatedAt:     ToPbTimestamp(c.CreatedAt),
		DeletedAt:     TimePtrToPbTime(c.DeletedAt),
	}
}
func ToPbMedicineCategoryList(cl []*MedicineCategory) []*pb.MedicineCategory {
	list := make([]*pb.MedicineCategory, len(cl))
	for i, item := range cl {
		if item != nil {
			list[i] = ToPbMedicineCategory(item)
		}
	}
	return list
}
