package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/middleware"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/chat/app"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/chat/app/constants"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func writeMessage(ws *websocket.Conn, flag string, body interface{}) error {
	response := app.WSResponse{
		Flag: flag,
		Body: body,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		return err
	}

	return ws.WriteMessage(websocket.TextMessage, responseJSON)
}

func sendToClients(clients []*websocket.Conn, flag string, msgData interface{}) (err error) {
	for _, receiverWsClient := range clients {
		if currErr := writeMessage(receiverWsClient, flag, msgData); currErr != nil {
			err = currErr
		}
	}
	return
}

func (server *Server) sendAll(wsMsgData app.WSMsgData) (err error) {
	for _, receiverUserId := range wsMsgData.UserIds {
		server.mutex.Lock()
		clientsByUser, found := server.userIdClients[receiverUserId]
		server.mutex.Unlock()
		if !found {
			continue
		}

		err = sendToClients(clientsByUser, wsMsgData.Flag, wsMsgData.MsgData)
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), constants.LogPostMethod, constants.LogOnMessageHandler, true)
			return
		}
	}

	err = sendToClients(server.userIdClients[wsMsgData.SenderId], wsMsgData.Flag, wsMsgData.MsgData)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), constants.LogPostMethod, constants.LogOnMessageHandler, true)
		return
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
	userIdDB := r.Context().Value(middleware.ContextUserKey)
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

	var input app.WSInput
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				logger.Log(http.StatusInternalServerError, err.Error(), constants.LogPostMethod, constants.LogWSClose, false)
			} else {
				logger.Log(http.StatusInternalServerError, err.Error(), constants.LogPostMethod, constants.LogOnMessageHandler, true)
			}
			return
		}

		err = json.Unmarshal(message, &input)
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), constants.LogPostMethod, constants.LogOnMessageHandler, true)
			continue
		}

		switch input.Flag {
		case "READ":
			var readData app.WSReadRequest
			err = json.Unmarshal(input.ReadData, &readData)
			if err != nil {
				logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
				continue
			}

			wsMsgData := app.WSMsgData{
				Flag:     "READ",
				SenderId: userId,
				UserIds:  readData.UserIds,
				MsgData: app.WSReadResponse{
					SenderId: userId,
					ChatId:   readData.ChatId,
				},
			}

			err = server.sendAll(wsMsgData)
			if err != nil {
				logger.Log(http.StatusInternalServerError, err.Error(), constants.LogPostMethod, constants.LogOnMessageHandler, true)
				continue
			}
		}
	}
}
