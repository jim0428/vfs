package main

import (
	"bufio"
	"os"
	"strings"

	uDB "vfs/internal/entity"
	cm "vfs/internal/pkg/command_manger"
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
		//fmt.Println(command_sli, reflect.TypeOf(command_sli), len(command_sli))
		command_sli = cm.Parse(command_sli)
		//fmt.Println(command_sli, reflect.TypeOf(command_sli), len(command_sli))

		//if message == exit then end this program
		if len(command_sli) <= 0 {
			continue
		} else if command_sli[0] == "exit" {
			break
		} else {
			//deal command
			cm.DealCommand(db, command_sli, command_sli[0])
		}
	}
}
