package command

import (
	"fmt"
	DB "vfs/internal/entity"
)

type Delete_folder struct {
	username  string
	folder_id string
	//where folder at,and need to delete that position
	del_pos int
}

func NewDelFolder(command_sli []string) *Delete_folder {
	if len(command_sli) != 3 {
		return nil
	} else {
		return &Delete_folder{
			username:  command_sli[1],
			folder_id: command_sli[2],
			del_pos:   -1,
		}
	}
}
func (dfd *Delete_folder) Execute_command(db *DB.UserDB) bool {
	if dfd.del_pos <= -1 {
		return false
	}
	//delete folder and the files of the folder
	db.DelFolder(dfd.username, dfd.del_pos, dfd.folder_id)
	db.DelAllFiles(dfd.folder_id)

	for _, data := range db.GetFolder(dfd.username) {
		if data.Folder_id == dfd.folder_id {
			return false
		}
	}

	return true
}

func (df *Delete_folder) Check_command(db *DB.UserDB) (bool, string) {
	// “Success” , “Error - folder doesn’t exist” , “Error - folder owner not match”
	//The wrong parameters
	//delete_folder {username} {folder_id}

	//have this user
	if db.CheckUser(df.username) {
		//check the folder id,if have this user give the change pos
		//if not, return false
		if idx, ok := db.CheckFolderByID(df.username, df.folder_id); ok {
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
