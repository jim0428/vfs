package model

import (
	"fmt"
)

//--------------------------------------------
// Register class
type Register struct {
	Info File
}

func (c *Register) Execute_command() {
	fmt.Println("Register:", c.Info)
	fmt.Println("= =")
}

func (c *Register) Check_command(length int) string {
	//check whether the register command have follow the rule
	//1. check parameter length
	//2. check exist user
	if length > 1 {
		return "Have to much parameters"
	} else {
		//check exist user
		//Error - user already existing
		return "Success"
	}
}
