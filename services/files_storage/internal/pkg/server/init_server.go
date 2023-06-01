package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal/app"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal/app/middleware"
	"github.com/gorilla/mux"
)

type ServerOptions struct {
	Host           string
	Port           string
	MaxHeaderBytes int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

type Server struct {
	service    app.IService
	httpServer *http.Server
}

func (server Server) getRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(func(h http.Handler) http.Handler {
		return middleware.JaegerMW(h)
	})

	router.HandleFunc("/api/v1/files/upload", server.UploadPhoto).Methods("POST")
	router.HandleFunc("/api/v2/files/upload", server.UploadFile).Methods("POST")
	router.HandleFunc("/api/files/{id}", server.GetFile).Methods("GET")
	router.HandleFunc("/api/files/{pathToFile:.*}", server.GetFileByPath).Methods("GET")
	return router
}

func InitServer(opts ServerOptions, service app.IService) Server {
	var server = Server{
		service: service,
	}
	return Server{
		httpServer: &http.Server{
			Addr:           fmt.Sprintf("%s:%s", opts.Host, opts.Port),
			Handler:        server.getRouter(),
			MaxHeaderBytes: opts.MaxHeaderBytes,
			ReadTimeout:    opts.ReadTimeout,
			WriteTimeout:   opts.WriteTimeout,
		},
	}
}

func (server *Server) RunServer() error {
	log.Printf("Сервер микросервис файлов успешно запущен на %s\n", server.httpServer.Addr)
	return server.httpServer.ListenAndServe()
}
