package belajar_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := "<html><body><h1>{{.}}</h1></body></html>"

	// Retrun Err
	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	// Must () Not return err
	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(w, "SIMPLE", "HTML TEMPLATE")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	// return err
	// t, err := template.ParseFiles("./template/simple.html")
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.ParseFiles("./template/simple.html"))

	t.ExecuteTemplate(w, "simple.html", "HTML Template File")
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}

func SimpleHTMLDirectory(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./template/*.html"))

	t.ExecuteTemplate(w, "simple.html", "HTML Template File")
}

func TestSimpleHTMLDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLDirectory(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}

//go:embed template/*.html
var templates embed.FS

func SimpleHTMLEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "template/*.html"))

	t.ExecuteTemplate(w, "simple.html", "HTML Template File")
}

func TestSimpleHTMLEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLEmbed(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
