package main

import (
	"io"
	"net/http"
)

func SendHTTPRequest(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "Something went wrong...", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "Something went wrong...", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Something went wrong...", err
	}

	return string(body), nil
}
