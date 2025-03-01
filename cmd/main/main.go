package main

import (
	"github.com/XaviFelix/go-cache-server.git/cmd/cache"
)

func main() {
	lruCache := cache.New[string, string](5)
	lruCache.Put("1", "Hi")
	lruCache.Put("2", "No")
	lruCache.Put("3", "SUP")
}
