package belajar_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIF(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/if.html"))

	t.ExecuteTemplate(w, "if.html", map[string]interface{}{
		"Title": "Template Data Struct",
		"Name":  "Heryan",
	})
}

func TestTemplateActionIF(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIF(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}

func TemplateActionIFOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/if_operator.html"))

	t.ExecuteTemplate(w, "if_operator.html", map[string]interface{}{
		"Title":      "Template Data Struct",
		"FinalValue": 50,
	})
}

func TestTemplateActionIFOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIFOperator(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/range.html"))

	t.ExecuteTemplate(w, "range.html", map[string]interface{}{
		"Title":   "Template Data Struct",
		"Hobbies": []string{"Makan", "Tidur"},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
