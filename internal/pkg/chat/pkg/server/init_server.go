package server

import (
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"sync"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"

	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/chat"
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

	upgrader      websocket.Upgrader
	userIdClients map[uint64][]*websocket.Conn
	mutex         *sync.Mutex
}

func (server Server) InitClient() (chatClient chatpc.ChatServiceClient, chatClientConn *grpc.ClientConn, err error) {
	chatClientConn, err = grpc.Dial(server.clientTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}

	chatClient = chatpc.NewChatServiceClient(chatClientConn)

	return
}

func (server *Server) InitRouter(pathPrefix string) {
	server.router = mux.NewRouter().PathPrefix(pathPrefix).Subrouter()

	// С префиксом /meetme/chats все ниже
	server.router.HandleFunc("/create", server.CreateChatHandler).Methods("POST")
	server.router.HandleFunc("/{chat_id}/send", server.SendMessageHandler).Methods("POST")
	server.router.HandleFunc("/list", server.GetChatsListHandler).Methods("GET")
	server.router.HandleFunc("/{chat_id}/messages", server.GetChatHandler).Methods("GET")
	server.router.HandleFunc("/subscribe", server.ConnectionHandler).Methods("GET")
}

func InitServer(opts ServerOptions) *mux.Router {
	server := Server{
		clientTarget: fmt.Sprintf("%s:%d", opts.Addr, opts.Port),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		userIdClients: make(map[uint64][]*websocket.Conn),
		mutex:         &sync.Mutex{},
	}

	server.InitRouter(opts.PathPrefix)

	return server.router
}
