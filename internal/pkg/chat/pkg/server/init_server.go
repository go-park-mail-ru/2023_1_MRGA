package server

import (
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/chat/pkg/service"
	"github.com/gorilla/mux"
)

type Server struct {
	router  *mux.Router
	service service.Service
}

func InitServer(service service.Service) *mux.Router {
	server := Server{
		router:  mux.NewRouter().PathPrefix("/chat").Subrouter(),
		service: service,
	}

	// С префиксом /chat все ниже
	server.router.HandleFunc("/send/message", server.SendMessage).Methods("POST")
	server.router.HandleFunc("/recent/messages/{user_id}", server.GetRecentMessages).Methods("GET")
	server.router.HandleFunc("/{single_user_id}/{other_user_id}/messages", server.GetConversationMessages).Methods("GET")

	return server.router
}
