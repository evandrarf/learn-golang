package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName string
	LastName string
	Age int
}

func TestJsonObject(t *testing.T) {
	customer := Customer{"Evan", "Raditya", 17}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}