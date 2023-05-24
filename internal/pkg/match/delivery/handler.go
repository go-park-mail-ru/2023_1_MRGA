package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"


	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/middleware"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/chat/app/constants"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (h *Handler) GetMatches(w http.ResponseWriter, r *http.Request) {
	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	matches, err := h.useCase.GetMatches(uint(userId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	result := make(map[string]interface{})
	result["matches"] = matches

	logger.Log(http.StatusOK, "give user information", r.Method, r.URL.Path, false)
	writer.Respond(w, r, result)
}

func (h *Handler) AddReaction(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()

	var reaction match.ReactionInp
	err := json.NewDecoder(r.Body).Decode(&reaction)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value(middleware.ContextUserKey)
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	reactionResult, err := h.useCase.PostReaction(uint(userId), reaction)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	if reactionResult.ResultCode == match.NewMatch {
		notification := MatchNotification{Type: NewMatch}
		h.NotifyClients(UserID(userId), notification)
		h.NotifyClients(UserID(reaction.EvaluatedUserId), notification)
	}
	if reactionResult.ResultCode == match.MissedMatch {
		notification := MatchNotification{Type: MissedMatch}
		h.NotifyClients(UserID(userId), notification)
	}
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, map[string]interface{}{})
}

func (h *Handler) DeleteMatch(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	matchUserIdStr := params["userId"]
	matchUserId, err := strconv.Atoi(matchUserIdStr)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	err = h.useCase.DeleteMatch(uint(userId), uint(matchUserId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, map[string]interface{}{})
}

func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	userIDDB := r.Context().Value(middleware.ContextUserKey)
	gotUserID, ok := userIDDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, constants.ErrSessionExpired, r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, errors.New(constants.ErrSessionExpired), http.StatusBadRequest)
		return
	}

	wsConnection, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userID := UserID(gotUserID)
	go h.handleWebsocketConnection(wsConnection, userID)

}

func (h *Handler) addClient(userID UserID, connectionID uuid.UUID, ws *websocket.Conn) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	// Если map для данного пользователя еще не существует, создаем ее
	if _, ok := h.WebsocketClients[userID]; !ok {
		h.WebsocketClients[userID] = make(map[uuid.UUID]*websocket.Conn)
	}

	h.WebsocketClients[userID][connectionID] = ws
}

func (h *Handler) removeClient(userID UserID, connectionID uuid.UUID) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	if connections, ok := h.WebsocketClients[userID]; ok {
		if _, ok := connections[connectionID]; ok {
			// Удаление соединения
			delete(connections, connectionID)

			// Если это было последнее соединение, удалить map соединений пользователя
			if len(connections) == 0 {
				delete(h.WebsocketClients, userID)
			}
		}
	}
}

func (h *Handler) handleWebsocketConnection(ws *websocket.Conn, userID UserID) {
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			logger.Log(http.StatusBadRequest, err.Error(), "WEBSOCKET", "/api/auth/match/subscribe", true)
		}
	}(ws)

	connectionID := uuid.New()
	h.addClient(userID, connectionID, ws)
	defer h.removeClient(userID, connectionID)

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

func (h *Handler) NotifyClients(id UserID, notification MatchNotification) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	if clients, ok := h.WebsocketClients[id]; ok {
		for connectionID, client := range clients {
			err := client.WriteJSON(notification)
			if err != nil {
				logger.Log(http.StatusInternalServerError, fmt.Sprintf("couldn't write notification message: %v, userID: %v, connectionID: %v", err, id, connectionID), "", "", true)
			}
		}
	}
}
