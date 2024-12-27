package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	entry map[string]cacheEntry
	mux *sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{createdAt: time.Now(), val: val}
	c.entry[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, exists  := c.entry[key]
	if !exists {
		return nil, exists
	}
	return entry.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	for key, entry := range c.entry {
		if (entry.createdAt - time.Now()) > interval {
			delete(c.entry, key)
		}
	}
}

func NewCache(interval time.Duration) {
	mux := &sync.Mutex{}
	cache := Cache{mux: mux}
	cache.reapLoop(interval)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cache.reapLoop(interval)
		}
	}
}