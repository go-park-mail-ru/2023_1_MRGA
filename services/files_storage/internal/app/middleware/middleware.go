package middleware

import (
	"net/http"

	tracejaeger "github.com/go-park-mail-ru/2023_1_MRGA.git/utils/trace_jaeger"
)

func JaegerMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tracejaeger.HTTPHandlerFunc(next, r.URL.String())(w, r)
	})
}
