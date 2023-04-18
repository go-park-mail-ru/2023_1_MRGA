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
	router.HandleFunc("/api/hashtags", h.GetHashtags).Methods("GET")   //move
	router.HandleFunc("/api/cities", h.GetCities).Methods("GET")       //move
	router.HandleFunc("/api/zodiac", h.GetZodiac).Methods("GET")       //move
	router.HandleFunc("/api/job", h.GetJobs).Methods("GET")            //move
	router.HandleFunc("/api/education", h.GetEducation).Methods("GET") //move

}
