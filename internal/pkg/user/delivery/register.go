package delivery

import (
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/user"
)

type Handler struct {
	useCase user.UseCase
}

func NewHandler(useCase user.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func RegisterHTTPEndpoints(router *http.ServeMux, uc user.UseCase) {
	h := NewHandler(uc)

	router.HandleFunc("/meetme/user", h.GetCurrentUser)

}
