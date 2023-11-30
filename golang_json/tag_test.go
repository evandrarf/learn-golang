package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id 		 int    `json:"id"` 
	Name     string `json:"name"`
	Price    int    `json:"price"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       1,
		Name:     "Mobilio",
		Price:    220000000,
		ImageUrl: "http://example.com/mobilio.jpg",
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":1,"name":"Mobilio","price":220000000,"image_url":"http://example.com/mobilio.jpg"}`
	jsonBytes := []byte(jsonString)

	product := Product{}

	err := json.Unmarshal(jsonBytes, &product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product.Id)
}