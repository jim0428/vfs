package commandmanger

import (
	m "vfs/internal/model"
)

//--------------------------------------------

func DealCommand(command_sli interface{}, check_type string) {
	//I will create many struct in struct
	//and I will use different struct to implement interface
	//and I will pass the parameter e.g Register{command_sli}, create_folder{command_sli}

	//m.Execute(command_sli)

	switch check_type {
	//the Command: command_sli will be replaced by File struct from input
	case "register":
		m.Execute(command_sli.(*m.Register))
	case "create_folder":
		m.Execute(command_sli.(*m.Create_folder))
	}

}
