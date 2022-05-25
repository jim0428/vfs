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

func NewRnFolder(u string, fid string, nfn string) *Rename_folder {
	// u = username, fid = folder_id, nfn = new_folder_name
	return &Rename_folder{
		username:        u,
		folder_id:       fid,
		new_folder_name: nfn,
		chg_pos:         -1,
	}
}

func (rf *Rename_folder) Execute_command(db *DB.UserDB) bool {
	if rf.chg_pos <= -1 {
		return false
	}
	check := db.RnFolder(rf.username, rf.new_folder_name, rf.chg_pos)
	return check
}

func (rf *Rename_folder) Check_command(db *DB.UserDB, length int) (bool, string) {
	// “Success” , “Error - folder_id not found” , “Error - unknown user”
	//The wrong parameters
	// rename_folders {username} {folder_id} {new_folder_name}
	if length != 3 {
		fmt.Println("Wrong Parameters")
		return false, "Wrong Parameters"
	} else {
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
}
