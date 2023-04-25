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
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/chat/app"
	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/proto_services/proto_chat"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (server *Server) InitRouter(pathPrefix string) {
	server.router = mux.NewRouter().PathPrefix(pathPrefix).Subrouter()

	// С префиксом /chat все ниже
	server.router.HandleFunc("/send/message", server.SendMessage).Methods("POST")
	server.router.HandleFunc("/recent/messages/{user_id}", server.GetRecentMessages).Methods("GET")
	server.router.HandleFunc("/{single_user_id}/{other_user_id}/messages", server.GetConversationMessages).Methods("GET")
}

func (server *Server) SendMessage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var msg app.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	msg.SentAt = time.Now()
	grpcMsg := app.GetGRPCMessage(msg)

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
	writer.Respond(w, r, map[string]interface{}{})
}

func (server *Server) GetRecentMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	varUserId := vars["user_id"]
	if varUserId == "" {
		err := errors.New("Отсутствует userId в url")
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userId, err := strconv.ParseUint(varUserId, 10, 64)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	recentMessagesRequest := chatpc.ResentMessagesRequest{
		UserId: wrapperspb.UInt32(uint32(userId)),
	}

	client, conn, err := server.InitClient()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	stream, err := client.GetRecentMessages(context.Background(), &recentMessagesRequest)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	var messages []app.Message
	for {
		grpcMsg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}

		msg := app.GetStructMessage(grpcMsg)
		messages = append(messages, msg)
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, structs.Map(app.Messages{Messages: messages}))
}

func (server *Server) GetConversationMessages(w http.ResponseWriter, r *http.Request) {
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, map[string]interface{}{})
}
