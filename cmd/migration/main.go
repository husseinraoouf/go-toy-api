package main

import (
	"log"

	"scenario/internal/migration"
	"scenario/internal/repo"
	"scenario/internal/setting"
)

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

	if err := migration.Migrate(); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
