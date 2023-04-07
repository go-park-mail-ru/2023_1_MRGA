package delivery

import (
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func RegisterHTTPEndpoints(router *mux.Router, uc auth.UseCase) {
	h := NewHandler(uc)

	router.HandleFunc("/api/register", h.Register)
	router.HandleFunc("/api/login", h.Login)
	router.HandleFunc("/meetme/user", h.GetCurrentUser)
	//router.HandleFunc("/meetme/user", h.GetCurrentUser)
	router.HandleFunc("/meetme/logout", h.Logout)

}
