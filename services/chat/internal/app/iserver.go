package app

import (
	"context"

	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/proto_services/proto_chat"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IServer interface {
	SendMessage(context.Context, *chatpc.Message) (*emptypb.Empty, error)
	GetRecentMessages(*chatpc.ResentMessagesRequest, chatpc.ChatService_GetRecentMessagesServer) error
	GetConversationMessages(*chatpc.Message, chatpc.ChatService_GetConversationMessagesServer) error
	mustEmbedUnimplementedChatServiceServer()
}
