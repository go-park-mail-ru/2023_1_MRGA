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

func (server Server) CreateChat(ctx context.Context, initialChatData *chatpc.CreateChatRequest) (outputChatData *chatpc.CreateChatResponse, err error) {
	return
}

func (server Server) SendMessage(ctx context.Context, newMsg *chatpc.SendMessageRequest) (outputMsgData *emptypb.Empty, err error) {
	msg := app.GetMessageStruct(newMsg)

	err = server.repository.SendMessage(ctx, msg)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), "POST", "SendMessage")
		return &emptypb.Empty{}, err
	}

	logger.Log(http.StatusOK, "Success", "POST", "SendMessage")
	return &emptypb.Empty{}, nil
}

func (server Server) GetChatsList(userData *chatpc.GetChatsListRequest, recentMsgs chatpc.ChatService_GetChatsListServer) (err error) {
	resentMessagesRequest := app.GetInitialUserStruct(userData)

	recentMessages, err := server.repository.GetChatsList(resentMessagesRequest)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), "GET", "GetRecentMessages")
		return err
	}

	for _, msg := range recentMessages {
		grpcChatMsg := app.GetGrpcChatMessage(msg)
		if err := recentMsgs.Send(grpcChatMsg); err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), "GET", "GetRecentMessages")
			return err
		}
	}

	logger.Log(http.StatusOK, "Success", "GET", "GetRecentMessages")
	return
}

func (server Server) GetChat(chatData *chatpc.GetChatRequest, chatMsgs chatpc.ChatService_GetChatServer) (err error) {
	// message := app.GetStructMessage(chatData)

	// log.Printf("Структура Message в запросе GetConversationMessages: %+v\n", message)
	return
}

func (server Server) mustEmbedUnimplementedChatServiceServer() {
	log.Printf("Не реализованный метод")
}
