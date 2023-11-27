package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(writer, "Hello Bang")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=supri", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}