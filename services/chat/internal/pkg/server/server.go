package server

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/protobuf/types/known/emptypb"

	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/proto_services/proto_chat"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/app"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
)

func (server Server) CreateChat(ctx context.Context, initialChatData *chatpc.CreateChatRequest) (outputChatData *chatpc.CreateChatResponse, err error) {
	var userIds []app.ChatUser
	for _, grpcUserId := range initialChatData.GetUserIds() {
		userId := uint(grpcUserId)
		userIds = append(userIds, app.ChatUser{UserId: userId})
	}

	var createdChat app.CreateChatResponse
	createdChat, err = server.repository.CreateChat(ctx, userIds)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), "POST", "CreateChat")
		return
	}

	outputChatData = app.GetGrpcInitialChatData(createdChat)

	logger.Log(http.StatusOK, "Success", "POST", "CreateChat")
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

func (server Server) GetChatsList(userData *chatpc.GetChatsListRequest, streamRecentMsgs chatpc.ChatService_GetChatsListServer) (err error) {
	resentMessagesRequest := app.GetInitialUserStruct(userData)

	recentMessages, err := server.repository.GetChatsList(resentMessagesRequest)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), "GET", "GetChatsList")
		return err
	}

	for _, msg := range recentMessages {
		grpcChatMsg := app.GetGrpcChatMessage(msg)
		if err := streamRecentMsgs.Send(grpcChatMsg); err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), "GET", "GetChatsList")
			return err
		}
	}

	logger.Log(http.StatusOK, "Success", "GET", "GetChatsList")
	return
}

func (server Server) GetChat(chatData *chatpc.GetChatRequest, streamChatMsgs chatpc.ChatService_GetChatServer) (err error) {
	initialChatData := app.GetInitialChatStruct(chatData)

	var chatMsgs []app.Message
	chatMsgs, err = server.repository.GetChat(initialChatData)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), "GET", "GetChat")
		return err
	}

	for _, msg := range chatMsgs {
		grpcMsg := app.GetGrpcMessage(msg)
		if err := streamChatMsgs.Send(grpcMsg); err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), "GET", "GetChat")
			return err
		}
	}

	logger.Log(http.StatusOK, "Success", "GET", "GetChat")
	return
}

func (server Server) mustEmbedUnimplementedChatServiceServer() {
	log.Printf("Нереализованный метод")
}
