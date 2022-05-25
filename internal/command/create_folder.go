package command

import (
	"fmt"
	DB "vfs/internal/entity"
	m "vfs/internal/model"
)

type Create_folder struct {
	fData m.Folder
}

func NewCreateFolder(u string, f string, d string) *Create_folder {
	//u => username, f => foldername, d => description
	folder, _ := m.NewFolder(u, f, d)
	return &Create_folder{
		fData: folder,
	}
}

func (c *Create_folder) Execute_command(db *DB.UserDB) bool {
	response := db.AddFolder(c.fData)
	fmt.Println(response)

	for _, data := range db.GetFolder(c.fData.Username) {
		if response == data.Folder_id {
			return true
		}
	}
	return false
}

func (c *Create_folder) Check_command(db *DB.UserDB, length int) (bool, string) {
	if length != 3 {
		//the command is wrong
		fmt.Println("Wrong parameters")
		return false, "Wrong Parameters"
	} else {
		//ckeck if have exist user
		if db.CheckUser(c.fData.Username) {
			//Check if have exist folder
			if ok := db.CheckFolderByname(c.fData.Username, c.fData.Folder_name); ok {
				return true, "Success"
			} else {
				fmt.Println("Already have this folder")
				return false, "Already have this folder"
			}
		} else {
			//don't have this user
			fmt.Println("unknown user")
			return false, "unknown user"
		}
	}
}
