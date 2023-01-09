package syncCrawl

import (
	"fmt"
	"github.com/aczietlow/scrael/links"
	"log"
	"net/url"
)

func Init(startingUrl string) {
	// list of urls
	var urls = make([]string, 0)

	u, err := url.Parse(startingUrl)
	if err != nil {
		panic(err)
	}
	host := u.Host
	// provide init urls.
	for _, v := range []string{u.String()} {
		urls = append(urls, v)
	}

	// crawl url(s).
	breadthFirst(crawl, urls, host)

}

// Selects the next to crawl based on what's already been visited.
func breadthFirst(f func(item string) []string, worklist []string, host string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			u, _ := url.Parse(item)
			// Adds new links to slice of all discovered from the given host.
			if !seen[item] && host == u.Host {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Fetch(url)
	if err != nil {
		log.Println(err)
	}
	return list
}
