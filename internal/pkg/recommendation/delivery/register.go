package delivery

import (
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation"
)

type Handler struct {
	useCase recommendation.UseCase
}

func NewHandler(useCase recommendation.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func RegisterHTTPEndpoints(router *mux.Router, uc recommendation.UseCase) {
	h := NewHandler(uc)

	router.HandleFunc("/meetme/recommendation", h.GetRecommendations).Methods("GET")
	router.HandleFunc("/meetme/filters", h.AddFilter).Methods("POST")
	router.HandleFunc("/meetme/filters", h.GetFilter).Methods("GET")
	router.HandleFunc("/meetme/filters", h.ChangeFilter).Methods("PUT")
	router.HandleFunc("/api/reason", h.GetReasons).Methods("GET")
}
