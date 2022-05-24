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

func (DB *UserDB) AddUser(u string) bool {
	if _, ok := DB.username[u]; ok {
		return false
	}

	DB.username[u] = true

	return true
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

func (DB *UserDB) RnFolder(u string, nfn string, chg_pos int) bool {
	DB.folder[u][chg_pos].Folder_name = nfn

	if DB.folder[u][chg_pos].Folder_name == nfn {
		return true
	}
	return false
}

func remove(slice []m.Folder, pos int) []m.Folder {
	return append(slice[:pos], slice[pos+1:]...)
}

func (DB *UserDB) DelFolder(u string, pos int, fid string) bool {
	DB.folder[u] = remove(DB.folder[u], pos)

	for _, data := range DB.folder[u] {
		if data.Folder_id == fid {
			return false
		}
	}

	return true
}
