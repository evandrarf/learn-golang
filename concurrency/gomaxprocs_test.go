package concurrency

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetGoMaxProcs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)
}