package command

import (
	"errors"
	"fmt"
	DB "vfs/internal/entity"
)

type Get_files struct {
	username    string
	folder_id   string
	sortby      string
	sort_direct string
}

//the default parameters have not finished yet
//because I don't implement about time sorting part
func NewGetFiles(command_sli []string) (*Get_files, error) {
	length := len(command_sli)

	if length == 3 { //get_files {username} {folder_id}
		return &Get_files{
			username:    command_sli[1],
			folder_id:   command_sli[2],
			sortby:      "",
			sort_direct: "",
		}, nil
	} else if length == 5 { //get_files {username} {folder_id} {sort_name|sort_time|sort_extension} {asc|dsc}
		return &Get_files{
			username:    command_sli[1],
			folder_id:   command_sli[2],
			sortby:      command_sli[3],
			sort_direct: command_sli[4],
		}, nil
	}

	return nil, errors.New("Command Error!")
}

func (getfl *Get_files) Execute_command(db *DB.UserDB) bool {
	fls := db.GetFile(getfl.folder_id, getfl.sortby, getfl.sort_direct)
	for _, fl := range fls {
		fmt.Printf("%s|%s|%s|%s|%s\n", fl.Filename, fl.Description, fl.Filename_extension, fl.Create_time, fl.Username)
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
