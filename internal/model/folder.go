package model

import (
	"time"

	"github.com/google/uuid"
)

type Folder struct {
	Username    string
	Folder_name string
	Description string
	Folder_id   string
	Create_time time.Time
}

func NewFolder(Username string, Folder_name string, Description string) (Folder, string) {
	file := Folder{}
	file.Username = Username
	file.Folder_name = Folder_name
	file.Description = Description
	file.Folder_id = uuid.New().String()
	file.Create_time = time.Now()
	return file, file.Folder_id
}
