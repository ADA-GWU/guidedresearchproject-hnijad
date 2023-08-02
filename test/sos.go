package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func sosTest() {
	dirEntries, err := os.ReadDir(readPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}
	i := 0
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			continue
		}

		filePath := filepath.Join(readPath, dirEntry.Name())
		fid := fmt.Sprintf("%v,%v", 16, i)
		err = uploadFile(filePath, "http://localhost:8081/data/"+fid)
		//fmt.Println(err)
		i++
	}

}

func uploadFile(filePath, url string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return fmt.Errorf("error creating form file: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("error copying file content: %v", err)
	}

	writer.Close()

	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}

	//fmt.Println(response.Body, "erad")

	//responseBody, err := ioutil.ReadAll(response.Body)
	//fmt.Println("response =", string(responseBody), url)
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("server returned an error: %s", response.Status)
	}

	return nil
}
