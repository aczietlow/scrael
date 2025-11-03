### When adding to maps

when adding values to a map with int, go has built in guards.
eg.
```go
pages := make(map[string]int)

// This is the same as... 
if _, exists := pages[urlRef]; exists {
    pages[urlRef]++
} else {
    pages[urlRef] = 1
}

// is the same as 
pages[urlRef]++
```

### Writing to stderr

go's log package Println() method will write to stderr by default. It's a nice QoL

## Concurrency

### Fatal errors and defer

When relying on defer to close open connections, close waitgroups, or clear buffered channels be careful with how error handling is approached. For example:

```go
defer func() {
    fmt.Println("defer was triggered")
    // Resolve open connections.
}()

fmt.Println("biz logic")

check := 1
if check == 1 {
    log.Fatalf("throwing os exit")
}
```

log.Fatalf() calls os.Exit() which exits the current program and returns a status code 0-125. This will skip `defer`, `panic()`, and any other resource cleanup.

### Maps

Obviously writing to maps isn't safe for conncurrency and a mutex lock should be used in this case. Reading can be thread safe, but only if it assured that no modifications could be performed while a thread is attempting to read. If thread is attempting to read a map while it is updated by another thread, this will result in a panic().

In a purely "hypothetical" example

```
fatal error: concurrent map read and map write

goroutine 181 [running]:
internal/runtime/maps.fatal({0x759767?, 0xc000481808?})
	/usr/lib/golang/src/runtime/panic.go:1058 +0x18
main.(*config).crawlPage(0xc0000d33b0, {0xc0002091c0, 0x1e})
	/home/aczietlow/Projects/scrael/crawler.go:55 +0x5d3
created by main.(*config).crawlPage in goroutine 7
	/home/aczietlow/Projects/scrael/crawler.go:57 +0x677
```

This can be solved by implementing logic to ensure that reading is a thread safe operation. e.g.

```go
// Unsafe
if _, exists := cfg.pages[urlNormalized]; !exists {}

// Safe
if cfg.hasPageAlreadyBeenCrawled(urlNormalized) {}

func (cfg *config) hasPageAlreadyBeenCrawled(url string) bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	_, exists := cfg.pages[url]
	return exists
}
```

###
