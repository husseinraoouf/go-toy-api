package context_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"scenario/internal/context"

	"github.com/stretchr/testify/assert"
)

func TestText(t *testing.T) {
	w := httptest.NewRecorder()

	ctx := context.NewContext(w, nil)

	ctx.Text(http.StatusCreated, "pong")

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "pong", strings.TrimSpace(w.Body.String()))
}

func TestJSON(t *testing.T) {
	w := httptest.NewRecorder()

	ctx := context.NewContext(w, nil)

	ctx.JSON(http.StatusCreated, map[string]string{"ping": "pong"})

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, `{"ping":"pong"}`, strings.TrimSpace(w.Body.String()))
}

func TestError(t *testing.T) {
	w := httptest.NewRecorder()

	ctx := context.NewContext(w, nil)

	ctx.Error(http.StatusInternalServerError, errors.New("internal server error"))

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, `{"message":"internal server error"}`, strings.TrimSpace(w.Body.String()))
}

func TestErrorString(t *testing.T) {
	w := httptest.NewRecorder()

	ctx := context.NewContext(w, nil)

	ctx.Error(http.StatusInternalServerError, "internal server error")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, `{"message":"internal server error"}`, strings.TrimSpace(w.Body.String()))
}
