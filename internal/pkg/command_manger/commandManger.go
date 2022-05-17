package commandmanger

import (
	"strings"

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
	if check, _ := command.Check_command(db, length); check {
		command.Execute_command(db)
	}

	//command.Execute_command()
}

func create_command(command_sli []string, check_type string) (cmd.Commandmanger, int) {

	length := len(command_sli) - 1

	switch check_type {
	case "register":
		return &cmd.Register{Username: command_sli[1]}, length
	case "create_folder":
		return &cmd.Create_folder{FData: m.Folder{Username: command_sli[1], Folder_name: command_sli[2], Description: command_sli[3]}}, length
	}

	return nil, 0
}

func Parse(command_sli []string) []string {
	result := make([]string, 0)
	tmp := ""
	link := false
	for _, data := range command_sli {
		if strings.Count(data, "'") == 2 {
			result = append(result, data)
		} else if link && strings.Contains(data, "'") {
			tmp += " " + data
			result = append(result, tmp)
			link = false
			tmp = ""
		} else if link {
			tmp += " " + data
		} else if !link && strings.Contains(data, "'") {
			tmp += data
			link = true
		} else if !link {
			result = append(result, data)
		}
	}
	return result
}
