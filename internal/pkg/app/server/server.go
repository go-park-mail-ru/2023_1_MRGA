package server

import (
	"context"
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
}

type ServerOptions struct {
	Host           string
	Port           string
	MaxHeaderBytes int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

func GetServerOptions() (opts ServerOptions) {
	flag.StringVar(&opts.Host, "h", "0.0.0.0", "set the server's host")
	flag.StringVar(&opts.Port, "p", "8080", "set the server's port")
	flag.IntVar(&opts.MaxHeaderBytes, "m", 1, "set the server's max header bytes in MB")
	readTimeout := flag.Int64("rt", 10, "set the server's read timeout in seconds")
	writeTimout := flag.Int("wt", 10, "set the server's read timeout in seconds")
	flag.Parse()
	opts.MaxHeaderBytes = opts.MaxHeaderBytes << 20 // MB to Bytes
	opts.ReadTimeout = time.Duration(*readTimeout) * time.Second
	opts.WriteTimeout = time.Duration(*writeTimout) * time.Second

	return opts
}

func loadTLS() tls.Certificate {
	_, err := os.Open("/etc/letsencrypt/live/meetme-app.ru/privkey.pem")
	if err != nil {
		log.Println(err)
		log.Fatalln(err)
	}
	_, err = os.Open("/etc/letsencrypt/live/meetme-app.ru/fullchain.pem")
	if err != nil {
		log.Println(err)
		log.Fatalln(err)
	}
	cert, err := tls.LoadX509KeyPair("/etc/letsencrypt/live/meetme-app.ru/fullchain.pem", "/etc/letsencrypt/live/meetme-app.ru/privkey.pem")
	if err != nil {
		panic(err)
	}
	return cert
}

func (s *Server) Run(opts ServerOptions, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           opts.Host + ":" + opts.Port,
		Handler:        handler,
		MaxHeaderBytes: opts.MaxHeaderBytes,
		ReadTimeout:    opts.ReadTimeout,
		WriteTimeout:   opts.WriteTimeout,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{
				loadTLS(),
			},
		},
	}
	log.Println("server starts on ", s.httpServer.Addr)

	return s.httpServer.ListenAndServeTLS("", "")
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
