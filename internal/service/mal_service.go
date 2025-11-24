package service

import (
	"encoding/json"
	"fmt"
	"io"
	"nerokuma-api/internal/cache"
	"nerokuma-api/internal/config"
	"nerokuma-api/internal/model"
	"net/http"
	"net/url"
	"time"

	"github.com/rs/zerolog/log"
)

type MalServiceInterface interface {
	Search(q string, limit int) (*model.AnimeSearchResponse, error)
	GetDetail(id string) (*model.Anime, error)
	GetTrending(limit int) (*model.AnimeTrendingResponse, error)
}

type MalService struct {
	clientID string
	cache    cache.CacheInterface
	cacheTTL time.Duration
	client   *http.Client
}

func NewMalService(cfg *config.Config, c cache.CacheInterface) *MalService {
	return &MalService{
		clientID: cfg.MalClientID,
		cache:    c,
		cacheTTL: time.Duration(cfg.CacheTTLSeconds) * time.Second,
		client:   &http.Client{Timeout: 10 * time.Second},
	}
}

func (s *MalService) doGet(apiURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-MAL-CLIENT-ID", s.clientID)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Warning: failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s - %s", resp.Status, string(body))
	}

	return io.ReadAll(resp.Body)
}

func (s *MalService) Search(q string, limit int) (*model.AnimeSearchResponse, error) {
	cacheKey := fmt.Sprintf("search:%s:%d", q, limit)
	var resp model.AnimeSearchResponse
	if found, _ := s.cache.Get(cacheKey, &resp); found {
		return &resp, nil
	}

	apiURL := fmt.Sprintf("https://api.myanimelist.net/v2/anime?q=%s&limit=%d&fields=id,title,main_picture,mean,synopsis,genres", url.QueryEscape(q), limit)
	b, err := s.doGet(apiURL)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}
	_ = s.cache.Set(cacheKey, resp, s.cacheTTL)
	return &resp, nil
}

func (s *MalService) GetDetail(id string) (*model.Anime, error) {
	cacheKey := fmt.Sprintf("detail:%s", id)
	var a model.Anime
	if found, _ := s.cache.Get(cacheKey, &a); found {
		return &a, nil
	}

	apiURL := fmt.Sprintf("https://api.myanimelist.net/v2/anime/%s?fields=id,title,main_picture,mean,synopsis,genres", id)
	b, err := s.doGet(apiURL)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &a); err != nil {
		return nil, err
	}
	_ = s.cache.Set(cacheKey, a, s.cacheTTL)
	return &a, nil
}

func (s *MalService) GetTrending(limit int) (*model.AnimeTrendingResponse, error) {
	cacheKey := fmt.Sprintf("trending:%d", limit)
	var resp model.AnimeTrendingResponse
	if found, _ := s.cache.Get(cacheKey, &resp); found {
		return &resp, nil
	}

	apiURL := fmt.Sprintf("https://api.myanimelist.net/v2/anime/ranking?ranking_type=all&limit=%d&fields=id,title,main_picture,mean,synopsis,genres", limit)
	b, err := s.doGet(apiURL)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}
	_ = s.cache.Set(cacheKey, resp, s.cacheTTL)
	return &resp, nil
}
