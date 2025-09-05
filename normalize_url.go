package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	u.Scheme = ""
	u.Fragment = ""
	u.RawQuery = ""
	stringURL := strings.Trim(u.String(), "/")
	return stringURL, nil
}
