package command

import (
	"fmt"
	DB "vfs/internal/entity"
)

type delete_folder struct {
	username  string
	folder_id string
	del_pos   int
}

func NewDelFolder(u string, fid string) *delete_folder {
	return &delete_folder{
		username:  u,
		folder_id: fid,
		del_pos:   -1,
	}
}
func (df *delete_folder) Execute_command(db *DB.UserDB) bool {
	if df.del_pos <= -1 {
		return false
	}
	check := db.DelFolder(df.username, df.del_pos, df.folder_id)
	return check
}

func (df *delete_folder) Check_command(db *DB.UserDB, length int) (bool, string) {
	// “Success” , “Error - folder doesn’t exist” , “Error - folder owner not match”
	//The wrong parameters
	//delete_folder {username} {folder_id}
	if length != 2 {
		fmt.Println("Wrong Parameters")
		return false, "Wrong Parameters"
	} else {
		//have this user
		if db.CheckUser(df.username) {
			//check the folder id,if have this user give the change pos
			//if not, return false
			if ok, idx := db.CheckFolderByID(df.username, df.folder_id); ok {
				df.del_pos = idx
				fmt.Println("Success")
				return true, "Success"
			} else {
				fmt.Println("Error - folder_id not found")
				return false, "folder_id not found"
			}
		} else {
			//Don't have this user
			fmt.Println("Error - folder owner not match")
			return false, "Error - folder owner not match"
		}
	}
}
