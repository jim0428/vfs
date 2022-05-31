package command

import (
	"fmt"
	DB "vfs/internal/entity"
)

type Get_folders struct {
	username string
}

func NewGetFolder(command_sli []string) *Get_folders {
	if len(command_sli) != 2 {
		return nil
	} else {
		return &Get_folders{
			username: command_sli[1],
		}
	}
}

func (c *Get_folders) Execute_command(db *DB.UserDB) bool {
	folders := db.GetFolder(c.username)
	for _, folder := range folders {
		fmt.Printf("%s|%s|%s|%s|%s\n", folder.Folder_id, folder.Folder_name, folder.Description, folder.Create_time, folder.Username)
	}
	return true
}

func (c *Get_folders) Check_command(db *DB.UserDB) (bool, string) {

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
