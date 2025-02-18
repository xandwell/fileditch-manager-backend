package handlers

import (
	"encoding/json"
	"github.com/sqweek/dialog"
	"github.com/xandwell/fileditch-manager-backend/pkg/utils"
	"io"
	"log"
	"os"
)

func Upload() {
	method := "POST"
	url := "https://up1.fileditch.com/upload.php"

	path, err := dialog.File().Load()
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	file, err := os.Open(path)
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	res, err := utils.HTTPRequest(file, method, url)
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}
	if res.StatusCode != 200 {
		log.Printf("%s", err.Error())
		return
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	r := &utils.UploadResponse{}

	jsonErr := json.Unmarshal(bytes, r)
	if jsonErr != nil {
		log.Printf("%s", jsonErr.Error())
		return
	}

	log.Printf("Hash  %s", r.Files[0].Hash)
	log.Printf("Name  %s", r.Files[0].Name)
	log.Printf("URL   %s", r.Files[0].URL)
	log.Printf("Size  %d b", r.Files[0].Size)
	return
}
