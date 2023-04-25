package server

import (
	"fmt"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/chat/pkg/service"
	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/proto_services/proto_chat"
)

type Server struct {
	clientConn *grpc.ClientConn
	router     *mux.Router
	service    service.Service

	target string
}

func (server *Server) InitClient() (chatClient chatpc.ChatServiceClient, chatClientConn *grpc.ClientConn, err error) {
	chatClientConn, err = grpc.Dial(server.target, grpc.WithInsecure())
	if err != nil {
		return
	}

	chatClient = chatpc.NewChatServiceClient(chatClientConn)

	return
}

func InitServer(service service.Service, pathPrefix string) (*mux.Router) {
	server := Server{
		service: service,
	}

	server.InitRouter(pathPrefix)

	const (
		addr string = "localhost"
		port uint   = 3000
	)
	server.target = fmt.Sprintf("%s:%d", addr, port)

	return server.router
}
