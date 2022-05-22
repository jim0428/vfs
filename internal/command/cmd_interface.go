package command

import DB "vfs/internal/entity"

type Commandmanger interface {
	Execute_command(*DB.UserDB) bool
	Check_command(*DB.UserDB, int) (bool, string)
}

//func Check_command(length int) string { return "" }
