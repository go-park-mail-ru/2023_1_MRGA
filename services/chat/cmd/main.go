package main

import (
	"log"

	repositoryPackage "github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/pkg/repository"
	serverPackage "github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/pkg/server"
	servicePackage "github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/pkg/service"
)

func main() {
	repo, err := repositoryPackage.InitRepository()
	if err != nil {
		log.Fatal(err.Error())
	}

	service := servicePackage.InitService()

	server := serverPackage.InitServer(repo, service)

	if err = server.RunServer(); err != nil {
		log.Fatal("Не удалось запустить сервер: " + err.Error())
	}
}
