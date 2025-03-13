package geolocation

import (
	"context"
	"fmt"

	"github.com/nrednav/cuid2"
	"github.com/yash91989201/superfast-delivery-api/common/types"
	"googlemaps.github.io/maps"
)

const GOOGLE_GEOCODING_URL = "https://maps.googleapis.com/maps/api/geocode/json"

type Service interface {
	ReverseGeocode(ctx context.Context, latitude, longitude float64, addressId string) (*types.AddressDetail, error)
	DeleteAddressDetail(ctx context.Context, addressId string) error
}

type geolocationService struct {
	r         Repository
	mapClient *maps.Client
}

func New(mapClient *maps.Client, r Repository) Service {
	return &geolocationService{
		r:         r,
		mapClient: mapClient,
	}
}

func (s *geolocationService) ReverseGeocode(ctx context.Context, latitude, longitude float64, addressId string) (*types.AddressDetail, error) {
	address, err := s.r.GetAddress(ctx, addressId)
	if err == nil {
		return address, nil
	}

	geocodeRequest := &maps.GeocodingRequest{
		LatLng: &maps.LatLng{
			Lat: latitude,
			Lng: longitude,
		},
	}

	results, err := s.mapClient.ReverseGeocode(ctx, geocodeRequest)

	if err != nil || len(results) == 0 {
		return nil, fmt.Errorf("failed to fetch address from Google Maps: %w", err)
	}

	firstResult := results[0]

	newAddress := &types.AddressDetail{
		Id:               cuid2.Generate(),
		PlaceId:          firstResult.PlaceID,
		FormattedAddress: firstResult.FormattedAddress,
		Latitude:         latitude,
		Longitude:        longitude,
		AddressId:        addressId,
	}

	// Extract address components
	for _, component := range firstResult.AddressComponents {
		for _, typ := range component.Types {
			switch typ {
			case "plus_code":
				newAddress.PlusCode = component.LongName
			case "route":
				newAddress.Route = component.LongName
			case "locality":
				newAddress.Town = component.LongName
			case "administrative_area_level_3":
				newAddress.District = component.LongName
			case "administrative_area_level_1":
				newAddress.State = component.LongName
			case "country":
				newAddress.Country = component.LongName
			case "postal_code":
				newAddress.PostalCode = component.LongName
			}
		}
	}

	err = s.r.SetAddress(ctx, newAddress)
	if err != nil {
		return nil, err
	}

	return newAddress, nil
}

func (s *geolocationService) DeleteAddressDetail(ctx context.Context, addressId string) error {
	return s.r.DeleteAddressDetail(ctx, addressId)
}
