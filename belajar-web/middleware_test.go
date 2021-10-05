package belajar_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before executed handler")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After executed handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi ERROR")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "ERROR : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(w, r)
}
func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("MAIN HANDLER")
		fmt.Fprint(w, "This is main handler")
	})
	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("PANIC")
		panic("ups")
	})

	// logMiddleware := new(LogMiddleware)
	// logMiddleware.Handler = mux

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	server.ListenAndServe()
}
