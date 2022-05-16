package entity

import (
	m "vfs/internal/model"
)

//var USER
type UserDB struct {
	username map[string]bool
	data     []m.Data
}

func NewUserDB() *UserDB {
	return &UserDB{
		username: map[string]bool{},
		data:     []m.Data{},
	}
}

func (uDB *UserDB) AddUser(u string) {
	uDB.username[u] = true
}

func (uDB *UserDB) CheckUser(u string) bool {
	if _, ok := uDB.username[u]; ok {
		return ok
	}
	return false
}

// type User struct {
// 	username string
// }

//CRUD
//Create
//Read
//Update
//Delete

//get set
