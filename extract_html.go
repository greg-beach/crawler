package main

import (
	"net/url"
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

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	reader := strings.NewReader(htmlBody)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	urls := []string{}

	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		parsedHref, err := url.Parse(href)
		if err != nil {
			return
		}
		absoluteURL := baseURL.ResolveReference(parsedHref)
		urls = append(urls, absoluteURL.String())
	})

	return urls, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	reader := strings.NewReader(htmlBody)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	images := []string{}

	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		img, _ := s.Attr("src")
		parsedImg, err := url.Parse(img)
		if err != nil {
			return
		}
		absoluteImageURL := baseURL.ResolveReference(parsedImg)
		images = append(images, absoluteImageURL.String())
	})

	return images, nil
}
