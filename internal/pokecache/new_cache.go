package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		data:     make(map[string]cacheEntry),
		interval: interval,
	}

	go cache.reapLoop()

	return cache
}
