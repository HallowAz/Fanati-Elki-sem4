package cors

import (
	"net/http"
)

const allowedOrigin = "http://83.166.237.142"

func CorsMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
		if r.Method == http.MethodOptions {
			w.Header().Add("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PATCH, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}

// CorsCredentionalsMiddleware provides header allow credentials
func CorsCredentionalsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
