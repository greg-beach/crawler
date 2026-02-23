package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return PageData{}
	}

	h1 := getH1FromHTML(html)
	p := getFirstParagraphFromHTML(html)
	links, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		return PageData{}
	}
	images, err := getImagesFromHTML(html, baseURL)
	if err != nil {
		return PageData{}
	}

	return PageData{
		URL:            baseURL.String(),
		H1:             h1,
		FirstParagraph: p,
		OutgoingLinks:  links,
		ImageURLs:      images,
	}
}

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
