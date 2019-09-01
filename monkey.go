package main

import (
	"flag"
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

	debugMode := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	opts := repl.Options{
		Debug: *debugMode,
	}

	fmt.Println("Welcome to Monkey")
	repl.Start(os.Stdin, os.Stdout, opts)
}
