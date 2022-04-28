package model

type File struct {
	Username    string
	Folder_id   string
	Folder_name string
	Description string
}

func NewFile() File {
	file := File{}
	file.Username = ""
	file.Folder_id = ""
	file.Folder_name = ""
	file.Description = ""
	return file
}
