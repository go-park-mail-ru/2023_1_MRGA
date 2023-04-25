package server

import (
	"context"
	"log"
	"net/http"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/proto_services/proto_chat"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/app"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
)

func (server Server) SendMessage(ctx context.Context, msgReq *chatpc.Message) (*emptypb.Empty, error) {
	msg := app.GetStructMessage(msgReq)

	err := server.repository.SendMessage(ctx, msg)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), "POST", "SendMessage")
		return &emptypb.Empty{}, err
	}

	logger.Log(http.StatusOK, "Success", "POST", "SendMessage")
	return &emptypb.Empty{}, nil
}

func (server Server) GetRecentMessages(req *chatpc.ResentMessagesRequest, stream chatpc.ChatService_GetRecentMessagesServer) error {
	resentMessagesRequest := app.GetStructResentMessagesRequest(req)

	recentMessages, err := server.repository.GetRecentMessages(resentMessagesRequest.UserId)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), "GET", "GetRecentMessages")
		return err
	}
	log.Println(recentMessages)

	for _, msg := range recentMessages {
		grpcMsg := app.GetGRPCMessage(msg)
		if err := stream.Send(grpcMsg); err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), "GET", "GetRecentMessages")
			return err
		}
	}

	logger.Log(http.StatusOK, "Success", "GET", "GetRecentMessages")
	return nil
}

func (server Server) GetConversationMessages(msgReq *chatpc.Message, stream chatpc.ChatService_GetConversationMessagesServer) error {
	message := app.GetStructMessage(msgReq)

	log.Printf("Структура Message в запросе GetConversationMessages: %+v\n", message)
	return nil
}

func (server Server) mustEmbedUnimplementedChatServiceServer() {
	log.Printf("Не реализованный метод")
}
