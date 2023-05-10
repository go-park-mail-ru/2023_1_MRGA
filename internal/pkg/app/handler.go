package app

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/middleware"
	authDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/delivery"
	ChatServerPackage "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/chat/pkg/server"
	compDel "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/complaints/delivery"
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

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/authProto"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaintProto"
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
	"http://meetme-app.ru:3000",
	"http://meetme-app.ru:80",
	"http://meetme-app.ru",
	"http://localhost",
	"http://localhost:8080",
	"http://localhost:80",
	"https://localhost:8080",
	"https://localhost:3000",
	"https://5.159.100.59:3000",
	"https://5.159.100.59:8080",
	"https://192.168.0.2:3000",
	"https://192.168.0.2:8080",
	"https://5.159.100.59:8080",
	"https://192.168.0.45:3000",
	"https://95.163.180.8:3000",
	"https://meetme-app.ru:3000",
	"https://meetme-app.ru:80",
	"https://meetme-app.ru",
	"https://localhost",
	"https://localhost:8080",
	"https://localhost:80",
}

func (a *Application) InitRoutes(db *gorm.DB, authServ authProto.AuthClient, compServ complaintProto.ComplaintsClient) {

	a.Router.Use(func(h http.Handler) http.Handler {
		return middleware.CorsMiddleware(frontendHosts, h)
	})

	a.Router.Use(func(h http.Handler) http.Handler {
		return middleware.AuthMiddleware(authServ, h)
	})

	photoRepo := PhotoRepository.NewPhotoRepo(db)
	ucPhoto := photoUC.NewPhotoUseCase(photoRepo)
	photoDel.RegisterHTTPEndpoints(a.Router, ucPhoto)

	infoRepo := InfoRepository.NewInfoRepo(db)
	ucInfo := infoUC.NewInfoUseCase(infoRepo)
	InfoDel.RegisterHTTPEndpoints(a.Router, ucInfo)

	infoUserRepo := InfoUserRepository.NewInfoRepo(db)
	ucUser := infoUserUC.NewInfoUseCase(infoUserRepo, ucInfo, ucPhoto)
	infoUserDel.RegisterHTTPEndpoints(a.Router, ucUser, compServ)

	filterRepo := FilterRepository.NewFilterRepo(db)
	ucFilter := filterUC.NewFilterUseCase(filterRepo, ucInfo, ucUser)
	filterDel.RegisterHTTPEndpoints(a.Router, ucFilter)

	recRepo := RecRepository.NewRepo(db)
	ucRec := recUC.NewRecUseCase(recRepo, ucFilter, ucPhoto, ucUser)
	recDel.RegisterHTTPEndpoints(a.Router, ucRec)

	matchRepo := MatchRepository.NewMatchRepo(db)
	ucMatch := matchUC.NewMatchUseCase(matchRepo)
	matchDel.RegisterHTTPEndpoints(a.Router, ucMatch)

	authDel.RegisterHTTPEndpoints(a.Router, authServ)
	compDel.RegisterHTTPEndpoints(a.Router, compServ)

	chatServerOptions := ChatServerPackage.ServerOptions{
		Addr:       "chat-service",
		Port:       3030,
		PathPrefix: "/meetme/chats",
	}
	chatRouter := ChatServerPackage.InitServer(chatServerOptions)

	a.Router.PathPrefix(chatServerOptions.PathPrefix).Handler(chatRouter)
}
