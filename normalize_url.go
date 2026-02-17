package main

import (
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	if strings.HasSuffix(inputURL, "/") {
		inputURL = inputURL[:len(inputURL) - 1]
	}
	
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	return parsedURL.Host + parsedURL.Path, nil
}
