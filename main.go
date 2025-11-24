package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"

	"nerokuma-api/internal/cache"
	"nerokuma-api/internal/config"
	"nerokuma-api/internal/handlers"
	"nerokuma-api/internal/middleware"
	"nerokuma-api/internal/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Tidak bisa membaca file .env")
	}

	cfg := config.Load()

	log.Printf("MAL Client ID yang dimuat: '%s'\n", cfg.MalClientID)

	var c cache.CacheInterface
	if cfg.UseRedis {
		rdb := redis.NewClient(&redis.Options{
			Addr:     cfg.RedisAddr,
			Password: cfg.RedisPass,
		})
		c = cache.NewRedisCache(rdb)
		log.Println("Using Redis cache")
	} else {
		c = cache.NewMemoryCache()
		log.Println("Using in-memory cache")
	}

	malService := service.NewMalService(cfg, c)

	animeHandler := handlers.NewAnimeHandler(malService)

	r := mux.NewRouter()
	r.Use(middleware.Logging())
	r.Use(middleware.RateLimit(cfg.RateLimitPerMinute))

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/anime/search", animeHandler.SearchAnime).Methods("GET")
	api.HandleFunc("/anime/trending", animeHandler.GetTrending).Methods("GET")
	api.HandleFunc("/anime/{id}", animeHandler.GetAnimeDetail).Methods("GET")

	addr := ":" + cfg.Port
	log.Printf("Server starting on %s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
