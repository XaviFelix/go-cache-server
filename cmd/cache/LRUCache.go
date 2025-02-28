package cache

import (
	"container/list"
	"fmt"
)

// TODO: Put() and Evict(), then test them
// TODO: Add more functionality to this interface
type CacheADT[K comparable, V any] interface {
	Get(key K) (V, bool)
	Put(key K, value V)
	Evict() (V, bool)
}

type CacheEntry[V any] struct {
	value   V
	element *list.Element
}

type LRUCache[K comparable, V any] struct {
	items    map[K]CacheEntry[V]
	list     *list.List
	size     int
	capacity int
}

func New[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		items:    make(map[K]CacheEntry[V]),
		list:     list.New(),
		size:     0,
		capacity: capacity,
	}
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	item, ok := c.items[key]

	if !ok {
		var zero V
		return zero, false
	}
	c.list.MoveToFront(item.element)

	return item.value, true
}

func (c *LRUCache[K, V]) Put(key K, value V) {
	if item, exists := c.items[key]; exists {
		item.value = value
		c.items[key] = item
		//item.value = value
		c.list.MoveToFront(item.element)
	} else {
		if c.size >= c.capacity {
			c.Evict()
		}

		element := c.list.PushFront(key)
		c.items[key] = CacheEntry[V]{value, element}
		c.size++
	}
}

// THis removes the least recently used (LRU)
func (c *LRUCache[K, V]) Evict() (V, bool) {
	e := c.list.Back()
	if e == nil {
		var zero V
		return zero, false
	}

	key := e.Value.(K)
	c.list.Remove(e)

	entry := c.items[key]
	delete(c.items, key)

	// Debug:
	fmt.Print("This is the entry's element should be nil: ")
	fmt.Println(entry.element)
	c.size--

	return entry.value, true
}

func (c *LRUCache[K, V]) GetSize() int {
	return c.size
}
