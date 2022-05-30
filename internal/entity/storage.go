package entity

import (
	m "vfs/internal/model"
)

//var USER
type UserDB struct {
	//
	username map[string]bool
	//a => {} b=> {}
	//folder map[string](map[m.Folder][]m.File) ??
	//key: username, value: folder struct
	//username => folder
	folder map[string]([]m.Folder)
	//key: folded_id, value: file struct
	//folder_id => file
	file map[string]([]m.File)
}

func NewUserDB() *UserDB {
	return &UserDB{
		username: map[string]bool{},
		folder:   map[string]([]m.Folder){},
		file:     map[string]([]m.File){},
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

func (DB *UserDB) CheckFolderByname(u string, fn string) bool {
	//for create_folder use
	//check, the folder must not have this folder name
	//because the folder name is unique
	for _, data := range DB.folder[u] {
		if data.Folder_name == fn {
			return false
		}
	}

	return true
}

func (DB *UserDB) CheckFolderByID(u string, fid string) (int, bool) {
	//for upload_file use
	//check,the folder need to exist because the file can be stored in folder
	for idx, data := range DB.folder[u] {
		if data.Folder_id == fid {
			return idx, true
		}
	}
	return -1, false
}

func (DB *UserDB) CheckFileByName(fid string, fln string) (int, bool) {
	for idx, data := range DB.file[fid] {
		if data.Filename == fln {
			return idx, true
		}
	}
	return -1, false
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

func (DB *UserDB) GetFile(fid string) []m.File {
	return DB.file[fid]
}

func (DB *UserDB) RnFolder(u string, nfn string, chg_pos int) bool {
	DB.folder[u][chg_pos].Folder_name = nfn

	if DB.folder[u][chg_pos].Folder_name == nfn {
		return true
	}
	return false
}

func rmFolder(slice []m.Folder, pos int) []m.Folder {
	return append(slice[:pos], slice[pos+1:]...)
}

func (DB *UserDB) DelFolder(u string, pos int, fid string) bool {
	DB.folder[u] = rmFolder(DB.folder[u], pos)

	for _, data := range DB.folder[u] {
		if data.Folder_id == fid {
			return false
		}
	}

	return true
}

func rmFile(slice []m.File, pos int) []m.File {
	return append(slice[:pos], slice[pos+1:]...)
}

func (DB *UserDB) DelFile(fid string, pos int, fln string) bool {
	DB.file[fid] = rmFile(DB.file[fid], pos)

	for _, data := range DB.file[fid] {
		if data.Filename == fln {
			return false
		}
	}

	return true
}

func (DB *UserDB) AddFile(fl m.File) {
	DB.file[fl.Folder_id] = append(DB.file[fl.Folder_id], fl)
}

func (DB *UserDB) CheckFileisEmpty(fid string) bool {
	if len(DB.file[fid]) > 0 {
		return false
	} else {
		return true
	}
}
