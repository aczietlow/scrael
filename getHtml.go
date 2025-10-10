package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHtml(rawUrl string) (string, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", rawUrl, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "ScraelCrawler/1.0")
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		err := fmt.Errorf("Encountered an http error from server: %v", resp.Status)
		return "", err
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		err := fmt.Errorf("Server returned the wrong content type: %v", contentType)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
