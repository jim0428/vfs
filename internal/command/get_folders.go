package command

import (
	"errors"
	"fmt"
	DB "vfs/internal/entity"
)

type Get_folders struct {
	username    string
	sort_by     string
	sort_direct string
}

func NewGetFolder(command_sli []string) (*Get_folders, error) {
	length := len(command_sli)

	if length == 2 {
		return &Get_folders{
			username:    command_sli[1],
			sort_by:     "",
			sort_direct: "",
		}, nil

	} else if length == 4 {
		return &Get_folders{
			username:    command_sli[1],
			sort_by:     command_sli[2],
			sort_direct: command_sli[3],
		}, nil
	}

	return nil, errors.New("Command Error!")
}

func (gfd *Get_folders) Execute_command(db *DB.UserDB) bool {

	folders := db.GetFolder(gfd.username, gfd.sort_by, gfd.sort_direct)
	//db.GetFolder(gfd.username, gfd.sortby, gfd.sort_direct)
	for _, folder := range folders {
		fmt.Printf("%s|%s|%s|%s|%s\n", folder.Folder_id, folder.Folder_name, folder.Description, folder.Create_time, folder.Username)
	}
	return true
}

func (gfd *Get_folders) Check_command(db *DB.UserDB) (bool, string) {
	if gfd.sort_by != "" && gfd.sort_by != "sort_name" && gfd.sort_by != "sort_time" {
		fmt.Println("Wrong Parameters")
		return false, "Wrong Parameters"
	}
	if gfd.sort_direct != "" && gfd.sort_direct != "asc" && gfd.sort_direct != "dsc" {
		fmt.Println("Wrong Parameters")
		return false, "Wrong Parameters"
	}
	//ckeck if have exist user
	if db.CheckUser(gfd.username) {
		//Check exist folder if not, then Create folder
		return true, "Have this user"
	} else {
		//don't have this user
		fmt.Println("unknown user")
		return false, "unknown user"
	}

}
