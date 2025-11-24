package cache

import (
	"encoding/json"
	"sync"
	"time"
)

type InMemoryCacheItem struct {
	Value      []byte
	Expiration time.Time
}

// MemoryCache adalah implementasi cache di memori.
type MemoryCache struct {
	mu   sync.RWMutex
	data map[string]InMemoryCacheItem
}

// NewMemoryCache membuat instance baru dari MemoryCache.
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		data: make(map[string]InMemoryCacheItem),
	}
}

// Set menyimpan nilai ke dalam cache memori.
func (c *MemoryCache) Set(key string, value interface{}, ttl time.Duration) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = InMemoryCacheItem{
		Value:      b,
		Expiration: time.Now().Add(ttl),
	}
	return nil
}

// Get mengambil nilai dari cache memori.
func (c *MemoryCache) Get(key string, dest interface{}) (bool, error) {
	c.mu.RLock()
	item, ok := c.data[key]
	c.mu.RUnlock()

	if !ok {
		return false, nil
	}

	if time.Now().After(item.Expiration) {
		c.mu.Lock()
		delete(c.data, key)
		c.mu.Unlock()
		return false, nil
	}

	return true, json.Unmarshal(item.Value, dest)
}
