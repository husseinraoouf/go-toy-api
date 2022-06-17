package testutils

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"scenario/internal/api"

	"github.com/go-chi/chi/v5"
)

type TestServer struct {
	router chi.Router
}

func NewTestServer() *TestServer {
	return &TestServer{
		router: api.Routes(),
	}
}

func (server *TestServer) Get(path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", path, nil)

	server.router.ServeHTTP(w, req)

	return w
}

func (server *TestServer) Post(path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", path, strings.NewReader(""))

	server.router.ServeHTTP(w, req)

	return w
}
