package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

func crawlPage(rawBaseUrl, rawCurrentUrl string, pages map[string]int) {
	if rawCurrentUrl == "" {
		rawCurrentUrl = rawBaseUrl
	}

	baseUrl, err := url.Parse(rawBaseUrl)
	if err != nil {
		log.Println("Failed parsing the base Url", err)
		return
	}

	currentUrl, err := url.Parse(rawCurrentUrl)
	if err != nil {
		log.Println("Failed parsing the current Url", err)
		return
	}

	if baseUrl.Hostname() != currentUrl.Hostname() {
		return
	}

	normalizedUrl, err := normalizeURL(rawCurrentUrl)
	if err != nil {
		log.Println("Failed normalizing the Url", err)
	}

	pages[normalizedUrl]++
	// if _, exists := pages[normalizedUrl]; exists {
	// 	pages[normalizedUrl]++
	// 	return
	// }

	fmt.Printf("Crawling %s \n", rawCurrentUrl)

	h, err := getHtml(rawCurrentUrl)
	if err != nil {
		log.Println("HTML parsing Error:", err)
		return
	}

	urls, err := getURLsFromHtml(h, currentUrl)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, url := range urls {
		urlRef, err := normalizeURL(url)
		if err != nil {
			log.Println("Failed normalizing the Url", err)
			continue
		}

		if _, exists := pages[urlRef]; exists {
			pages[urlRef]++
		} else {
			crawlPage(rawBaseUrl, url, pages)
		}
	}
}
