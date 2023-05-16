package app

import "net/http"

type IServer interface {
	CreateChatHandler(http.ResponseWriter, *http.Request)
	SendMessageHandler(http.ResponseWriter, *http.Request)
	GetChatsListHandler(http.ResponseWriter, *http.Request)
	GetChatHandler(http.ResponseWriter, *http.Request)
	ConnectionHandler(http.ResponseWriter, *http.Request)
}
