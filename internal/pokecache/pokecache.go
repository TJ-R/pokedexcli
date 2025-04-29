package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu 			 *sync.Mutex
}

type cacheEntry struct {
	createdAt 	time.Time
	val 		[]byte	
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheEntries: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}

	cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.cacheEntries[key] = cacheEntry {
		createdAt: time.Now(),
		val: val,
	}
	return
}

func (c *Cache) Remove(key string) ([]byte, bool) {
	val, ok := c.cacheEntries[key]

	if !ok {
		return nil, false
	}
		
	delete(c.cacheEntries, key)
	return val.val, true
}

func (c *Cache) Get(key string) ([]byte, bool) {
	val, ok := c.cacheEntries[key]

	if !ok {
		return nil, false
	}
		
	return val.val, true
}
	

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <- ticker.C:
				c.mu.Lock()
				for key, entry := range c.cacheEntries {
					timePassed := time.Since(entry.createdAt);

					if timePassed > interval {
						c.Remove(key)	
					}
				}
				c.mu.Unlock()
			}
		}
	}()
}
