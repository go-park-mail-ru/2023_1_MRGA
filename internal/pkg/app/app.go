package app

import (
	"github.com/gorilla/mux"
)

type Application struct {
	Router *mux.Router
}

func New() *Application {
	router := mux.NewRouter()

	a := &Application{Router: router}

	return a
}
