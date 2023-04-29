package delivery

import (
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaints"
)

type Handler struct {
	useCase     info_user.UseCase
	compService complaints.ComplaintsClient
}

func NewHandler(useCase info_user.UseCase, compServ complaints.ComplaintsClient) *Handler {
	return &Handler{
		useCase:     useCase,
		compService: compServ,
	}
}

func RegisterHTTPEndpoints(router *mux.Router, ic info_user.UseCase, compServ complaints.ComplaintsClient) {
	h := NewHandler(ic, compServ)
	router.HandleFunc("/meetme/info-user/{userId}", h.GetInfoById).Methods("GET")

	router.HandleFunc("/meetme/user", h.GetCurrentUser).Methods("GET")

	router.HandleFunc("/meetme/info-user", h.GetInfo).Methods("GET")
	router.HandleFunc("/meetme/info-user", h.ChangeInfo).Methods("PUT", "OPTIONS")
	router.HandleFunc("/meetme/info-user", h.CreateInfo).Methods("POST")

	router.HandleFunc("/meetme/hashtags-user", h.AddUserHashtags).Methods("POST")
	router.HandleFunc("/meetme/hashtags-user", h.GetUserHashtags).Methods("GET")
	router.HandleFunc("/meetme/hashtags-user", h.ChangeUserHashtags).Methods("PUT", "OPTIONS")

}
