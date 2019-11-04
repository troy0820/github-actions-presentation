package hello_test

import (
	"fmt"
	"testing"

	"github.com/troy0820/github-actions-presentation/hello"
)

func TestHello(t *testing.T) {
	fmt.Println("Starting test")
	if want, got := "Hello World", hello.SayHello(); want != got {
		t.Errorf("We wanted %v but got %v", want, got)
	}
}
