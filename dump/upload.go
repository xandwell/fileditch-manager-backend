package dump

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sqweek/dialog"
)

func Upload() error {
	/* Define the arguments */
	method := "POST"
	url := "https://up1.fileditch.com/upload.php"

	/* Call the File picker dialog to retreive its path */
	path, err := dialog.File().Load()
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	/* IO's the file */
	file, err := os.Open(path)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	/* Call to send HTTP Request for file upload */
	res, err := Request(file, method, url)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	/* Useless check for status code */
	if res.StatusCode != 200 {
		log.Printf("Status Code %d", res.StatusCode)
		return fmt.Errorf("Server responded with code other than 200, Returned %d", res.StatusCode)
	}
	defer res.Body.Close()

	/* Store the response body to bytes */
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	// fmt.Printf("%s\n", string(bytes))

	/* Initializes an empty buffer */
	r := &UploadResponse{}

	jsonErr := json.Unmarshal(bytes, r)
	if jsonErr != nil {
		log.Printf("%s", jsonErr.Error())
		return jsonErr
	}

	/* TODO: Store results in JSON and do additional cleanup (partially done) */

	/* Prints the file contents */
	log.Printf("Hash  %s", r.Files[0].Hash)
	log.Printf("Name  %s", r.Files[0].Name)
	log.Printf("URL   %s", r.Files[0].URL)
	log.Printf("Size  %d bytes", r.Files[0].Size)

	/* Check if the stored.yaml file exists */
	_, exists := os.Open("stored.yaml")
	if exists != nil {
		/* Creates the file since it doesn't exists */
		fmt.Printf("Creating stored.yaml..")
		_, err := os.Create("stored.yaml")
		if err != nil {
			return err
		}

		/* Writes a temporary field for the parser */
		writeErr := os.WriteFile("stored.yaml", []byte("stored: []"), 0666)
		if writeErr != nil {
			return writeErr
		}
	}

	Save(r.Files[0], "stored.yaml")
	return nil
}
