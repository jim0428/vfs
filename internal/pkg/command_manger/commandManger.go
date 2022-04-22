package commandmanger

import (
	m "vfs/internal/model"
)

//--------------------------------------------

func DealCommand(command_sli []string) {
	//I will create many struct in struct
	//and I will use different struct to implement interface
	//and I will pass the parameter e.g Register{command_sli}, create_folder{command_sli}

	switch command_sli[0] {
	case "register":
		m.Execute(&m.Register{Command: command_sli})
	case "create_folder":
		m.Execute(&m.Create_folder{Command: command_sli})
	}

}
