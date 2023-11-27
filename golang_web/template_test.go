package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type User struct {
	Id string
	Name string
}


func TemplateFile(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./template/template.html")) // ./template.html
	t.ExecuteTemplate(writer, "template.html", "Supri")
}

func TestTemplateFile(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/template", nil)
	recorder := httptest.NewRecorder()

	TemplateFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./template/name.html")) // ./name.html
	t.ExecuteTemplate(writer, "name.html", map[string]interface{}{
		"Title": "Template Data Struct",
		"Name": "Supri",
	})
}

func TestTemplateData(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/template", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}