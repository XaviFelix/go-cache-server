package main

import (
	"fmt"

	"github.com/XaviFelix/go-cache-server.git/cmd/cache"
)

func main() {
	lruCache := cache.New[string, string]()

	fmt.Println(lruCache.GetSize())

}
