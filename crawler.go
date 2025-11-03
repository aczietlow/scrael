package main

import (
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentUrl string) {

	cfg.concurrencyControl <- struct{}{}

	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	currentUrl, err := url.Parse(rawCurrentUrl)
	if err != nil {
		log.Println("Failed parsing the current Url", err)
		return
	}

	if cfg.baseUrl.Hostname() != currentUrl.Hostname() {
		// Only crawl pages that belong to the same host
		return
	}

	normalizedUrl, err := normalizeURL(rawCurrentUrl)
	if err != nil {
		log.Println("Failed normalizing the Url", err)
	}

	isFirstVisit := cfg.addPageVist(normalizedUrl)
	if !isFirstVisit {
		// don't fetch pages more than once
		return
	}

	h, err := getHtml(rawCurrentUrl)
	if err != nil {
		log.Println("HTML parsing Error:", err)
		return
	}

	pd := extractPageData(h, rawCurrentUrl)
	cfg.setpageData(normalizedUrl, pd)

	for _, url := range pd.outgoingLinks {
		urlNormalized, err := normalizeURL(url)
		if err != nil {
			log.Println("Failed normalizing the Url", err)
			continue
		}

		if cfg.hasPageAlreadyBeenCrawled(urlNormalized) {
			cfg.wg.Add(1)
			go cfg.crawlPage(url)
		}
	}
}

func (cfg *config) hasPageAlreadyBeenCrawled(url string) bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	_, exists := cfg.pages[url]
	return exists
}

func (cfg *config) setpageData(normalizedUrl string, pd PageData) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	cfg.pages[normalizedUrl] = pd
}

func (cfg *config) addPageVist(normalizedUrl string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, exists := cfg.pages[normalizedUrl]; exists {
		return false
	}
	cfg.pages[normalizedUrl] = PageData{}

	return true
}
