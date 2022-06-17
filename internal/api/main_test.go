package api_test

import (
	"log"
	"os"
	"scenario/internal/repo"
	"scenario/internal/setting"
	"testing"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
)

func init() {
	httpin.UseGochiURLParam("path", chi.URLParam)
}

func TestMain(m *testing.M) {
	var err error

	err = setting.LoadTestConfigAtPath("../..")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	err = repo.InitDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	code := m.Run()

	os.Exit(code)
}
