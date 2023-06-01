package tracejaeger

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

func HTTPHandlerFunc(handler http.Handler, name string) http.HandlerFunc {
	return otelhttp.NewHandler(handler, name, otelhttp.WithTracerProvider(otel.GetTracerProvider())).ServeHTTP
}
