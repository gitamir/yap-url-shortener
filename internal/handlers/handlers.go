package handlers

import (
	"io"
	"net/http"
	"net/url"

	"github.com/gitamir/yap-url-shortener/internal/config"
)

func (serv *Server) GetFullURL(w http.ResponseWriter, r *http.Request, id string) {
	url, ok := serv.storage.Get(id)
	if !ok {
		http.Error(w, "url for ID not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (serv *Server) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	urlString := string(body)

	hash := serv.keyGenerator.Generate()
	str, _ := serv.storage.Get(hash)
	if str != "" {
		hash = serv.keyGenerator.Generate()
	}
	serv.storage.Set(hash, urlString)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	value, err := url.JoinPath(serv.Config.ResolvedHost, hash)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	w.Write([]byte(value))
}

type Server struct {
	storage      Repository
	keyGenerator KeyGenerator
	Config       config.Options
}

func NewServer(s Repository, k KeyGenerator) *Server {
	return &Server{
		storage:      s,
		keyGenerator: k,
		Config:       *config.NewConfig(),
	}
}

type KeyGenerator interface {
	Generate() string
}

type Repository interface {
	Set(string, string)
	Get(string) (string, bool)
}
