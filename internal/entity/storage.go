package entity

import (
	"sort"
	m "vfs/internal/model"
)

//var USER
type UserDB struct {
	//check user whether exist
	//a => true
	username map[string]bool
	//a user correspond to folder struct
	//a => {} b=> {}
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

func (DB *UserDB) GetFolder(u string, sortby string, sort_direct string) []m.Folder {
	//According to username, and create a folder struct as folder
	//need to check folder name
	result := make([]m.Folder, len(DB.folder[u]))
	copy(result, DB.folder[u])

	if sortby == "sort_name" && sort_direct == "asc" {
		sort.Slice(result, func(i, j int) bool {
			return result[i].Folder_name < result[j].Folder_name
		})
	} else if sortby == "sort_name" && sort_direct == "dsc" {
		sort.Slice(result, func(i, j int) bool {
			return result[i].Folder_name > result[j].Folder_name
		})
	} else if sortby == "sort_time" && sort_direct == "asc" {
		sort.Slice(result, func(i, j int) bool {
			return result[i].Create_time.Before(result[j].Create_time)
		})
	} else if sortby == "sort_time" && sort_direct == "dsc" {
		sort.Slice(result, func(i, j int) bool {
			return result[i].Create_time.After(result[j].Create_time)
		})
	}
	return result
}

func (DB *UserDB) GetFile(fid string, sortby string, sort_direct string) []m.File {
	result := make([]m.File, len(DB.file[fid]))
	copy(result, DB.file[fid])

	if sortby == "sort_name" && sort_direct == "asc" {
		sort.Slice(result, func(i, j int) bool {
			return result[i].Filename < result[j].Filename
		})
	} else if sortby == "sort_name" && sort_direct == "dsc" {
		sort.Slice(result, func(i, j int) bool {
			return result[i].Filename > result[j].Filename
		})
	} else if sortby == "sort_time" && sort_direct == "asc" {
		sort.Slice(result, func(i, j int) bool {
			return result[i].Create_time.Before(result[j].Create_time)
		})
	} else if sortby == "sort_time" && sort_direct == "dsc" {
		sort.Slice(result, func(i, j int) bool {
			return result[i].Create_time.After(result[j].Create_time)
		})
	} else if sortby == "sort_extension" && sort_direct == "asc" {
		sort.Slice(result, func(i, j int) bool {
			return result[i].Filename_extension < result[j].Filename_extension
		})
	} else if sortby == "sort_extension" && sort_direct == "dsc" {
		sort.Slice(result, func(i, j int) bool {
			return result[i].Filename_extension > result[j].Filename_extension
		})
	}

	return result
}

func (DB *UserDB) RnFolder(u string, nfn string, chg_pos int) bool {
	//change the folder name depend on "Username" and this username's "change position"
	DB.folder[u][chg_pos].Folder_name = nfn

	if DB.folder[u][chg_pos].Folder_name == nfn {
		return true
	}
	return false
}

func rmFolder(slice []m.Folder, pos int) []m.Folder {
	return append(slice[:pos], slice[pos+1:]...)
}

func (DB *UserDB) DelFolder(u string, pos int, fid string) {
	DB.folder[u] = rmFolder(DB.folder[u], pos)
}

func (DB *UserDB) DelAllFiles(uid string) {
	delete(DB.file, uid)
}

func rmFile(slice []m.File, pos int) []m.File {
	return append(slice[:pos], slice[pos+1:]...)
}

func (DB *UserDB) DelFile(fid string, pos int, fln string) {
	DB.file[fid] = rmFile(DB.file[fid], pos)
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
