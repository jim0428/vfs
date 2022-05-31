package command

import (
	"fmt"
	DB "vfs/internal/entity"
)

type Get_folders struct {
	username    string
	sort_by     string
	sort_direct string
}

func NewGetFolder(command_sli []string) *Get_folders {
	length := len(command_sli)
	u := command_sli[1]

	if length == 2 {
		return &Get_folders{
			username:    u,
			sort_by:     "",
			sort_direct: "",
		}
	} else if length == 4 {
		sort_by := command_sli[2]
		direct := command_sli[3]

		return &Get_folders{
			username:    u,
			sort_by:     sort_by,
			sort_direct: direct,
		}
	}

	return nil
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
