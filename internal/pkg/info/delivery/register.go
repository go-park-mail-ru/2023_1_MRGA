package delivery

import (
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info"
)

type Handler struct {
	useCase info.UseCase
}

func NewHandler(useCase info.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func RegisterHTTPEndpoints(router *mux.Router, ic info.UseCase) {
	h := NewHandler(ic)

	router.HandleFunc("/meetme/info", h.GetInfo).Methods("GET")
	router.HandleFunc("/meetme/info", h.ChangeInfo).Methods("PUT")
	router.HandleFunc("/meetme/info", h.CreateInfo).Methods("POST")

	///getters
	router.HandleFunc("/api/cities", h.GetCities).Methods("GET")
	router.HandleFunc("/api/zodiac", h.GetZodiac).Methods("GET")
	router.HandleFunc("/api/job", h.GetJobs).Methods("GET")
	router.HandleFunc("/api/education", h.GetEducation).Methods("GET")
}
