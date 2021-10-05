package belajar_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.PostFormValue("first_name")) //Without parsing

	//With parsing
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, r.PostForm.Get("first_name"))

	fmt.Fprintln(w, r.PostForm.Get("last_name"))
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=lili&last_name=dudud")
	request := httptest.NewRequest(http.MethodPost, "http://localhost/", requestBody)
	recorder := httptest.NewRecorder()
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}
