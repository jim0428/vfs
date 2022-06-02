package command

import (
	"errors"
	"fmt"
	DB "vfs/internal/entity"
	m "vfs/internal/model"
)

type Upload_file struct {
	fl m.File
}

func NewULFile(command_sli []string) (*Upload_file, error) {
	if len(command_sli) != 5 {
		return nil, errors.New("Command Error!")
	} else {
		u := command_sli[1]
		fid := command_sli[2]
		fln := command_sli[3]
		d := command_sli[4]

		return &Upload_file{
			fl: m.NewFile(u, fid, fln, d),
		}, nil
	}

}

func (ULfile *Upload_file) Execute_command(db *DB.UserDB) bool {
	db.AddFile(ULfile.fl)

	for _, data := range db.GetFile(ULfile.fl.Folder_id, "", "") {
		if ULfile.fl == data {
			return true
		}
	}
	return false
}

func (ULfile *Upload_file) Check_command(db *DB.UserDB) (bool, string) {
	// “Success” , “Error - folder_id not found” , “Error - unknown user”
	//The wrong parameters
	//upload_file {username} {folder_id} {file_name} {description}

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
