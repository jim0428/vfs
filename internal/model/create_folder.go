package model

import (
	"fmt"
)

//--------------------------------------------
//create_folder
type Create_folder struct {
	Command []string
}

func (c *Create_folder) Execute_command() {
	fmt.Println("create_folder:", c.Command)
}

//--------------------------------------------
