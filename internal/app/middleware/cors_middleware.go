package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"golang.org/x/exp/slices"
	"golang.org/x/net/context"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/cookie"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/default"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func CorsMiddleware(allowedHosts []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if slices.Contains(allowedHosts, origin) {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		next.ServeHTTP(w, r)
	})

}
var ContextUserKey = "userId"
var ProtectedPath = "/meetme/"

func AuthMiddleware(client *redis.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, ProtectedPath) {
			next.ServeHTTP(w, r)
			return
		}
		token, err := cookie.GetValueCookie(r, _default.SessionTokenCookieName)
		if err != nil {
			logger.Log(http.StatusUnauthorized, err.Error(), r.Method, r.URL.Path)
			writer.ErrorRespond(w, r, err, http.StatusUnauthorized)
			return
		}
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
