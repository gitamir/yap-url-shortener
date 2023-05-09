package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

var serverFlags = flag.NewFlagSet("server", flag.ExitOnError)
var host *string
var resolvedHost *string

func init() {
	host = serverFlags.String("a", "localhost:8080", "server address")
	resolvedHost = serverFlags.String("b", "http://localhost:8080", "generated url address")
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "server":
			serverFlags.Parse(os.Args[2:])
		default:
			// PrintDefaults выводит параметры командной строки
			flag.PrintDefaults()
			os.Exit(1)
		}
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
