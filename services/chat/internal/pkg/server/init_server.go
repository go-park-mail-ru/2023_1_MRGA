package server

import (
	"log"
	"net"

	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/proto_services/proto_chat"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/app"
	"google.golang.org/grpc"
)

type Server struct {
	repository app.IRepository
	service    app.IService
	chatpc.UnimplementedChatServiceServer
}

func InitServer(repository app.IRepository, service app.IService) Server {
	var server = Server{
		repository: repository,
		service:    service,
	}

	return server
}

func (server *Server) RunServer() error {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Ошибка в создании tpc-соединения сервера: %v", err)
	}

	s := grpc.NewServer()
	chatpc.RegisterChatServiceServer(s, server)

	log.Printf("gRPC-микросервис чатов успешно запущен\n")
	return s.Serve(lis)
}
