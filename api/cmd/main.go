package main

import (
	"log"

	"github.com/arunvm/jsHarp/api/pkg/server"
)

func main() {
	log.Println("Starting Server")
	server.StartServer()

	log.Println("Server Shutdown")
}
