package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/rs/cors"
)

func StartServer() {
	routes := initRoutes()

	handler := cors.AllowAll().Handler(routes)

	http.Handle("/", handler)

	if err := http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, handler)); err != nil {
		log.Fatal("Shutting Down", err)
	}
}
