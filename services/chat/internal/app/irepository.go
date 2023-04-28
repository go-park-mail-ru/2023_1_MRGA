package app

import (
	"context"
)

type IRepository interface {
	CreateChat(context.Context, []ChatUser) (CreateChatResponse, error)
	SendMessage(context.Context, Message) error
	GetChatsList(GetChatsListRequest) ([]Message, error)
	GetChat(GetChatRequest) ([]Message, error)
}
