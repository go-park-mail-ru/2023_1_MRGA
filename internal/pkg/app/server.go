package app

import (
	"fmt"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const frontendHost = "http://localhost:8081"

func (a *Application) StartServer() {
	log.Println("Server start up")
	router := a.Router

	server := &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: router,
	}

	handler := mux.NewRouter()

	handlerWithCorsMiddleware := middleware.CorsMiddleware(frontendHost, handler)
	router.Handle("/", handlerWithCorsMiddleware)
	handler.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Addr:", ":8080", "URL:", r.URL.String())
		})
	handler.HandleFunc("/meetme/register", a.Register)
	handler.HandleFunc("/meetme/login", a.Login)
	handler.HandleFunc("/meetme/logout", a.Logout)
	handler.HandleFunc("/meetme/cities", a.GetCities)
	handler.HandleFunc("/meetme/user", a.GetCurrentUser)
	handler.HandleFunc("/meetme/recommendations", a.GetRecommendations)
	err := server.ListenAndServe()
	if err != nil {
		log.Println("ListenServer failed")
	}

	log.Println("Server down")
}
