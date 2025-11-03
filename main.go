package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Printf("no website provided")
		os.Exit(1)
	} else if len(args) > 3 {
		fmt.Printf("too many arguments provided")
		os.Exit(1)
	}

	conf, err := newConfig(args[0], args[1], args[2])

	if err != nil {
		log.Println("Failed to bootsrap app.", err)
		os.Exit(1)
	}

	conf.wg.Add(1)
	go conf.crawlPage(conf.baseUrl.String())
	conf.wg.Wait()

	// prettyList, _ := json.MarshalIndent(conf.pages, "", "  ")
	// fmt.Printf("%s", string(prettyList))

	//
	// count := 1
	// for normalizedURL, page := range conf.pages {
	// 	fmt.Printf("[%d] urls on page: %s\n", count, normalizedURL)
	// 	count++
	// 	for i, url := range page.outgoingLinks {
	// 		fmt.Printf("%d - %s\n", i+1, url)
	// 	}
	// }

	if err := writeCsvReport(conf.pages, "report.csv"); err != nil {
		log.Println("Failed generating report", err)
	}

	fmt.Printf("total pages crawled %d\n", len(conf.pages))
}
