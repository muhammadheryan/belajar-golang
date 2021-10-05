package belajar_web

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "upload_form.html", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	file, fileHeader, err := r.FormFile("document")
	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resource/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	myTemplate.ExecuteTemplate(w, "upload_success.html", map[string]interface{}{
		"Name": name,
		"Path": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resource"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
