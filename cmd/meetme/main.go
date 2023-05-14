package main

import (
	"log"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/dsn"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/period_function"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/servicedefault"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/app"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/app/server"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/authProto"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaintProto"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
)

// @title MRGA
// @version 1.0
// @description Meetme backend documentation

// @contact.name API Support
// @contact.url mrga.com
// @contact.email mrga@mail.com

// @license.name AS IS (NO WARRANTY)

// @host 5.159.100.59
// @schemes http
// @BasePath /meetme/
func main() {
	logger.Init(servicedefault.NameService)
	log.Println("Application is starting")

	a := app.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to connect env" + err.Error())
	}
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect db" + err.Error())
	}

	go period_function.RunCronJobs(db)
	connAuth, err := grpc.Dial(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer connAuth.Close()
	authClient := authProto.NewAuthClient(connAuth)

	connComp, err := grpc.Dial(":8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connComp.Close()
	compClient := complaintProto.NewComplaintsClient(connComp)

	serv := new(server.Server)
	opts := server.GetServerOptions()
	a.InitRoutes(db, authClient, compClient)
	err = serv.Run(opts, a.Router)
	if err != nil {
		log.Fatalf("error occured while server starting: %v", err)
	}
	log.Println("Application terminate")
}
