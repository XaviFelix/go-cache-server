package main

import (
	"fmt"

	"github.com/XaviFelix/go-cache-server.git/cmd/cache"
)

func main() {
	lruCache := cache.New[string, string]()

	lruCache.Put("hello", "world")
	lruCache.Put("Xavier", "xavi")

	lruCache.Put("Hi", "there")
	lruCache.Put("fish", "shark")

	fmt.Println(lruCache.GetSize())

}
