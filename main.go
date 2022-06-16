package main

import (
	"log"
	"os"
	"scenario/context"
	"scenario/repo"
	"scenario/setting"

	"net/http"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	_ "github.com/lib/pq"
)

type H map[string]any

func setupRouter() chi.Router {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Ping test
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		c := context.NewContext(w, r)

		c.Text(http.StatusOK, "pong")
	})

	r.Route("/deck", func(r chi.Router) {
		// Create a new Deck
		r.With(httpin.NewInput(CreateDeckInput{})).Post("/", CreateDeck)

		r.Route("/{id}", func(r chi.Router) {
			// Open a Deck
			r.With(httpin.NewInput(OpenDeckInput{})).Get("/", OpenDeck)

			// Draw Cards From Deck
			r.With(httpin.NewInput(DrawCardInput{})).Post("/draw", DrawCard)
		})
	})

	return r
}

func init() {
	httpin.UseGochiURLParam("path", chi.URLParam)
}

func main() {
	var err error

	err = setting.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
		os.Exit(1)
	}

	err = repo.InitDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		os.Exit(1)
	}

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
		os.Exit(1)
	}
}
