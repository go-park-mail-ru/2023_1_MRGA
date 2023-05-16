package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/app"
	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/chat"
)

type ServerOptions struct {
	Port int
}

type Server struct {
	repository app.IRepository
	chatpc.UnimplementedChatServiceServer

	opts ServerOptions
}

func InitServer(opts ServerOptions, repository app.IRepository) Server {
	var server = Server{
		repository: repository,
		opts:       opts,
	}

	return server
}

func (server *Server) RunServer() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", server.opts.Port))
	if err != nil {
		log.Fatalf("Ошибка в создании tpc-соединения сервера: %v", err)
	}

	s := grpc.NewServer()
	chatpc.RegisterChatServiceServer(s, server)

	log.Printf("gRPC-микросервис чатов успешно запущен\n")
	return s.Serve(lis)
}
