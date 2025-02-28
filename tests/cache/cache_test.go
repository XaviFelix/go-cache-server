package cache_test

import (
	"testing"

	"github.com/XaviFelix/go-cache-server.git/cmd/cache"
)

func TestLRUCache(t *testing.T) {
	// Intance of LRUCache
	t.Log("Testing LRU Cache")
	c := cache.New[string, string](2)

	// Testing Put and Get method
	c.Put("sport1", "boxing")
	c.Put("sport2", "basketball")
	item, exist := c.Get("sport1")

	if !exist {
		t.Errorf("Expected an existing Value (boxing), returned: %t", exist)
	} else {
		t.Logf("Value of sport1 (boxing): %v\n", item)
	}

	// Testing Evict()
	lru, exist := c.Evict()
	if !exist {
		t.Errorf("Expected a least recently used item (basketball), instead got %t", exist)
	} else {
		t.Logf("LRU result (basketball) in default fromat: %v\n", lru)
		t.Logf("LRU result's (basketball) field names: %+v\n", lru) // this format is used for structs
		t.Logf("LRU result (basketball) with Go-syntax representation: %#v\n", lru)
		t.Logf("The size of the cache is now 1: %d", c.GetSize())
	}

	// TODO: Test when size reaches capacity
	c.Put("sport3", "swimming")
	c.Put("sport4", "hokey")

	// Checking if LRU still exists in the cache:
	// boxing should be gone and the order should be (LRU) swimming -> hokey(MRU)
	leastRecently, _ := c.Evict()
	mostRecently, _ := c.Evict()

	t.Logf("LRU should be swimming: %v\n", leastRecently)
	t.Logf("MRU should be hockey: %v\n", mostRecently)
	t.Logf("Size of the cache should be zero: %d\n", c.GetSize())
}
