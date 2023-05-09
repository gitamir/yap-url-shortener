package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gitamir/yap-url-shortener/cmd/config"
	"github.com/go-chi/chi/v5"
)

func (serv *Server) GetHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	url, ok := serv.s.Get(id)
	if !ok {
		http.Error(w, "url for ID not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (serv *Server) PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	url := string(body)

	defer r.Body.Close()

	hash := serv.k.Generate()
	serv.s.Set(hash, url)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("%s/%s", serv.c.ResolvedHost, hash)))
}

type Server struct {
	s Repository
	k KeyGenerator
	c config.Options
}

func NewServer(s Repository, k KeyGenerator) *Server {
	return &Server{
		s: s,
		k: k,
		c: config.Options{
			Host:         *host,
			ResolvedHost: *resolvedHost,
		},
	}
}
