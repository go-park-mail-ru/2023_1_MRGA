package main

import (
	"log"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/repository"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/app"
)

func main() {
	log.Println("Application is starting")
	r := repository.NewRepo()
	a := app.New(r)
	a.StartServer()
	log.Println("Application terminate")
}
