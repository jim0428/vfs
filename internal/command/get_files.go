package command

import (
	"fmt"
	DB "vfs/internal/entity"
)

type Get_files struct {
	username       string
	folder_id      string
	sort_name      bool
	sort_time      bool
	sort_extension bool
	mode           string
}

//the default parameters have not finished yet
//because I don't implement about time sorting part
func NewGetFiles(command_sli []string) *Get_files {
	if len(command_sli) != 3 {
		return nil
	} else {
		u := command_sli[1]
		fid := command_sli[2]

		return &Get_files{
			username:       u,
			folder_id:      fid,
			sort_name:      false,
			sort_time:      false,
			sort_extension: false,
			mode:           "",
		}
	}

}

func (getfl *Get_files) Execute_command(db *DB.UserDB) bool {
	fls := db.GetFile(getfl.folder_id)
	for _, fl := range fls {
		fmt.Printf("%s|%s|%s|%s\n", fl.Filename, fl.Description, fl.Create_time, fl.Username)
	}

	return true
}

func (getfl *Get_files) Check_command(db *DB.UserDB) (bool, string) {

	if db.CheckUser(getfl.username) {
		//check the folder id,if have this user give the change pos
		//if not, return false
		if _, ok := db.CheckFolderByID(getfl.username, getfl.folder_id); ok {
			if check := db.CheckFileisEmpty(getfl.folder_id); !check {
				return true, "Success"
			} else {
				fmt.Println("Warning - empty files")
				return true, "Success"
			}
		} else {
			fmt.Println("Error - folder_id not found")
			return false, "Error - folder_id not found"
		}
	} else {
		//Don't have this user
		fmt.Println("unknown users")
		return false, "unknown users"
	}

}
