package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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
	fmt.Printf("starting crawl of: %s", args[0])

	h, err := getHtml(args[0])
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("\n%v", h)
}

func getHtml(rawUrl string) (string, error) {
	c := &http.Client{}
	req, err := http.NewRequest("Get", rawUrl, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("User-Agent", "ScraelCrawler/1.0")
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		err := errors.New("Encountered an http error from server")
		return "", err
	}

	if resp.Header.Get("Content-Type") != "text/html" {
		err := errors.New("Server returned the wrong content type")
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}
