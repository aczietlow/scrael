package main

import (
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseUrl            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func newConfig(baseRawUrl string, maxCurrency int) (config, error) {

	baseUrl, err := url.Parse(baseRawUrl)
	if err != nil {
		return config{}, err
	}

	conf := config{
		pages:              make(map[string]PageData),
		baseUrl:            baseUrl,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxCurrency),
		wg:                 &sync.WaitGroup{},
	}

	return conf, nil
}
