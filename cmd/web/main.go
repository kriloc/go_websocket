package main

import (
	"go_websocket/internal/handlers"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting web server on port 8080")
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()
	_ = http.ListenAndServe(":8080", mux)
}
