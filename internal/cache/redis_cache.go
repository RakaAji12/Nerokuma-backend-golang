package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
		ctx:    context.Background(),
	}
}

func (c *RedisCache) Set(key string, value interface{}, ttl time.Duration) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(c.ctx, key, b, ttl).Err()
}

func (c *RedisCache) Get(key string, dest interface{}) (bool, error) {
	r, err := c.client.Get(c.ctx, key).Bytes()
	if err == redis.Nil {
		return false, nil // Tidak ditemukan
	}
	if err != nil {
		return false, err // Error lain
	}
	return true, json.Unmarshal(r, dest)
}
