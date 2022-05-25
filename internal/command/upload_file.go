package command

import (
	"fmt"
	DB "vfs/internal/entity"
	m "vfs/internal/model"
)

type Upload_file struct {
	fl m.File
}

func NewULFile(u string, fid string, fln string, d string) *Upload_file {
	return &Upload_file{
		fl: m.NewFile(u, fid, fln, d),
	}
}

func (ULfile *Upload_file) Execute_command(db *DB.UserDB) bool {
	db.AddFile(ULfile.fl)

	for _, data := range db.GetFile(ULfile.fl.Folder_id) {
		if ULfile.fl == data {
			return true
		}
	}
	return false
}

func (ULfile *Upload_file) Check_command(db *DB.UserDB, length int) (bool, string) {
	// “Success” , “Error - folder_id not found” , “Error - unknown user”
	//The wrong parameters
	//upload_file {username} {folder_id} {file_name} {description}
	if length != 3 {
		//the command is wrong
		fmt.Println("Wrong parameters")
		return false, "Wrong Parameters"
	} else {
		//ckeck if have exist user
		if db.CheckUser(ULfile.fl.Username) {
			//Check if have exist folder
			if _, ok := db.CheckFolderByID(ULfile.fl.Username, ULfile.fl.Folder_id); ok {
				fmt.Println("Success")
				return true, "Success"
			} else {
				fmt.Println("folder_id not found")
				return false, "folder_id not found"
			}
		} else {
			//don't have this user
			fmt.Println("unknown user")
			return false, "unknown user"
		}
	}
}
