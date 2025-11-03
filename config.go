package main

import (
	"net/url"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseUrl            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	maxPages           int
	wg                 *sync.WaitGroup
}

func newConfig(baseRawUrl string, rawMaxCurrency, rawMaxPages string) (config, error) {

	baseUrl, err := url.Parse(baseRawUrl)
	if err != nil {
		return config{}, err
	}

	maxCurrency, err := strconv.Atoi(rawMaxCurrency)
	if err != nil {
		return config{}, err
	}

	maxPages, err := strconv.Atoi(rawMaxPages)
	if err != nil {
		return config{}, err
	}

	conf := config{
		pages:              make(map[string]PageData),
		baseUrl:            baseUrl,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxCurrency),
		maxPages:           maxPages,
		wg:                 &sync.WaitGroup{},
	}

	return conf, nil
}
