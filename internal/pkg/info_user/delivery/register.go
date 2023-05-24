package delivery

import (
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaintProto"
)

type Handler struct {
	useCase     info_user.UseCase
	compService complaintProto.ComplaintsClient
}

func NewHandler(useCase info_user.UseCase, compServ complaintProto.ComplaintsClient) *Handler {
	return &Handler{
		useCase:     useCase,
		compService: compServ,
	}
}

func RegisterHTTPEndpoints(router *mux.Router, ic info_user.UseCase, compServ complaintProto.ComplaintsClient) {
	h := NewHandler(ic, compServ)
	router.HandleFunc("/api/auth/info-user/{userId}", h.GetInfoById).Methods("GET")

	router.HandleFunc("/api/auth/user", h.GetCurrentUser).Methods("GET")

	router.HandleFunc("/api/auth/info-user", h.GetInfo).Methods("GET")
	router.HandleFunc("/api/auth/info-user", h.ChangeInfo).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/auth/info-user", h.CreateInfo).Methods("POST")

	router.HandleFunc("/api/auth/hashtags-user", h.AddUserHashtags).Methods("POST")
	router.HandleFunc("/api/auth/hashtags-user", h.GetUserHashtags).Methods("GET")
	router.HandleFunc("/api/auth/hashtags-user", h.ChangeUserHashtags).Methods("PUT", "OPTIONS")

	router.HandleFunc("/api/auth/status-user", h.GetUserStatus).Methods("GET")
	router.HandleFunc("/api/auth/status-user", h.ChangeUserStatus).Methods("PUT", "OPTIONS")
}
