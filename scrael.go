package main

import (
	"fmt"
	"github.com/aczietlow/scrael/crawler/sequential"
	"time"
)

func main() {
	start := time.Now()
	crawl.Init()

	elapsed := time.Since(start)
	fmt.Printf("\n\nComplete in %s", elapsed)
}
