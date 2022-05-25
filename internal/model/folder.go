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
	folder := Folder{}
	folder.Username = Username
	folder.Folder_name = Folder_name
	folder.Description = Description
	folder.Folder_id = uuid.New().String()
	folder.Create_time = time.Now()
	return folder, folder.Folder_id
}
