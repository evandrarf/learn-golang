package test

import (
	"fmt"
	"restful_api/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializeService(false)
	fmt.Println(simpleService, err)
}