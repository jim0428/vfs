package model

type Data struct {
	Username    string
	Folder_id   string
	Folder_name string
	Description string
}

func NewFile(Username string, Folder_id string, Folder_name string, Description string) Data {
	file := Data{}
	file.Username = Username
	file.Folder_id = Folder_id
	file.Folder_name = Folder_name
	file.Description = Description
	return file
}
