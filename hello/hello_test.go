package hello_test

import (
	"testing"

	"github.com/troy0820/github-actions-presentation/hello"
)

func TestHello(t *testing.T) {
	tests := map[string]struct {
		input  string
		answer string
	}{
		"Hello World":         {"World", "Hello World"},
		"Hello 757ColorCoded": {"757ColorCoded", "Hello 757ColorCoded"},
		"Hello Fail":          {"Fail", "Hello Fail"},
	}
	for name, test := range tests {
		got := hello.SayHello(test.input)
		t.Run(name, func(t *testing.T) {
			if test.answer != got {
				t.Fatalf("expected %#v, got %#v", test.answer, got)
			}
		})
	}
}
