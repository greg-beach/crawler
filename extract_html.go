package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHTML(html string) string {
	reader := strings.NewReader(html)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return ""
	}

	h1 := doc.Find("h1").First().Text()

	return h1
}

func getFirstParagraphFromHTML(html string) string {
	reader := strings.NewReader(html)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return ""
	}

	main := doc.Find("main")
	if main != nil {
		p := main.Find("p").First().Text()
		if p != "" {
			return p
		}
	}

	p := doc.Find("p").First().Text()

	return p
}
