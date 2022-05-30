package command

import (
	"fmt"

	DB "vfs/internal/entity"
)

type delete_file struct {
	username  string
	folder_id string
	file_name string
	del_pos   int
}

func NewDLFile(u string, fid string, fln string) *delete_file {
	return &delete_file{
		username:  u,
		folder_id: fid,
		file_name: fln,
		del_pos:   -1,
	}
}

func (df *delete_file) Execute_command(db *DB.UserDB) bool {
	if df.del_pos <= -1 {
		return false
	}
	check := db.DelFile(df.folder_id, df.del_pos, df.file_name)
	return check
}

func (df *delete_file) Check_command(db *DB.UserDB, length int) (bool, string) {
	// “Success” , “Error - folder_id not found” , “Error - file_name not found” , “Error - unknown user”
	//The wrong parameters
	//delete_file  {username} {folder_id} {file_name}
	if length != 3 {
		fmt.Println("Wrong Parameters")
		return false, "Wrong Parameters"
	} else {
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
}
