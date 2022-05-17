package main

import (
	"bufio"
	"os"
	"strings"

	uDB "vfs/internal/entity"
	commandmanger "vfs/internal/pkg/command_manger"
)

func main() {
	db := uDB.NewUserDB()

	for {
		//consoleReader := bufio.NewReader(os.Stdin)
		consoleReader := bufio.NewReader(os.Stdin)

		command, _ := consoleReader.ReadString('\n')

		command = strings.ToLower(command)
		// command looks like this: register Jim
		// TODO command -> tokens = ["register", "Jim"]
		command_sli := strings.Fields(command)

		//if message == exit then end this program
		if len(command_sli) <= 0 {
			continue
		} else if command_sli[0] == "exit" {
			break
		} else {
			//deal command
			commandmanger.DealCommand(db, command_sli, command_sli[0])
		}
	}
}
