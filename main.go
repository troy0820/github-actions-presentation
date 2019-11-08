package main

import (
	"fmt"
	"os"

	"bufio"
	"flag"

	"github.com/troy0820/github-actions-presentation/hello"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string to say hello: ")
	reader.Scan()
	var greeting string
	fmt.Sscanf(reader.Text(), "%s", &greeting)
	flag.Parse()
	fmt.Println(hello.SayHello(greeting))
}
