package model

import (
	"strings"
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
	//give time format and time zone
	timeZone, _ := time.LoadLocation("Asia/Taipei")
	nowTime, _ := time.ParseInLocation("2006.01.02 15:04:05", time.Now().Format("2006.01.02 15:04:05"), timeZone)

	// filename = 1.tc
	// fileExtenArr = [1 , tc] fileExtenStr = '.' + 'tc'
	// filename = abc
	// fileExtenArr = [abc] fileExtenStr = ''
	fileExtenStr := ""

	if ok := strings.Contains(fln, "."); ok {
		fileExtenArr := strings.Split(fln, ".")
		fileExtenStr = "." + fileExtenArr[len(fileExtenArr)-1]
	} else {
		fileExtenStr = ""
	}

	return File{
		Username:           u,
		Folder_id:          fid,
		Filename:           fln,
		Filename_extension: fileExtenStr,
		Description:        d,
		Create_time:        nowTime,
	}
}
