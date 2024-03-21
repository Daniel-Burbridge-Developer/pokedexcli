package pokecache

import (
	"fmt"
	"sync"
	"time"
)

// Got this from GPT to help change colors for console -- not entirely sure what the const () syntax is creating.
// It's the same as the import syntax, seems just like a way to do bulk variables?
// loading consts and build time I think
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

type Cache struct {
	CacheEntries map[string]cacheEntry
	mu           sync.RWMutex
}

func NewCache() *Cache {
	c := Cache{
		CacheEntries: make(map[string]cacheEntry),
	}
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.CacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.CacheEntries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) ReapLoop(d int) {
	// fmt.Println(Red, "We be looping", Reset)
	// fmt.Print("Pokedex > ")
	ticker := time.NewTicker(time.Duration(d) * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		//log.Println("I'm reaping the things")
		fmt.Println(Green, "Reaping Location Data Cache", Reset)
		fmt.Print("Pokedex > ")
		c.mu.Lock()
		keysToDelete := []string{}
		for key := range c.CacheEntries {
			if time.Now().After(c.CacheEntries[key].createdAt.Add(time.Duration(d))) {
				keysToDelete = append(keysToDelete, key)
			}
		}

		for _, key := range keysToDelete {
			delete(c.CacheEntries, key)
		}
		c.mu.Unlock()
	}
}
