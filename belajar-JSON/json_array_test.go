package belajar_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodeJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Muh",
		MiddleName: "Her",
		LastName:   "Cha",
		Age:        18,
		Married:    false,
		Hobbies:    []string{"Tidur", "Makan"},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestDecodeJSONArray(t *testing.T) {
	JsonString := `{"FirstName":"Muh","MiddleName":"Her","LastName":"Cha","Age":18,"Married":false,"Hobbies":["Tidur","Makan"]}`
	JsonBytes := []byte(JsonString)

	customer := &Customer{}

	err := json.Unmarshal(JsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Hobbies)
}

func TestEncodeJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName:  "Muh",
		MiddleName: "Her",
		LastName:   "Cha",
		Age:        18,
		Married:    false,
		Hobbies:    []string{"Tidur", "Makan"},
		Address: []Address{
			{
				City:    "Malang",
				ZipCode: "98989",
			},
			{
				City:    "Jakarta",
				ZipCode: "12345",
			},
		},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestDecodeJSONArrayComplex(t *testing.T) {
	JsonString := `[{"City":"Malang","ZipCode":"98989"},{"City":"Jakarta","ZipCode":"12345"}]`
	JsonBytes := []byte(JsonString)

	Address := &[]Address{}

	err := json.Unmarshal(JsonBytes, Address)
	if err != nil {
		panic(err)
	}

	fmt.Println(Address)
}
