package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/chat/app"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/chat/app/constants"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
	"github.com/gorilla/websocket"
)

func writeMessage(ws *websocket.Conn, flag string, body app.WSMessageResponse) error {
	response := app.WSSendResponse{
		Flag: flag,
		Body: body,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		return err
	}

	return ws.WriteMessage(websocket.TextMessage, responseJSON)
}

func sendToClients(clients []*websocket.Conn, senderId, chatId uint64, msg, sentAt string) (err error) {
	for _, receiverWsClient := range clients {
		msgData := app.WSMessageResponse{
			ChatId:   chatId,
			SenderId: senderId,
			Msg:      msg,
			SentAt:   sentAt,
		}

		if currErr := writeMessage(receiverWsClient, "SEND", msgData); currErr != nil {
			err = currErr
		}
	}
	return
}

func (server *Server) addClient(userId uint64, ws *websocket.Conn) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	server.userIdClients[userId] = append(server.userIdClients[userId], ws)
}

func (server *Server) removeClient(userId uint64, ws *websocket.Conn) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	clients := server.userIdClients[userId]
	for i, client := range clients {
		if client == ws {
			clients[i] = clients[len(clients)-1]
			clients = clients[:len(clients)-1]
			break
		}
	}

	if len(clients) > 0 {
		server.userIdClients[userId] = clients
	} else {
		delete(server.userIdClients, userId)
	}

	logger.Log(http.StatusOK, constants.LogSuccess, constants.LogPostMethod, constants.LogCloseHandler, false)
}

func (server Server) ConnectionHandler(w http.ResponseWriter, r *http.Request) {
	userIdDB := r.Context().Value("userId")
	gotUserId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, constants.ErrSessionExpired, r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, errors.New(constants.ErrSessionExpired), http.StatusBadRequest)
		return
	}
	userId := uint64(gotUserId)

	ws, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	defer ws.Close()

	server.addClient(userId, ws)
	defer server.removeClient(userId, ws)

	logger.Log(http.StatusOK, constants.LogSuccess, constants.LogGetMethod, constants.LogConnectionHandler, false)

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				logger.Log(http.StatusInternalServerError, err.Error(), constants.LogPostMethod, constants.LogWSClose, false)
			} else {
				logger.Log(http.StatusInternalServerError, err.Error(), constants.LogPostMethod, constants.LogOnMessageHandler, true)
			}
			return
		}
	}
}
