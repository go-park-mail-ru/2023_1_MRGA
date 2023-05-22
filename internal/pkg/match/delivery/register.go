package delivery

import (
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/upgrader"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match"
)

type UserID uint64

const (
	NewMatch    = "new_match"
	MissedMatch = "missed_match"
)

type MatchNotification struct {
	Type string `json:"type"`
}

type Handler struct {
	useCase  match.UseCase
	upgrader websocket.Upgrader
	mutex    *sync.RWMutex
	//  ключ – id подписавшегося пользователя,
	// значение – словарь с уникальным значением для каждого соединения
	WebsocketClients map[UserID]map[uuid.UUID]*websocket.Conn
}

func NewHandler(useCase match.UseCase) *Handler {
	return &Handler{
		useCase:          useCase,
		upgrader:         upgrader.Upgrader,
		WebsocketClients: make(map[UserID]map[uuid.UUID]*websocket.Conn),
		mutex:            &sync.RWMutex{},
	}
}

func RegisterHTTPEndpoints(router *mux.Router, uc match.UseCase) {
	h := NewHandler(uc)

	router.HandleFunc("/meetme/match", h.GetMatches).Methods("GET")
	router.HandleFunc("/meetme/match/{userId}", h.DeleteMatch).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/meetme/reaction", h.AddReaction).Methods("POST")
	router.HandleFunc("/meetme/match/subscribe", h.Subscribe).Methods("GET")
	router.HandleFunc("/meetme/chat/{userId}", h.GetChatByUserId).Methods("GET")
}
