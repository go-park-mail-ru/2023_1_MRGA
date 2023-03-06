package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/middleware"
	"github.com/gorilla/mux"
)

const frontendHost = "http://localhost:8080"

func (a *Application) StartServer(host, port string) {
	log.Println("Server start up")
	router := a.Router
	h := host + ":" + port
	server := &http.Server{
		Addr:    h,
		Handler: router,
	}

	handler := mux.NewRouter()

	handlerWithCorsMiddleware := middleware.CorsMiddleware(frontendHost, handler)
	router.Handle("/", handlerWithCorsMiddleware)
	handler.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Addr:", ":8080", "URL:", r.URL.String())
			Respond(w, r, Result{status: http.StatusOK, err: ""}, map[string]interface{}{})
		})
	handler.HandleFunc("/meetme/register", a.Register)
	handler.HandleFunc("/meetme/login", a.Login)
	handler.HandleFunc("/meetme/logout", a.Logout)
	handler.HandleFunc("/meetme/cities", a.GetCities)
	handler.HandleFunc("/meetme/user", a.GetCurrentUser)
	handler.HandleFunc("/meetme/recommendations", a.GetRecommendations)
	err := server.ListenAndServe()
	if err != nil {
		log.Println("ListenServer failed", err)
	}

	log.Println("Server down")
}
