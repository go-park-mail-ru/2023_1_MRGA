package delivery

import (
	"net/http"

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

func RegisterHTTPEndpoints(router *http.ServeMux, uc auth.UseCase) {
	h := NewHandler(uc)

	router.HandleFunc("/meetme/register", h.Register)
	router.HandleFunc("/meetme/login", h.Login)
	//router.HandleFunc("/meetme/logout", h.Logout)

}
