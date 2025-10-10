package main

import (
	"net/url"
	"strings"
)

func crawlPage(rawBaseUrl, rawCurrentUrl string, pages map[string]int) {
	if rawCurrentUrl == "" {
		rawCurrentUrl = rawBaseUrl
	}

	if !doUrlsMatchDomain(rawBaseUrl, rawCurrentUrl) {
		return
	}

}

func doUrlsMatchDomain(rawBaseUrl, rawCurrentUrl string) bool {

	u, err := url.Parse(rawBaseUrl)
	if err != nil {
		return false
	}
	u.Fragment = ""
	u.RawQuery = ""
	u.Path = ""
	baseHost := strings.Trim(u.String(), "/")

	u2, err := url.Parse(rawCurrentUrl)
	if err != nil {
		return false
	}
	u2.Fragment = ""
	u2.RawQuery = ""
	u2.Path = ""
	currentHost := strings.Trim(u2.String(), "/")

	if currentHost != baseHost {
		return false
	}

	return true
}
