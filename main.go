package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[1]

	fmt.Printf("starting crawl of: %s\n", baseURL)

	pages := map[string]int{}
	crawlPage(baseURL, baseURL, pages)

	for page, count := range pages {
		fmt.Printf("page: %s, visited %d times\n", page, count)
	}
}
