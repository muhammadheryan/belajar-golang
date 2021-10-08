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

func TestRouterPattern(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id/item/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ProductId := p.ByName("id")
		ItemId := p.ByName("itemId")
		fmt.Fprintf(w, "Product %s - ItemId %s", ProductId, ItemId)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/product/2/item/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 2 - ItemId 1", string(body))
}

func TestRouterPatternCatchAll(t *testing.T) {
	router := httprouter.New()
	router.GET("/image/*dir", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		Dir := p.ByName("dir")
		fmt.Fprintf(w, "Image Directory : %s", Dir)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/image/c/asset/img.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image Directory : /c/asset/img.png", string(body))
}
