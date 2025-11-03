# Scrael

Web Crawler specifically intended use for load testing.

***Scrael** - spider-like creatures about the size of a wagon wheel with razor like feet. Often considered demons by commonfolk*

![CI](https://github.com/aczietlow/scrael/actions/workflows/ci.yml/badge.svg?branch=v2)

## Usage 

arg 0 - base url to crawl
arg 1 - number of conrrent workers
arg 2 - max number of pages to crawl

```
go run . https://zietlow.io 5 100
```

### Sequential Mode

![gophish logo](https://raw.github.com/aczietlow/scrael/main/static/images/sequential.png)


### Concurrency

![gophish logo](https://raw.github.com/aczietlow/scrael/main/static/images/concurrency.png)

## Resources
* [personal miro board notes](https://miro.com/app/board/uXjVPYZIqT0=/?share_link_id=198914050024)

### More resources for future musings
* https://www.ardanlabs.com/blog/2015/09/composition-with-go.html
* https://github.com/golang/go/wiki/LearnConcurrency
* https://jdanger.com/build-a-web-crawler-in-go.html
* https://stackoverflow.com/questions/5834808/designing-a-web-crawler/5834890#5834890%  
