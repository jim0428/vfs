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

	timeZone, _ := time.LoadLocation("Asia/Taipei")
	nowTime, _ := time.ParseInLocation("2006.01.02 15:04:05", time.Now().Format("2006.01.02 15:04:05"), timeZone)
	folder.Create_time = nowTime

	return folder, folder.Folder_id
}
