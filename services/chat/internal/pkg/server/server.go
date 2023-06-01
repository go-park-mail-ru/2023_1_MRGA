package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/app"
	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/chat"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	tracejaeger "github.com/go-park-mail-ru/2023_1_MRGA.git/utils/trace_jaeger"
)

func (server Server) CreateChat(ctx context.Context, initialChatData *chatpc.CreateChatRequest) (outputChatData *chatpc.CreateChatResponse, err error) {
	ctx, span := tracejaeger.NewSpan(ctx, "chatServer", "CreateChat", nil)
	defer span.End()

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

func (server Server) SendMessage(ctx context.Context, newMsg *chatpc.SendMessageRequest) (outputMsgData *chatpc.SendMessageResponse, err error) {
	ctx, span := tracejaeger.NewSpan(ctx, "chatServer", "SendMessage", nil)
	defer span.End()

	msg := app.GetMessageStruct(newMsg)

	msgId, err := server.repository.SendMessage(ctx, msg)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), "POST", "SendMessage", true)
		return &chatpc.SendMessageResponse{}, err
	}

	logger.Log(http.StatusOK, "Success", "POST", "SendMessage", false)
	return &chatpc.SendMessageResponse{
		MsgId: uint32(msgId),
	}, nil
}

func (server Server) GetChatsList(userData *chatpc.GetChatsListRequest, streamRecentMsgs chatpc.ChatService_GetChatsListServer) (err error) {
	ctx, span := tracejaeger.NewSpan(streamRecentMsgs.Context(), "chatServer", "GetChatsList", nil)
	defer span.End()

	resentMessagesRequest := app.GetInitialUserStruct(userData)

	recentMessages, err := server.repository.GetChatsList(ctx, resentMessagesRequest)
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
	ctx, span := tracejaeger.NewSpan(streamChatMsgs.Context(), "chatServer", "GetChat", nil)
	defer span.End()

	initialChatData := app.GetInitialChatStruct(chatData)

	var isMemberOfChat bool
	isMemberOfChat, err = server.repository.IsMemberOfChat(ctx, initialChatData.UserId, initialChatData.ChatId)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), "GET", "GetChat", true)
		return err
	}
	if !isMemberOfChat {
		err = errors.New("У пользователя нет доступа к чату")
		logger.Log(http.StatusBadRequest, err.Error(), "GET", "GetChat", false)
		return err
	}

	var chatMsgs []app.ChatMessage
	chatMsgs, err = server.repository.GetChat(ctx, initialChatData)
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

func (server Server) GetChatParticipants(ctx context.Context, chatData *chatpc.GetChatParticipantsRequest) (participants *chatpc.GetChatParticipantsResponse, err error) {
	ctx, span := tracejaeger.NewSpan(ctx, "chatServer", "GetChatParticipants", nil)
	defer span.End()

	initialChatData := app.GetInitialChatForParticipantsStruct(chatData)

	var participantsStruct app.GetChatParticipantsResponse
	participantsStruct, err = server.repository.GetChatParticipants(ctx, initialChatData)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), "POST", "GetChatParticipants", true)
		return
	}

	participants = &chatpc.GetChatParticipantsResponse{
		ChatUserIds: participantsStruct.ChatUserIds,
	}

	logger.Log(http.StatusOK, "Success", "POST", "GetChatParticipants", false)
	return
}
