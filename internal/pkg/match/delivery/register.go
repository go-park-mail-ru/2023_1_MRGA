package delivery

import (
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match"
)

type Handler struct {
	useCase match.UseCase
}

func NewHandler(useCase match.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func RegisterHTTPEndpoints(router *mux.Router, uc match.UseCase) {
	h := NewHandler(uc)

	router.HandleFunc("/meetme/match", h.GetMatches).Methods("GET")
	router.HandleFunc("/meetme/reaction", h.AddReaction).Methods("POST")

	router.HandleFunc("/meetme/chat/{userId}", h.GetChatByEmail).Methods("GET")
}
