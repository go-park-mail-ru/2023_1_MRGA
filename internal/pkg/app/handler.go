package app

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"gorm.io/gorm"

	authDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/delivery"
	AuthRepository "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/repository"
	authUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/usecase"
	//recDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation/delivery"
	//RecRepository "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation/repository"
	//recUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation/usecase"
	//userDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/user/delivery"
	//userRepository "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/user/repository"
	//userUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/user/usecase"
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

func (a *Application) InitRoutes(db *gorm.DB, client *redis.Client) *http.ServeMux {
	router := a.Router

	handler := mux.NewRouter()

	handlerWithCorsMiddleware := middleware.CorsMiddleware(frontendHosts, handler)
	router.Handle("/", handlerWithCorsMiddleware)

	authRepo := AuthRepository.NewRepo(db, client)
	ucAuth := authUC.NewAuthUseCase(authRepo, "0123", 1233)
	authDel.RegisterHTTPEndpoints(a.Router, ucAuth)

	//recRepo := RecRepository.NewRepo(db)
	//ucRec := recUC.NewRecUseCase(recRepo)
	//recDel.RegisterHTTPEndpoints(a.Router, ucRec)

	//userRepo := userRepository.NewRepo(db)
	//ucUser := userUC.NewUserUseCase(userRepo)
	//userDel.RegisterHTTPEndpoints(a.Router, ucUser)

	return router
}
