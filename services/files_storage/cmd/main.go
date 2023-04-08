package main

import (
	"log"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal/pkg/repository"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal/pkg/server"
)

func main() {
	err := repository.InitRepository()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = server.RunServer()
	if err != nil {
		log.Fatal("Не удалось запустить сервер: " + err.Error())
	}
}
