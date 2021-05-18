package main

import (
	"fmt"
	"sync"
)

// Embedding the mutex in the Cache makes it easier to see
// what is being guarded by the mutex.
type Cache struct {
	sync.Mutex
	mappings map[string]string
}

var cache = Cache{
	mappings: make(map[string]string),
}

func (c *Cache) Set(key string, value string) {
	c.Lock()
	defer c.Unlock()

	c.mappings[key] = value
}

func (c *Cache) Has(key string) bool {
	_, ok := c.mappings[key]

	return ok
}

func (c *Cache) Lookup(key string) (string, bool) {
	c.Lock()
	defer c.Unlock()

	value, ok := c.mappings[key]

	return value, ok
}

func main() {
	fmt.Println(cache.Has("a")) // false

	cache.Set("a", "1")

	fmt.Println(cache.Has("a")) // true

	fmt.Println(cache.Lookup("a"))
	fmt.Println(cache.Lookup("b"))
}
