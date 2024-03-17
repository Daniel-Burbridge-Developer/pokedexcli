package pokecache

import (
	"time"
)

type Cache struct {
	CacheEntries map[string]cacheEntry
}

func NewCache(d int) {
	c := Cache{}

	ticker := time.NewTicker(time.Duration(d))

	// PROBABLY NEEDS TO BE IN A GO ROUTINE OR SOMETHING?
	for {
		<-ticker.C
		ticker.Reset(time.Duration(d))
		c.ReapLoop()
	}

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

func (c *Cache) ReapLoop() {
	// DELETE ALL THE OLD THINGS!
}
