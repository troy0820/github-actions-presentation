package hello_test

import (
	"testing"

	"github.com/troy0820/github-actions-presentation/hello"
)

func TestHello(t *testing.T) {

	if want, got := "Hello World", hello.SayHello(); want != got {
		t.Errorf("We wanted %v but got %v", want, got)
	}
}
