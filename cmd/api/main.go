package main

import (
	"log"
	"net/http"

	"scenario/internal/api"
	"scenario/internal/repo"
	"scenario/internal/setting"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
)

func init() {
	httpin.UseGochiURLParam("path", chi.URLParam)
}

func main() {
	var err error

	err = setting.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	err = repo.InitDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	r := api.Routes()

	log.Println("Starting server on port 8080")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
