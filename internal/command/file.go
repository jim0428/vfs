package model

type File struct {
	Username    string
	Folder_id   string
	Folder_name string
	Description string
}

func NewFile(Username string, Folder_id string, Folder_name string, Description string) File {
	file := File{}
	file.Username = Username
	file.Folder_id = Folder_id
	file.Folder_name = Folder_name
	file.Description = Description
	return file
}
