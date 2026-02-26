package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) < 4 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usages: crawler <baseURL> <maxConcurrency> <maxPages>")
		os.Exit(1)
	}

	if len(args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := args[1]
	maxConcurrencyArg := args[2]
	maxPagesArg := args[3]

	maxConcurrency, err := strconv.Atoi(maxConcurrencyArg)
	if err != nil {
		fmt.Printf("max concurrency error: %v\n", err)
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(maxPagesArg)
	if err != nil {
		fmt.Printf("max pages error: %v\n", err)
		os.Exit(1)
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("error setting configuration: %v\n", err)
		return
	}

	fmt.Printf("starting crawl of: %s\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	writeCSVReport(cfg.pages, "report.csv")
}
