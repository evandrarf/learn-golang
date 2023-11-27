package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstname := request.PostForm.Get("firstname")
	lastname := request.PostForm.Get("lastname")

	fmt.Fprintf(writer, "Hello %s %s", firstname, lastname)
}

func TestPostForm(t *testing.T) {
	requestBody := strings.NewReader("firstname=Supri&lastname=Yanfei")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}