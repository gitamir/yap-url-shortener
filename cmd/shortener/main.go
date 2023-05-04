package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
)

var urls = map[string]string{}

func getHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	idSlice := []byte(path)[1:]

	url, ok := urls[string(idSlice)]
	if !ok {
		http.Error(w, "url for ID not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	url := string(body)

	defer r.Body.Close()

	hasher := md5.New()
	hasher.Write([]byte(url))
	hash := fmt.Sprintf("%x", hasher.Sum(nil))
	urls[hash] = url

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("http://127.0.0.1:8080/%s", hash)))
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getHandler(w, r)
	} else if r.Method == http.MethodPost {
		postHandler(w, r)
	} else {
		http.Error(w, "Only GET/POST requsts are currently supported", http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", http.HandlerFunc(handle))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
