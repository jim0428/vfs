package command

import (
	"errors"
	"fmt"
	DB "vfs/internal/entity"
)

//--------------------------------------------
// Register class
type Register struct {
	username string
}

func NewRegister(command_sli []string) (*Register, error) {
	//fmt.Println(len(command_sli))
	if len(command_sli) != 2 {
		return nil, errors.New("Command Error!")
	} else {
		return &Register{
			username: command_sli[1],
		}, nil
	}
}

func (c *Register) Execute_command(db *DB.UserDB) bool {
	check := db.AddUser(c.username)

	return check
}

func (c *Register) Check_command(db *DB.UserDB) (bool, string) {
	//check whether the register command have follow the rule
	//1. check parameter length
	//2. check exist user

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
