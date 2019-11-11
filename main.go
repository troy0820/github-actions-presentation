package main

import (
	"fmt"
	"os"

	"github.com/troy0820/github-actions-presentation/hello"
)

func main() {
	fmt.Println(hello.SayHello("Troy's friends who are listening to his talk"))
	fmt.Fprintf(os.Stdout, hello.SayHello("from the stdout \n"))
}
