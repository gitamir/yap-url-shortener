package routing

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server interface {
	GetFullURL(w http.ResponseWriter, r *http.Request, id string)
	ShortenURL(w http.ResponseWriter, r *http.Request)
}

func SetupRouting(server Server) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		server.ShortenURL(w, r)
	})
	router.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		server.GetFullURL(w, r, chi.URLParam(r, "id"))
	})
	return router
}
