package app

import (
	"context"
)

type IRepository interface {
	GetDialogIfExists(context.Context, []ChatUser) (CreateChatResponse, bool, error)
	CreateChat(context.Context, []ChatUser) (CreateChatResponse, error)
	SendMessage(context.Context, ChatMessage) (uint, error)
	GetChatsList(context.Context, GetChatsListRequest) ([]MessageWithChatUsers, error)
	GetChat(context.Context, GetChatRequest) ([]ChatMessage, error)
	GetChatParticipants(context.Context, GetChatParticipantsRequest) (GetChatParticipantsResponse, error)
	IsMemberOfChat(context.Context, uint, uint) (bool, error)
}
