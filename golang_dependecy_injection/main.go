package main

import (
	"net/http"
	"restful_api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: authMiddleware,
	}
}

func main() {
//  server := InitializeServer()


// 	err := server.ListenAndServe()
// 	helper.PanicIfError(err)
}