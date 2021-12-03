package main

import (
	"github.com/bmizerany/pat"
	"go_websocket/internal/handlers"
	"net/http"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	return mux
}
