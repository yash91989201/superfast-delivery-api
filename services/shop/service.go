package shop

import (
	"context"

	"github.com/nrednav/cuid2"
	"github.com/yash91989201/superfast-delivery-api/common/types"
)

type (
	Service interface {
		InsertShop(context.Context, *types.CreateShop) (*types.Shop, error)

		GetShopInfo(ctx context.Context, id string) (*types.ShopInfo, error)
		GetShopInfoByOwnerId(ctx context.Context, ownerId string) (*types.ShopInfo, error)
		GetShop(ctx context.Context, id string) (*types.Shop, error)
		GetShopByOwnerId(ctx context.Context, ownerId string) (*types.Shop, error)
	}

	shopService struct {
		r Repository
	}
)

func New(r Repository) Service {
	return &shopService{r: r}
}

func (s *shopService) InsertShop(ctx context.Context, shop *types.CreateShop) (*types.Shop, error) {
	shopId := cuid2.Generate()

	shopTiming := make([]*types.ShopTiming, len(shop.Timing))
	for i, timing := range shop.Timing {
		shopTiming[i] = &types.ShopTiming{
			ID:       cuid2.Generate(),
			Day:      timing.Day,
			OpensAt:  timing.OpensAt,
			ClosesAt: timing.ClosesAt,
			ShopID:   shopId,
		}
	}

	shopImage := make([]*types.ShopImage, len(shop.Image))
	for i, image := range shop.Image {
		shopImage[i] = &types.ShopImage{
			ID:          cuid2.Generate(),
			ImageUrl:    image.ImageUrl,
			Description: image.Description,
			ShopID:      shopId,
		}
	}

	newShop := &types.Shop{
		ID:         shopId,
		Name:       shop.Name,
		ShopType:   shop.ShopType,
		ShopStatus: shop.ShopStatus,
		OwnerID:    shop.OwnerId,
		Address: &types.ShopAddress{
			ID:             cuid2.Generate(),
			Address1:       shop.Address.Address1,
			Address2:       shop.Address.Address2,
			Longitude:      shop.Address.Longitude,
			Latitude:       shop.Address.Latitude,
			NearbyLandmark: shop.Address.NearbyLandmark,
			City:           shop.Address.City,
			State:          shop.Address.State,
			Pincode:        shop.Address.Pincode,
			Country:        shop.Address.Country,
			ShopID:         shopId,
		},
		Contact: &types.ShopContact{
			ID:          cuid2.Generate(),
			Name:        shop.Contact.Name,
			PhoneNumber: shop.Contact.PhoneNumber,
			Email:       shop.Contact.Email,
			ShopID:      shopId,
		},
		Timing: shopTiming,
		Image:  shopImage,
	}

	if err := s.r.InsertShop(ctx, newShop); err != nil {
		return nil, err
	}

	return newShop, nil
}

func (s *shopService) GetShopInfo(ctx context.Context, id string) (*types.ShopInfo, error) {
	return s.r.GetShopInfo(ctx, id)
}

func (s *shopService) GetShopInfoByOwnerId(ctx context.Context, ownerId string) (*types.ShopInfo, error) {
	return s.r.GetShopInfoByOwnerId(ctx, ownerId)
}

func (s *shopService) GetShop(ctx context.Context, id string) (*types.Shop, error) {
	shopInfo, err := s.r.GetShopInfo(ctx, id)
	if err != nil {
		return nil, err
	}

	shopAddress, err := s.r.GetShopAddressByShopId(ctx, id)
	if err != nil {
		return nil, err
	}

	shopContact, err := s.r.GetShopContactByShopId(ctx, id)
	if err != nil {
		return nil, err
	}

	shopTimings, err := s.r.GetShopTimings(ctx, id)
	if err != nil {
		return nil, err
	}

	shopImages, err := s.r.GetShopImages(ctx, id)
	if err != nil {
		return nil, err
	}

	return &types.Shop{
		ID:         shopInfo.ID,
		Name:       shopInfo.Name,
		ShopType:   shopInfo.ShopType,
		ShopStatus: shopInfo.ShopStatus,
		OwnerID:    shopInfo.OwnerID,
		CreatedAt:  shopInfo.CreatedAt,
		UpdatedAt:  shopInfo.UpdatedAt,
		DeletedAt:  shopInfo.DeletedAt,
		Address:    shopAddress,
		Contact:    shopContact,
		Timing:     shopTimings,
		Image:      shopImages,
	}, nil
}

func (s *shopService) GetShopByOwnerId(ctx context.Context, ownerId string) (*types.Shop, error) {
	shopInfo, err := s.r.GetShopInfoByOwnerId(ctx, ownerId)
	if err != nil {
		return nil, err
	}

	shopAddress, err := s.r.GetShopAddressByShopId(ctx, shopInfo.ID)
	if err != nil {
		return nil, err
	}

	shopContact, err := s.r.GetShopContactByShopId(ctx, shopInfo.ID)
	if err != nil {
		return nil, err
	}

	shopTimings, err := s.r.GetShopTimings(ctx, shopInfo.ID)
	if err != nil {
		return nil, err
	}

	shopImages, err := s.r.GetShopImages(ctx, shopInfo.ID)
	if err != nil {
		return nil, err
	}

	return &types.Shop{
		ID:         shopInfo.ID,
		Name:       shopInfo.Name,
		ShopType:   shopInfo.ShopType,
		ShopStatus: shopInfo.ShopStatus,
		OwnerID:    shopInfo.OwnerID,
		CreatedAt:  shopInfo.CreatedAt,
		UpdatedAt:  shopInfo.UpdatedAt,
		DeletedAt:  shopInfo.DeletedAt,
		Address:    shopAddress,
		Contact:    shopContact,
		Timing:     shopTimings,
		Image:      shopImages,
	}, nil
}
