package entity

import (
	"fmt"
	m "vfs/internal/model"
)

//var USER
type UserDB struct {
	username map[string]bool
	folder   map[string]([]m.Folder)
}

func NewUserDB() *UserDB {
	return &UserDB{
		username: map[string]bool{},
		folder:   map[string]([]m.Folder){},
	}
}

func (DB *UserDB) CheckUser(u string) bool {
	if _, ok := DB.username[u]; ok {
		return ok
	}
	return false
}

func (DB *UserDB) AddUser(u string) {
	DB.username[u] = true
}

func (DB *UserDB) CheckFolder(u string, fn string) bool {
	//check
	//need to interate
	for _, data := range DB.folder[u] {
		if data.Folder_name == fn {
			return false
		}
	}
	return true
}

func (DB *UserDB) AddFolder(u string, fn string, d string) string {
	//u => username, f => foldername, d => description
	//other parameters, Create time , folder id will be auto created
	//so it don't need to pass
	folder, folder_id := m.NewFolder(u, fn, d)
	//According to username, and create a folder struct as folder
	DB.folder[u] = append(DB.folder[u], folder)
	fmt.Println("In AddFolder:", DB.folder[u])
	return folder_id
}

func (DB *UserDB) GetFolder(u string) []m.Folder {
	//According to username, and create a folder struct as folder
	//need to check folder name
	return DB.folder[u]
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
