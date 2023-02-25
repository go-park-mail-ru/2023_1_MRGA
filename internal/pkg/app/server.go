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

	router.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Addr:", ":8080", "URL:", r.URL.String())
		})

	router.HandleFunc("/register", a.register)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("server.ListenAndServe: %s", err))
	}
}
