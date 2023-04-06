package middleware

import (
	"net/http"
	"strconv"

	"github.com/go-redis/redis"
	"golang.org/x/exp/slices"
	"golang.org/x/net/context"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func CorsMiddleware(allowedHosts []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Context()
		origin := r.Header.Get("Origin")
		if slices.Contains(allowedHosts, origin) {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		next.ServeHTTP(w, r)
	})

}

var ContextUserKey = "userId"

func AuthMiddleware(client *redis.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("session_token")
		userIdStr, err := client.Get(token).Result()
		if err != nil {
			logger.Log(http.StatusUnauthorized, err.Error(), r.Method, r.URL.Path)
			writer.ErrorRespond(w, r, err, http.StatusUnauthorized)
			return
		}

		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), ContextUserKey, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
