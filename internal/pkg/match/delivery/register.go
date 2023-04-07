package delivery

import (
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match"
)

type Handler struct {
	useCase match.UseCase
}

func NewHandler(useCase match.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func RegisterHTTPEndpoints(router *http.ServeMux, uc match.UseCase) {
	h := NewHandler(uc)

	router.HandleFunc("/meetme/match", h.GetMatches)

}
