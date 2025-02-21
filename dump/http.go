package dump

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func Request(file *os.File, method string, address string) (*http.Response, error) {
	/* Sets up the HTTP Client*/
	client := &http.Client{}

	/* Prepare an empty buffer that we will write the contents of the multipart later*/
	body := &bytes.Buffer{}

	/* Initialize the multipart writer */
	writer := multipart.NewWriter(body)

	/* Write the contents of the requested file to the body buffer */
	formWriter, err := writer.CreateFormFile("files[]", file.Name())
	if err != nil {
		return nil, err
	}
	io.Copy(formWriter, file)

	writer.Close()

	/* Now we make an empty request that will be completed later */
	req, err := http.NewRequest(method, address, body)
	if err != nil {
		return nil, err
	}

	/* Now we complete the HTTP Request by setting the correct header */
	req.Header.Set("Content-Type", writer.FormDataContentType())

	/* Sends the HTTP Request using the client */
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
