package migration

import (
	"fmt"
	"log"
	"os"
	"scenario/internal/repo"
	"scenario/internal/setting"
)

func Migrate() error {

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

	if err := syncAllTables(); err != nil {
		return fmt.Errorf("sync database struct error: %v", err)
	}

	if err := seedDatabase(); err != nil {
		return fmt.Errorf("seed database: %v", err)
	}

	return nil
}

// syncAllTables sync the schemas of all tables, is required by unit test code
func syncAllTables() error {

	db := repo.GetDatabase()
	tables := repo.Tables()

	err := db.AutoMigrate(tables...)
	if err != nil {
		return fmt.Errorf("models sync: %v", err)
	}

	return nil
}

func seedDatabase() error {

	seedFuncs := repo.SeedFuncs()

	for _, seedFunc := range seedFuncs {
		if err := seedFunc(); err != nil {
			return fmt.Errorf("seed function failed: %v", err)
		}
	}

	return nil
}
