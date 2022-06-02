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
	command, err := create_command(command_sli, check_type)

	if err != nil {
		fmt.Println(err)
		return
	}

	if check, _ := command.Check_command(db); check {
		if msg := command.Execute_command(db); !msg {
			fmt.Println("Execute_command occur error")
		}
	}

	//command.Execute_command()
}

func create_command(command_sli []string, check_type string) (cmd.Commandmanger, error) {
	switch check_type {
	case "register":
		return cmd.NewRegister(command_sli)
	case "create_folder":
		return cmd.NewCreateFolder(command_sli)
	case "get_folders":
		return cmd.NewGetFolder(command_sli)
	case "rename_folder":
		return cmd.NewRnFolder(command_sli)
	case "delete_folder":
		return cmd.NewDelFolder(command_sli)
	case "upload_file":
		return cmd.NewULFile(command_sli)
	case "get_files":
		return cmd.NewGetFiles(command_sli)
	case "delete_file":
		return cmd.NewDLFile(command_sli)
	}

	return nil, nil
}
