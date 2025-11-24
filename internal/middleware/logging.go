package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func Logging() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Info().Str("method", r.Method).Str("path", r.URL.Path).Dur("t", time.Since(start)).Msg("")
		})
	}
}
