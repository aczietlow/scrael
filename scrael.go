package main

import (
	"fmt"
	"github.com/aczietlow/scrael/crawler/asynchronus"
	"github.com/aczietlow/scrael/crawler/sequential"
	"time"
)

func main() {
	start := time.Now()

	// "Run time config"
	//url := "https://spinningcode.org/"
	url := "https://zietlow.io/"
	async := false
	maxConcurrency := 20

	// Choose to run synchronously or asynchronously
	if async {
		asyncCrawl.Init(url, maxConcurrency)
	} else {
		syncCrawl.Init(url)
	}

	elapsed := time.Since(start)
	fmt.Printf("\n\nComplete in %s", elapsed)
}
