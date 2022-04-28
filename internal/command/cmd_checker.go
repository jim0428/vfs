package command

import (
	m "vfs/internal/model"
)

func check(cmd_sli []string) (interface{}, string) {
	switch cmd_sli[0] {
	case "exit":
		return "nothing", "exit"
	case "register":
		return &m.Register{Command: cmd_sli[1:len(cmd_sli)]}, "register"
	case "create_folder":
		return &m.Create_folder{Command: cmd_sli[1:len(cmd_sli)]}, "create_folder"
	default:
		return "wrong cmd", "error"
	}
}
