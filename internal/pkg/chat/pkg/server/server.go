package server

import (
	"log"
	"net/http"
)

func (server *Server) SendMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("В ручке SendMessage")
}

func (server *Server) GetRecentMessages(w http.ResponseWriter, r *http.Request) {
	log.Println("В ручке GetRecentMessages")
}

func (server *Server) GetConversationMessages(w http.ResponseWriter, r *http.Request) {
	log.Println("В ручке GetConversationMessages")
}
