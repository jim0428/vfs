package command

import DB "vfs/internal/entity"

type Commandmanger interface {
	//Every command will implement this two functions
	Execute_command(*DB.UserDB) bool
	Check_command(*DB.UserDB) (bool, string)
}
