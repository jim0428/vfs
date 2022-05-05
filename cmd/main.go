package main

import (
	"bufio"
	"os"
	"strings"

	commandmanger "vfs/internal/pkg/command_manger"
)

func main() {
	for {
		consoleReader := bufio.NewReader(os.Stdin)

		command, _ := consoleReader.ReadString('\n')

		// command looks like this: register jim
		// TODO command -> tokens = ["register", "jim"]
		// if tokens[0] == "register" then ...
		// else if tokens[0] == "create_folder" then ...
		command_sli := strings.Fields(command)

		//if message == error then don't go down
		//if message == exit then end this program
		if command_sli[0] == "error" {
			println("error")
		} else if command_sli[0] == "exit" {
			break
		} else {
			//deal command
			commandmanger.DealCommand(command_sli, command_sli[0])
		}
	}
}
