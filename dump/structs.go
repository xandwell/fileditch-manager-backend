package dump

import "time"

type UploadResponse struct {
	Success     bool          `json:"success"`
	ErrorCode   int           `json:"errorcode"`
	Description string        `json:"description"`
	Files       []FileDetails `json:"files"`
}

type FileDetails struct {
	Hash string `json:"hash"`
	Name string `json:"name"`
	URL  string `json:"url"`
	Size int64  `json:"size"`
}

type File struct {
	Hash     string    `yaml:"hash"`
	FileName string    `yaml:"file_name"`
	URL      string    `yaml:"url"`
	Size     int64     `yaml:"size"`
	Date     time.Time `yaml:"date"`
}

type Files struct {
	Stored []File `yaml:stored_urls`
}
