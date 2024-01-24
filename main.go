package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Germanchrystan/Interpreter-In-Go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Mate programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
