package cache

import "time"

type CacheInterface interface {
	Get(key string, dest interface{}) (bool, error)
	Set(key string, value interface{}, ttl time.Duration) error
}
