package main

import (
	"fmt"

	"test/internal/command"
	"test/internal/foo"
	"test/internal/voo"
)

func main() {
	fmt.Println("123")
	foo.Foo()
	voo.Voo()

	command.Run()
}
