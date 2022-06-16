package main

import (
	"log"
	"os"
	"scenario/internal/api"
	"scenario/internal/repo"
	"scenario/internal/setting"

	"net/http"

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
		os.Exit(1)
	}

	err = repo.InitDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		os.Exit(1)
	}

	r := api.Routes()
	// Listen and Server in 0.0.0.0:8080
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
		os.Exit(1)
	}
}
