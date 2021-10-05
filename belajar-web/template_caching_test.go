package belajar_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Content struct {
	Title string
	Name  string
}

var myTemplate = template.Must(template.ParseGlob("./template/*.html"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "name.html", Content{
		Title: "Template Data Struct",
		Name:  "Heryan",
	})
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
