package middleware

import (
	"net/http"
	"os"
)

func OrderServiceAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		expectedApiKey := os.Getenv("ORDER_SERVICE_API_KEY")

		if apiKey == "" || apiKey != expectedApiKey {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
