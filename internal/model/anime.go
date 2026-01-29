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
	ImageURL string  `json:"image_url"`
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

type AnimeSearchResponse struct {
	Data []struct {
		Node Anime `json:"node"`
	} `json:"data"`
}

type AnimeTrendingResponse struct {
	Data []struct {
		Node Anime `json:"node"`
	} `json:"data"`
}
