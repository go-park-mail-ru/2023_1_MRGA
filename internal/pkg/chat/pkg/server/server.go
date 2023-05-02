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
	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/proto_services/proto_chat"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (server *Server) InitRouter(pathPrefix string) {
	server.router = mux.NewRouter().PathPrefix(pathPrefix).Subrouter()

	// С префиксом /meetme/chats все ниже
	server.router.HandleFunc("/create", server.CreateChatHandler).Methods("POST")
	server.router.HandleFunc("/{chat_id}/send", server.SendMessageHandler).Methods("POST")
	server.router.HandleFunc("/list", server.GetChatsListHandler).Methods("GET")
	server.router.HandleFunc("/{chat_id}/messages", server.GetChatHandler).Methods("GET")
}

func (server Server) CreateChatHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var initialChatData app.CreateChatRequest
	err := json.NewDecoder(r.Body).Decode(&initialChatData)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	grpcInitialChatData := app.GetGRPCInitialChatData(initialChatData)

	client, conn, err := server.InitClient()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	var grpcCreatedChatData *chatpc.CreateChatResponse
	grpcCreatedChatData, err = client.CreateChat(context.Background(), grpcInitialChatData)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	createdChatData := app.GetCreatedChatDataStruct(grpcCreatedChatData)

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, structs.Map(createdChatData))
}

func (server Server) SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var msgData app.SendMessageRequest
	err := json.NewDecoder(r.Body).Decode(&msgData)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	strChatId := vars["chat_id"]
	uint64ChatId, err := strconv.ParseUint(strChatId, 10, 64)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(int)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, errors.New("Нет кук пользователя"), http.StatusBadRequest)
		return
	}

	msg := app.Message{
		SenderId:   uint(userId),
		Content:    msgData.Content,
		SentAt:     time.Now(),
		ReadStatus: false,
	}

	grpcMsg := app.GetGRPCChatMessage(msg, uint(uint64ChatId))

	client, conn, err := server.InitClient()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	_, err = client.SendMessage(context.Background(), grpcMsg)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, structs.Map(app.SendMessageResponse{
		SentAt: msg.SentAt.Format("15:04 02.01.2006"),
	}))
}

func (server Server) GetChatsListHandler(w http.ResponseWriter, r *http.Request) {
	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(int)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, errors.New("Нет кук пользователя"), http.StatusBadRequest)
		return
	}

	grpcChatsListRequest := chatpc.GetChatsListRequest{
		UserId: uint32(userId),
	}

	client, conn, err := server.InitClient()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	streamChatsList, err := client.GetChatsList(context.Background(), &grpcChatsListRequest)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
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
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}

		chatMsg := app.GetChatMessageStruct(grpcChatMsg)
		chatsMessages = append(chatsMessages, chatMsg)
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, structs.Map(app.GetChatsListResponse{
		ChatsList: chatsMessages,
	}))
}

func (server Server) GetChatHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strChatId := vars["chat_id"]
	uint64ChatId, err := strconv.ParseUint(strChatId, 10, 64)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(int)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, errors.New("Нет кук пользователя"), http.StatusBadRequest)
		return
	}

	grpcChatRequest := chatpc.GetChatRequest{
		ChatId: uint32(uint64ChatId),
		UserId: uint32(userId),
	}

	client, conn, err := server.InitClient()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	streamChat, err := client.GetChat(context.Background(), &grpcChatRequest)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
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
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}

		msgData := app.GetMessageDataStruct(grpcMsgData)
		messagesData = append(messagesData, msgData)
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, structs.Map(app.GetChatResponse{
		Chat: messagesData,
	}))
}
