package main

import (
	"fmt"
	"io"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request, s Repository) {
	path := r.URL.Path
	idSlice := []byte(path)[1:]

	url, ok := s.Get(string(idSlice))
	if !ok {
		http.Error(w, "url for ID not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func PostHandler(w http.ResponseWriter, r *http.Request, s Repository, k KeyGenerator) {
	body, _ := io.ReadAll(r.Body)
	url := string(body)

	defer r.Body.Close()

	hash := k.Generate()
	s.Set(hash, url)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("%s/%s", Host, hash)))
}

func MainHandler(s Repository, k KeyGenerator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetHandler(w, r, s)
		} else if r.Method == http.MethodPost {
			PostHandler(w, r, s, k)
		} else {
			http.Error(w, "Only GET/POST requsts are currently supported", http.StatusMethodNotAllowed)
		}
	}
}
