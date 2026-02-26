package main

import (
	"encoding/csv"
	"os"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	writer.Write([]string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"})

	for page, data := range pages {
		writer.Write([]string{page, data.H1, data.FirstParagraph, strings.Join(data.OutgoingLinks, ";"), strings.Join(data.ImageURLs, ";")})
	}

	return nil
}
