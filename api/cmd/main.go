package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/arunvm/jsHarp/api/pkg/operations"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
)

func main() {

	// render.html is where the code is inserted and later
	// displayed in the iframe.
	// Creating the file if it does not exist
	if _, err := os.Stat("render.html"); err != nil {
		os.Create("render.html")
	}

	// Cleanup: Removes the 'render.html' when killing the server
	interruptSignalChan := make(chan os.Signal, 1)
	signal.Notify(interruptSignalChan, os.Interrupt)
	go func() {
		<-interruptSignalChan
		os.Remove("render.html")
		os.Exit(0)
	}()

	port := flag.String("port", "8000", "Port to listen on")
	flag.Parse()

	log.Printf("Server listening on port:%s", *port)

	routes := operations.InitRoutes()

	handler := cors.AllowAll().Handler(routes)

	http.Handle("/", handler)

	if err := http.ListenAndServe(":"+*port, handlers.LoggingHandler(os.Stdout, handler)); err != nil {
		log.Fatal("Error starting server:", err)
	}

}
