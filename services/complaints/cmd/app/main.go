package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/complaints/internal/app/dsn"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/complaints/internal/pkg/repository"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/complaints/internal/pkg/server"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaints"
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

	log.Println("Application terminate")
	s := grpc.NewServer()
	compRepo := repository.NewRepo(db)
	srv := server.NewGPRCServer(compRepo)
	complaints.RegisterComplaintsServer(s, srv)

	l, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("listener failed " + err.Error())
	}

	err = s.Serve(l)
	if err != nil {
		log.Fatalf("Serve failed" + err.Error())
	}
}
