package app

import (
	"fmt"
	"log"
	"net/http"
)

func (a *Application) StartServer() {
	log.Println("Server start up")
	router := a.Router

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	router.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Addr:", ":8080", "URL:", r.URL.String())
		})

	router.HandleFunc("/register", a.Register)
	router.HandleFunc("/login", a.Login)
	router.HandleFunc("/logout", a.Logout)

	err := server.ListenAndServe()
	if err != nil {
		log.Println("ListenServer failed")
	}

	log.Println("Server down")
}
