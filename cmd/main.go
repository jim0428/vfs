package main

import (
	"bufio"
	"os"
	"strings"

	shlex "github.com/buildkite/shellwords"

	DB "vfs/internal/entity"
	cm "vfs/internal/pkg/command_manger"
)

func main() {
	db := DB.NewUserDB()

	for {
		//origin version
		// consoleReader := bufio.NewReader(os.Stdin)
		// command, _ := consoleReader.ReadString('\n')

		consoleScanner := bufio.NewScanner(os.Stdin)

		consoleScanner.Scan()

		command := consoleScanner.Text()
		command_sli, _ := shlex.Split(command)

		//only enter then skip
		if len(command_sli) == 1 && command_sli[0] == "exit" {
			break
		} else if len(command_sli) <= 1 {
			continue
		} else {
			command_sli[0] = strings.ToLower(command_sli[0])
			//deal command
			cm.DealCommand(db, command_sli, command_sli[0])
		}
	}
}
