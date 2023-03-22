package app

import (
	"net/http"

	"github.com/gorilla/mux"

	authDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/delivery"
	authUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/usecase"
	recDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation/delivery"
	recUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation/usecase"
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
	//handler.HandleFunc("/meetme/register", a.Register)
	//handler.HandleFunc("/meetme/login", a.Login)
	//handler.HandleFunc("/meetme/logout", a.Logout)
	//handler.HandleFunc("/meetme/cities", a.GetCities)
	//handler.HandleFunc("/meetme/user", a.GetCurrentUser)
	//handler.HandleFunc("/meetme/recommendations", a.GetRecommendations)

	return router
}
