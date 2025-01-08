package handlers

type UploadResponse struct {
	Success     bool   `json:"success"`
	ErrorCode   int    `json:"errorcode"`
	Description string `json:"description"`
}
