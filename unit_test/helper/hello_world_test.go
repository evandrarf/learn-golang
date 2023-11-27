package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		request string
		want string
	}{
		{
			name: "Eko",
			request: "Eko",
			want: "Hello Eko",
		},
		{
			name: "Kurniawan",
			request: "Kurniawan",
			want: "Hello Kurniawan",
		},
		{
			name: "Khannedy",
			request: "Khannedy",
			want: "Hello Khannedy",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.want, result)
		})
	}
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		panic("Result is not Hello Eko")
	}
}