package app

import (
	"net/http"

	"github.com/go-redis/redis"
	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/middleware"
	authDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/delivery"
	AuthRepository "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/repository"
	authUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/usecase"
	infoDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user/delivery"
	InfoRepository "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user/repository"
	infoUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user/usecase"
	recDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation/delivery"
	RecRepository "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation/repository"
	recUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation/usecase"
)

var frontendHosts = []string{
	"http://localhost:8080",
	"http://localhost:3000",
	"http://5.159.100.59:3000",
	"http://5.159.100.59:8080",
	"http://192.168.0.2:3000",
	"http://192.168.0.2:8080",
	"",
}

func (a *Application) InitRoutes(db *gorm.DB, client *redis.Client) {

	a.Router.Use(func(h http.Handler) http.Handler {
		return middleware.CorsMiddleware(frontendHosts, h)
	})

	a.Router.Use(func(h http.Handler) http.Handler {
		return middleware.AuthMiddleware(client, h)
	})
	authRepo := AuthRepository.NewRepo(db, client)
	ucAuth := authUC.NewAuthUseCase(authRepo, "0123", 1233)
	authDel.RegisterHTTPEndpoints(a.Router, ucAuth)

	infoRepo := InfoRepository.NewInfoRepo(db)
	ucInfo := infoUC.NewInfoUseCase(infoRepo)
	infoDel.RegisterHTTPEndpoints(a.Router, ucInfo)

	recRepo := RecRepository.NewRepo(db)
	ucRec := recUC.NewRecUseCase(recRepo)
	recDel.RegisterHTTPEndpoints(a.Router, ucRec)

	//userRepo := userRepository.NewRepo(db)
	//ucUser := userUC.NewUserUseCase(userRepo)
	//userDel.RegisterHTTPEndpoints(a.Router, ucUser)

}