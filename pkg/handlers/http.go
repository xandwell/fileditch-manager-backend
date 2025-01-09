package handlers

import (
	"bytes"
	"mime/multipart"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"github.com/sqweek/dialog"
	"os"
)

func Request() {

	filePath, err := dialog.File().Title("Select a file").Load()
	if err != nil {
		log.Fatalf("Failed to select file: %v", err)
	}
	
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	part, err := writer.CreateFormFile("files[]", filePath)
	if err != nil {
		log.Fatalf("Error creating form file: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		log.Fatalf("Error copying file data: %v", err)
	}

	err = writer.Close()
	if err != nil {
		log.Fatalf("Error closing writer: %v", err)
	}

	req, err := http.NewRequest("POST", "https://up1.fileditch.com/upload.php", &requestBody)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var jsonResponse UploadResponse
	jsonErr := json.Unmarshal(body, &jsonResponse)
	if jsonErr != nil {
		log.Fatalf("Error parsing JSON response: %v", jsonErr)
	}
	
	log.Printf("Success: %s", jsonResponse.Description)
	log.Printf("File name: %s", jsonResponse.Files[0].Name)
	log.Printf("File name: %s", jsonResponse.Files[0].URL)
}

