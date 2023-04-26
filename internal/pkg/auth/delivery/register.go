package delivery

import (
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto"
)

type Handler struct {
	AuthService authProto.AuthClient
}

func NewHandler(authService authProto.AuthClient) *Handler {
	return &Handler{
		AuthService: authService,
	}
}

func RegisterHTTPEndpoints(router *mux.Router, authServ authProto.AuthClient) {
	h := NewHandler(authServ)

	router.HandleFunc("/api/register", h.Register).Methods("POST")
	router.HandleFunc("/api/login", h.Login).Methods("POST")
	router.HandleFunc("/meetme/user", h.ChangeUser).Methods("PUT")
	router.HandleFunc("/meetme/logout", h.Logout).Methods("POST")
}
