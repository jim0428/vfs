package command

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
	commandmanger "vfs/internal/pkg/command_manger"
)

func Run() {
	//it need to let user know what step need to do
	//1. Register 2. log in
	for {
		consoleReader := bufio.NewReader(os.Stdin)

		command, _ := consoleReader.ReadString('\n')

		// command looks like this: register jim
		// TODO command -> tokens = ["register", "jim"]
		// if tokens[0] == "register" then ...
		// else if tokens[0] == "create_folder" then ...
		command_sli := strings.Fields(command)

		//check the command whether is right first
		//if right, then get the command in a struct
		//if wrong, then get the error message
		//the return array will be like (struct,message)
		cmd, message := check(command_sli)

		//if message == error then don't go down
		//if message == exit then end this program
		if message == "error" {
			println("error")
		} else if message == "exit" {
			break
		} else {
			//deal command
			fmt.Println("test:", cmd, reflect.TypeOf(cmd))
			commandmanger.DealCommand(cmd, message)
		}
	}

}
