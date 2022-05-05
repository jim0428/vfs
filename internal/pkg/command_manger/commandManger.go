package commandmanger

import (
	"fmt"
	m "vfs/internal/command"
)

//--------------------------------------------

func DealCommand(command_sli []string, check_type string) {
	//I will create many struct in struct
	//and I will use different struct to implement interface
	//and I will pass the parameter e.g Register{command_sli}, create_folder{command_sli}

	//m.Execute(command_sli)

	command, length := create_command(command_sli, check_type)

	message := command.Check_command(length)

	fmt.Println(message)
	//command.Execute_command()
}

func create_command(command_sli []string, check_type string) (m.Commandmanger, int) {

	length := len(command_sli) - 1

	switch check_type {
	//the Command: command_sli will be replaced by File struct from input
	// Username    string
	// Folder_id   string
	// Folder_name string
	// Description string
	case "register":
		return &m.Register{Info: m.File{Username: command_sli[1], Folder_id: "", Folder_name: "", Description: ""}}, length
	case "create_folder":
		return &m.Create_folder{Info: m.File{Username: command_sli[1], Folder_id: "", Folder_name: command_sli[2], Description: command_sli[3]}}, length
	}

	return nil, 0
}
