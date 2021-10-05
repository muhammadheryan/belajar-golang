package belajar_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/name.html"))

	t.ExecuteTemplate(w, "name.html", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Ryan",
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}

type Page struct {
	Title   string
	Name    string
	Address Address
}

type Address struct {
	Street string
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/name.html"))

	t.ExecuteTemplate(w, "name.html", Page{
		Title: "Template Data Struct",
		Name:  "Heryan",
		Address: Address{
			Street: "Jl Praji",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
