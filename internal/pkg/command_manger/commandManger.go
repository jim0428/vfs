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

	if command == err {
		fmt.Println("Error command")
		return
	}

	if check, _ := command.Check_command(db); check {
		if msg := command.Execute_command(db); !msg {
			fmt.Println("Execute_command occur error")
		}
	}

	//command.Execute_command()
}

func create_command(command_sli []string, check_type string) (cmd.Commandmanger, cmd.Commandmanger) {
	switch check_type {
	case "register":
		return cmd.NewRegister(command_sli), (*cmd.Register)(nil)
	case "create_folder":
		return cmd.NewCreateFolder(command_sli), (*cmd.Create_folder)(nil)
	case "get_folders":
		return cmd.NewGetFolder(command_sli), (*cmd.Get_folders)(nil)
	case "rename_folder":
		return cmd.NewRnFolder(command_sli), (*cmd.Rename_folder)(nil)
	case "delete_folder":
		return cmd.NewDelFolder(command_sli), (*cmd.Delete_folder)(nil)
	case "upload_file":
		return cmd.NewULFile(command_sli), (*cmd.Upload_file)(nil)
	case "get_files":
		return cmd.NewGetFiles(command_sli), (*cmd.Get_files)(nil)
	case "delete_file":
		return cmd.NewDLFile(command_sli), (*cmd.Delete_file)(nil)
	}

	return nil, nil
}
