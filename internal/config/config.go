package config

import (
	"fmt"
	"os"
)

type Config struct {
	MalClientID        string
	Port               string
	UseRedis           bool
	RedisAddr          string
	RedisPass          string
	CacheTTLSeconds    int
	RateLimitPerMinute int
}

func Load() *Config {
	// default values
	cfg := &Config{
		MalClientID:        os.Getenv("MAL_CLIENT_ID"),
		Port:               getEnv("PORT", "8080"),
		UseRedis:           getEnv("USE_REDIS", "false") == "true",
		RedisAddr:          getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPass:          os.Getenv("REDIS_PASS"),
		CacheTTLSeconds:    atoiWithDefault(getEnv("CACHE_TTL_SECONDS", "300"), 300),
		RateLimitPerMinute: atoiWithDefault(getEnv("RATE_LIMIT_PER_MINUTE", "60"), 60),
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func atoiWithDefault(s string, d int) int {
	var v int
	_, err := fmt.Sscanf(s, "%d", &v)
	if err != nil {
		return d
	}
	return v
}
