package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonRequest := `{"FirstName": "Evandra", "LastName": "Raditya", "Age": 17}`
	jsonBytes := []byte(jsonRequest)

	customer := Customer{}

	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}