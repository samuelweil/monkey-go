package main

import (
	"fmt"
	"monkey-go/repl"
	"os"
)

func errorHandler() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func main() {

	defer errorHandler()

	fmt.Println("Welcome to Monkey")
	repl.Start(os.Stdin, os.Stdout)
}
