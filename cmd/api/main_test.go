package main_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"scenario/internal/api"
	"scenario/internal/repo"
	"scenario/internal/setting"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	var err error

	err = setting.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	err = repo.InitDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	router := api.Routes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
