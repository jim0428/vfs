package model

import (
	"fmt"
)

type Create_folder struct {
	Info File
}

func (c *Create_folder) Execute_command() {
	fmt.Println("create_folder:", c.Info)
}

func (c *Create_folder) Check_command(length int) string {
	fmt.Println("Register:", c.Info)
	return "no develop"
}
