package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Post("/wallet", app.UpdateWallet)

	return mux
}
