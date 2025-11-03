package main

import (
	"encoding/csv"
	"os"
	"strings"
)

func writeCsvReport(pages map[string]PageData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)

	defer w.Flush()

	headers := []string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"}

	if err := w.Write(headers); err != nil {
		return err
	}

	for url, page := range pages {
		outgoingLinks := strings.Join(page.outgoingLinks, ";")
		imageLinks := strings.Join(page.imageUrl, ";")
		record := []string{url, page.H1, page.FirstParagraph, outgoingLinks, imageLinks}
		if err := w.Write(record); err != nil {
			return err
		}
	}

	return nil
}
