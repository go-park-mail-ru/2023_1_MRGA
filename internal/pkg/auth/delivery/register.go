package delivery

import (
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth"
)

func RegisterHTTPEndpoints(router *http.ServeMux, uc auth.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/register", h.Register)
		authEndpoints.POST("/login", h.Login)
		authEndpoints.POST("/logout", h.Logout)
	}
}
