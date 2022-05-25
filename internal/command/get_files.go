package command

type get_files struct {
	username       string
	folder_id      string
	sort_name      bool
	sort_time      bool
	sort_extension bool
	mode           string
}

func NewGetFiles(u string, fid string) *get_files {
	return &get_files{
		username:       u,
		folder_id:      fid,
		sort_name:      false,
		sort_time:      false,
		sort_extension: false,
		mode:           "",
	}
}
