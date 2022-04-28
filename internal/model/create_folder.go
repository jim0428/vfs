package model

import (
	"fmt"
)

type Create_folder struct {
	Command []string
}

func (c *Create_folder) Execute_command() {
	fmt.Println("create_folder:", c.Command)
}

func (c *Create_folder) Check_command() {
	fmt.Println("Register:", c.Command)
}
