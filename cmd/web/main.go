package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/gist/view", gistView)
	mux.HandleFunc("/gist/create", gistCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from gistbin"))
}

func gistView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gist view"))
}

func gistCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("ALLOW", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
	w.Write([]byte("Gist create"))
}
