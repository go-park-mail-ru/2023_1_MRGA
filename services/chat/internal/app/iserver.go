package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/chat"
)

type IServer interface {
	CreateChat(context.Context, *chatpc.CreateChatRequest) (*chatpc.CreateChatResponse, error)
	SendMessage(context.Context, *chatpc.SendMessageRequest) (*emptypb.Empty, error)
	GetChatsList(*chatpc.GetChatsListRequest, chatpc.ChatService_GetChatsListServer) error
	GetChat(chatData *chatpc.GetChatRequest, streamChatMsgs chatpc.ChatService_GetChatServer) (err error)
	mustEmbedUnimplementedChatServiceServer()
}
