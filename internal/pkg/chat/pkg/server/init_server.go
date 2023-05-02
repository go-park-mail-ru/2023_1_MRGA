package server

import (
	"fmt"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/proto_services/proto_chat"
)

type ServerOptions struct {
	Addr       string
	Port       int
	PathPrefix string
}

type Server struct {
	clientConn *grpc.ClientConn
	router     *mux.Router

	clientTarget string
}

func (server Server) InitClient() (chatClient chatpc.ChatServiceClient, chatClientConn *grpc.ClientConn, err error) {
	chatClientConn, err = grpc.Dial(server.clientTarget, grpc.WithInsecure())
	if err != nil {
		return
	}

	chatClient = chatpc.NewChatServiceClient(chatClientConn)

	return
}

func InitServer(opts ServerOptions) *mux.Router {
	var server Server

	server.clientTarget = fmt.Sprintf("%s:%d", opts.Addr, opts.Port)

	server.InitRouter(opts.PathPrefix)

	return server.router
}
