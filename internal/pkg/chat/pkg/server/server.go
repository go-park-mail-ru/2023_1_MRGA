package server

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/fatih/structs"
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/chat/app"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/chat/app/constants"
	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/chat"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (server Server) CreateChatHandler(w http.ResponseWriter, r *http.Request) {
	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, constants.ErrSessionExpired, r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, errors.New(constants.ErrSessionExpired), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var initialChatData app.CreateChatRequest
	err := json.NewDecoder(r.Body).Decode(&initialChatData)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	initialChatData.UserIds = append(initialChatData.UserIds, uint(userId))

	grpcInitialChatData := app.GetGRPCInitialChatData(initialChatData)

	client, conn, err := server.InitClient()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	var grpcCreatedChatData *chatpc.CreateChatResponse
	grpcCreatedChatData, err = client.CreateChat(context.Background(), grpcInitialChatData)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	createdChatData := app.GetCreatedChatDataStruct(grpcCreatedChatData)

	logger.Log(http.StatusOK, constants.LogSuccess, r.Method, r.URL.Path, false)
	writer.Respond(w, r, structs.Map(createdChatData))
}

func (server Server) SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var msgData app.SendMessageRequest
	err := json.NewDecoder(r.Body).Decode(&msgData)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	strChatId := vars["chat_id"]
	uint64ChatId, err := strconv.ParseUint(strChatId, 10, 64)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, constants.ErrSessionExpired, r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, errors.New(constants.ErrSessionExpired), http.StatusBadRequest)
		return
	}

	msg := app.InitialMessageData{
		Message: app.Message{
			SenderId:    uint(userId),
			Content:     msgData.Content,
			SentAt:      time.Now(),
			ReadStatus:  false,
			MessageType: msgData.MessageType,
			Path:        msgData.Path,
		},
	}

	grpcMsg := app.GetGRPCChatMessage(msg, uint(uint64ChatId))

	client, conn, err := server.InitClient()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	response, err := client.SendMessage(context.Background(), grpcMsg)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	msgId := uint64(response.GetMsgId())
	sentAt := msg.SentAt.Format(constants.FormatData)
	senderId := uint64(userId)

	wsMsgData := app.WSMsgData{
		SenderId: senderId,
		UserIds:  msgData.UserIds,
		MsgData: app.WSMessageResponse{
			Msg:         msgData.Content,
			ChatId:      uint64ChatId,
			SentAt:      sentAt,
			SenderId:    senderId,
			MsgId:       msgId,
			MessageType: string(msgData.MessageType),
			Path:        msgData.Path,
		},
	}

	err = server.sendAll(wsMsgData)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), constants.LogPostMethod, constants.LogOnMessageHandler, true)
		return
	}

	logger.Log(http.StatusOK, constants.LogSuccess, r.Method, r.URL.Path, false)
	writer.Respond(w, r, structs.Map(app.SendMessageResponse{
		SentAt: sentAt,
		MsgId:  msgId,
	}))
}

func (server Server) GetChatsListHandler(w http.ResponseWriter, r *http.Request) {
	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, constants.ErrSessionExpired, r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, errors.New(constants.ErrSessionExpired), http.StatusBadRequest)
		return
	}

	grpcChatsListRequest := chatpc.GetChatsListRequest{
		UserId: userId,
	}

	client, conn, err := server.InitClient()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	streamChatsList, err := client.GetChatsList(context.Background(), &grpcChatsListRequest)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	var chatsMessages []app.ChatMessage
	for {
		grpcChatMsg, err := streamChatsList.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}

		chatMsg := app.GetChatMessageStruct(grpcChatMsg)
		chatsMessages = append(chatsMessages, chatMsg)
	}

	logger.Log(http.StatusOK, constants.LogSuccess, r.Method, r.URL.Path, false)
	writer.Respond(w, r, structs.Map(app.GetChatsListResponse{
		ChatsList: chatsMessages,
	}))
}

func (server Server) GetChatHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strChatId := vars["chat_id"]
	uint64ChatId, err := strconv.ParseUint(strChatId, 10, 64)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, constants.ErrSessionExpired, r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, errors.New(constants.ErrSessionExpired), http.StatusBadRequest)
		return
	}

	grpcChatRequest := chatpc.GetChatRequest{
		ChatId: uint32(uint64ChatId),
		UserId: userId,
	}

	client, conn, err := server.InitClient()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	streamChat, err := client.GetChat(context.Background(), &grpcChatRequest)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	var messagesData []app.MessageData
	for {
		grpcMsgData, err := streamChat.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}

		msgData := app.GetMessageDataStruct(grpcMsgData)
		messagesData = append(messagesData, msgData)
	}

	logger.Log(http.StatusOK, constants.LogSuccess, r.Method, r.URL.Path, false)
	writer.Respond(w, r, structs.Map(app.GetChatResponse{
		Chat: messagesData,
	}))

	grpcChatParticipantsRequest := chatpc.GetChatParticipantsRequest{
		ChatId: uint32(uint64ChatId),
		UserId: userId,
	}

	grpcChatParticipantsResponse, err := client.GetChatParticipants(context.Background(), &grpcChatParticipantsRequest)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	userIdsUint32 := grpcChatParticipantsResponse.GetChatUserIds()
	userIds := make([]uint64, len(userIdsUint32))

	for idx, userIdUint32 := range userIdsUint32 {
		userIds[idx] = uint64(userIdUint32)
	}

	wsMsgData := app.WSMsgData{
		SenderId: uint64(userId),
		UserIds:  userIds,
		MsgData: app.WSReadDataStruct{
			ChatId: uint64ChatId,
		},
	}

	err = server.sendAll(wsMsgData)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), constants.LogPostMethod, constants.LogOnMessageHandler, true)
	}
}
