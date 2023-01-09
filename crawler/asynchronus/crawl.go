package asyncCrawl

import (
	"fmt"
	"github.com/aczietlow/scrael/links"
	"log"
	"net/url"
)

func Init(startingUrl string, maxConcurrency int) {
	var urls = make([]string, 0)
	var concurrencyLocks = make(chan struct{}, maxConcurrency)

	u, err := url.Parse(startingUrl)
	if err != nil {
		panic(err)
	}

	for _, v := range []string{u.String()} {
		urls = append(urls, v)
	}

	// crawl url(s).
	work(urls, u.Host, concurrencyLocks)

}

// Selects the next to crawl based on what's already been visited.
func work(urls []string, host string, locks chan struct{}) {
	// Unbuffered, omnidirectional channel of link tree items to process
	worklist := make(chan []string)
	// Number of remaining items in worklist to be sent.
	var n int

	// Add starting urls slice to the channel.
	n++
	go func(urls []string) { worklist <- urls }(urls)

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			u, _ := url.Parse(link)
			if !seen[link] && host == u.Host {
				seen[link] = true
				n++
				// Concurrently crawl through link tree.
				go func(link string) {
					worklist <- crawl(link, locks)
				}(link)
			}
		}
	}
}

func crawl(url string, concurrencyLocks chan struct{}) []string {
	fmt.Println(url)
	// Pass empty struct to claim one of the available locks.
	concurrencyLocks <- struct{}{}
	list, err := links.Fetch(url)
	// Release the lock.
	<-concurrencyLocks
	if err != nil {
		log.Println(err)
	}
	return list
}
