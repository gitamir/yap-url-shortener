package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/gitamir/yap-url-shortener/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

type TestStorage struct {
	Urls map[string]string
}

func NewTestStorage() *TestStorage {
	return &TestStorage{
		Urls: map[string]string{},
	}
}

func (s *TestStorage) Set(_, _ string) {
	s.Urls["test"] = "Test"
}

func (s *TestStorage) Get(key string) (string, bool) {
	if key == "test" {
		return "Test", true
	} else {
		return "", false
	}
}

type TestGenerator struct{}

func NewTestGenerator() *TestGenerator {
	return &TestGenerator{}
}

func (s *TestGenerator) Generate() string {
	return "test"
}

func NewTestServer() *Server {
	store := NewTestStorage()
	keyGenerator := NewGenerator()
	return &Server{
		storage:      store,
		keyGenerator: keyGenerator,
		Config: config.Options{
			Host:         "localhost:8080",
			ResolvedHost: "http://localhost:8080",
		},
	}
}

func TestGetHandler(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		path         string
		expectedCode int
	}{
		{
			name:         "Invalid Request",
			method:       http.MethodGet,
			path:         "/",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Invalid Request",
			method:       http.MethodGet,
			path:         "Some",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Valid Request",
			method:       http.MethodGet,
			path:         "test",
			expectedCode: http.StatusTemporaryRedirect,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/{id}", nil)

			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("id", tt.path)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

			s := NewTestServer()

			s.GetFullURL(w, r, tt.path)

			assert.Equal(t, tt.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
		})
	}
}

func TestPostHandler(t *testing.T) {
	tests := []struct {
		name               string
		method             string
		path               string
		body               string
		expectedCode       int
		expectedBodyRegexp string
	}{
		{
			name:               "Invalid Path",
			method:             http.MethodPost,
			path:               "/invalid",
			body:               "http://practicum.ru",
			expectedBodyRegexp: `Invalid path\n`,
			expectedCode:       http.StatusBadRequest,
		},
		{
			name:               "Invalid Body",
			method:             http.MethodPost,
			path:               "/",
			body:               "",
			expectedBodyRegexp: `Invalid request\n`,
			expectedCode:       http.StatusBadRequest,
		},
		{
			name:               "Valid Request",
			method:             http.MethodPost,
			path:               "/",
			body:               "http://practicum.ru",
			expectedBodyRegexp: `^http\:\/\/localhost\:\d{1,8}\/\w{8}$`,
			expectedCode:       http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			r := httptest.NewRequest(tt.method, tt.path, strings.NewReader(tt.body))
			w := httptest.NewRecorder()

			server := NewTestServer()
			server.ShortenURL(w, r)

			assert.Equal(t, tt.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			assert.Regexp(t, regexp.MustCompile(tt.expectedBodyRegexp), w.Body.String(), "Тело ответа не совпадает с ожидаемым")
		})
	}
}
