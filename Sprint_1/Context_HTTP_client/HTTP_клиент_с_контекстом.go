package main

import (
	"context"
	"io"
	"net/http"
)

func SendHTTPRequestWithContext(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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
