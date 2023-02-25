package app

import (
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/repository"
)

type Application struct {
	Router *http.ServeMux
	repo   *repository.Repository
}

func New() *Application {
	router := http.NewServeMux()
	repo := repository.New()
	a := &Application{repo: repo, Router: router}

	return a
}
