package model

import "fmt"

//--------------------------------------------
// Register class
type Register struct {
	Command []string
}

func (c *Register) Execute_command() {
	fmt.Println("Register:", c.Command)
	fmt.Println("= =")
}

func (c *Register) Check_command() {
	fmt.Println("Register:", c.Command)
}
