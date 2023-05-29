package middleware

import (
	"bufio"
	"errors"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     0,
	}

}

func (rw *ResponseWriter) WriteHeader(code int) {
	if rw.statusCode != 0 {
		return
	}
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *ResponseWriter) Write(data []byte) (int, error) {
	return rw.ResponseWriter.Write(data)
}
func (rw *ResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := rw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("hijack not supported")
	}
	return h.Hijack()
}

var fooCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "foo_total",
	Help: "Number of foo successfully processed.",
})

var hits = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "hits",
}, []string{"status", "path"})

var httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_response_time_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

var authHints = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "auth_hits",
}, []string{"error", "path"})

var authHttpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "auth_http_response_time_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

var compHints = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "comp_hits",
}, []string{"error", "path"})

var compHttpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "comp_http_response_time_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

func init() {
	err := prometheus.Register(fooCount)
	if err != nil {
		logger.Log(0, err.Error(), "init metrics", "/", true)
	}
	err = prometheus.Register(hits)
	if err != nil {
		logger.Log(0, err.Error(), "init metrics", "/", true)
	}
	err = prometheus.Register(httpDuration)
	if err != nil {
		logger.Log(0, err.Error(), "init metrics", "/", true)
	}
	err = prometheus.Register(authHints)
	if err != nil {
		logger.Log(0, err.Error(), "init metrics", "/", true)
	}
	err = prometheus.Register(authHttpDuration)
	if err != nil {
		logger.Log(0, err.Error(), "init metrics", "/", true)
	}
	err = prometheus.Register(compHints)
	if err != nil {
		logger.Log(0, err.Error(), "init metrics", "/", true)
	}
	err = prometheus.Register(compHttpDuration)
	if err != nil {
		logger.Log(0, err.Error(), "init metrics", "/", true)
	}
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

	err := invoker(ctx, method, req, reply, cc, opts...)
	duration := time.Since(start)
	authHttpDuration.WithLabelValues(method).Observe(duration.Seconds())

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
	err := invoker(ctx, method, req, reply, cc, opts...)
	duration := time.Since(start)
	compHttpDuration.WithLabelValues(method).Observe(duration.Seconds())

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
