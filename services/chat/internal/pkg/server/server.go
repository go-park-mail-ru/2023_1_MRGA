package server

import (
	"context"
	"log"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/proto_services/proto_chat"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/pkg/service"
)

func (server *Server) SendMessage(ctx context.Context, msgReq *chatpc.Message) (*emptypb.Empty, error) {
	message := service.GetStructMessage(msgReq)

	log.Printf("Структура Message в запросе SendMessage: %+v\n", message)
	return &emptypb.Empty{}, nil
}

func (server *Server) GetRecentMessages(req *chatpc.GetResentMessagesRequest, stream chatpc.ChatService_GetRecentMessagesServer) error {
	// message := service.GetStructMessage(msgReq)

	// log.Printf("Структура Message в запросе GetRecentMessages: %+v\n", message)
	return nil
}

func (server *Server) GetConversationMessages(msgReq *chatpc.Message, stream chatpc.ChatService_GetConversationMessagesServer) error {
	message := service.GetStructMessage(msgReq)

	log.Printf("Структура Message в запросе GetConversationMessages: %+v\n", message)
	return nil
}

func (server *Server) mustEmbedUnimplementedChatServiceServer() {
	log.Printf("Не реализованный метод")
}
