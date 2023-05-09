package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/app"
	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/chat"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
)

func (server Server) CreateChat(ctx context.Context, initialChatData *chatpc.CreateChatRequest) (outputChatData *chatpc.CreateChatResponse, err error) {
	var userIds []app.ChatUser
	for _, grpcUserId := range initialChatData.GetUserIds() {
		userId := uint(grpcUserId)
		userIds = append(userIds, app.ChatUser{UserId: userId})
	}

	var countUsers int = len(userIds)

	if countUsers == 1 {
		err = errors.New("Чат должен состоять, как минимум, из двух участников")
		logger.Log(http.StatusBadRequest, err.Error(), "POST", "CreateChat", true)
		return
	}

	var (
		createdChat app.CreateChatResponse
		found       = false
	)
	if countUsers == 2 {
		createdChat, found, err = server.repository.GetDialogIfExists(ctx, userIds)
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), "POST", "CreateChat", true)
			return
		}
	}

	if !found {
		createdChat, err = server.repository.CreateChat(ctx, userIds)
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), "POST", "CreateChat", true)
			return
		}
	}

	outputChatData = app.GetGrpcInitialChatData(createdChat)

	logger.Log(http.StatusOK, "Success", "POST", "CreateChat", false)
	return
}

func (server Server) SendMessage(ctx context.Context, newMsg *chatpc.SendMessageRequest) (outputMsgData *emptypb.Empty, err error) {
	msg := app.GetMessageStruct(newMsg)

	err = server.repository.SendMessage(ctx, msg)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), "POST", "SendMessage", true)
		return &emptypb.Empty{}, err
	}

	logger.Log(http.StatusOK, "Success", "POST", "SendMessage", false)
	return &emptypb.Empty{}, nil
}

func (server Server) GetChatsList(userData *chatpc.GetChatsListRequest, streamRecentMsgs chatpc.ChatService_GetChatsListServer) (err error) {
	resentMessagesRequest := app.GetInitialUserStruct(userData)

	recentMessages, err := server.repository.GetChatsList(resentMessagesRequest)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), "GET", "GetChatsList", true)
		return err
	}

	for _, msg := range recentMessages {
		grpcChatMsg := app.GetGrpcChatMessage(msg)
		if err := streamRecentMsgs.Send(grpcChatMsg); err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), "GET", "GetChatsList", true)
			return err
		}
	}

	logger.Log(http.StatusOK, "Success", "GET", "GetChatsList", false)
	return
}

func (server Server) GetChat(chatData *chatpc.GetChatRequest, streamChatMsgs chatpc.ChatService_GetChatServer) (err error) {
	initialChatData := app.GetInitialChatStruct(chatData)

	var isMemberOfChat bool
	isMemberOfChat, err = server.repository.IsMemberOfChat(initialChatData.UserId, initialChatData.ChatId)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), "GET", "GetChat", true)
		return err
	}
	if !isMemberOfChat {
		err = errors.New("У пользователя нет доступа к чату")
		logger.Log(http.StatusBadRequest, err.Error(), "GET", "GetChat", false)
		return err
	}

	var chatMsgs []app.Message
	chatMsgs, err = server.repository.GetChat(initialChatData)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), "GET", "GetChat", true)
		return err
	}
	for _, msg := range chatMsgs {
		grpcMsg := app.GetGrpcMessage(msg)
		if err := streamChatMsgs.Send(grpcMsg); err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), "GET", "GetChat", true)
			return err
		}
	}

	logger.Log(http.StatusOK, "Success", "GET", "GetChat", false)
	return
}

func (server Server) mustEmbedUnimplementedChatServiceServer() {
	log.Printf("Нереализованный метод")
}
