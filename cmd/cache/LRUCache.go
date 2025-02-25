package cache

import (
	"container/list"
)

// TODO: Put() and Evict(), then test them
type CacheADT[K comparable, V any] interface {
	Get(key K) (V, bool)
	Put(key K, value V)
	Evict(key K)
}

type LRUCache[K comparable, V any] struct {
	items map[K]struct {
		value   V
		element *list.Element
	}
	list *list.List
}

func New[K comparable, V any]() *LRUCache[K, V] {
	return &LRUCache[K, V]{
		items: make(map[K]struct {
			value   V
			element *list.Element
		}),
		list: list.New(),
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

// TODO: Finish this
func (c *LRUCache[K, V]) Put(key K, value V) {
	// if key exists then update its value and move it to the front (MRU)
	// else if key doesn't exist, create a new entry and add it to map and list (push to front on list for LRU)

}
