package main

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"

	tracejaeger "github.com/go-park-mail-ru/2023_1_MRGA.git/utils/trace_jaeger"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/complaints/internal/app/dsn"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/complaints/internal/pkg/repository"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/complaints/internal/pkg/server"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaintProto"
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
		ServiceName:    "complaintsServer",
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

	log.Println("Application terminate")
	s := grpc.NewServer(grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))
	compRepo := repository.NewRepo(db)
	srv := server.NewGPRCServer(compRepo)
	complaintProto.RegisterComplaintsServer(s, srv)

	l, err := net.Listen("tcp", "0.0.0.0:8083")
	if err != nil {
		log.Fatalf("listener failed " + err.Error())
	}

	err = s.Serve(l)
	if err != nil {
		log.Fatalf("Serve failed" + err.Error())
	}
}
