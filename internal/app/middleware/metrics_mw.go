package middleware

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

var fooCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "foo_total",
	Help: "Number of foo successfully processed.",
})

var hits = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "hits",
}, []string{"status", "path"})

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_response_time_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

var authHints = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "auth_hits",
}, []string{"error", "path"})

var authHttpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "auth_http_response_time_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

var compHints = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "comp_hits",
}, []string{"error", "path"})

var compHttpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "comp_http_response_time_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

func init() {
	prometheus.Register(fooCount)
	prometheus.Register(hits)
	prometheus.Register(httpDuration)
	prometheus.Register(authHints)
	prometheus.Register(authHttpDuration)
	prometheus.Register(compHints)
	prometheus.Register(compHttpDuration)
}

func MetricsMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/metrics") {
			next.ServeHTTP(w, r)
			return
		}
		fooCount.Add(1)
		rw := NewResponseWriter(w)
		start := time.Now()
		next.ServeHTTP(rw, r)
		duration := time.Since(start)
		st := rw.statusCode
		httpDuration.WithLabelValues(r.URL.String()).Observe(duration.Seconds())
		hits.WithLabelValues(strconv.Itoa(st), r.URL.String()).Inc()
	})
}

func authClientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	// Calls the invoker to execute RPC
	err := invoker(ctx, method, req, reply, cc, opts...)
	duration := time.Since(start)
	authHttpDuration.WithLabelValues(method).Observe(duration.Seconds())
	// Logic after invoking the invoker
	log.Println("hi",
		time.Since(start), err)
	message := ""
	if err != nil {
		message = err.Error()
	}
	authHints.WithLabelValues(message, method).Inc()
	return err
}

func AuthWithClientUnaryInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(authClientInterceptor)
}

func compClientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	// Calls the invoker to execute RPC
	err := invoker(ctx, method, req, reply, cc, opts...)
	duration := time.Since(start)
	compHttpDuration.WithLabelValues(method).Observe(duration.Seconds())
	// Logic after invoking the invoker
	log.Println("hi",
		time.Since(start), err)
	message := ""
	if err != nil {
		message = err.Error()
	}
	compHints.WithLabelValues(message, method).Inc()
	return err
}

func CompWithClientUnaryInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(compClientInterceptor)
}
