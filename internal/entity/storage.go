package entity

import (
	m "vfs/internal/model"
)

//var USER
type UserDB struct {
	//
	username map[string]bool
	//a => {} b=> {}
	folder map[string]([]m.Folder)
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

func (DB *UserDB) AddFolder(fdata m.Folder) string {
	//u => username, f => foldername, d => description
	//other parameters, Create time , folder id will be auto created
	//so it don't need to pass
	//According to username, and create a folder struct as folder
	DB.folder[fdata.Username] = append(DB.folder[fdata.Username], fdata)
	//fmt.Println("In storage's AddFolder:", DB.folder[u])
	return fdata.Folder_id
}

func (DB *UserDB) GetFolder(u string) []m.Folder {
	//According to username, and create a folder struct as folder
	//need to check folder name
	return DB.folder[u]
}
