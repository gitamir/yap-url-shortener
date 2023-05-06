package main

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

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
			path:         "/Some",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Valid Request",
			method:       http.MethodGet,
			path:         "/test",
			expectedCode: http.StatusTemporaryRedirect,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()
			s := NewTestStorage()

			GetHandler(w, r, s)

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
			method:             http.MethodGet,
			path:               "/invalid",
			body:               "http://practicum.ru",
			expectedBodyRegexp: `Invalid path\n`,
			expectedCode:       http.StatusBadRequest,
		},
		{
			name:               "Invalid Body",
			method:             http.MethodGet,
			path:               "/",
			body:               "",
			expectedBodyRegexp: `Invalid request\n`,
			expectedCode:       http.StatusBadRequest,
		},
		{
			name:               "Valid Request",
			method:             http.MethodGet,
			path:               "/",
			body:               "http://practicum.ru",
			expectedBodyRegexp: `^http\:\/\/localhost\:\d{1,8}\/\w{8}$`,
			expectedCode:       http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.path, strings.NewReader(tt.body))
			w := httptest.NewRecorder()
			s := NewTestStorage()
			k := NewGenerator(s)

			PostHandler(w, r, s, k)

			assert.Equal(t, tt.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			assert.Regexp(t, regexp.MustCompile(tt.expectedBodyRegexp), w.Body.String(), "Тело ответа не совпадает с ожидаемым")
		})
	}
}
