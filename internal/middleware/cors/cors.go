package cors

import (
	"github.com/gorilla/handlers"
	"net/http"
)

func CorsMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.CORS(
			handlers.AllowedOrigins([]string{"http://83.166.237.142"}),
			handlers.AllowedHeaders([]string{"Content-Type", "X-Auth-Token"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS", "HEAD", "DELETE"}),
			handlers.AllowCredentials(),
		)
		next.ServeHTTP(w, r)
	})
}
