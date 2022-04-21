package command

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
	commandmanger "test/internal/pkg/command_manger"
)

func Run() {
	consoleReader := bufio.NewReader(os.Stdin)

	command, _ := consoleReader.ReadString('\n')

	// command looks like this: register jim
	// TODO command -> tokens = ["register", "jim"]
	// if tokens[0] == "register" then ...
	// else if tokens[0] == "create_folder" then ...
	command_sli := strings.Fields(command)

	fmt.Println(command_sli, len(command_sli), reflect.TypeOf(command_sli))

	//deal command
	//I will add chekcer class for the command
	commandmanger.DealCommand(command_sli)

}
