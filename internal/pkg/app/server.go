package app

import (
	"fmt"
	"net/http"
)

func (a *Application) StartServer() {
	router := a.Router
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("server.ListenAndServe: %s", err))
	}
}
