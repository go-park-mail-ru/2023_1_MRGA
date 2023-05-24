package app

import (
	"context"
)

type IRepository interface {
	GetDialogIfExists(context.Context, []ChatUser) (CreateChatResponse, bool, error)
	CreateChat(context.Context, []ChatUser) (CreateChatResponse, error)
	SendMessage(context.Context, ChatMessage) (uint, error)
	GetChatsList(GetChatsListRequest) ([]MessageWithChatUsers, error)
	GetChat(GetChatRequest) ([]ChatMessage, error)
	GetChatParticipants(context.Context, GetChatParticipantsRequest) (GetChatParticipantsResponse, error)
	IsMemberOfChat(uint, uint) (bool, error)
}
