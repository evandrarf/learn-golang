package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecode(t *testing.T) {
	file, _ := os.Open("product.json")

	decoder := json.NewDecoder(file)

	product := Product{}

	decoder.Decode(&product)
	
	fmt.Println(product)
}

func TestStreamEcndeo(t *testing.T) {
	product := Product{
		Id:       1,
		Name:     "Mobilio",
		Price:    220000000,
		ImageUrl: "http://example.com/mobilio.jpg",
	}

	file, _ := os.Create("product_encode.json")
	encoder := json.NewEncoder(file)
	encoder.Encode(product)
}