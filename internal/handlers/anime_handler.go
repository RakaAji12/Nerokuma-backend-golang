package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"nerokuma-api/internal/service"

	"github.com/gorilla/mux"
)

type AnimeHandler struct {
	malSvc service.MalServiceInterface
}

func NewAnimeHandler(svc service.MalServiceInterface) *AnimeHandler {
	return &AnimeHandler{malSvc: svc}
}

func (h *AnimeHandler) writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *AnimeHandler) getLimitFromQuery(r *http.Request) int {
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		return 10
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		return 10
	}
	return limit
}

func (h *AnimeHandler) SearchAnime(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}
	limit := h.getLimitFromQuery(r)
	result, err := h.malSvc.Search(query, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.writeJSON(w, http.StatusOK, result)
}

func (h *AnimeHandler) GetAnimeDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Anime ID is required", http.StatusBadRequest)
		return
	}
	result, err := h.malSvc.GetDetail(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.writeJSON(w, http.StatusOK, result)
}

func (h *AnimeHandler) GetTrending(w http.ResponseWriter, r *http.Request) {
	limit := h.getLimitFromQuery(r)
	result, err := h.malSvc.GetTrending(limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.writeJSON(w, http.StatusOK, result)
}
