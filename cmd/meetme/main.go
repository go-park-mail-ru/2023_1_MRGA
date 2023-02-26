package main

import (
	"log"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/app"
)

func main() {
	log.Println("Application is starting")
	a := app.New()
	a.StartServer()
	log.Println("Application terminate")
}
