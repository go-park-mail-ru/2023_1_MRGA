package delivery

import (
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation"
)

type Handler struct {
	useCase recommendation.UseCase
}

func NewHandler(useCase recommendation.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func RegisterHTTPEndpoints(router *http.ServeMux, uc recommendation.UseCase) {
	h := NewHandler(uc)

	router.HandleFunc("/meetme/recommendation", h.GetRecommendations)

}
