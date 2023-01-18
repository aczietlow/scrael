# TODO List


* [x] Build a **BASIC** web crawler
    * [ ] Stop getting distracted by actually making a working web crawler
* [x] Refactor code to use basic runtime config to change settings more easily.
* [x] Refactor the crawler to support concurrency
* [ ] Unit tests.
* [ ] Add output formats

## Web Crawler

* Pick back up from the miro board as a design of what I was thinking for concurrency
    * ~~Not having access to maps or slices due to being unsafe and causing race conditions is going to be a challenge for future Chris to figure out~~
      * I handled this by having a go routine specifically handle the link tree; adding new nodes to the tree, traversing to the next node, and sending it to the link extractor.
* Look more into: https://jdanger.com/build-a-web-crawler-in-go.html