package main

import (
	"fmt"
	"os"

	"bufio"

	"github.com/troy0820/github-actions-presentation/hello"
)

func main() {
	fmt.Fprintf(os.Stdout, hello.SayHello("from the stdout \n"))
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string to say hello: ")
	reader.Scan()
	var greeting string
	fmt.Sscanf(reader.Text(), "%s", &greeting)
	fmt.Println(hello.SayHello(greeting))
}
