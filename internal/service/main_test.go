package service_test

import (
	"log"
	"math/rand"
	"os"
	"testing"

	"scenario/internal/migration"
	"scenario/internal/repo"
	"scenario/internal/setting"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func init() {
	httpin.UseGochiURLParam("path", chi.URLParam)
}

func TestMain(m *testing.M) {
	var err error

	if err = setting.LoadTestConfigAtPath("../.."); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err = repo.InitDatabase(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err = migration.Migrate(); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	//nolint:gosec // it is used for deterministic uuid
	rnd := rand.New(rand.NewSource(1))
	uuid.SetRand(rnd)

	code := m.Run()

	if err = repo.ResetDatabase(); err != nil {
		log.Fatalf("failed to reset database: %v", err)
	}

	os.Exit(code)
}
