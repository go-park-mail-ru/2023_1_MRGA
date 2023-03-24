package app

import (
	"net/http"
)

type Application struct {
	Router *http.ServeMux
}

func New() *Application {
	router := http.NewServeMux()

	a := &Application{Router: router}

	return a
}
