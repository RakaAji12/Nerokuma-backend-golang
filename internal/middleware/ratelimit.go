package middleware

import (
	"net/http"
	"time"

	"github.com/juju/ratelimit"
)

func RateLimit(perMinute int) func(http.Handler) http.Handler {
	bucket := ratelimit.NewBucketWithQuantum(time.Minute, int64(perMinute), int64(perMinute))
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if bucket.TakeAvailable(1) == 0 {
				http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
