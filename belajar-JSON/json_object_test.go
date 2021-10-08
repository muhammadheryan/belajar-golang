package belajar_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Address    []Address
}

type Address struct {
	City    string
	ZipCode string
}

func TestEncodeJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Muh",
		MiddleName: "Her",
		LastName:   "Cha",
		Age:        18,
		Married:    false,
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestDecodeJSONObject(t *testing.T) {
	JsonString := `{"FirstName":"Muh","MiddleName":"Her","LastName":"Cha","Age":18,"Married":false}`
	JsonBytes := []byte(JsonString)

	customer := Customer{}

	err := json.Unmarshal(JsonBytes, &customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.Age)
}
