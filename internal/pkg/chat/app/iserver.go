package app

import "net/http"

type IServer interface {
	CreateChatHander(http.ResponseWriter, *http.Request)
	SendMessageHander(http.ResponseWriter, *http.Request)
	GetChatsListHander(http.ResponseWriter, *http.Request)
	GetChatHander(http.ResponseWriter, *http.Request)
}
