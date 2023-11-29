package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hello, World!")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Hello, World!", string(body), "response body should match")
}

func TestParams(t *testing.T) {
	router := httprouter.New()

	router.GET("/products/:id", func(writer http.ResponseWriter, response *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Product #%s", params.ByName("id"))
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Product #1", string(body), "response body should match")
}

//go:embed resources/*.txt
var files embed.FS

func TestServeFiles(t *testing.T) {
	router := httprouter.New()
	direcotry, _ := fs.Sub(files, "resources")
	router.ServeFiles("/static/*filepath", http.FS(direcotry))

	request := httptest.NewRequest("GET", "http://localhost:8080/static/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Hello World", string(body), "response body should match")
}