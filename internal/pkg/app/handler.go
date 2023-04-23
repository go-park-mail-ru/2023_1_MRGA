package app

import (
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/middleware"
	authDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/delivery"
	filterDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/filter/delivery"
	FilterRepository "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/filter/repository"
	filterUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/filter/usecase"
	InfoDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info/delivery"
	InfoRepository "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info/repository"
	infoUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info/usecase"
	infoUserDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user/delivery"
	InfoUserRepository "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user/repository"
	infoUserUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user/usecase"
	matchDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match/delivery"
	MatchRepository "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match/repository"
	matchUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match/usecase"
	photoDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/photo/delivery"
	PhotoRepository "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/photo/repository"
	photoUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/photo/usecase"
	recDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation/delivery"
	RecRepository "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation/repository"
	recUC "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation/usecase"
	auth "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto"
)

var frontendHosts = []string{
	"http://localhost:8080",
	"http://localhost:3000",
	"http://5.159.100.59:3000",
	"http://5.159.100.59:8080",
	"http://192.168.0.2:3000",
	"http://192.168.0.2:8080",
	"http://5.159.100.59:8080",
	"http://192.168.0.45:3000",
	"http://95.163.180.8:3000",
}

func (a *Application) InitRoutes(db *gorm.DB, client *redis.Client) {

	a.Router.Use(func(h http.Handler) http.Handler {
		return middleware.CorsMiddleware(frontendHosts, h)
	})

	a.Router.Use(func(h http.Handler) http.Handler {
		return middleware.AuthMiddleware(client, h)
	})

	conn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()

	authClient := auth.NewAuthClient(conn)

	photoRepo := PhotoRepository.NewPhotoRepo(db)
	ucPhoto := photoUC.NewPhotoUseCase(photoRepo)
	photoDel.RegisterHTTPEndpoints(a.Router, ucPhoto)

	infoRepo := InfoRepository.NewInfoRepo(db)
	ucInfo := infoUC.NewInfoUseCase(infoRepo)
	InfoDel.RegisterHTTPEndpoints(a.Router, ucInfo)

	infoUserRepo := InfoUserRepository.NewInfoRepo(db)
	ucUser := infoUserUC.NewInfoUseCase(infoUserRepo, ucInfo, ucPhoto)
	infoUserDel.RegisterHTTPEndpoints(a.Router, ucUser)

	filterRepo := FilterRepository.NewFilterRepo(db)
	ucFilter := filterUC.NewFilterUseCase(filterRepo, ucInfo, ucUser)
	filterDel.RegisterHTTPEndpoints(a.Router, ucFilter)

	recRepo := RecRepository.NewRepo(db)
	ucRec := recUC.NewRecUseCase(recRepo, ucFilter, ucPhoto, ucUser)
	recDel.RegisterHTTPEndpoints(a.Router, ucRec)

	matchRepo := MatchRepository.NewMatchRepo(db)
	ucMatch := matchUC.NewMatchUseCase(matchRepo)
	matchDel.RegisterHTTPEndpoints(a.Router, ucMatch)

	authDel.RegisterHTTPEndpoints(a.Router, authClient)

}
