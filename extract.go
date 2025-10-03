package main

import (
	"log"
	"net/url"
)

type PageData struct {
	Url            string
	H1             string
	FirstParagraph string
	outgoingLinks  []string
	imageUrl       []string
}

func extractPageData(html, pageUrl string) PageData {
	u, err := url.Parse(pageUrl)
	if err != nil {
		log.Fatalf("Failed to parse url %v", err)
	}

	heading, err := getH1FromHtml(html)
	if err != nil {
		log.Fatalf("Failed fetching heading %v", err)
	}
	paragraph, err := getFirstParagraphFromHTML(html)
	if err != nil {
		log.Fatalf("Failed fetching paragraph %v", err)
	}
	links, err := getURLsFromHtml(html, u)
	if err != nil {
		log.Fatalf("Failed fetching links %v", err)
	}
	images, err := getImagesFromHtml(html, u)
	if err != nil {
		log.Fatalf("Failed fetching images %v", err)
	}

	pd := PageData{
		Url:            pageUrl,
		H1:             heading,
		FirstParagraph: paragraph,
		outgoingLinks:  links,
		imageUrl:       images,
	}

	return pd
}
