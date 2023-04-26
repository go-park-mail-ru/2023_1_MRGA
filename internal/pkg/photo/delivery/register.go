package delivery

import (
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/photo"
)

type Handler struct {
	useCase photo.UseCase
}

func NewHandler(useCase photo.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func RegisterHTTPEndpoints(router *mux.Router, uc photo.UseCase) {
	h := NewHandler(uc)
	router.HandleFunc("/meetme/photos/upload", h.AddPhoto).Methods("POST")
	router.HandleFunc("/meetme/photo/{photo}", h.GetPhoto).Methods("GET")
	router.HandleFunc("/meetme/photo/{photo}", h.DeletePhoto).Methods("DELETE")
	router.HandleFunc("/meetme/photo/{photo}", h.ChangePhoto).Methods("PUT")
}
