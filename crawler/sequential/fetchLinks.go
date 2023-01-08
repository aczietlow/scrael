package crawl

import (
	"github.com/aczietlow/scrael/links"
	"log"
	"net/url"
)

func Init() {
	// list of urls
	var urls = make([]string, 0)
	// @TODO get the host from the initialized working list of urls.
	var host string = "zietlow.io"

	// provide init urls.
	for _, v := range []string{"https://zietlow.io/"} {
		urls = append(urls, v)
	}

	// crawl url(s).
	breadthFirst(crawl, urls, host)

}

func breadthFirst(f func(item string) []string, worklist []string, host string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			u, _ := url.Parse(item)
			if !seen[item] && host == u.Host {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	list, err := links.Fetch(url)
	if err != nil {
		log.Println(err)
	}
	return list
}
