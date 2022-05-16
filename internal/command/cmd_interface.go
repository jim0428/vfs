package command

import uDB "vfs/internal/entity"

type Commandmanger interface {
	Execute_command(*uDB.UserDB)
	Check_command(*uDB.UserDB, int) (bool, string)
}

//func Check_command(length int) string { return "" }
