package main

import (
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
