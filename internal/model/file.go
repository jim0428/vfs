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
	timeZone, _ := time.LoadLocation("Asia/Taipei")
	nowTime, _ := time.ParseInLocation("2006.01.02 15:04:05", time.Now().Format("2006.01.02 15:04:05"), timeZone)
	return File{
		Username:           u,
		Folder_id:          fid,
		Filename:           fln,
		Filename_extension: "",
		Description:        d,
		Create_time:        nowTime,
	}
}
