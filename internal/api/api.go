// Package classification Scenario API.
//
// REST API for Scenario Card Game
//
//     Schemes: http
//     Host: localhost:8080
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: ElHussein Abdelraouf<hussein@raoufs.me>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package api

import (
	"net/http"

	"scenario/internal/api/deck"
	"scenario/internal/context"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Routes is to get api routes.
func Routes() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// allow cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}))

	// Ping test
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		c := context.NewContext(w, r)

		c.Text(http.StatusOK, "pong")
	})

	router.Route("/deck", func(r chi.Router) {
		// Create a new Deck
		r.With(httpin.NewInput(deck.CreateDeckInput{})).Post("/", deck.CreateDeck)

		r.Route("/{id}", func(r chi.Router) {
			// Open a Deck
			r.With(httpin.NewInput(deck.OpenDeckInput{})).Get("/", deck.OpenDeck)

			// Draw Cards From Deck
			r.With(httpin.NewInput(deck.DrawCardInput{})).Post("/draw", deck.DrawCard)
		})
	})

	return router
}
