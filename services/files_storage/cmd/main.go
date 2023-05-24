package main

import (
	"flag"
	"log"
	"time"

	repositoryPackage "github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal/pkg/repository"
	serverPackage "github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal/pkg/server"
	servicePackage "github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal/pkg/service"
)

func getServerOptions() (opts serverPackage.ServerOptions) {
	opts.Host = "0.0.0.0"
	flag.StringVar(&opts.Port, "p", "8081", "set the server's port")
	flag.IntVar(&opts.MaxHeaderBytes, "m", 1, "set the server's max header bytes in MB")

	readTimeout := flag.Int64("rt", 10, "set the server's read timeout in seconds")
	writeTimout := flag.Int("wt", 10, "set the server's read timeout in seconds")

	flag.Parse()

	opts.MaxHeaderBytes = opts.MaxHeaderBytes << 20 // MB to Bytes
	opts.ReadTimeout = time.Duration(*readTimeout) * time.Second
	opts.WriteTimeout = time.Duration(*writeTimout) * time.Second

	return opts
}

func main() {
	repo, err := repositoryPackage.InitRepository()
	if err != nil {
		log.Fatal(err.Error())
	}

	service, err := servicePackage.InitService(repo)
	if err != nil {
		log.Fatal(err.Error())
	}

	server := serverPackage.InitServer(getServerOptions(), service)
	err = server.RunServer()
	if err != nil {
		log.Fatal("Не удалось запустить сервер: " + err.Error())
	}
}
