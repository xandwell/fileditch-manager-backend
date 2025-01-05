package internal

import (
	"net/http"
	"log"
)

const port = ":8080"

func Start() {
	log.Printf("%s%s","Server starting at localhost", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
