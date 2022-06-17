package deck_test

import (
	"log"
	"math/rand"
	"os"
	"scenario/internal/migration"
	"scenario/internal/repo"
	"scenario/internal/setting"
	"testing"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func init() {
	httpin.UseGochiURLParam("path", chi.URLParam)
	rnd := rand.New(rand.NewSource(1))

	uuid.SetRand(rnd)
}

func TestMain(m *testing.M) {
	var err error

	if err = setting.LoadTestConfigAtPath("../../.."); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err = repo.InitDatabase(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err = migration.Migrate(); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	code := m.Run()

	os.Exit(code)
}
