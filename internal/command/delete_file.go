package command

import (
	"fmt"

	DB "vfs/internal/entity"
)

type Delete_file struct {
	username  string
	folder_id string
	file_name string
	del_pos   int
}

func NewDLFile(command_sli []string) *Delete_file {
	if len(command_sli) != 4 {
		return nil
	} else {
		u := command_sli[1]
		fid := command_sli[2]
		fln := command_sli[3]

		return &Delete_file{
			username:  u,
			folder_id: fid,
			file_name: fln,
			del_pos:   -1,
		}
	}
}

func (df *Delete_file) Execute_command(db *DB.UserDB) bool {
	//if the delete position doesn't change
	//this function will return false
	if df.del_pos <= -1 {
		return false
	}
	//Delete the File
	db.DelFile(df.folder_id, df.del_pos, df.file_name)
	//check whether delete
	for _, data := range db.GetFile(df.folder_id, "", "") {
		if data.Filename == df.file_name {
			return false
		}
	}

	return true
}

func (df *Delete_file) Check_command(db *DB.UserDB) (bool, string) {
	// “Success” , “Error - folder_id not found” , “Error - file_name not found” , “Error - unknown user”
	//The wrong parameters
	//delete_file  {username} {folder_id} {file_name}

	//have this user
	if db.CheckUser(df.username) {
		//check the folder id,if have this user give the change pos
		//if not, return false
		if _, ok := db.CheckFolderByID(df.username, df.folder_id); ok {
			if idx, check := db.CheckFileByName(df.folder_id, df.file_name); check {
				df.del_pos = idx
				fmt.Println("Success")
				return true, "Success"
			} else {
				fmt.Println("Error - file_name not found")
				return false, "Error - file_name not found"
			}
		} else {
			fmt.Println("Error - folder_id not found")
			return false, "Error - folder_id not found"
		}
	} else {
		//Don't have this user
		fmt.Println("Error - unknown users")
		return false, "Error - unknown users"
	}

}
