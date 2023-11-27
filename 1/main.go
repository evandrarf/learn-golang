package main

import (
	"1/database"
	"fmt"
)

func main() {
	// helper.sayHello("Michael")
	// fmt.Println("Hello World")
	// helper.SayHello("Michael")
	connection := database.GetDatabase()

	fmt.Println(connection)
}