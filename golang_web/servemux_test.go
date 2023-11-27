package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func (writer http.ResponseWriter, request *http.Request)  {
		fmt.Fprint(writer, "Hello World")
	})

	mux.HandleFunc("/hi", func (writer http.ResponseWriter, request *http.Request)  {
		fmt.Fprint(writer, "Hi")
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	
}