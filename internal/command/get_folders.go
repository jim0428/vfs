package command

import (
	"fmt"
	DB "vfs/internal/entity"
)

type Get_folders struct {
	username string
}

func NewGetFolder(name string) *Get_folders {
	return &Get_folders{
		username: name,
	}
}

func (c *Get_folders) Execute_command(db *DB.UserDB) {
	folders := db.GetFolder(c.username)
	for _, folder := range folders {
		fmt.Printf("%s|%s|%s|%s|%s\n", folder.Folder_id, folder.Folder_name, folder.Description, folder.Create_time, folder.Username)
	}
}

func (c *Get_folders) Check_command(db *DB.UserDB, length int) (bool, string) {
	if length > 1 {
		//the command is wrong
		fmt.Println("Have too parameters")
		return false, "Have too parameters"
	} else {
		//ckeck if have exist user
		if db.CheckUser(c.username) {
			//Check exist folder if not, then Create folder
			return true, "Have this user"
		} else {
			//don't have this user
			fmt.Println("unknown user")
			return false, "unknown user"
		}
	}
}
