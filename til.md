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

