package routing

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server interface {
	GetHandler(w http.ResponseWriter, r *http.Request)
	PostHandler(w http.ResponseWriter, r *http.Request)
}

func SetupRouting(server Server) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		server.PostHandler(w, r)
	})
	router.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		server.GetHandler(w, r)
	})
	return router
}
