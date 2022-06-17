package api_test

import (
	"net/http"
	"scenario/internal/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {

	server := testutils.NewTestServer()

	w := server.Get("/ping")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
