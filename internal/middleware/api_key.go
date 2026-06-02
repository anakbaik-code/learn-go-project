package middleware

import (
	"crypto/subtle"
	"net/http"
)

func ApiKeyMiddleware(apiKey string) func(http.Handler) http.Handler{
	handle := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request,
		) {
			key := r.Header.Get("X-API-Key")

			if subtle.ConstantTimeCompare(
				[]byte(key),
				[]byte(apiKey),
			) != 1 {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
	return handle
}
