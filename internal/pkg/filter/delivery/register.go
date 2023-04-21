package delivery

import (
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/filter"
)

type Handler struct {
	useCase filter.UseCase
}

func NewHandler(useCase filter.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func RegisterHTTPEndpoints(router *mux.Router, uc filter.UseCase) {
	h := NewHandler(uc)
	router.HandleFunc("/meetme/filters", h.AddFilter).Methods("POST")
	router.HandleFunc("/meetme/filters", h.GetFilter).Methods("GET")
	router.HandleFunc("/meetme/filters", h.ChangeFilter).Methods("PUT")

}
