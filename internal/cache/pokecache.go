package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CacheEntries map[string]cacheEntry
	mu           sync.RWMutex
}

func NewCache(d int) Cache {
	c := Cache{
		CacheEntries: make(map[string]cacheEntry),
	}
	c.reapLoop(d)
	return c
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

func (c *Cache) reapLoop(d int) {
	ticker := time.NewTicker(time.Duration(d) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		//log.Println("I'm reaping the things", t)

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
