package model

import (
	"time"
)

type File struct {
	Username           string
	Folder_id          string
	Filename           string
	Filename_extension string
	Description        string
	Create_time        time.Time
}

func NewFile(u string, fid string, fln string, d string) File {
	return File{
		Username:           u,
		Folder_id:          fid,
		Filename:           fln,
		Filename_extension: "",
		Description:        d,
		Create_time:        time.Now(),
	}
}
