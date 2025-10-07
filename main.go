package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf("no website provided")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Printf("too many arguments provided")
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %s\n", args[0])

	h, err := getHtml(args[0])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Printf("\n%v", h)
}

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
