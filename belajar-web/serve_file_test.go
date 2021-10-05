package belajar_web

import (
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name != "" {
		http.ServeFile(w, r, "./resource/ok.html")
	} else {
		http.ServeFile(w, r, "./resource/not_found.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	server.ListenAndServe()
}
