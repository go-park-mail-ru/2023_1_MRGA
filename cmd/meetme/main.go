package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/repository"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/app"
)

// @title MRGA
// @version 1.0
// @description Meetme backend documentation

// @contact.name API Support
// @contact.url mrga.com
// @contact.email mrga@mail.com

// @license.name AS IS (NO WARRANTY)

// @host 127.0.0.1
// @schemes http
// @BasePath /meetme/
func main() {
	host := "localhost"
	port := "8080"
	if len(os.Args) == 3 {
		host = os.Args[1]
		port = os.Args[2]
	} else if len(os.Args) == 2 {
		host = os.Args[1]
	}
	fmt.Println(os.Args[0])
	log.Println("Application is starting")
	r := repository.NewRepo()
	a := app.New(r)
	a.StartServer(host, port)
	log.Println("Application terminate")
}
