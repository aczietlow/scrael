package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHtml(html string) ([]string, error) {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return []string{}, err
	}

	header := doc.Find("h1").Text()

	return []string{header}, nil
}

func getFirstParagraphFromHTML(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}
	firstPara := doc.Find("p").First().Text()

	return firstPara, nil
}

func getURLsFromHtml(htmlBody string, baseURL *url.URL) ([]string, error) {
	urls := make([]string, 0)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		if href, exists := s.Attr("href"); exists {
			fmt.Println("found an anchor")
			// TODO: normalize rel and abs urls
			u, err := url.Parse(href)
			if err != nil {
				// TODO: figure out how to error here
				log.Println("Failed to parse url")
			}

			if u.IsAbs() {
				urls = append(urls, href)
			} else {
				baseURL.Path = href
				urls = append(urls, baseURL.String())
			}

		}
	})

	return urls, nil
}
