package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/gitamir/yap-url-shortener/cmd/shortener/routing"
)

var flagHost string
var flagResolvedHost string

func main() {
	setupFlags()

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	storage := NewStorage()
	keyGenerator := NewGenerator(storage)
	server := NewServer(storage, keyGenerator)
	router := routing.SetupRouting(server)
	return http.ListenAndServe(server.c.Host, router)
}

func setupFlags() {
	flag.StringVar(&flagHost, "a", "localhost:8080", "server address")
	flag.StringVar(&flagResolvedHost, "b", "http://localhost:8080", "generated url address")
	flag.Parse()

	if envServerAddr := os.Getenv("SERVER_ADDRESS"); envServerAddr != "" {
		flagHost = envServerAddr
	}

	if envBaseAddr := os.Getenv("BASE_URL"); envBaseAddr != "" {
		flagResolvedHost = envBaseAddr
	}
}
