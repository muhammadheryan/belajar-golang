package belajar_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	//response Header
	w.Header().Add("Author", "hrynch")

	//Request Header
	contentType := r.Header.Get("content-type")

	fmt.Fprint(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	//Request Header
	fmt.Println(string(body))

	//Response Header
	author := recorder.Header().Get("Author")
	fmt.Println(author)
}
