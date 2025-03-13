package geolocation

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yash91989201/superfast-delivery-api/common/types"
)

type Repository interface {
	Close() error

	GetAddress(ctx context.Context, addressId string) (*types.AddressDetail, error)
	SetAddress(ctx context.Context, addressDetail *types.AddressDetail) error
	DeleteAddressDetail(ctx context.Context, addressId string) error
}

type redisRepository struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisRepository(redisUrl string, ttl time.Duration) (Repository, error) {
	opts, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Redis URL: %w", err)
	}

	client := redis.NewClient(opts)

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &redisRepository{
		client: client,
		ttl:    ttl,
	}, nil
}

func (r *redisRepository) Close() error {
	return r.client.Close()
}

func (r *redisRepository) GetAddress(ctx context.Context, addressId string) (*types.AddressDetail, error) {
	key := fmt.Sprintf("address:%s", addressId)

	data, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("address not found")
	} else if err != nil {
		return nil, fmt.Errorf("failed to get address from Redis: %w", err)
	}

	var addressDetail types.AddressDetail
	if err := json.Unmarshal([]byte(data), &addressDetail); err != nil {
		return nil, fmt.Errorf("failed to unmarshal addressDetail: %w", err)
	}

	return &addressDetail, nil
}

func (r *redisRepository) SetAddress(ctx context.Context, addressDetail *types.AddressDetail) error {
	data, err := json.Marshal(addressDetail)
	if err != nil {
		return fmt.Errorf("failed to marshal addressDetail: %w", err)
	}

	key := fmt.Sprintf("address:%s", addressDetail.AddressId)
	if err := r.client.Set(ctx, key, data, r.ttl).Err(); err != nil {
		return fmt.Errorf("failed to save address to Redis: %w", err)
	}

	return nil
}

func (r *redisRepository) DeleteAddressDetail(ctx context.Context, addressId string) error {
	key := fmt.Sprintf("address:%s", addressId)

	if err := r.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete address from Redis: %w", err)
	}

	return nil
}
