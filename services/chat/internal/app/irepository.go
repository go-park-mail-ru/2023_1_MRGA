package app

import (
	"context"
)

type IRepository interface {
	GetDialogIfExists(context.Context, []ChatUser) (CreateChatResponse, bool, error)
	CreateChat(context.Context, []ChatUser) (CreateChatResponse, error)
	SendMessage(context.Context, Message) error
	GetChatsList(GetChatsListRequest) ([]MessageWithChatUsers, error)
	GetChat(GetChatRequest) ([]Message, error)
	IsMemberOfChat(uint, uint) (bool, error)
}
