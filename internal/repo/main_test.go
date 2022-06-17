package repo_test

import (
	"math/rand"
	"os"
	"testing"

	"github.com/google/uuid"
)

func TestMain(m *testing.M) {
	//nolint:gosec // it is used for deterministic uuid
	rnd := rand.New(rand.NewSource(1))
	uuid.SetRand(rnd)

	code := m.Run()

	os.Exit(code)
}
