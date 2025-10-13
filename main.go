package main

import (
	"encoding/json"
	"fmt"
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
	rawUrl := args[0]

	// fmt.Printf("starting crawl of: %s\n", rawUrl)

	pageMap := make(map[string]int)

	crawlPage(rawUrl, "", pageMap)

	prettyList, _ := json.MarshalIndent(pageMap, "", "  ")
	fmt.Printf("%s", string(prettyList))

}
