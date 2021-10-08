package belajar_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	ImageUrl string `json:"image_url"`
}

func TestEncodeJSONTag(t *testing.T) {
	product := Product{
		Id:       "P01",
		Name:     "Smartwatch",
		Price:    999800,
		ImageUrl: "http://file/image.png",
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}

func TestDecodeJSONTag(t *testing.T) {
	Json := []byte(`{"id":"P01","name":"Smartwatch","price":999800,"image_url":"http://file/image.png"}`)

	product := Product{}

	_ = json.Unmarshal(Json, &product)
	fmt.Println(product)
}
