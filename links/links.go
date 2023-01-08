// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package links

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Fetch(url string) ([]string, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("gettings %s: %s", url, response.Status)
	}

	htmlDoc, err := html.Parse(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attribute := range node.Attr {
				if attribute.Key != "href" {
					continue
				}
				if strings.HasPrefix(attribute.Val, "#") || strings.HasPrefix(attribute.Val, "tel") {
					continue
				}
				link, err := response.Request.URL.Parse(attribute.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(htmlDoc, visitNode, nil)

	return links, nil
}

func forEachNode(node *html.Node, pre, post func(node *html.Node)) {
	if pre != nil {
		pre(node)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		forEachNode(child, pre, post)
	}
	if post != nil {
		post(node)
	}
}
