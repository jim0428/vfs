package command

import (
	"fmt"
	uDB "vfs/internal/entity"
	m "vfs/internal/model"
)

type Create_folder struct {
	Info m.Data
}

func (c *Create_folder) Execute_command(db *uDB.UserDB) {
	fmt.Println("create_folder:", c.Info)
}

func (c *Create_folder) Check_command(db *uDB.UserDB, length int) (bool, string) {
	fmt.Println("Register:", c.Info)
	return false, ""
}
