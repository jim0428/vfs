package commandmanger

import (
	"fmt"
	cmd "vfs/internal/command"
	DB "vfs/internal/entity"
)

//--------------------------------------------

func DealCommand(db *DB.UserDB, command_sli []string, check_type string) {
	//I will create many struct in struct
	//and I will use different struct to implement interface
	//and I will pass the parameter e.g Register{command_sli}, create_folder{command_sli}
	command, length := create_command(command_sli, check_type)

	if command == nil {
		return
	}
	if check, _ := command.Check_command(db, length); check {
		if msg := command.Execute_command(db); !msg {
			fmt.Println("Execute_command occur error")
		}
	}

	//command.Execute_command()
}

func create_command(command_sli []string, check_type string) (cmd.Commandmanger, int) {

	length := len(command_sli) - 1

	switch check_type {
	case "register":
		return cmd.NewRegister(command_sli[1]), length
	case "create_folder":
		return cmd.NewCreateFolder(command_sli[1], command_sli[2], command_sli[3]), length
	case "get_folders":
		return cmd.NewGetFolder(command_sli[1]), length
	}

	return nil, 0
}
