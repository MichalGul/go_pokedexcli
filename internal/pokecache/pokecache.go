package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entryData  map[string]CacheEntry
	mutex      *sync.RWMutex
	timeToLive time.Duration
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var entry = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	fmt.Printf("Adding Entry %s to Cache \n", key)
	c.entryData[key] = entry

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	element, exists := c.entryData[key]
	if exists != true {
		fmt.Printf("Entry %s not found in Cache \n", key)
		return make([]byte, 0), false
	}
	fmt.Printf("Cache hit for %s !\n", key)
	return element.val, exists

}

func (c *Cache) reapLoop() {

	ticker := time.NewTicker(c.timeToLive)
	defer ticker.Stop()

	for {
		<-ticker.C // wait for tick
		c.mutex.Lock()
		for key, entry := range c.entryData {
			if time.Now().Sub(entry.createdAt) > c.timeToLive {
				delete(c.entryData, key)
				fmt.Printf("Delete entry %s from cache \n", key)
			}
		}
		c.mutex.Unlock()
	}

}

func NewCache(interval time.Duration) Cache {
	mu := &sync.RWMutex{}
	cacheMap := map[string]CacheEntry{}

	var cacheData = Cache{
		entryData:  cacheMap,
		mutex:      mu,
		timeToLive: interval,
	}
	go cacheData.reapLoop()

	return cacheData
}
