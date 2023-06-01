package main

import (
	"context"
	"log"
	"os"
	"strconv"

	tracejaeger "github.com/go-park-mail-ru/2023_1_MRGA.git/utils/trace_jaeger"
	"github.com/joho/godotenv"

	repositoryPackage "github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/pkg/repository"
	serverPackage "github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/pkg/server"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
)

func main() {
	logger.Init("ChatService")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to connect env" + err.Error())
	}

	ctx := context.Background()

	tracingDisabled, err := strconv.ParseBool(os.Getenv("TRACING_DISABLED"))
	if err != nil {
		log.Fatal(err)
	}

	prv, err := tracejaeger.NewProvider(tracejaeger.ProviderConfig{
		JaegerEndpoint: "http://localhost:14268/api/traces",
		ServiceName:    "chatServer",
		Disabled:       tracingDisabled,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer prv.Close(ctx)

	repo, err := repositoryPackage.InitRepository()
	if err != nil {
		log.Fatal(err.Error())
	}

	serverOpts := serverPackage.ServerOptions{
		Port: 3030,
	}

	server := serverPackage.InitServer(serverOpts, repo)

	if err = server.RunServer(); err != nil {
		log.Fatal("Не удалось запустить сервер: " + err.Error())
	}
}
