package main

import (
	"fmt"
	"github.com/aczietlow/scrael/crawler/sequential"
	"time"
)

func main() {
	start := time.Now()

	//url := "https://spinningcode.org/"
	url := "https://zietlow.io/"

	crawl.Init(url)

	elapsed := time.Since(start)
	fmt.Printf("\n\nComplete in %s", elapsed)
}
