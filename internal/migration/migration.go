package migration

import (
	"fmt"

	"scenario/internal/repo"
)

// Migrate syncs the schema and seed the database.
func Migrate() error {
	if err := syncAllTables(); err != nil {
		return fmt.Errorf("sync database struct error: %v", err)
	}

	if err := seedDatabase(); err != nil {
		return fmt.Errorf("seed database: %v", err)
	}

	return nil
}

// syncAllTables syncs the schemas of all tables.
func syncAllTables() error {
	db := repo.GetDatabase()
	tables := repo.Tables()

	err := db.AutoMigrate(tables...)
	if err != nil {
		return fmt.Errorf("models sync: %v", err)
	}

	return nil
}

// seedDatabase seeds the database with init data.
func seedDatabase() error {
	seedFuncs := repo.SeedFuncs()

	for _, seedFunc := range seedFuncs {
		if err := seedFunc(); err != nil {
			return fmt.Errorf("seed function failed: %v", err)
		}
	}

	return nil
}
