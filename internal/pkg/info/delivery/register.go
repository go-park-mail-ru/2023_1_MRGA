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
	router.HandleFunc("/api/hashtags", h.GetHashtags).Methods("GET")
	router.HandleFunc("/api/cities", h.GetCities).Methods("GET")
	router.HandleFunc("/api/zodiac", h.GetZodiac).Methods("GET")
	router.HandleFunc("/api/job", h.GetJobs).Methods("GET")
	router.HandleFunc("/api/education", h.GetEducation).Methods("GET")
	router.HandleFunc("/api/reason", h.GetReasons).Methods("GET")
	router.HandleFunc("/api/status", h.GetStatuses).Methods("GET")
}
