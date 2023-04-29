package delivery

import (
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaints"
)

type Handler struct {
	CompService complaints.ComplaintsClient
}

func NewHandler(compService complaints.ComplaintsClient) *Handler {
	return &Handler{
		CompService: compService,
	}
}

func RegisterHTTPEndpoints(router *mux.Router, compServ complaints.ComplaintsClient) {
	h := NewHandler(compServ)

	router.HandleFunc("/meetme/complain", h.Complain).Methods("POST")
}
