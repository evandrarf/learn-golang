package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./template/layout.gohtml", 
		"./template/header.gohtml", 
		"./template/footer.gohtml"),
	)
	t.ExecuteTemplate(writer, "layout.gohtml", map[string]interface{}{
		"Title": "Template Hei",
		"Name":  "Jupri",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/template", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}