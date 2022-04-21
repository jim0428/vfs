package commandmanger

import "fmt"

type Commandmanger interface {
	Execute_command()
}

func Execute(commandmanger Commandmanger) {
	commandmanger.Execute_command()
}

//--------------------------------------------
// Register class
type register struct {
	command []string
}

func (c *register) Execute_command() {
	fmt.Println("register:", c.command)
}

//--------------------------------------------
//create_folder
type create_folder struct {
	command []string
}

func (c *create_folder) Execute_command() {
	fmt.Println("create_folder:", c.command)
}

//--------------------------------------------

func DealCommand(command_sli []string) {
	//I will create many struct
	//and I will use different struct to implement interface
	//and I will pass the parameter e.g Register{command_sli}, create_folder{command_sli}

	switch command_sli[0] {
	case "register":
		Execute(&register{command_sli})
	case "create_folder":
		Execute(&create_folder{command_sli})
	}

}
