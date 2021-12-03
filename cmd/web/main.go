package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting web server on port 8080")
	mux := routes()
	_ = http.ListenAndServe(":8080", mux)
}
