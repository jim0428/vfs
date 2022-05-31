package command

import (
	"fmt"
	DB "vfs/internal/entity"
	m "vfs/internal/model"
)

type Create_folder struct {
	fData m.Folder
}

func NewCreateFolder(command_sli []string) *Create_folder {
	//u => username, f => foldername, d => description
	if len(command_sli) != 4 {
		return nil
	} else {
		folder, _ := m.NewFolder(command_sli[1], command_sli[2], command_sli[3])
		return &Create_folder{
			fData: folder,
		}
	}
}

func (c *Create_folder) Execute_command(db *DB.UserDB) bool {
	//
	//
	response := db.AddFolder(c.fData)
	fmt.Println(response)

	for _, data := range db.GetFolder(c.fData.Username) {
		if response == data.Folder_id {
			return true
		}
	}
	return false
}

func (c *Create_folder) Check_command(db *DB.UserDB) (bool, string) {
	//{folder id} , “Error - unknown user”
	//The wrong parameters
	//create_folder {username} {folder_name} {description}

	//ckeck if have exist user
	if db.CheckUser(c.fData.Username) {
		//Check if have exist folder already
		//if exist , then can not create this folder
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
