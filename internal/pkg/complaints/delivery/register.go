package delivery

import (
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaintProto"
)

type Handler struct {
	CompService complaintProto.ComplaintsClient
}

func NewHandler(compService complaintProto.ComplaintsClient) *Handler {
	return &Handler{
		CompService: compService,
	}
}

func RegisterHTTPEndpoints(router *mux.Router, compServ complaintProto.ComplaintsClient) {
	h := NewHandler(compServ)

	router.HandleFunc("/api/auth/complain", h.Complain).Methods("POST")
}
