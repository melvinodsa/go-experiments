package cache

import "sync"

/*
 * Thsi file cotains the definitions for the high performant cache
 */

type rwMutexCache struct {
	cache map[string]interface{}
	mux   sync.RWMutex
}

func NewRWMutexCache() *rwMutexCache {
	return &rwMutexCache{
		cache: map[string]interface{}{},
	}

}

//GetCacheValue gets the cache value
func (c *rwMutexCache) Get(key string) (interface{}, bool) {
	c.mux.RLock()
	v, ok := internalCache.cache[key]
	c.mux.RUnlock()
	return v, ok
}

//PutCacheValue puts in the cache value
func (c *rwMutexCache) Put(key string, value interface{}) {
	c.mux.Lock()
	c.cache[key] = value
	c.mux.Unlock()
}
