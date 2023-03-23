package app

import (
	"net/http"

	"github.com/gorilla/mux"

	authDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/delivery"
	authUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/usecase"
	recDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation/delivery"
	recUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation/usecase"
	userDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/user/delivery"
	userUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/user/usecase"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/middleware"
)

var frontendHosts = []string{
	"http://localhost:8080",
	"http://localhost:3000",
	"http://5.159.100.59:3000",
	"http://5.159.100.59:8080",
	"http://192.168.0.2:3000",
	"http://192.168.0.2:8080",
}

func (a *Application) InitRoutes() *http.ServeMux {
	router := a.Router

	handler := mux.NewRouter()

	handlerWithCorsMiddleware := middleware.CorsMiddleware(frontendHosts, handler)
	router.Handle("/", handlerWithCorsMiddleware)

	ucAuth := authUC.NewAuthUseCase(a.repo, "0123", []byte("0123"), 1233)
	authDel.RegisterHTTPEndpoints(a.Router, ucAuth)

	ucRec := recUC.NewRecUseCase(a.repo)
	recDel.RegisterHTTPEndpoints(a.Router, ucRec)

	ucUser := userUC.NewUserUseCase(a.repo)
	userDel.RegisterHTTPEndpoints(a.Router, ucUser)

	return router
}
