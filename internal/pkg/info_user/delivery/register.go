package delivery

import (
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user"
)

type Handler struct {
	useCase info_user.UseCase
}

func NewHandler(useCase info_user.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func RegisterHTTPEndpoints(router *mux.Router, ic info_user.UseCase) {
	h := NewHandler(ic)

	router.HandleFunc("/meetme/info_user", h.GetInfo).Methods("GET")
	router.HandleFunc("/meetme/info_user", h.ChangeInfo).Methods("PUT")
	router.HandleFunc("/meetme/info_user", h.CreateInfo).Methods("POST")

	///getters
	router.HandleFunc("/api/cities", h.GetCities).Methods("GET")
	router.HandleFunc("/api/zodiac", h.GetZodiac).Methods("GET")
	router.HandleFunc("/api/job", h.GetJobs).Methods("GET")
	router.HandleFunc("/api/education", h.GetEducation).Methods("GET")
}
