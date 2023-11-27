package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld()  {
	fmt.Println("Hello, World!")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Hello, from main!")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}