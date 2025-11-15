package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/maziyar-redox/Go-NScript/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is Go-NScript programming language!\n", user.Username)
	fmt.Printf("Feel free to type commands\n")
	repl.Start(os.Stdin, os.Stdout)
}