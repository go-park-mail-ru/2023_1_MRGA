package middleware

import (
	"net/http"
)

func CorsMiddleware(allowedHost string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", allowedHost)
		next.ServeHTTP(w, r)
	},
	)
}
