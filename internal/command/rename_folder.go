package command

import (
	"fmt"
	DB "vfs/internal/entity"
)

type Rename_folder struct {
	username        string
	folder_id       string
	new_folder_name string
	chg_pos         int
}

//priority queue
//function
func NewRnFolder(command_sli []string) *Rename_folder {
	if len(command_sli) != 4 {
		return nil
	} else {
		// u = username, fid = folder_id, nfn = new_folder_name
		return &Rename_folder{
			username:        command_sli[1],
			folder_id:       command_sli[2],
			new_folder_name: command_sli[3],
			chg_pos:         -1,
		}
	}

}

func (rf *Rename_folder) Execute_command(db *DB.UserDB) bool {
	if rf.chg_pos <= -1 {
		return false
	}
	check := db.RnFolder(rf.username, rf.new_folder_name, rf.chg_pos)
	return check
}

func (rf *Rename_folder) Check_command(db *DB.UserDB) (bool, string) {
	// “Success” , “Error - folder_id not found” , “Error - unknown user”
	//The wrong parameters
	// rename_folders {username} {folder_id} {new_folder_name}

	//have this user
	if db.CheckUser(rf.username) {
		//check the folder id,if have this user give the change pos
		//if not, return false
		if idx, ok := db.CheckFolderByID(rf.username, rf.folder_id); ok {
			rf.chg_pos = idx
			fmt.Println("Success")
			return true, "Success"
		} else {
			fmt.Println("Error - folder_id not found")
			return false, "folder_id not found"
		}
	} else {
		//Don't have this user
		fmt.Println("Error - unknown user")
		return false, "Error - unknown user"
	}

}
