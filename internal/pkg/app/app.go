package app

import (
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/dsn"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/repository"
)

type Application struct {
	Router *http.ServeMux
	repo   *repository.Repository
}

func main() {
	router := http.NewServeMux()
	dsnStr := dsn.FromEnv()
	repo, err := repository.New(dsnStr)
}
