package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	postBody, _ := json.Marshal(map[string]string{
		"url":        "https://pkg.go.dev/net/http#Post",
		"visibility": "public",
	})

	responseBody := bytes.NewBuffer(postBody)

	res, _ := http.NewRequest(
		http.MethodPost,
		"https://urlscan.io/api/v1/scan/", responseBody)
	res.Header.Add("API-Key", "ad3aa8ef-251b-4ba5-873c-c49e793c8379")
	res.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(res)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	sb := string(body)
	fmt.Print(sb)
}
