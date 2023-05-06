package main

import (
	"net/http"
)

const Host = "http://localhost:8080"

func main() {
	storage := NewStorage()
	keyGenerator := NewGenerator(storage)
	mux := http.NewServeMux()
	mux.HandleFunc("/", http.HandlerFunc(MainHandler(storage, keyGenerator)))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
