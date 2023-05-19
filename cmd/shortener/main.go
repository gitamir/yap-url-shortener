package main

import (
	"log"
	"net/http"

	"github.com/gitamir/yap-url-shortener/internal/handlers"
	"github.com/gitamir/yap-url-shortener/internal/routing"
	"github.com/gitamir/yap-url-shortener/internal/storage"
)

func main() {
	if err := run(); err != nil {
		log.Fatal("failed to run server")
	}
}

func run() error {
	storage := storage.NewStorage()
	keyGenerator := handlers.NewGenerator(storage)
	server := handlers.NewServer(storage, keyGenerator)
	router := routing.SetupRouting(server)
	return http.ListenAndServe(server.Config.Host, router)
}
