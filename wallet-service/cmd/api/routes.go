package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

const uuidRegexp = `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}`

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Route("/api/v1", func(r chi.Router) {
		r.Get("/wallet/{id:"+uuidRegexp+"}", app.GetWallet)
		r.Post("/wallet", app.UpdateWallet)
	})

	return mux
}
