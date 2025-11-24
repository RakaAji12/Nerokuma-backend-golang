// file: internal/model/anime.go
package model

import (
	"encoding/json"
)

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Anime struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	ImageURL string  `json:"-"`
	Score    float64 `json:"mean"`
	Synopsis string  `json:"synopsis"`
	Genres   []Genre `json:"genres"`
}

func (a *Anime) UnmarshalJSON(data []byte) error {
	type Alias Anime

	temp := &struct {
		MainPicture map[string]string `json:"main_picture"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	if temp.MainPicture != nil {
		// Prioritaskan gambar 'large', jika tidak ada gunakan 'medium'
		if img, ok := temp.MainPicture["large"]; ok {
			a.ImageURL = img
		} else if img, ok := temp.MainPicture["medium"]; ok {
			a.ImageURL = img
		}
	}

	return nil
}

// AnimeSearchResponse adalah model untuk respons pencarian dari MAL API.
type AnimeSearchResponse struct {
	Data []struct {
		Node Anime `json:"node"`
	} `json:"data"`
}

// AnimeTrendingResponse adalah model untuk respons trending/ranking dari MAL API.
type AnimeTrendingResponse struct {
	Data []struct {
		Node Anime `json:"node"`
	} `json:"data"`
}
