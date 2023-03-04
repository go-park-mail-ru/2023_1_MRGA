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

	router.HandleFunc("/meetme/register", a.Register)
	router.HandleFunc("/meetme/login", a.Login)
	router.HandleFunc("/meetme/logout", a.Logout)
	router.HandleFunc("/meetme/cities", a.GetCities)
	router.HandleFunc("/meetme/return500", a.Return500)
	router.HandleFunc("/meetme/user", a.GetCurrentUser)
	router.HandleFunc("/meetme/recommendations", a.GetRecommendations)
	err := server.ListenAndServe()
	if err != nil {
		log.Println("ListenServer failed")
	}

	log.Println("Server down")
}
