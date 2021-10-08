package belajar_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodeJSONMap(t *testing.T) {
	product := map[string]interface{}{
		"id":        "P02",
		"name":      "Classic Watch",
		"price":     1200000,
		"image_url": "http://file/image2.png",
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}

func TestDecodeJSONMap(t *testing.T) {
	Json := []byte(`{"id":"P01","name":"Smartwatch","price":999800,"image_url":"http://file/image.png"}`)

	var product map[string]interface{}

	_ = json.Unmarshal(Json, &product)
	fmt.Println(product)
	fmt.Println(product["name"])
}
