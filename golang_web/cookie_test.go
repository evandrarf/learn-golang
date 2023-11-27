package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Cookie created")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("user")
	if err != nil {
		fmt.Fprint(writer, "No cookie found")
	} else {
		fmt.Fprintf(writer, "Hello %s", cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	
}