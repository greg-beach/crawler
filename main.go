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

	rawBaseURL := args[1]

	const maxConcurreny = 3
	cfg, err := configure(rawBaseURL, maxConcurreny)
	if err != nil {
		fmt.Printf("error setting configuration: %v\n", err)
		return
	}

	fmt.Printf("starting crawl of: %s\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for normalizedURL, _ := range cfg.pages {
		fmt.Printf("visited %s\n", normalizedURL)
	}
}
