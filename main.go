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
	rawUrl := args[0]
	u, err := normalizeURL(rawUrl)
	if err != nil {
		log.Println("Error when parsing URL:", err)
	}

	fmt.Printf("starting crawl of: %s\n", u)

	h, err := getHtml(rawUrl)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Printf("\n%v", h)
}
