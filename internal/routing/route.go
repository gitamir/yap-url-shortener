package routing

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server interface {
	FullURLForIDHandler(w http.ResponseWriter, r *http.Request, id string)
	ShortenURLHandler(w http.ResponseWriter, r *http.Request)
}

func SetupRouting(server Server) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		server.ShortenURLHandler(w, r)
	})
	router.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		server.FullURLForIDHandler(w, r, chi.URLParam(r, "id"))
	})
	return router
}
