package main

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"

	tracejaeger "github.com/go-park-mail-ru/2023_1_MRGA.git/utils/trace_jaeger"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/auth/internal/app/dsn"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/auth/internal/pkg/repository"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/auth/internal/pkg/server"
	api "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/authProto"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/env_getter"
)

func main() {
	log.Println("Application is starting")

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
		JaegerEndpoint: "http://meetme-app.ru:14268/api/traces",
		ServiceName:    "authServer",
		Disabled:       tracingDisabled,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer prv.Close(ctx)

	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect db" + err.Error())
	}
	redisHost := env_getter.GetHostFromEnv("REDIS_HOST")

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":6379",
		Password: "password",
		DB:       0,
	})

	_, err = client.Ping().Result()
	if err != nil {
		log.Fatalf("failed to connect redis" + err.Error())
	}

	log.Println("Application terminate")
	s := grpc.NewServer(grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))

	authRepo := repository.NewRepo(db, client)
	srv := server.NewGPRCServer(authRepo)
	api.RegisterAuthServer(s, srv)

	l, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("listener failed " + err.Error())
	}

	err = s.Serve(l)
	if err != nil {
		log.Fatalf("Serve failed" + err.Error())
	}
}
