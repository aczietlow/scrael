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
			// TODO: normalize rel and abs urls
			u, err := url.Parse(href)
			if err != nil {
				log.Fatalf("Failed to parse url: %s", href)
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

func getImagesFromHtml(htmlBody string, baseURL *url.URL) ([]string, error) {
	imgs := make([]string, 0)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		fmt.Println("found image")
		if src, exists := s.Attr("src"); exists {
			u, err := url.Parse(src)
			if err != nil {
				log.Fatalf("Failed to parse url: %s", src)
			}
			if u.IsAbs() {
				imgs = append(imgs, src)
			} else {
				baseURL.Path = src
				imgs = append(imgs, baseURL.String())
			}
		}
	})

	return imgs, nil
}
