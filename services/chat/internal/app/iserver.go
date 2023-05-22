package app

import (
	"context"

	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/chat"
)

type IServer interface {
	CreateChat(context.Context, *chatpc.CreateChatRequest) (*chatpc.CreateChatResponse, error)
	SendMessage(context.Context, *chatpc.SendMessageRequest) (*chatpc.SendMessageResponse, error)
	GetChatsList(*chatpc.GetChatsListRequest, chatpc.ChatService_GetChatsListServer) error
	GetChat(*chatpc.GetChatRequest, chatpc.ChatService_GetChatServer) error
	GetChatParticipants(context.Context, *chatpc.GetChatParticipantsRequest) (*chatpc.GetChatParticipantsResponse, error)
	mustEmbedUnimplementedChatServiceServer()
}
