package belajar_web

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from redirected page")
}

func RedirectFrom(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
