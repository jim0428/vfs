package commandmanger

import (
	"fmt"

	cmd "vfs/internal/command"
	uDB "vfs/internal/entity"
	m "vfs/internal/model"
)

//--------------------------------------------

func DealCommand(db *uDB.UserDB, command_sli []string, check_type string) {
	//I will create many struct in struct
	//and I will use different struct to implement interface
	//and I will pass the parameter e.g Register{command_sli}, create_folder{command_sli}
	command, length := create_command(command_sli, check_type)

	if command == nil {
		return
	}

	if check, message := command.Check_command(db, length); check {
		fmt.Println(message)
		command.Execute_command(db)
	} else {
		fmt.Println(message)
	}

	//command.Execute_command()
}

func create_command(command_sli []string, check_type string) (cmd.Commandmanger, int) {

	length := len(command_sli) - 1

	switch check_type {
	//the Command: command_sli will be replaced by File struct from input
	// Username    string
	// Folder_id   string
	// Folder_name string
	// Description string
	case "register":
		return &cmd.Register{Username: command_sli[1]}, length
	case "create_folder":
		return &cmd.Create_folder{FData: m.Folder{Username: command_sli[1], Folder_name: command_sli[2], Description: command_sli[3]}}, length
	}

	return nil, 0
}
