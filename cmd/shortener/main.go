package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

var flagHost string
var flagResolvedHost string

func main() {
	flag.StringVar(&flagHost, "a", "localhost:8080", "server address")
	flag.StringVar(&flagResolvedHost, "b", "http://localhost:8080", "generated url address")
	flag.Parse()

	if envServerAddr := os.Getenv("SERVER_ADDRESS"); envServerAddr != "" {
		flagHost = envServerAddr
	}

	if envBaseAddr := os.Getenv("BASE_URL"); envBaseAddr != "" {
		flagResolvedHost = envBaseAddr
	}

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	storage := NewStorage()
	keyGenerator := NewGenerator(storage)
	router := chi.NewRouter()
	server := NewServer(storage, keyGenerator)
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		server.PostHandler(w, r)
	})
	router.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		server.GetHandler(w, r)
	})
	return http.ListenAndServe(server.c.Host, router)
}
