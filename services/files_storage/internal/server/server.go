package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type serverOptions struct {
	host           string
	port           string
	maxHeaderBytes int
	readTimeout    time.Duration
	writeTimeout   time.Duration
}

func getServerOptions() (opts serverOptions) {
	flag.StringVar(&opts.port, "p", "8081", "set the server's port")
	flag.IntVar(&opts.maxHeaderBytes, "m", 1, "set the server's max header bytes in MB")

	readTimeout := flag.Int64("rt", 10, "set the server's read timeout in seconds")
	writeTimout := flag.Int("wt", 10, "set the server's read timeout in seconds")

	flag.Parse()

	opts.maxHeaderBytes = opts.maxHeaderBytes << 20 // MB to Bytes
	opts.readTimeout = time.Duration(*readTimeout) * time.Second
	opts.writeTimeout = time.Duration(*writeTimout) * time.Second

	return opts
}

func RunServer() error {
	serverOptions := getServerOptions()

	httpServer := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", "localhost", serverOptions.port),
		Handler:        getRouter(),
		MaxHeaderBytes: serverOptions.maxHeaderBytes,
		ReadTimeout:    serverOptions.readTimeout,
		WriteTimeout:   serverOptions.writeTimeout,
	}

	log.Printf("Сервер микросервис файлов успешно запущен на %s:%s\n", "localhost", serverOptions.port)
	return httpServer.ListenAndServe()	
}
