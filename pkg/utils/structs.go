package utils

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
