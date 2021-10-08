package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		panic("error bro")
	})

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		fmt.Fprint(w, "Error Message : ", err)
	}

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Error Message : error bro", string(body))
}
