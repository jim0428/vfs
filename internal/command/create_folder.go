package command

import (
	"fmt"
	DB "vfs/internal/entity"
	m "vfs/internal/model"
)

type Create_folder struct {
	FData m.Folder
}

func (c *Create_folder) Execute_command(db *DB.UserDB) {
	if ok := db.CheckFolder(c.FData.Username, c.FData.Folder_name); ok {
		resonse := db.AddFolder(c.FData.Username, c.FData.Folder_name, c.FData.Description)
		fmt.Println(resonse)
	} else {
		fmt.Println("Already have this folder")
	}
}

func (c *Create_folder) Check_command(db *DB.UserDB, length int) (bool, string) {
	if length > 3 {
		//the command is wrong
		fmt.Println("Have too parameters")
		return false, "Have too parameters"
	} else {
		//ckeck if have exist user
		if db.CheckUser(c.FData.Username) {
			//Check exist folder if not, then Create folder
			return true, "Have this user"
		} else {
			//don't have this user
			fmt.Println("unknown user")
			return false, "unknown user"
		}
	}
}
