package main

import (
	"log"

	repositoryPackage "github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/pkg/repository"
	serverPackage "github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/pkg/server"
)

func main() {
	repo, err := repositoryPackage.InitRepository()
	if err != nil {
		log.Fatal(err.Error())
	}

	serverOpts := serverPackage.ServerOptions{
		Port:       3030,
	}

	server := serverPackage.InitServer(serverOpts, repo)
	
	if err = server.RunServer(); err != nil {
		log.Fatal("Не удалось запустить сервер: " + err.Error())
	}
}
