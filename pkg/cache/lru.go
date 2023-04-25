package cache

import "container/list"

type node struct {
	value string
	ll    *list.Element
}

type lruCache struct {
	m   map[string]*node
	cap int
	l   *list.List
}

// I might want to inject the cache interface here
// I might not want this to be a pointer because then I can have multiple caches
func New(cap int) lruCache {
	return lruCache{
		m:   make(map[string]*node),
		cap: cap,
		l:   list.New(),
	}
}

func (c *lruCache) Get(key string) (string, bool) {
	if node, ok := c.m[key]; ok {
		c.l.MoveToFront(node.ll)
		return node.value, true
	}
	return "", false
}

func (c *lruCache) Set(key, value string) {
	// If the key is already in the map, I want to move it to the front of the list
	if node, ok := c.m[key]; ok {
		c.l.MoveToFront(node.ll)
		node.value = value
		return
	}

	//Element is not there so I need to do a check to see if the map is full
	if c.l.Len() == c.cap {
		//remove last element from list and map
		last := c.l.Back()
		delete(c.m, last.Value.(string))
		c.l.Remove(last)
	}

	// If the key is not in the map, I want to add it to the map and the list and move it to the front
	c.m[key] = &node{
		value: value,
		ll:    c.l.PushFront(key),
	}
}
