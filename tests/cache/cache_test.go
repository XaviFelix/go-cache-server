package cache_test

import (
	"testing"

	"github.com/XaviFelix/go-cache-server.git/cmd/cache"
)

func TestLRUCache(t *testing.T) {
	// Intance of LRUCache
	t.Log("Testing LRU Cache")
	c := cache.New[string, string]()

	// Testing Put and Get method
	c.Put("sport1", "boxing")
	c.Put("sport2", "basketball")
	item, exist := c.Get("sport2")

	if !exist {
		t.Errorf("Expected an existing Value (basketball), returned: %t", exist)
	} else {
		t.Logf("Value of sport2: %v\n", item)
	}

	// Testing Evict()
	lru, exist := c.Evict()
	if !exist {
		t.Errorf("Expected a least recently used item (boxing), instead got %t", exist)
	} else {
		t.Logf("LRU result (boxing) in default fromat: %v\n", lru)
		t.Logf("LRU result's (boxing) field names: %+v\n", lru) // this format is used for structs
		t.Logf("LRU result (boxing) with Go-syntax representation: %#v\n", lru)
	}

	// TODO: Test when size reaches capacity

}
