package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf("no website provided")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Printf("too many arguments provided")
		os.Exit(1)
	}

	conf, err := newConfig(args[0], 5)

	if err != nil {
		log.Println("Failed to bootsrap app.", err)
		os.Exit(1)
	}

	conf.wg.Add(1)
	go conf.crawlPage(conf.baseUrl.String())
	conf.wg.Wait()

	// prettyList, _ := json.MarshalIndent(conf.pages, "", "  ")
	// fmt.Printf("%s", string(prettyList))
	for normalizedURL, page := range conf.pages {
		fmt.Printf("urls on page: %s\n", normalizedURL)
		for i, url := range page.outgoingLinks {
			fmt.Printf("%d - %s\n", i, url)
		}
	}
}
