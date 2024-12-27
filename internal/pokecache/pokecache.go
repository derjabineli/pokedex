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
	c.mux.Lock()
	defer c.mux.Unlock()
    
    entry := cacheEntry{createdAt: time.Now(), val: val}
    c.entry[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	entry, exists  := c.entry[key]
	if !exists {
		return nil, exists
	}
	return entry.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for key, entry := range c.entry {
		expiryTime := entry.createdAt.Add(interval)
		if currentTime := time.Now(); currentTime.After(expiryTime) {
			delete(c.entry, key)
		}
	}
}

func NewCache(interval time.Duration) *Cache {
    c := &Cache{
        entry: make(map[string]cacheEntry),
        mux:   &sync.Mutex{},
    }
    go func() {
         ticker := time.NewTicker(interval)
         defer ticker.Stop()
         for {
             select {
             case <-ticker.C:
                 c.reapLoop(interval)
             }
         }
    }()
    return c
}