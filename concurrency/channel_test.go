package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	ch := make(chan string)

	defer close(ch)

	go func(){
		time.Sleep(2 * time.Second)
		ch <- "Hello, World!"
		fmt.Println("Sent message to channel")
	}()

	data := <- ch

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Hello, World!"
}

func TestChannelAsParameter(t *testing.T) {
	ch := make(chan string)

	defer close(ch)

	go GiveMeResponse(ch)

	data := <- ch
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello, World!"
}

func OnlyOut(channel <-chan string) {
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	ch := make(chan string)

	defer close(ch)

	go OnlyIn(ch)
	go OnlyOut(ch)

	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	ch := make(chan string, 3)

	defer close(ch)

	go func() {
		ch <- "Hello, World!"
		ch <- "Hello, World!"
		time.Sleep(2 * time.Second)
	}()

	fmt.Println(len(ch))

	go func() {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}()

	time.Sleep(2 * time.Second)
}