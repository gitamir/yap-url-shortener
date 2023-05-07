package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

const Host = "http://localhost:8080"

func main() {
	storage := NewStorage()
	keyGenerator := NewGenerator(storage)
	router := chi.NewRouter()
	server := Server{
		s: storage,
		k: keyGenerator,
	}
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		server.PostHandler(w, r)
	})
	router.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		server.GetHandler(w, r)
	})
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
