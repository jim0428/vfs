package command

import (
	"fmt"
	DB "vfs/internal/entity"
)

//--------------------------------------------
// Register class
type Register struct {
	username string
}

func NewRegister(name string) *Register {
	return &Register{
		username: name,
	}
}

func (c *Register) Execute_command(db *DB.UserDB) {
	db.AddUser(c.username)
}

func (c *Register) Check_command(db *DB.UserDB, length int) (bool, string) {
	//check whether the register command have follow the rule
	//1. check parameter length
	//2. check exist user
	if length > 1 {
		//the command is wrong
		fmt.Println("Have too parameters")
		return false, "Have too parameters"
	} else {
		//check exist user
		//Error - user already existing
		if !(db.CheckUser(c.username)) {
			//don't have this user
			fmt.Println("Success")
			return true, "Success"
		} else {
			//already have this user
			fmt.Println("Error -user already existing")
			return false, "Error -user already existing"
		}
	}
}
