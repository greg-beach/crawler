package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing %s\n", rawBaseURL)
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing %s\n", rawCurrentURL)
	}

	if baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing URL: %s\n", rawCurrentURL)
		return
	}

	_, ok := pages[normalizedURL]
	if ok {
		pages[normalizedURL]++
		return
	} else {
		pages[normalizedURL] = 1
	}

	fmt.Printf("extracting html from %s\n", rawCurrentURL)
	currentHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting HTML from %s\n", rawCurrentURL)
		return
	}

	pageData := extractPageData(currentHTML, rawCurrentURL)
	for _, url := range pageData.OutgoingLinks {
		crawlPage(rawBaseURL, url, pages)
	}
}
