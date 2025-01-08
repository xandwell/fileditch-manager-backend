package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Request() {
	log.Printf("%s", "Real")
	res, err := http.Get("https://up1.fileditch.com/upload.php")
	if err != nil {
		log.Fatalf("Fatal: %v", err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Fatal: %v", err.Error())
	}

	var jsonResponse UploadResponse

	jsonErr := json.Unmarshal(body, &jsonResponse)
	if jsonErr != nil {
		log.Fatalf("Fatal: %s", err.Error())
	}
	log.Printf("Success: %s", jsonResponse.Description)
}
