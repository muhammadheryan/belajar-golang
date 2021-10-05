package belajar_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var tmpl = template.Must(template.ParseFiles("./template/post.html"))

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "post.html", map[string]interface{}{
		"Title": "XSS Test",
		"Body":  "<p>Ini di escape<p>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}

func TemplateAutoEscapeDisabled(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "post.html", map[string]interface{}{
		"Title": "XSS Test Unescaped",
		"Body":  template.HTML("<h2>Ini di escape<h2>"),
	})
}

func TestTemplateAutoEscapeDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
