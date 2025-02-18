package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sqweek/dialog"
	"github.com/xandwell/fileditch-manager-backend/pkg/utils"
)

func Upload() error {
	method := "POST"
	url := "https://up1.fileditch.com/upload.php"

	path, err := dialog.File().Load()
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	file, err := os.Open(path)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	res, err := utils.HTTPRequest(file, method, url)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}
	if res.StatusCode != 200 {
		log.Printf("Status Code %d", res.StatusCode)
		return fmt.Errorf("Server responded with code other than 200, Returned %d", res.StatusCode)
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	r := &utils.UploadResponse{}

	jsonErr := json.Unmarshal(bytes, r)
	if jsonErr != nil {
		log.Printf("%s", jsonErr.Error())
		return jsonErr
	}

	/* TODO: Store results in JSON and do additional cleanup */
	log.Printf("Hash  %s", r.Files[0].Hash)
	log.Printf("Name  %s", r.Files[0].Name)
	log.Printf("URL   %s", r.Files[0].URL)
	log.Printf("Size  %d b", r.Files[0].Size)
	return nil
}
