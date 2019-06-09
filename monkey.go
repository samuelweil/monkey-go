package main

import (
	"fmt"
	"monkey-go/repl"
	"os"
)

func main() {
	fmt.Println("Welcome to Monkey")
	repl.Start(os.Stdin, os.Stdout)
}
