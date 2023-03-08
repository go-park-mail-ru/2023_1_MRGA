package app

import (
	"log"
	"net/http"
)

func (a *Application) StartServer(host, port string) {
	log.Println("Server start up")
	router := a.Router

	h := host + ":" + port
	server := &http.Server{
		Addr:    h,
		Handler: router,
	}

	router.HandleFunc("/meetme/register", a.Register)
	router.HandleFunc("/meetme/login", a.Login)
	router.HandleFunc("/meetme/logout", a.Logout)
	router.HandleFunc("/meetme/cities", a.GetCities)
	router.HandleFunc("/meetme/user", a.GetCurrentUser)
	router.HandleFunc("/meetme/recommendations", a.GetRecommendations)

	err := server.ListenAndServe()
	if err != nil {
		log.Println("ListenServer failed")
	}

	log.Println("Server down")
}
