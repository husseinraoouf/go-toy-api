package main

import (
	"log"
	"os"
	"scenario/internal/migration"
)

func main() {

	if err := migration.Migrate(); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
		os.Exit(1)
	}

}
