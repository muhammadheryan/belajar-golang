package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hello Httprouter")
	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}
