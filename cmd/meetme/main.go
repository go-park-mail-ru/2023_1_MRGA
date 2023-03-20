package main

import (
	"log"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/repository"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/app"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/app/server"
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
	log.Println("Application is starting")
	r := repository.NewRepo()
	a := app.New(r)

	serv := new(server.Server)
	opts := server.GetServerOptions()
	err := serv.Run(opts, a.InitRoutes())
	if err != nil {
		log.Fatalf("error occured while server starting: %v", err)
	}
	log.Println("Application terminate")
}
