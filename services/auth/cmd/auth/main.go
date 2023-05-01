package main

import (
	"log"
	"net"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/auth/internal/app/dsn"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/auth/internal/pkg/repository"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/auth/internal/pkg/server"
<<<<<<< HEAD
	api "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/auth"
=======
	api "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/authProto"
>>>>>>> 244f467 (add logger)
)

func main() {
	log.Println("Application is starting")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to connect env" + err.Error())
	}

	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect db" + err.Error())
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "password",
		DB:       0,
	})

	_, err = client.Ping().Result()
	if err != nil {
		log.Fatalf("failed to connect redis" + err.Error())
	}

	log.Println("Application terminate")
	s := grpc.NewServer()
	authRepo := repository.NewRepo(db, client)
	srv := server.NewGPRCServer(authRepo)
	api.RegisterAuthServer(s, srv)

	l, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("listener failed " + err.Error())
	}

	err = s.Serve(l)
	if err != nil {
		log.Fatalf("Serve failed" + err.Error())
	}
}
