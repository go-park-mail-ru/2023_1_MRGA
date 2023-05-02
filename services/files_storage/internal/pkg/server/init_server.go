package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal/app"
)

type ServerOptions struct {
	Host           string
	Port           string
	MaxHeaderBytes int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

type Server struct {
	repository app.IRepository
	service    app.IService
	httpServer *http.Server
}

func InitServer(opts ServerOptions, repository app.IRepository, service app.IService) Server {
	var server = Server{
		repository: repository,
		service:    service,
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
