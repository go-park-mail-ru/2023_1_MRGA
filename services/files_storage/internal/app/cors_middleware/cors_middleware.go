package cors

import "net/http"

var allowedOrigins = []string{
	"http://5.159.100.59:8080",
	"http://localhost:8080",
	"http://192.168.0.2:8080",
}

func SetCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" || !contains(allowedOrigins, origin) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		next.ServeHTTP(w, r)
	})
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
